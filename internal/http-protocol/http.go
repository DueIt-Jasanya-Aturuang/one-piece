package httpprotocol

import (
	"fmt"

	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/config"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
)

type HttpImpl struct {
	HttpHandler *handlers.HttpHandlerImpl
}

func NewHttpImpl(
	httpHandler *handlers.HttpHandlerImpl,
) *HttpImpl {
	return &HttpImpl{
		HttpHandler: httpHandler,
	}
}

func (h *HttpImpl) Listen() {
	app := fiber.New()

	app.Use(
		logger.New(),
	)

	h.HttpHandler.HandlerRouter(app)

	servePort := fmt.Sprintf(":%s", config.Get().Application.Port)
	if err := app.Listen(servePort); err != nil {
		panic(err)
	}

	log.Info().Msgf("serve started on port %s", servePort)
}
