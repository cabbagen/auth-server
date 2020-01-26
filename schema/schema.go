package schema

import "time"

type BaseSchema struct {
	ID               int            `gorm:"column:id; type:int; primary key; not null" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at; type:datetime; default now()" json:"createdAt"`
	UpdatedAt        time.Time      `gorm:"column:updated_at; type:datetime; default now()" json:"updatedAt"`
}
