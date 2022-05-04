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
		person := main.Group("person")
		{
			person.GET("/:id", service.ShowFamily)
			person.GET("/", service.ShowPeoples)
			person.POST("/", service.CreatePerson)
			person.DELETE("/", service.DeletePerson)
			person.PUT("/")
			person.GET("/genealogy/:id", service.ShowGenealogy)
			person.POST("/genealogy", service.CreateGenealogy)
			person.POST("/family", service.CreateFamily)
			person.DELETE("/family", service.DeleteFamily)
		}

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
