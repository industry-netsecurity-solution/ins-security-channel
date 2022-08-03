package cecho

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/industry-netsecurity-solution/ins-security-channel/logger"
	echo "github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	CustomContext struct {
		echo.Context
		Ud interface{}
	}
)

func (c *CustomContext) GetUserData() interface{} {
	return c.Ud
}

func Start(config *ins.ServiceConfigurations, ud interface{}, callback func(*echo.Echo)) *echo.Echo {
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
			certpool := x509.NewCertPool()
			if 0 < len(config.CaCert) {
				pemCerts, err := ioutil.ReadFile(config.CaCert)
				if err == nil {
					certpool.AppendCertsFromPEM(pemCerts)
				}
			}

			// 2022-08-03
			address := fmt.Sprintf("%s:%d", config.Address, config.Port)
			//e.Logger.Fatal(e.StartTLS(address, config.TlsCert, config.TlsKey))

			cert, err := tls.LoadX509KeyPair(config.TlsCert, config.TlsKey)
			if err != nil {
				logger.Error(err)
				return
			}
			if len(cert.Certificate) == 0 {
				return
			}
			// Just to print out the client certificate..
			cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
			if err != nil {
				logger.Error(err)
				return
			}

			config := &tls.Config{
				MinVersion:       tls.VersionTLS12,
				CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_CHACHA20_POLY1305_SHA256,
					tls.TLS_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
				RootCAs:      certpool,
				Certificates: []tls.Certificate{cert},
			}
			server := &http.Server{
				Addr:         address,
				TLSConfig:    config,
				TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			}

			e.Logger.Fatal(e.StartServer(server))
		} else {
			e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Address, config.Port)))
		}
	}()

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
