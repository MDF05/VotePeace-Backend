package database

import (
	"log"
	"time"
	"votepeace/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("votepeace.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Could not connect to the database")
	}

	DB.AutoMigrate(&models.User{}, &models.Campaign{}, &models.Candidate{}, &models.Vote{})

	seedAdmin()
	seedCampaignsAndCandidates()
}

func seedAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "ADMIN").Count(&count)
	if count == 0 {
		password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
		admin := models.User{
			NIK:      "0000000000000000",
			Name:     "Super Admin",
			Password: string(password),
			Role:     "ADMIN",
		}
		DB.Create(&admin)
		log.Println("Admin seeded!")
	}
}

func seedCampaignsAndCandidates() {
	var campaignCount int64
	DB.Model(&models.Campaign{}).Count(&campaignCount)

	if campaignCount == 0 {
		// Create Default Campaign
		campaign := models.Campaign{
			Title:       "Pemilihan Ketua BEM 2024",
			Description: "Pemilihan umum raya untuk memilih ketua Badan Eksekutif Mahasiswa periode 2024/2025.",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(0, 1, 0), // 1 month from now
			IsActive:    true,
		}
		DB.Create(&campaign)
		log.Println("Default Campaign seeded!")

		// Create Candidates for this Campaign
		candidates := []models.Candidate{
			{
				CampaignID: campaign.ID,
				Number:     1,
				Name:       "Calon Damai Satu",
				Vision:     "Mewujudkan kedamaian abadi.",
				Mission:    "Membangun komunikasi yang baik.",
				Photo:      "https://via.placeholder.com/300?text=Kandidat+1",
			},
			{
				CampaignID: campaign.ID,
				Number:     2,
				Name:       "Calon Sentosa Dua",
				Vision:     "Kesejahteraan untuk semua.",
				Mission:    "Meningkatkan ekonomi rakyat.",
				Photo:      "https://via.placeholder.com/300?text=Kandidat+2",
			},
		}
		DB.Create(&candidates)
		log.Println("Candidates seeded for default campaign!")
	}
}
