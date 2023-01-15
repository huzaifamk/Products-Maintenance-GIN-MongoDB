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
 
func Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "forms")
 
	Id := c.Param("Id")
	var post model.Object
 
	defer cancel()
 
	objId, _ := primitive.ObjectIDFromHex(Id)
 
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	edited := bson.M{"Name": post.Name, "Fields": post.Form_Fields}
 
	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})
 
	res := map[string]interface{}{"data": result}
 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
 
	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}
 
	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}