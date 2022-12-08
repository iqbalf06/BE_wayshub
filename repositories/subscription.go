package repositories

import (
	"wayshub/models"

	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	AddSubscription(subcription models.Subscription) (models.Subscription, error)
	Unsubscribe(subcription models.Subscription) (models.Subscription, error)
}

func RepositorySubscription(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddSubscription(subscription models.Subscription) (models.Subscription, error) {
	err := r.db.Preload("Channel").Create(&subscription).Error

	return subscription, err
}

func (r *repository) Unsubscribe(subscription models.Subscription) (models.Subscription, error) {
	err := r.db.Delete(&subscription).Error

	return subscription, err
}
