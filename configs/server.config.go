package configs

import (
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App

	//db database.Service
}

func CreateServer() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "yu-yummy-be-core",
			AppName:      "yu-yummy-be-core",
		}),

		//db: database.New(),
	}

	return server
}
