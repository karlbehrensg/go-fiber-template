package http

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/config/database"
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http/routes"
	"github.com/karlbehrensg/go-fiber-template/internal/domain"
	"github.com/karlbehrensg/go-fiber-template/pkg/logger"
	"github.com/karlbehrensg/go-fiber-template/pkg/utils"
)

// start run app
func Start() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			logger.Error(fmt.Sprintf("Error godotenv %s", err.Error()))
		}
	}

	// validate env
	utils.CheckEnv()

	// get client db
	dbClient := database.GetDbClient()
	// run migration
	database.Migrate(dbClient, &domain.Author{}, &domain.Book{}, &domain.User{})

	// instantiating fiber
	app := fiber.New()
	// added middleware
	app.Use(recover.New())
	app.Use(fiberLogger.New())

	// define routes
	routes.SwaggerRoutes(app)
	routes.AuthRoutes(app, dbClient)
	routes.AuthorRoutes(app, dbClient)
	routes.BookRoutes(app, dbClient)
	routes.NotFoundRoute(app)

	// run server
	err := app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		logger.Fatal(err.Error())
	}
}
