package main

import (
	routes "maintenance-products/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	// calles as localhost:3000/api/v1/postForm
	v1.POST("/postForm", routes.Create)

	// called as localhost:3000/api/v1/getOne/:id
	v1.GET("getOne/:Id", routes.ReadOne)

	// called as localhost:3000/api/v1/getAll
	v1.GET("getAll", routes.ReadAll)

	// called as localhost:3000/api/v1/update/:id
	v1.PUT("/update/:Id", routes.Update)

	// called as localhost:3000/api/v1/delete/:id
	v1.DELETE("/delete/:Id", routes.Delete)

	router.Run("localhost: 3000")

}
