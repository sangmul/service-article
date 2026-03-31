package main

import (
	"article/internal/config"
	"article/internal/handler"
	"article/internal/repository"
	"article/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := config.InitDB()

	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)
	h := handler.NewPostHandler(svc)

	r := gin.Default()

	article := r.Group("/article")
	{
		article.POST("", h.Create)       //end point 1
		article.GET("", h.GetAll)        //end point 2
		article.GET("/:id", h.GetByID)   //end point 3
		article.PUT("/:id", h.Update)    //end point 4a
		article.PATCH("/:id", h.Update)  //end point 4b
		article.DELETE("/:id", h.Delete) //end point 5

	}

	r.Run(":8080")
}
