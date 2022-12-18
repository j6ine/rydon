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

	// check
	var drivers map[string]Driver = map[string]Driver{}
	drivers = getDrivers()
	fmt.Println(drivers)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers/{driverid}", driver).Methods("GET", "POST", "PATCH", "PUT")
	router.HandleFunc("/api/v1/drivers", alldrivers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func driver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["driverid"]); !ok {
					fmt.Println(data)
					insertDriver(params["driverid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "Driver ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data Driver

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["driverid"]); ok {
					fmt.Println(data)
					updateDriver(params["driverid"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				if orig, ok := isExist(params["driverid"]); ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "FirstName":
							orig.FirstName = v.(string)
						case "LastName":
							orig.LastName = v.(string)
						case "MobileNum":
							orig.MobileNum = v.(string)
						case "Email":
							orig.Email = v.(string)
						case "ID":
							orig.ID = v.(string)
						case "LicenseNum":
							orig.LicenseNum = v.(string)
						}
					}
					updateDriver(params["driverid"], orig)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "Driver ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Invalid Driver ID")
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

func isExist(id string) (Driver, bool) {
	var d Driver

	result := db.QueryRow("select * from driver where driverid=?", id)
	err := result.Scan(&id, &d.FirstName, &d.LastName, &d.MobileNum, &d.Email, &d.ID, &d.LicenseNum)
	if err == sql.ErrNoRows {
		return d, false
	}

	return d, true
}

func insertDriver(id string, d Driver) {
	_, err := db.Exec("insert into driver values(?,?,?,?,?,?,?)", id, d.FirstName, d.LastName, d.MobileNum, d.Email, d.ID, d.LicenseNum)
	if err != nil {
		panic(err.Error())
	}
}

func updateDriver(id string, d Driver) {
	_, err := db.Exec("update driver set firstname=?, lastname=?, mobilenum=?, email=?, id=?, licensenum=? where driverid=?", d.FirstName, d.LastName, d.MobileNum, d.Email, d.ID, d.LicenseNum, id)
	if err != nil {
		panic(err.Error())
	}
}
