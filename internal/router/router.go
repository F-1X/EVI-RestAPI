package router

import (
	docs "advertisement-rest-api-http-service/cmd/docs"
	v1 "advertisement-rest-api-http-service/internal/handler/api/v1"
	"advertisement-rest-api-http-service/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
type GinRouter struct {
	engine *gin.Engine
	server *http.Server
	quit   chan os.Signal
}

func NewGinRouter() *GinRouter {

	return &GinRouter{
		engine: gin.Default(),
		server: &http.Server{},
		quit:   make(chan os.Signal, 1),
	}
}

func (r *GinRouter) Run(port string) {
	r.server.Handler = r.engine
	r.server.Addr = ":" + port

	signal.Notify(r.quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	<-r.quit

	log.Println("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %s\n", err)
	}

	log.Println("Server stopped gracefully")
}

func (r *GinRouter) Shutdown() {
	r.quit <- syscall.SIGTERM
}

func (r *GinRouter) AddHandlers(service service.AdServicer) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := v1.NewAdHandler(service)
	apiv1 := r.engine.Group("/api/v1")
	{
		apiv1.GET("/ads", v1.GetAds)
		apiv1.GET("/ad/:id", v1.GetAd)
		apiv1.POST("/ad", v1.CreateAd)
	}
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
}
