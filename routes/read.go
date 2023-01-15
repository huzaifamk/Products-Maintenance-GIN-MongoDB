package routes

import (
	"context"
	getcollection "maintenance-products/collections"
	database "maintenance-products/database"
	model "maintenance-products/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
 
func ReadOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "forms")
 
	Id := c.Param("Id")
	var result model.Object
 
	defer cancel()
 
	objId, _ := primitive.ObjectIDFromHex(Id)
 
	err := postCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&result)
 
	res := map[string]interface{}{"data": result}
 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"message": "success!", "Data": res})
}

func ReadAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "forms")
 
	var results []*model.Object
 
	defer cancel()
 
	cursor, err := postCollection.Find(ctx, bson.D{})
 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	for cursor.Next(ctx) {
		var result model.Object
		cursor.Decode(&result)
		results = append(results, &result)
	}
 
	res := map[string]interface{}{"data": results}
 
	c.JSON(http.StatusOK, gin.H{"message": "success!", "Data": res})
}