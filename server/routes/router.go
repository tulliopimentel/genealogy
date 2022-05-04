package routes

import (
	"example/desafio/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "example/desafio/docs"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{

	main := router.Group("/v1")
	{
		people := main.Group("person")
		{
			people.GET("/:id", service.ShowFamily)
			people.GET("/", service.ShowPeoples)
			people.POST("/", service.CreatePerson)
		

			people.GET("/genealogy/:id", service.ShowGenealogy)
			people.POST("/genealogy", service.CreateGenealogy)

			people.POST("/family", service.CreateFamily)
		}

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
