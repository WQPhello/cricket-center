package model

import "time"

type Position struct {
	Coordinate
	UpdatedAt time.Time `json:"updated_at"` // The time when the coordinate was last updated.
}

type Coordinate struct {
	X float64 `json:"x" xorm:"x"` // X-coordinate
	Y float64 `json:"y" xorm:"y"` // Y-coordinate
	Z float64 `json:"z" xorm:"z"` // Z-coordinate
}
