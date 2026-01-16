package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NIK       string    `json:"nik" gorm:"unique"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Role      string    `json:"role" gorm:"default:'USER'"` // USER | ADMIN
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Campaign struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	StartDate   time.Time   `json:"startDate"`
	EndDate     time.Time   `json:"endDate"`
	IsActive    bool        `json:"isActive" gorm:"default:true"`
	Candidates  []Candidate `json:"candidates" gorm:"foreignKey:CampaignID"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Candidate struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CampaignID uint      `json:"campaignId"`
	Number     int       `json:"number"`
	Name       string    `json:"name"`
	Vision     string    `json:"vision"`
	Mission    string    `json:"mission"`
	Photo      string    `json:"photo"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Vote struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"userId"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	CampaignID  uint      `json:"campaignId"`
	CandidateID uint      `json:"candidateId"`
	Candidate   Candidate `json:"candidate" gorm:"foreignKey:CandidateID"`
	Timestamp   time.Time `json:"timestamp" gorm:"autoCreateTime"`
}
