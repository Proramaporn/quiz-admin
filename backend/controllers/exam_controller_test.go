package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/user/quiz-admin/backend/database"
	"github.com/user/quiz-admin/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() {
	_ = godotenv.Load("../../.env")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME_TEST")

	if dbName == "" {
		dbName = "quiz_admin_test"
	}

	// Connect to MySQL server without selecting a database
	dsnServer := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort)

	db, err := gorm.Open(mysql.Open(dsnServer), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to MySQL server: %v\n", err)
		os.Exit(1)
	}

	// Create test database if it doesn't exist
	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))

	// Reconnect to the test database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to test database: %v\n", err)
		os.Exit(1)
	}

	database.DB.AutoMigrate(&models.Exam{}, &models.Choice{})

	// Clean up tables before each test
	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	database.DB.Exec("TRUNCATE TABLE choices")
	database.DB.Exec("TRUNCATE TABLE exams")
	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func setupApp() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	api.Get("/exams", GetExams)
	api.Post("/exams", CreateExam)
	api.Delete("/exams/:id", DeleteExam)
	return app
}

func TestGetExams(t *testing.T) {
	setupTestDB()
	app := setupApp()

	// Seed
	exam := models.Exam{
		Question: "What is 2+2?",
		Choices: []models.Choice{
			{ChoiceText: "3"}, {ChoiceText: "4"}, {ChoiceText: "5"}, {ChoiceText: "6"},
		},
	}
	database.DB.Create(&exam)

	req := httptest.NewRequest("GET", "/api/exams", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var exams []models.Exam
	json.NewDecoder(resp.Body).Decode(&exams)
	assert.Len(t, exams, 1)
	assert.Equal(t, "What is 2+2?", exams[0].Question)
}

func TestCreateExam(t *testing.T) {
	setupTestDB()
	app := setupApp()

	input := CreateExamInput{
		Question: "Is Go fast?",
		Choices:  []string{"Yes", "No", "Maybe", "Depends"},
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest("POST", "/api/exams", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var exam models.Exam
	database.DB.Preload("Choices").First(&exam)
	assert.Equal(t, "Is Go fast?", exam.Question)
	assert.Len(t, exam.Choices, 4)
}

func TestDeleteExam(t *testing.T) {
	setupTestDB()
	app := setupApp()

	exam := models.Exam{Question: "Delete me?"}
	database.DB.Create(&exam)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/exams/%d", exam.ID), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	var count int64
	database.DB.Model(&models.Exam{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
