package configuration

import (
	"blog/domains/posts/controller"
	"blog/domains/posts/repository"
	"blog/domains/posts/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

type Router struct {
	Engine *gin.Engine
	Port   string
}

func CreateRouter() (*Router, error) {
	engine := gin.Default()
	config := LoadConfig()
	if config == nil {
		return nil, fmt.Errorf("Failed to load configuration")
	}

	db := ConnectDB(*config)
	if db == nil {
		return nil, fmt.Errorf("Failed to connect to the database")
	}

	createEndPoints(engine, db)

	return &Router{
		Engine: engine,
		Port:   config.Port,
	}, nil
}

func (router *Router) Run() {
	fmt.Println("Server is running on", router.Port)
	if err := router.Engine.Run(":" + router.Port); err != nil {
		log.Fatalln("Error running server: ", err)
	}
}

func createEndPoints(engine *gin.Engine, db *pg.DB) {
	postsController := setUpPostsLayers(db)

	addPostsHandler(engine, postsController)
}

func setUpPostsLayers(db *pg.DB) *controller.PostsController {
	postsRepo, err := repository.CreatePostsRepository(db)

	if err != nil {
		log.Fatalf("failed to create posts repository: %v", err)
	}

	postsService := service.CreatePostsService(postsRepo)

	postsController := controller.CreatePostsController(postsService)

	return postsController
}

func addPostsHandler(engine *gin.Engine, postsController *controller.PostsController) {
	postsGroup := engine.Group("/posts")
	{
		postsGroup.POST("", postsController.CreatePost)
		postsGroup.PUT("/:id", postsController.UpdatePost)
		postsGroup.GET("/:id", postsController.GetPost)
		postsGroup.GET("", postsController.GetPosts)
		postsGroup.DELETE("/:id", postsController.DeletePost)
	}
}
