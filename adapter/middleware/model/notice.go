package middlewaremodel

import "time"

type Notice struct {
	ID          string    `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Text        string    `gorm:"column:text"`
	PublishFrom int       `gorm:"column:publish_from"`
	PublishTo   int       `gorm:"column:publish_to"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (n *Notice) TableName() string {
	return "notice"
}
