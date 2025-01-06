package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type RoutiqClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// アクセストークンの発行
func GenerateAccessToken(userID string, role string, expiresIn time.Duration) (string, error) {
	// クレームを設定
	claims := &RoutiqClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)), // 有効期限
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // 発行時刻
			NotBefore: jwt.NewNumericDate(time.Now()),                // 有効開始時刻
			Issuer:    "routiq",                                      // 発行者名
		},
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// リフレッシュトークンの発行
func GenerateRefreshToken() (string, error) {
	// UUIDを使用してリフレッシュトークンを生成
	refreshToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return refreshToken.String(), nil
}

// 暗号学的ランダムトークン生成の例（オプション）
func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func VerifyToken(tokenStr string) (*RoutiqClaims, error) {
	claims := &RoutiqClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, xerrors.New("invalid token")
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, xerrors.New("token has expired")
	}
	return claims, nil
}
