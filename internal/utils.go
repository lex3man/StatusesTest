package internal

import (
	"fountcore.ru/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fmt"
)

type DataBase struct {
	name   string
	login  string
	passwd string
	host   string
	port   int16
	Base   *gorm.DB
}

func (db *DataBase) Init(drv string) {
	db.name = "test"
	db.login = "test"
	db.passwd = "test"
	db.host = "localhost"
	db.port = 3306
	switch drv {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.login, db.passwd, db.host, db.port, db.name)
		dbc, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.Base = dbc
	case "sqlite":
		dsn := fmt.Sprintf("%s.db", db.name)
		dbc, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.Base = dbc
	}
	db.Base.AutoMigrate(models.User{}, models.Item{}, models.Vehicle{}, models.Order{}, models.Trip{}, models.ObjectStatus{})

}

func (db *DataBase) Save(item models.Statusable) {
	switch item.GetTable() {
	case "items":
		entity := models.Item{
			ID:       item.GetID(),
			Caption:  item.GetCaption(),
			Statuses: item.GetStatuses(),
		}
		db.Base.Create(&entity)
	case "users":
		entity := models.User{
			ID:       item.GetID(),
			Caption:  item.GetCaption(),
			Statuses: item.GetStatuses(),
		}
		db.Base.Create(&entity)
	case "vehicles":
		entity := models.Vehicle{
			ID:       item.GetID(),
			Caption:  item.GetCaption(),
			Statuses: item.GetStatuses(),
		}
		db.Base.Create(&entity)
	case "orders":
		entity := models.Order{
			ID:       item.GetID(),
			Caption:  item.GetCaption(),
			Statuses: item.GetStatuses(),
		}
		db.Base.Create(&entity)
	case "trips":
		entity := models.Trip{
			ID:       item.GetID(),
			Caption:  item.GetCaption(),
			Statuses: item.GetStatuses(),
		}
		db.Base.Create(&entity)
	}

}

func (db *DataBase) Find(item models.Statusable, id string) models.Statusable {
	switch item.GetTable() {
	case "items":
		var entity models.Item
		db.Base.First(&entity, "id = ?", id)
		db.Base.Model(&entity).Association("Statuses").Find(&entity.Statuses)
		return &entity
	case "users":
		var entity models.User
		db.Base.First(&entity, "id = ?", id)
		db.Base.Model(&entity).Association("Statuses").Find(&entity.Statuses)
		return &entity
	case "vehicles":
		var entity models.Vehicle
		db.Base.First(&entity, "id = ?", id)
		db.Base.Model(&entity).Association("Statuses").Find(&entity.Statuses)
		return &entity
	case "orders":
		var entity models.Order
		db.Base.First(&entity, "id = ?", id)
		db.Base.Model(&entity).Association("Statuses").Find(&entity.Statuses)
		return &entity
	case "trips":
		var entity models.Trip
		db.Base.First(&entity, "id = ?", id)
		db.Base.Model(&entity).Association("Statuses").Find(&entity.Statuses)
		return &entity
	}
	return nil
}
