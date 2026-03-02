package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/user/quiz-admin/backend/controllers"
	"github.com/user/quiz-admin/backend/database"
)

func main() {
	database.Connect()

	app := fiber.New()

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api")
	{
		api.Get("/exams", controllers.GetExams)
		api.Post("/exams", controllers.CreateExam)
		api.Delete("/exams/:id", controllers.DeleteExam)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
