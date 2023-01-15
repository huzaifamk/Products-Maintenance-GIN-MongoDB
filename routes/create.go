package routes

import (
	"context"
	"log"
	getcollection "maintenance-products/collections"
	database "maintenance-products/database"
	model "maintenance-products/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
 
func Create(c *gin.Context) {
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "forms")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Object)
	defer cancel()
 
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
 
	postPayload := model.Object{
		ID:      primitive.NewObjectID(),
		Name:   post.Name,
		Form_Fields: post.Form_Fields,
	}
 
	result, err := postCollection.InsertOne(ctx, postPayload)
 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}