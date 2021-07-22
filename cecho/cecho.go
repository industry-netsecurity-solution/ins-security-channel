package cecho

import (
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/labstack/echo/v4"
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
