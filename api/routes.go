package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"practice-go/database"
	"practice-go/model"
)

func GenerateRoutes(db *sql.DB) *gin.Engine {
	// Create the engine object.
	routes := gin.Default()

	// Create all routes.

	// ROUTES - GET

	routes.GET("/events", func(context *gin.Context) {
		//fmt.Println("Fetching events...")
		//chapter_id := context.Param("chapter_id")

		events, err := database.GetEvents(db, "")

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "List of events",
			"events":  events,
		})
		return
		//context.JSON(200, gin.H{"message": "List of events"})
	})
	routes.GET("/events/:id", func(context *gin.Context) {
		eventID := context.Param("id")
		events, err := database.GetEvents(db, "id = "+eventID)

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		for _, event := range events {
			context.JSON(200, gin.H{
				"message": "Event found",
				"data":    event,
			})
			return
		}
		// No value found.
		context.JSON(404, gin.H{
			"message": "Event not found.",
		})
		return
	})
	routes.GET("/chapters", func(context *gin.Context) {
		chapters, err := database.GetChapters(db)

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message":  "List of chapters",
			"chapters": chapters,
		})
		return
	})

	// ROUTES - POST

	routes.POST("/events", func(context *gin.Context) {
		var newEvent model.Event

		// Gather variables.
		var jsonFields map[string]interface{}

		// Read the JSON body
		body, err := io.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(400, gin.H{"error": "Failed to read body"})
			return
		}

		// Decode JSON into map
		if err := json.Unmarshal(body, &jsonFields); err != nil {
			context.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		// Alright, we have the fields. Christ.

		name := jsonFields["name"].(string)
		fmt.Println("name is " + name)

		// TODO: You got the fields. Now just like... Use them for a mysql call. Good night.

		// Save to database.
		eventID, err := database.CreateEvent(db, newEvent)
		if err != nil {
			context.JSON(400, gin.H{"error": "mysql error - " + err.Error()})
			return
		}

		// Success.
		newEvent.ID = eventID

		// Respond.
		context.JSON(201, gin.H{
			"message": "Creation successful.",
			"data":    newEvent,
		})
		return
	})

	// ROUTES - PUT - REPLACEMENT

	// ROUTES - PATCH - UPDATES

	routes.PATCH("/events/:id", func(context *gin.Context) {
		//eventID := context.Param("id")
		//var partialEvent model.Event

	})

	// ROUTES - DELETE

	// Done creating routes. Return.
	return routes
}
