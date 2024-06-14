package models

import (
	"gorm.io/gorm"
)

type ObjectStatus struct {
	gorm.Model
	StatusGroupCode string
	StatusCode      string
}

type Statusable interface {
	SetStatus(caption string, stat string)
	GetStatus(caption string) ObjectStatus
	New(name string)
	GetTable() string
	GetID() string
	GetCaption() string
	GetStatuses() []ObjectStatus
	SetTable(table string)
}
