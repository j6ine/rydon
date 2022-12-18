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

type Passenger struct {
	PassengerID string `json:"PassengerID"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	MobileNum   string `json: "Mobile Number`
	Email       string `json: "Email Address"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "ju:a$$123@tcp(127.0.0.1:3306)/rydon_passenger_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// print passengers in terminal
	var passengers map[string]Passenger = map[string]Passenger{}
	passengers = getPassengers()
	fmt.Println(passengers)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passengers/{email}", updatepassenger).Methods("PUT")
	router.HandleFunc("/api/v1/passengers", createpassenger).Methods("POST")
	router.HandleFunc("/api/v1/passengers", allpassengers)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func createpassenger(w http.ResponseWriter, r *http.Request) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		var data Passenger

		if err := json.Unmarshal(body, &data); err == nil {
			_, emailExists := isEmailExist(data.Email)

			if emailExists == false {
				fmt.Println(data)
				insertPassenger(data)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Passenger email exists")
			}
		} else {
			fmt.Println(err)
		}
	}
}

func updatepassenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		var data Passenger

		if err := json.Unmarshal(body, &data); err == nil {
			p, emailExists := isEmailExist(params["email"])

			if emailExists == true {
				fmt.Println(p)
				fmt.Println(p.PassengerID)
				fmt.Println(data)
				updatePassenger(p.PassengerID, data)
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Passenger email exists")
			}
		} else {
			fmt.Println(err)
		}
	}
}

func allpassengers(w http.ResponseWriter, r *http.Request) {
	passengerWrapper := struct {
		Passengers map[string]Passenger `json:"Passengers"`
	}{getPassengers()}
	json.NewEncoder(w).Encode(passengerWrapper)
	return
}

func getPassengers() map[string]Passenger {
	results, err := db.Query("select * from Passenger")
	if err != nil {
		panic(err.Error())
	}

	var passengers map[string]Passenger = map[string]Passenger{}

	for results.Next() {
		var p Passenger
		var id string = p.PassengerID
		err = results.Scan(&id, &p.FirstName, &p.LastName, &p.MobileNum, &p.Email)
		if err != nil {
			panic(err.Error())
		}

		passengers[id] = p
	}

	return passengers
}

func isIdExist(id string) (Passenger, bool) {
	var p Passenger

	result := db.QueryRow("select * from passenger where passengerid=?", id)
	err := result.Scan(&id, &p.FirstName, &p.LastName, &p.MobileNum, &p.Email)
	if err == sql.ErrNoRows {
		return p, false
	}

	return p, true
}

func isEmailExist(email string) (Passenger, bool) {
	var p Passenger

	result := db.QueryRow("select * from passenger where email=?", email)
	err := result.Scan(&p.PassengerID, &p.FirstName, &p.LastName, &p.MobileNum, &p.Email)
	if err == sql.ErrNoRows {
		return p, false
	}

	return p, true
}

func insertPassenger(p Passenger) {
	_, err := db.Exec("insert into passenger (firstname, lastname, mobilenum, email) values(?,?,?,?)", p.FirstName, p.LastName, p.MobileNum, p.Email)
	if err != nil {
		panic(err.Error())
	}
}

func updatePassenger(id string, p Passenger) {
	_, err := db.Exec("update passenger set firstname=?, lastname=?, mobilenum=?, email=? where passengerid=?", p.FirstName, p.LastName, p.MobileNum, p.Email, id)
	if err != nil {
		panic(err.Error())
	}
}
