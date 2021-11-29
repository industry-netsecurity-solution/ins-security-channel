package cecho

import (
	"encoding/json"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type (
	CustomContext struct  {
		echo.Context
		Ud interface{}
	}
)

func (c *CustomContext) GetUserData() interface{} {
	return c.Ud
}

func Start(config *ins.ServiceConfigurations, ud interface{}, callback func (*echo.Echo)) (*echo.Echo) {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{Context: c, Ud: ud}
			return next(cc)
		}
	})

	callback(e)

	go func() {
		if config.EnableTls {
			address := fmt.Sprintf("%s:%d", config.Address, config.Port)
			e.Logger.Fatal(e.StartTLS(address, config.TlsCert, config.TlsKey))
		} else {
			e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Address, config.Port)))
		}
	} ()

	return e
}

func Unmarshal(c echo.Context, v interface{}) (err error) {
/*
	err := c.Bind(v)
	return err
*/
	req := c.Request()
	ctype := req.Header.Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationJSON) == false {
		return echo.ErrUnsupportedMediaType
	}

	if err = json.NewDecoder(req.Body).Decode(v); err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
		} else if se, ok := err.(*json.SyntaxError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	return nil
}