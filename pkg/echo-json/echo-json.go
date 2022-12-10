package echoJSON

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mailru/easyjson"

	"github.com/labstack/echo/v4"
)

// Serializer implements JSON encoding using easyjson.
type Serializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (s Serializer) Serialize(c echo.Context, i any, _ string) error {
	marshal, err := json.Marshal(i)
	if err != nil {
		return err
	}
	_, err = c.Response().Write(marshal)

	return err
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (s Serializer) Deserialize(c echo.Context, i any) error {
	el, ok := i.(easyjson.Unmarshaler)
	if !ok {
		err := json.NewDecoder(c.Request().Body).Decode(i)
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
		} else if se, ok := err.(*json.SyntaxError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
		}
		return err
	}
	err := easyjson.UnmarshalFromReader(c.Request().Body, el)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unmarshal type error").SetInternal(err)
	}
	return nil
}
