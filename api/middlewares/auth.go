package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/rainbow777/todolist/envconfig"
	"github.com/rainbow777/todolist/myerrors"
	"google.golang.org/api/idtoken"
)

func ValidateIDtoken(IDtoken string) (string, error) {
	// IDtokenを検証するメソッドを持つ構造体を生成
	validator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		err = myerrors.CannotMakeValidator.Wrap(err, "Failed to make Validator")
		return "", err
	}

	// IDtokenを検証し、ペイロードの情報を入手
	payload, err := validator.Validate(context.Background(), IDtoken, envconfig.AppConfig.GoogleClientID)
	if err != nil {
		err = myerrors.Unauthorizated.Wrap(err, "Failed to validate IDtoken")
		return "", err
	}

	// ペイロード内のユーザー名をリクエストのコンテキストに格納
	username, ok := payload.Claims["name"]
	if !ok {
		err = myerrors.Unauthorizated.
			Wrap(errors.New("invalid req header"), "Failed to get username from the ID token payload")
		return "", err
	}

	return username.(string), nil
}

func AuthHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// リクエストヘッダーからAuthorizationフィールドの情報を入手
		authStr := req.Header.Get("Authorization")

		// AuthorizationフィールドにBearerとIDtokenが格納されているかを確認
		splitAuth := strings.Split(authStr, " ")
		if len(splitAuth) != 2 {
			err := myerrors.RequiredAuthorizationHeader.
				Wrap(errors.New("invalid req header"), "Authorization Header must be two string 「Bearer IDtoken」")
			myerrors.ErrorHandler(w, err)
			return
		}

		if splitAuth[0] != "Bearer" || splitAuth[1] == "" {
			err := myerrors.RequiredAuthorizationHeader.
				Wrap(errors.New("invalid req header"), "Bearer and IDtoken are required")
			myerrors.ErrorHandler(w, err)
			return
		}

		IDtoken := splitAuth[1]

		// IDトークンの検証
		username, err := ValidateIDtoken(IDtoken)
		if err != nil {
			myerrors.ErrorHandler(w, err)
			return
		}

		req = SetUserName(req, username)
		next.ServeHTTP(w, req)
	})

}

func SetUserName(req *http.Request, username string) *http.Request {
	ctx := req.Context()

	ctx = context.WithValue(ctx, userNameKey{}, username)

	return req.WithContext(ctx)
}

func GetUserName(ctx context.Context) string {
	name := ctx.Value(userNameKey{})
	return name.(string)
}

type userNameKey struct{}
