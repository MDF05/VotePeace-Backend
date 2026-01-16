package controllers

import (
	"votepeace/database"
	"votepeace/models"

	"github.com/gofiber/fiber/v2"
)

func GetCandidates(c *fiber.Ctx) error {
	var candidates []models.Candidate
	database.DB.Find(&candidates)
	return c.JSON(candidates)
}

func GetStats(c *fiber.Ctx) error {
	var totalUsers int64
	var totalVotes int64
	// Participation logic can be refined later

	database.DB.Model(&models.User{}).Count(&totalUsers)
	database.DB.Model(&models.Vote{}).Count(&totalVotes)

	return c.JSON(fiber.Map{
		"users":         totalUsers,
		"votes":         totalVotes,
		"participation": 0.0, // Placeholder
	})
}
