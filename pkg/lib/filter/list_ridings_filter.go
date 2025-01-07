package filter

import (
	"log"
	"log/slog"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/operators"
	"golang.org/x/xerrors"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var celEnv *cel.Env

type RidingFilter struct {
	RouteID string
	UserID  string
}

func init() {
	var err error
	celEnv, err = cel.NewEnv(
		cel.Declarations(
			decls.NewVar("route_id", decls.String),
			decls.NewVar("user_id", decls.String),
		),
	)
	if err != nil {
		log.Panicf("failed creating cel environment: %v", err)
	}
}

func ParseListRidingsFilter(filterStr string) (*RidingFilter, error) {
	ast, issues := celEnv.Compile(filterStr)
	if issues != nil && issues.Err() != nil {
		return nil, xerrors.Errorf("Error compiling filter: %w", issues.Err())
	}

	// 評価結果を RidingFilter に変換
	// AST ノードからフィールドと値を抽出
	extracted := make(map[string]interface{})
	cel.AstToParsedExpr(ast)
	if err := parseAST(ast.Expr(), extracted); err != nil {
		return nil, xerrors.Errorf("failed to parse AST: %w", err)
	}
	filter := &RidingFilter{}
	if routeID, ok := extracted["route_id"]; ok {
		filter.RouteID = routeID.(string)
	}
	log.Printf("filter: %v", filter)
	slog.Info("filter", slog.Any("filter", filter))
	return filter, nil
}

// AST を再帰的に解析する
func parseAST(expr *exprpb.Expr, result map[string]interface{}) error {
	switch expr.ExprKind.(type) {
	case *exprpb.Expr_CallExpr:
		call := expr.GetCallExpr()
		if call.Function == operators.Equals { // "==" 演算子
			// 左辺がフィールド名、右辺が値と仮定
			field := call.Args[0]
			value := call.Args[1]

			// フィールド名を取得
			fieldName, err := extractFieldName(field)
			if err != nil {
				return err
			}

			// 値を取得
			fieldValue, err := extractValue(value)
			if err != nil {
				return err
			}

			// 結果に追加
			result[fieldName] = fieldValue
		}
	}
	return nil
}

// フィールド名を抽出
func extractFieldName(expr *exprpb.Expr) (string, error) {
	if selectExpr, ok := expr.ExprKind.(*exprpb.Expr_IdentExpr); ok {
		return selectExpr.IdentExpr.Name, nil
	}
	return "", xerrors.Errorf("unsupported field expression: %w", expr)
}

// 値を抽出
func extractValue(expr *exprpb.Expr) (interface{}, error) {
	if constExpr, ok := expr.ExprKind.(*exprpb.Expr_ConstExpr); ok {
		constant := constExpr.ConstExpr
		switch constant.ConstantKind.(type) {
		case *exprpb.Constant_StringValue:
			return constant.GetStringValue(), nil
		case *exprpb.Constant_Int64Value:
			return constant.GetInt64Value(), nil
		case *exprpb.Constant_BoolValue:
			return constant.GetBoolValue(), nil
		}
	}
	return nil, xerrors.Errorf("unsupported value expression: %w", expr)
}
