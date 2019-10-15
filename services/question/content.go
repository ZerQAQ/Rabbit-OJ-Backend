package question

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
)

func Content (tid string) (*models.QuestionContent, error) {
	content := &models.QuestionContent{}

	if err := db.DB.Where("tid = ?", tid).First(&content).Error; err != nil {
		return nil, err
	} else {
		return content, nil
	}
}