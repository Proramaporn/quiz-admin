package models

import (
	"time"
)

type Exam struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Question  string    `gorm:"not null" json:"question"`
	Choices   []Choice  `gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE;" json:"choices"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Choice struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ExamID     uint   `gorm:"not null" json:"exam_id"`
	ChoiceText string `gorm:"not null" json:"choice_text"`
}
