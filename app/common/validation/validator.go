package validation

import (
	"net/http"

	"github.com/go-playground/validator"

	"github.com/labstack/echo"
)

// validatorパッケージを使用することで構造体のデータのバリデーションが可能になります。
type CustomValidator struct {
	V *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.V.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
