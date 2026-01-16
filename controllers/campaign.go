package controllers

import (
	"fmt"
	"time"
	"votepeace/database"
	"votepeace/models"

	"github.com/gofiber/fiber/v2"
)

// GetCampaigns returns all campaigns
// GetCampaigns returns all campaigns with vote counts
func GetCampaigns(c *fiber.Ctx) error {
	var campaigns []models.Campaign
	// Preload Candidates
	result := database.DB.Preload("Candidates").Find(&campaigns)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not retrieve campaigns"})
	}

	// Create a response structure that matches what frontend expects + votesCast
	type CampaignResponse struct {
		models.Campaign
		VotesCast   int64 `json:"votesCast"`
		TotalVoters int64 `json:"totalVoters"` // This might be used for participation rate
	}

	var response []CampaignResponse

	for _, camp := range campaigns {
		var count int64
		// Count votes for this campaign
		database.DB.Model(&models.Vote{}).Where("campaign_id = ?", camp.ID).Count(&count)

		// For TotalVoters, we might check users count, but currently we just use Total Users globally or per campaign assignment.
		// For this simple app, let's assume TotalVoters is the global user count (as any user can vote).
		var totalUsers int64
		database.DB.Model(&models.User{}).Count(&totalUsers)

		response = append(response, CampaignResponse{
			Campaign:    camp,
			VotesCast:   count,
			TotalVoters: totalUsers,
		})
	}

	return c.JSON(response)
}

// GetCampaign returns a single campaign by ID
func GetCampaign(c *fiber.Ctx) error {
	id := c.Params("id")
	var campaign models.Campaign
	result := database.DB.Preload("Candidates").First(&campaign, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Campaign not found"})
	}
	return c.JSON(campaign)
}

// CreateCampaignInput defines the expected body for creating a campaign
type CreateCampaignInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"` // Receive as string, parse to time
	EndDate     string `json:"endDate"`   // Receive as string, parse to time
}

// CreateCampaign adds a new campaign (Admin only)
func CreateCampaign(c *fiber.Ctx) error {
	var input CreateCampaignInput
	if err := c.BodyParser(&input); err != nil {
		fmt.Println("Error parsing body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	fmt.Println("Received CreateCampaign Data:", input) // DEBUG

	if input.Title == "" || input.Description == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Title and Description are required"})
	}

	// Parse Dates
	layout := "2006-01-02" // YYYY-MM-DD format
	startDate, err := time.Parse(layout, input.StartDate)
	if err != nil {
		// Fallback or returned error if critical. For now default to Now if empty/invalid?
		// Better to return error if user explicitly wants to set it.
		// But if empty, we can default.
		if input.StartDate == "" {
			startDate = time.Now()
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid Start Date format (YYYY-MM-DD)"})
		}
	}

	endDate, err := time.Parse(layout, input.EndDate)
	if err != nil {
		if input.EndDate == "" {
			endDate = time.Now().AddDate(0, 1, 0)
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid End Date format (YYYY-MM-DD)"})
		}
	}

	campaign := models.Campaign{
		Title:       input.Title,
		Description: input.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		IsActive:    true,
	}

	if err := database.DB.Create(&campaign).Error; err != nil {
		fmt.Println("Error creating campaign in DB:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Could not create campaign"})
	}

	return c.JSON(campaign)
}

// DeleteCampaign removes a campaign (Admin only)
func DeleteCampaign(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Campaign{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete campaign"})
	}
	return c.JSON(fiber.Map{"message": "Campaign deleted successfully"})
}

// CreateCandidateInput defines the expected body for creating a candidate
type CreateCandidateInput struct {
	CampaignID uint   `json:"campaignId"`
	Number     int    `json:"number"`
	Name       string `json:"name"`
	Vision     string `json:"vision"`
	Mission    string `json:"mission"`
	Photo      string `json:"photo"`
}

// CreateCandidate adds a new candidate to a campaign
func CreateCandidate(c *fiber.Ctx) error {
	var input CreateCandidateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	// Basic Validation
	if input.CampaignID == 0 || input.Name == "" || input.Number == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "CampaignID, Name, and Number are required"})
	}

	candidate := models.Candidate{
		CampaignID: input.CampaignID,
		Number:     input.Number,
		Name:       input.Name,
		Vision:     input.Vision,
		Mission:    input.Mission,
		Photo:      input.Photo,
	}

	if err := database.DB.Create(&candidate).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create candidate"})
	}

	return c.JSON(candidate)
}

// GetCampaignVotes returns the list of voters for a specific campaign
func GetCampaignVotes(c *fiber.Ctx) error {
	id := c.Params("id")
	var votes []models.Vote
	// Preload User and Candidate to get names
	result := database.DB.Preload("User").Preload("Candidate").Where("campaign_id = ?", id).Order("timestamp desc").Find(&votes)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not retrieve votes"})
	}
	return c.JSON(votes)
}

// GetCampaignSummary returns aggregated results for a campaign
func GetCampaignSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	// 1. Get Campaign with Candidates
	var campaign models.Campaign
	if err := database.DB.Preload("Candidates").First(&campaign, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Campaign not found"})
	}

	// 2. Count votes for each candidate
	// We can use a Group By query or just simple counts if volume is low.
	// For "Quick Count", we want accurate numbers.

	type Result struct {
		CandidateID uint   `json:"candidateId"`
		Name        string `json:"name"`
		Number      int    `json:"number"`
		Photo       string `json:"photo"`
		Votes       int64  `json:"votes"`
	}

	var results []Result

	for _, candidate := range campaign.Candidates {
		var count int64
		database.DB.Model(&models.Vote{}).Where("candidate_id = ?", candidate.ID).Count(&count)

		results = append(results, Result{
			CandidateID: candidate.ID,
			Name:        candidate.Name,
			Number:      candidate.Number,
			Photo:       candidate.Photo,
			Votes:       count,
		})
	}

	// Calculate total for percentage
	var totalVotes int64
	database.DB.Model(&models.Vote{}).Where("campaign_id = ?", id).Count(&totalVotes)

	return c.JSON(fiber.Map{
		"campaign":   campaign.Title,
		"totalVotes": totalVotes,
		"results":    results,
	})
}
