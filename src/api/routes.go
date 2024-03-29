package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"practice-go/database"
	"practice-go/model"
	"strconv"
)

func GenerateRoutes(db *sql.DB) *gin.Engine {
	// Create the engine object.
	routes := gin.Default()

	// Create all routes.

	////////////////////////////////////////////////////////////////
	//// ROUTES - GET //////////////////////////////////////////////
	////////////////////////////////////////////////////////////////

	// HACKING ////////////////////////////////

	routes.GET("/hack/character/:id/credentials", func(context *gin.Context) {
		charID, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		data, err := database.GetHackCharacterCredentials(db, charID)

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "Credentials found.",
			"data":    data,
		})
		return
	})

	routes.GET("/hack/character/:id/servers", func(context *gin.Context) {
		charID, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		servers, err := database.GetHackCharacterServers(db, charID)

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "Servers found.",
			"data":    servers,
		})
		return
	})

	routes.GET("/hack/servers", func(context *gin.Context) {
		data, err := database.GetHackServers(db, "")

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error.",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "Servers found.",
			"data":    data,
		})
		return
	})

	//// NOT HACKING ////////////////////////////////

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
			"data":    events,
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

	//// HACKING ////////////////////////////////////

	routes.POST("/hack/command/connect", func(context *gin.Context) {
		// Gather variables.
		var credential model.HackCredential
		if err := context.ShouldBindJSON(&credential); err != nil {
			context.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}

		server, err := database.HackConnectToServer(db, credential)
		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// server found.
		context.JSON(200, gin.H{
			"message": "Server connection successful.",
			"data":    server,
		})
		return

	})

	routes.POST("/hack/command/transfer", func(context *gin.Context) {
		// Gather variables.
		var jsonFields map[string]interface{}

		// Read the JSON body
		body, err := io.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(400, gin.H{"error": "Failed to read JSON body"})
			return
		}

		// Decode JSON into map
		if err := json.Unmarshal(body, &jsonFields); err != nil {
			context.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		// Variables retrieved.

		fileId := int64(jsonFields["file_id"].(float64))
		serverId := int64(jsonFields["server_id"].(float64))

		// Get the file.
		file, err := database.GetHackServerFile(db, fileId)

		if err != nil {
			context.JSON(500, gin.H{"error": "File not found."})
			return
		}

		// Attempt the upload.
		result, err := database.HackTransferFile(db, serverId, file)

		if err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Transfer successful.
		context.JSON(201, gin.H{
			"message": "File transfer successful.",
			"data":    result,
		})
		return
	})

	//// NOT HACKING ////////////////////////////////

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
