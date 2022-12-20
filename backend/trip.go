package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Trip struct {
	TripID          string `json:"Trip ID"`
	TripDate        string `json: "Trip Date`
	TripStatus      string `json:"Trip Status"`
	PassengerID     string `json:"Trip ID"`
	PickupLocation  string `json:"Pickup Location"`
	DropoffLocation string `json:"Dropoff Location"`
	DriverID        string `json:"DriverID"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "ju:a$$123@tcp(127.0.0.1:3306)/rydon_trip_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// print trips in terminal
	var trips map[string]Trip = map[string]Trip{}
	trips = getTrips()
	fmt.Println(trips)

	// print passenger trips
	var ptrips map[string]Trip = map[string]Trip{}
	ptrips = getpassengertrips("2")
	fmt.Println(ptrips)

	router := mux.NewRouter()
	// router.HandleFunc("/api/v1/trips/{tripid}", updatetripstatus).Methods("PATCH")
	router.HandleFunc("/api/v1/trips", createtrip).Methods("POST")
	router.HandleFunc("/api/v1/trips", alltrips)
	router.HandleFunc("/api/v1/trips/{passengerid}", allpassengertrips)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func createtrip(w http.ResponseWriter, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		var data Trip
		fmt.Println(body)
		if err := json.Unmarshal(body, &data); err == nil {
			t, pidExists := isPassengerIdExist(data.PassengerID)

			if pidExists == true {
				fmt.Println(t)
				fmt.Println(data)
				insertTrip(data)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Passenger ID does not exist")
			}
		} else {
			fmt.Println(err)
		}

	}
}

func alltrips(w http.ResponseWriter, r *http.Request) {
	tripWrapper := struct {
		Trips map[string]Trip `json:"Trips"`
	}{getTrips()}
	json.NewEncoder(w).Encode(tripWrapper)
	return
}

func getTrips() map[string]Trip {
	results, err := db.Query("select * from Trip")
	if err != nil {
		panic(err.Error())
	}

	var trips map[string]Trip = map[string]Trip{}

	for results.Next() {
		var t Trip
		var id string = t.TripID
		err = results.Scan(&id, &t.TripDate, &t.TripStatus, &t.PassengerID, &t.PickupLocation, &t.DropoffLocation, &t.DriverID)
		if err != nil {
			panic(err.Error())
		}

		trips[id] = t
	}

	return trips
}

func allpassengertrips(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripWrapper := struct {
		Trips map[string]Trip `json:"Trips"`
	}{getpassengertrips(params["passengerid"])}
	json.NewEncoder(w).Encode(tripWrapper)
	return
}

func getpassengertrips(passengerid string) map[string]Trip {
	results, err := db.Query("select * from Trip where passengerid=?", passengerid)
	if err != nil {
		panic(err.Error())
	}

	var trips map[string]Trip = map[string]Trip{}

	for results.Next() {
		var t Trip
		var id string = t.TripID
		err = results.Scan(&id, &t.TripDate, &t.TripStatus, &t.PassengerID, &t.PickupLocation, &t.DropoffLocation, &t.DriverID)
		if err != nil {
			panic(err.Error())
		}

		trips[id] = t
	}

	return trips
}

func isPassengerIdExist(id string) (Trip, bool) {
	var t Trip

	result := db.QueryRow("select * from trip where passengerid=?", id, "order by tripdate asc")
	err := result.Scan(&id, &t.TripDate, &t.TripStatus, &t.PassengerID, &t.PickupLocation, &t.DropoffLocation, &t.DriverID)
	if err == sql.ErrNoRows {
		return t, false
	}

	return t, true
}

func insertTrip(t Trip) {
	tripdate := time.Now()
	tripstatus := "Finding a Driver"
	driverid := "-"

	_, err := db.Exec("insert into trip (tripdate, tripstatus, passengerid, pickuplocation, dropofflocation, driverid) values(?,?,?,?,?,?)", tripdate, tripstatus, t.PassengerID, t.PickupLocation, t.DropoffLocation, driverid)
	if err != nil {
		panic(err.Error())
	}
}
