package main

import (
	"fmt"
	"strconv"

	"realworld-backend/article"
	"realworld-backend/shared"

	"github.com/gin-gonic/gin"
)





var ArticleRepository = article.InMemoryArticleRepository()

var idCounter = 0

func generateId() string {
	idCounter++
	return strconv.Itoa(idCounter)
}


func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"status": "OK",
		})
	})

	r.GET("/api/articles", func (c *gin.Context) {
		c.JSON(200, ArticleRepository.FindAll())
	})

	r.GET("/api/articles/:slug", func (c *gin.Context) {
		slug := c.Param("slug");

		c.JSON(200, ArticleRepository.FindBySlug(slug))
	})	

	r.POST("/api/articles", func(c *gin.Context) {
		articleInput := new(article.ArticleInput)
		err := c.BindJSON(articleInput)
		if err != nil {
			fmt.Println("Error")
			c.Status(400)
			return
		}

		article, err := article.CreateArticle(ArticleRepository, shared.UUIDGenerator)(articleInput)
		
		if err != nil {
			c.Status(400)
		} else {
			c.JSON(201, article)
		}
		
		
	})

	r.Run()

}