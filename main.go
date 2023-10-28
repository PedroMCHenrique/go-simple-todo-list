package main

import (
	"github.com/pedromchenrique/todo-list/config"
	"github.com/pedromchenrique/todo-list/router"
)

// @title           Simple Todo-list API
// @version         1.0
// @description     This is a simple Todo-list api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Initializing error: %v", err)
		return
	}

	router.Init()
}
