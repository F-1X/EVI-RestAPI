package main

import (
	firebaseRepository "advertisement-rest-api-http-service/internal/repository/firebase" // docs is generated by Swag CLI, you have to import it.
	"advertisement-rest-api-http-service/internal/router"
	"advertisement-rest-api-http-service/internal/service"
	database "advertisement-rest-api-http-service/pkg/firebase"
	"context"
	"log"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8888
//	@BasePath	/api/v1

func main() {

	ctx := context.Background()

	// init database
	firestoreClient, err := database.InitFirestoreClient(ctx, "key-firebase.json")
	if err != nil {
		log.Fatal("err: firestoreClient", err)
	}
	defer firestoreClient.Close(ctx)

	// init repository
	repo := firebaseRepository.NewAdRepository(firestoreClient)

	// init service
	service := service.NewAdService(repo)

	// init router
	router := router.NewGinRouter()

	router.AddHandlers(service)

	router.Run("8888")
}