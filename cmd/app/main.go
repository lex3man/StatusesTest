package main

import (
	"fmt"
	"log"

	"fountcore.ru/cmd/models"
	"fountcore.ru/internal"
)

func main() {
	db := &internal.DataBase{}
	db.Init("sqlite")
	
	var user models.Statusable = &models.User{}
	user.New("Max")
	user.SetStatus("U_PLANNING_SG", "NOT PLANNED")
	user.SetStatus("U_ACTIVATION_SG", "ACTIVE")
	db.Save(user)
	
	var vehicle models.Statusable = &models.Vehicle{}
	vehicle.New("Car")
	vehicle.SetStatus("U_PLANNING_SG", "PLANNED")
	db.Save(vehicle)
	
	var order models.Statusable = &models.Order{}
	order.New("New order")
	order.SetStatus("U_PLANNING_SG", "IN PROGRESS")
	db.Save(order)
	
	var trip models.Statusable = &models.Trip{}
	trip.New("To west")
	trip.SetStatus("U_PLANNING_SG", "DONE")
	db.Save(trip)

	var getUser models.Statusable = &models.User{}
	getUser.SetTable("users")
	getUser = db.Find(getUser, user.GetID())
	log.Println(getUser.GetStatus("U_PLANNING_SG").StatusCode)
	log.Println(getUser.GetStatus("U_ACTIVATION_SG").StatusCode)

	var getVehicle models.Statusable = &models.Vehicle{}
	getVehicle.SetTable("vehicles")
	getVehicle = db.Find(getVehicle, vehicle.GetID())
	log.Println(getVehicle.GetStatus("U_PLANNING_SG").StatusCode)

	var getOrder models.Statusable = &models.Vehicle{}
	getOrder.SetTable("orders")
	getOrder = db.Find(getOrder, order.GetID())
	log.Println(getOrder.GetStatus("U_PLANNING_SG").StatusCode)

	var getTrip models.Statusable = &models.Vehicle{}
	getTrip.SetTable("trips")
	getTrip = db.Find(getTrip, trip.GetID())
	log.Println(getTrip.GetStatus("U_PLANNING_SG").StatusCode)


	fmt.Println("All done!")
}
