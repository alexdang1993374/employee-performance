package models

import (
	"time"
)

type Employee struct {
    Id              int64       `json:"id,omitempty"`
    Name            string      `json:"name,omitempty" validate:"required"`
    Performance     int         `json:"performance,omitempty" validate:"required"`
    Date            time.Time   `json:"Date"`
}