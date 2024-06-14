package models

import "github.com/gofrs/uuid"

type Trip struct {
	ID        string `gorm:"primary_key;"`
	TableName string `gorm:"-"`
	Caption   string
	Statuses  []ObjectStatus `gorm:"many2many:trip_statuses;"`
}

func (item *Trip) New(name string) {
	statuses := []ObjectStatus{
		{
			StatusGroupCode: "U_ACTIVATION_SG",
			StatusCode:      "NOT_ACTIVE",
		},
		{
			StatusGroupCode: "U_LIFECYCLE_SG",
			StatusCode:      "PREPARED",
		},
	}
	item.ID = uuid.Must(uuid.NewV4()).String()
	item.TableName = "trips"
	item.Caption = name
	item.Statuses = statuses
}

func (item *Trip) SetTable(table string) {
	item.TableName = table
}

func (item *Trip) GetID() string {
	return item.ID
}

func (item *Trip) GetCaption() string {
	return item.Caption
}

func (item *Trip) GetStatuses() []ObjectStatus {
	return item.Statuses
}

func (item *Trip) GetTable() string {
	return item.TableName
}

func (item *Trip) GetStatus(caption string) ObjectStatus {
	for _, status := range item.Statuses {
		if status.StatusGroupCode == caption {
			return status
		}
	}
	return ObjectStatus{}
}

func (item *Trip) SetStatus(caption string, stat string) {
	for i, status := range item.Statuses {
		if status.StatusGroupCode == caption {
			item.Statuses[i].StatusCode = stat
			return
		}
	}
	item.Statuses = append(item.Statuses, ObjectStatus{
		StatusGroupCode: caption,
		StatusCode:      stat,
	})
}
