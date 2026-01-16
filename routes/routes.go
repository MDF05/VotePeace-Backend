package routes

import (
	"votepeace/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/check", controllers.Check)

	app.Get("/candidates", controllers.GetCandidates)
	app.Post("/candidates", controllers.CreateCandidate)
	app.Get("/stats", controllers.GetStats)

	// Campaign Routes
	app.Get("/campaigns", controllers.GetCampaigns)
	app.Get("/campaigns/:id", controllers.GetCampaign)
	app.Post("/campaigns", controllers.CreateCampaign)
	app.Get("/campaigns/:id/votes", controllers.GetCampaignVotes)
	app.Get("/campaigns/:id/summary", controllers.GetCampaignSummary)
	app.Delete("/campaigns/:id", controllers.DeleteCampaign)
}
