package models

import "github.com/gofrs/uuid"

type Order struct {
	ID        string `gorm:"primary_key;"`
	TableName string `gorm:"-"`
	Caption   string
	Statuses  []ObjectStatus `gorm:"many2many:order_statuses;"`
}

func (item *Order) New(name string) {
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
	item.TableName = "orders"
	item.Caption = name
	item.Statuses = statuses
}

func (item *Order) SetTable(table string) {
	item.TableName = table
}

func (item *Order) GetID() string {
	return item.ID
}

func (item *Order) GetCaption() string {
	return item.Caption
}

func (item *Order) GetStatuses() []ObjectStatus {
	return item.Statuses
}

func (item *Order) GetTable() string {
	return item.TableName
}

func (item *Order) GetStatus(caption string) ObjectStatus {
	for _, status := range item.Statuses {
		if status.StatusGroupCode == caption {
			return status
		}
	}
	return ObjectStatus{}
}

func (item *Order) SetStatus(caption string, stat string) {
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
