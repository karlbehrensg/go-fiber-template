package main

import (
	"github.com/karlbehrensg/go-fiber-template/cmd/api/internal/http"
	_ "github.com/karlbehrensg/go-fiber-template/docs"
)

// @title Fiber Example API
// @version 0.0.1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/

// @contact.name Zeleri
// @contact.url http://www.swagger.io/support
// @contact.email edwyn.rangel.externo@zeleri.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	http.Start()
}
