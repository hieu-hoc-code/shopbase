package models

import "time"

type Order struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Total      int64     `json:"total"`
	PaymentId  uint      `json:"payment_id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
