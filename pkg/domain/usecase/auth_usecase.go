package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	pkgurl "net/url"
	"time"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"github.com/asaycle/routiq.git/pkg/lib/token"
	"golang.org/x/xerrors"
	"google.golang.org/api/idtoken"
)

type CreateUserResponse struct {
	RedirectURL  string
	AccessToken  string
	RefreshToken string
	Role         string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, code string) (*CreateUserResponse, error)
}

type AuthUsecaseImpl struct {
	cfg                   *config.AuthConfig
	txManager             repository.TransactionManager
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	sessionRepository     repository.SessionRepository
}

func NewAuthUsecaseImpl(
	cfg *config.AuthConfig,
	userRepository repository.UserRepository,
	userProfileRepository repository.UserProfileRepository,
	sessionRepository repository.SessionRepository,
	txManager repository.TransactionManager,
) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		cfg:                   cfg,
		txManager:             txManager,
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		sessionRepository:     sessionRepository,
	}
}

func (u *AuthUsecaseImpl) verifyIDToken(ctx context.Context, clientID string, token string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return nil, xerrors.Errorf("Error Validate: %w", err)
	}
	return payload, nil
}

func (u *AuthUsecaseImpl) CreateUser(ctx context.Context, code string) (*CreateUserResponse, error) {
	// Googleのトークンエンドポイント
	url := "https://oauth2.googleapis.com/token"

	// クライアントID、クライアントシークレット、リダイレクトURIを設定
	clientID := u.cfg.GoogleClientID
	clientSecret := u.cfg.GoogleClientSecret
	redirectURI := u.cfg.RedirectURI

	log.Printf("clientID: %s, clientSecret: %s, redirectURI: %s", clientID, clientSecret, redirectURI)

	data := pkgurl.Values{}
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("RESP: %v", resp)

	// ステータスコードを確認
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("failed to exchange code for token: %s", resp.Status)
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	slog.Info("token response", slog.Any("resp", tokenResponse))

	payload, err := u.verifyIDToken(ctx, clientID, tokenResponse.IdToken)
	if err != nil {
		return nil, xerrors.Errorf("Error verifyIDToken: %w", err)
	}
	slog.Info("payload", slog.Any("payload", payload.Claims["picture"]))
	user, err := u.findOrCreateUser(ctx, payload.Subject)
	if err != nil {
		return nil, xerrors.Errorf("Error findOrCreateUser: %w", err)
	}
	userProfile := entity.NewUserProfile(user.ID, "test", payload.Claims["picture"].(string))
	if err := u.createUserProfile(ctx, userProfile); err != nil {
		return nil, xerrors.Errorf("Error createUserProfile: %w", err)
	}
	session, err := u.createSession(ctx, user)
	if err != nil {
		return nil, xerrors.Errorf("Error createSession: %w", err)
	}
	accessToken, err := token.GenerateAccessToken(user.ID, "USER", time.Hour)
	if err != nil {
		return nil, xerrors.Errorf("Error GenerateAccessToken: %w", err)
	}
	return &CreateUserResponse{
		RedirectURL:  "/",
		AccessToken:  accessToken,
		RefreshToken: session.RefreshToken,
		Role:         "User",
	}, nil
}

func (u *AuthUsecaseImpl) createUserProfile(ctx context.Context, userProfile *entity.UserProfile) error {
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.userProfileRepository.CreateUserProfile(ctx, tx, userProfile); err != nil {
			return xerrors.Errorf("Error CreateUserProfile: %w", err)
		}
		return nil
	})
	if err != nil {
		return xerrors.Errorf("FailedReadWriteTransaction: %w", err)
	}
	return nil
}

func (u *AuthUsecaseImpl) createSession(ctx context.Context, user *entity.User) (*entity.Session, error) {
	reftoken, err := token.GenerateRefreshToken()
	if err != nil {
		return nil, xerrors.Errorf("Error GenerateRefreshToken: %w")
	}
	session := entity.NewSession(user.ID, reftoken, time.Now().Add(time.Hour*24*30))
	err = u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.sessionRepository.CreateSession(ctx, tx, session); err != nil {
			return xerrors.Errorf("Error CreateSession: %w", err)
		}
		return nil
	})
	return session, nil
}

func (u *AuthUsecaseImpl) findOrCreateUser(ctx context.Context, subject string) (*entity.User, error) {
	var user *entity.User
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		user, err = u.userRepository.FindUserByGoogleSub(ctx, tx, subject)
		if err == nil {
			return nil
		}
		if !xerrors.Is(err, repository.NotfoundError) {
			return xerrors.Errorf("Error FindUserByGoogleSub: %w", err)
		}
		// userが存在しなかった場合は新規作成
		user = entity.NewUser(subject)
		if err := u.userRepository.CreateUser(ctx, tx, user); err != nil {
			return xerrors.Errorf("Error CreateUser: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadWriteTx: %w", err)
	}
	return user, nil
}

// TokenResponseはGoogleのトークンエンドポイントのレスポンスを表す構造体
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"` // リフレッシュトークンは省略される場合がある
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token,omitempty"` // OpenID Connectで使用されるIDトークン
}

var _ AuthUsecase = &AuthUsecaseImpl{}
