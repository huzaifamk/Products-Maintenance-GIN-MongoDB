package routes

import (
	"context"
	getcollection "maintenance-products/collections"
	database "maintenance-products/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
 
func Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	Id := c.Param("Id")
 
	var postCollection = getcollection.GetCollection(DB, "forms") 
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(Id) 
	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId}) 
	res := map[string]interface{}{"data": result}
 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}
 
	c.JSON(http.StatusCreated, gin.H{"message": "Object deleted successfully", "Data": res})
}