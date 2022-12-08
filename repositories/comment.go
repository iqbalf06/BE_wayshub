package repositories

import (
	"wayshub/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindComments() ([]models.Comment, error)
	GetComment(ID int) (models.Comment, error)
	AddComment(comment models.Comment) (models.Comment, error)
	EditComment(comment models.Comment) (models.Comment, error)
	DeleteComment(comment models.Comment) (models.Comment, error)
}

func RepositoryComment(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindComments() ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Find(&comments).Error

	return comments, err
}

func (r *repository) GetComment(ID int) (models.Comment, error) {
	var comment models.Comment
	err := r.db.First(&comment, ID).Error

	return comment, err
}

func (r *repository) AddComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Preload("Channel").Create(&comment).Error

	return comment, err
}

func (r *repository) EditComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Save(&comment).Error

	return comment, err
}

func (r *repository) DeleteComment(comment models.Comment) (models.Comment, error) {
	err := r.db.Delete(&comment).Error

	return comment, err
}
