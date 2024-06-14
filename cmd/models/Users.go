package models

import "github.com/gofrs/uuid"

type User struct {
	ID        string `gorm:"primary_key;"`
	TableName string `gorm:"-"`
	Caption   string
	Statuses  []ObjectStatus `gorm:"many2many:user_statuses;"`
}

func (item *User) New(name string) {
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
	item.TableName = "users"
	item.Caption = name
	item.Statuses = statuses
}

func (item *User) SetTable(table string) {
	item.TableName = table
}

func (item *User) GetID() string {
	return item.ID
}

func (item *User) GetCaption() string {
	return item.Caption
}

func (item *User) GetStatuses() []ObjectStatus {
	return item.Statuses
}

func (item *User) GetTable() string {
	return item.TableName
}

func (item *User) GetStatus(caption string) ObjectStatus {
	for _, status := range item.Statuses {
		if status.StatusGroupCode == caption {
			return status
		}
	}
	return ObjectStatus{}
}

func (item *User) SetStatus(caption string, stat string) {
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
