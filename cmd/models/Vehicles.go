package models

import "github.com/gofrs/uuid"

type Vehicle struct {
	ID        string `gorm:"primary_key;"`
	TableName string `gorm:"-"`
	Caption   string
	Statuses  []ObjectStatus `gorm:"many2many:vehicle_statuses;"`
}

func (item *Vehicle) New(name string) {
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
	item.TableName = "vehicles"
	item.Caption = name
	item.Statuses = statuses
}

func (item *Vehicle) SetTable(table string) {
	item.TableName = table
}

func (item *Vehicle) GetID() string {
	return item.ID
}

func (item *Vehicle) GetCaption() string {
	return item.Caption
}

func (item *Vehicle) GetStatuses() []ObjectStatus {
	return item.Statuses
}

func (item *Vehicle) GetTable() string {
	return item.TableName
}

func (item *Vehicle) GetStatus(caption string) ObjectStatus {
	for _, status := range item.Statuses {
		if status.StatusGroupCode == caption {
			return status
		}
	}
	return ObjectStatus{}
}

func (item *Vehicle) SetStatus(caption string, stat string) {
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
