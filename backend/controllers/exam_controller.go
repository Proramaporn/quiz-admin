package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user/quiz-admin/backend/database"
	"github.com/user/quiz-admin/backend/models"
)

type CreateExamInput struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

func GetExams(c *fiber.Ctx) error {
	var exams []models.Exam
	if err := database.DB.Preload("Choices").Find(&exams).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch exams"})
	}
	return c.JSON(exams)
}

func CreateExam(c *fiber.Ctx) error {
	var input CreateExamInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validation
	if input.Question == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Question is required"})
	}

	if len(input.Choices) != 4 {
		return c.Status(400).JSON(fiber.Map{"error": "Exactly 4 choices are required"})
	}

	for _, choice := range input.Choices {
		if choice == "" {
			return c.Status(400).JSON(fiber.Map{"error": "All 4 choices must be non-empty"})
		}
	}

	// Transform input to models
	exam := models.Exam{
		Question: input.Question,
	}
	for _, choiceText := range input.Choices {
		exam.Choices = append(exam.Choices, models.Choice{ChoiceText: choiceText})
	}

	if err := database.DB.Create(&exam).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create exam"})
	}

	return c.Status(201).JSON(exam)
}

func DeleteExam(c *fiber.Ctx) error {
	id := c.Params("id")
	var exam models.Exam
	if err := database.DB.First(&exam, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Exam not found"})
	}

	if err := database.DB.Delete(&exam).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete exam"})
	}

	return c.SendStatus(204)
}
