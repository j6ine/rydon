package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Driver struct {
	DriverID   string `json:"DriverID"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	MobileNum  string `json: "Mobile Number`
	Email      string `json: "Email Address"`
	ID         string `json: "ID"`
	LicenseNum string `json: "License Number"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "ju:a$$123@tcp(127.0.0.1:3306)/rydon_driver_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// print drivers in terminal
	var drivers map[string]Driver = map[string]Driver{}
	drivers = getDrivers()
	fmt.Println(drivers)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers/{email}", updatedriver).Methods("PUT")
	router.HandleFunc("/api/v1/drivers", createdriver).Methods("POST")
	router.HandleFunc("/api/v1/drivers", alldrivers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func createdriver(w http.ResponseWriter, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		var data Driver

		if err := json.Unmarshal(body, &data); err == nil {
			_, emailExists := isEmailExist(data.Email)

			if emailExists == false {
				fmt.Println(data)
				insertDriver(data)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Driver email exists")
			}
		} else {
			fmt.Println(err)
		}
	}
}

func updatedriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		var data Driver

		if err := json.Unmarshal(body, &data); err == nil {
			d, emailExists := isEmailExist(params["email"])

			if emailExists == true {
				fmt.Println(d)
				fmt.Println(d.DriverID)
				fmt.Println(data)
				updateDriver(d.DriverID, data)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Driver email does not exist")
			}
		} else {
			fmt.Println(err)
		}
	}
}

func alldrivers(w http.ResponseWriter, r *http.Request) {
	driverWrapper := struct {
		Drivers map[string]Driver `json:"Drivers"`
	}{getDrivers()}
	json.NewEncoder(w).Encode(driverWrapper)
	return
}

func getDrivers() map[string]Driver {
	results, err := db.Query("select * from Driver")
	if err != nil {
		panic(err.Error())
	}

	var drivers map[string]Driver = map[string]Driver{}

	for results.Next() {
		var d Driver
		var id string = d.DriverID
		err = results.Scan(&id, &d.FirstName, &d.LastName, &d.MobileNum, &d.Email, &d.ID, &d.LicenseNum)
		if err != nil {
			panic(err.Error())
		}

		drivers[id] = d
	}

	return drivers
}

func isIdExist(id string) (Driver, bool) {
	var d Driver

	result := db.QueryRow("select * from driver where driverid=?", id)
	err := result.Scan(&id, &d.FirstName, &d.LastName, &d.MobileNum, &d.Email, &d.ID, &d.LicenseNum)
	if err == sql.ErrNoRows {
		return d, false
	}

	return d, true
}

func isEmailExist(email string) (Driver, bool) {
	var d Driver

	result := db.QueryRow("select * from driver where email=?", email)
	err := result.Scan(&d.DriverID, &d.FirstName, &d.LastName, &d.MobileNum, &d.Email, &d.ID, &d.LicenseNum)
	if err == sql.ErrNoRows {
		return d, false
	}

	return d, true
}

func insertDriver(d Driver) {
	_, err := db.Exec("insert into driver (firstname, lastname, mobilenum, email, id, licensenum) values(?,?,?,?,?,?)", d.FirstName, d.LastName, d.MobileNum, d.Email, d.ID, d.LicenseNum)
	if err != nil {
		panic(err.Error())
	}
}

func updateDriver(id string, d Driver) {
	_, err := db.Exec("update driver set firstname=?, lastname=?, mobilenum=?, email=?, licensenum=? where driverid=?", d.FirstName, d.LastName, d.MobileNum, d.Email, d.LicenseNum, id)
	if err != nil {
		panic(err.Error())
	}
}
