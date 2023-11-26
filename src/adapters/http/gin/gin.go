package gin

import (
	"back/src/core/config"
	"back/src/core/usecases"
	"fmt"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
)

type ServiceHTTPServer struct {
	engine *gin.Engine
	config config.HTTPConfig

	ucs        usecases.Usecases
	translator *i18n.Bundle
}

func NewServiceHttpServer(ucs usecases.Usecases, config config.HTTPConfig, translator *i18n.Bundle) *ServiceHTTPServer {
	server := ServiceHTTPServer{
		config:     config,
		ucs:        ucs,
		translator: translator,
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug().Str("Method", httpMethod).Str("Handler", handlerName).Int("Handlers", nuHandlers).Msg(absolutePath)
	}

	r := gin.New()
	r.Use(ginzerolog.Logger("backend"))
	server.engine = r

	server.SetRoutes()
	return &server
}

func (server *ServiceHTTPServer) Listen() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.config.ListeningAddress, server.config.ListeningPort))

	if err != nil {
		return err
	}

	log.Info().Str("Address", ln.Addr().String()).Msg("Now Listening !")

	err = http.Serve(ln, server.engine)
	if err != nil {
		return err
	}
	err = server.engine.Run()

	if err != nil {
		return err
	}

	return nil
}
