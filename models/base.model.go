package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// declaration timestamp
type Timestamp time.Time

// declaration model basemodel
type BaseModel struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	IsActive    bool       `gorm:"default:true"`
	CreatedBy   string     `json:"createdBy"`
	UpdatedBy   *string    `json:"updatedBy"`
	CreatedDate *Timestamp `json:"createdDate" gorm:"type:timestamp without time zone;"`
	UpdatedDate *Timestamp `json:"updatedDate" gorm:"type:timestamp without time zone;"`
}

// as interface from basemodel
type Tabler interface {
	TableName() string
}

// function for create ID
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	return nil
}

// MarshalJSON ...
func (t Timestamp) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
