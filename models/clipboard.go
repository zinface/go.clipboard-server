package models

import "time"

type ClipBoard struct {
	BaseData string    `json:"data"`
	MimeType string    `json:"mime"`
	CreateAt time.Time `json:"createAt"`
}
