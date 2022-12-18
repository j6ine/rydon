package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Passenger struct {
	PassengerID string `json:"PassengerID"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	MobileNum   string `json: "Mobile Number`
	Email       string `json: "Email Address"`
}
type Passengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

type Driver struct {
	DriverID   string `json:"DriverID"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	MobileNum  string `json: "Mobile Number`
	Email      string `json: "Email Address"`
	ID         string `json: "ID"`
	LicenseNum string `json: "License Number"`
}
type Drivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

func main() {

outer:
	for {
		fmt.Println("\n\n==============================")
		fmt.Println("Rydon: Ride Sharing Platform\n",
			"\nFor Passengers:\n",
			"1. Create Passenger Account\n",
			"2. Update Passenger Account\n",
			"4. Request a Trip\n",
			"5. Retrieve All Trips\n",
			"\nFor Drivers:\n",
			"6. Create Driver Account\n",
			"7. Update Driver Account\n",
			"8. Initiate a Trip\n",
			"9. End a Trip\n",
			"\nFor Admin (JuJu):\n",
			"10. Retrieve All Passengers\n",
			"11. Retrieve All Drivers\n",
			"\n 12. Quit")

		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Println("\n~ Creating a Passenger Account ~")
			createPassenger()
		case 2:
			fmt.Println("\n~ Updating a Passenger Account ~")
			updatePassenger()
		case 4:
			// requestTrip()
		case 5:
			// listAllTrips()
		case 6:
			fmt.Println("\n~ Creating a Driver Account ~")
			createDriver()
		case 7:
			fmt.Println("\n~ Updating a Driver Account ~")
			updateDriver()
		case 10:
			fmt.Println("\n~ Listing All Passengers ~")
			listAllPassengers()
		case 11:
			fmt.Println("\n~ Listing All Drivers ~")
			listAllDrivers()
		case 12:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}

func listAllPassengers() {
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/passengers", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Passengers
				json.Unmarshal(body, &res)
				println(body)

				for k, v := range res.Passengers {
					fmt.Println(v.PassengerID, "(", k, ")")
					fmt.Println("First Name:", v.FirstName)
					fmt.Println("Last Name:", v.LastName)
					fmt.Println("Mobile Number:", v.MobileNum)
					fmt.Println("Email:", v.Email)
					fmt.Println()
				}
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func createPassenger() {
	var passenger Passenger

	fmt.Scanf("%s", &(passenger.PassengerID))
	fmt.Println(passenger.PassengerID)
	fmt.Print("Enter the ID of the Passenger to be created: ")
	fmt.Scanf("%s", &(passenger.PassengerID))
	fmt.Println("You typed: " + passenger.PassengerID)

	fmt.Scanf("%s", &(passenger.FirstName))
	fmt.Println(passenger.FirstName)
	fmt.Print("Enter your First Name: ")
	fmt.Scanf("%s", &(passenger.FirstName))
	fmt.Println("You typed: " + passenger.FirstName)

	fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println(passenger.LastName)
	fmt.Print("Enter your Last Name: ")
	fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println("You typed: " + passenger.LastName)

	fmt.Scanf("%s", &(passenger.MobileNum))
	fmt.Println(passenger.MobileNum)
	fmt.Print("Enter your Mobile Number: ")
	fmt.Scanf("%s", &(passenger.MobileNum))
	fmt.Println("You typed: " + passenger.MobileNum)

	fmt.Scanf("%s", &(passenger.Email))
	fmt.Println(passenger.Email)
	fmt.Print("Enter your Email: ")
	fmt.Scanf("%s", &(passenger.Email))
	fmt.Println("You typed: " + passenger.Email)

	postBody, _ := json.Marshal(passenger)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passengers/"+passenger.PassengerID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("\nSucess! - Passenger", passenger.PassengerID, "created")
			} else if res.StatusCode == 409 {
				fmt.Println("\nOh No, Error! - Passenger", passenger.PassengerID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func updatePassenger() {
	var passenger Passenger

	fmt.Scanf("%s", &(passenger.PassengerID))
	fmt.Println(passenger.PassengerID)
	fmt.Print("Enter the ID of the Passenger to be updated: ")
	fmt.Scanf("%s", &(passenger.PassengerID))
	fmt.Println("You typed: " + passenger.PassengerID)

	fmt.Scanf("%s", &(passenger.FirstName))
	fmt.Println(passenger.FirstName)
	fmt.Print("Enter updated First Name: ")
	fmt.Scanf("%s", &(passenger.FirstName))
	fmt.Println("You typed: " + passenger.FirstName)

	fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println(passenger.LastName)
	fmt.Print("Enter updated Last Name: ")
	fmt.Scanf("%s", &(passenger.LastName))
	fmt.Println("You typed: " + passenger.LastName)

	fmt.Scanf("%s", &(passenger.MobileNum))
	fmt.Println(passenger.MobileNum)
	fmt.Print("Enter updated Mobile Number: ")
	fmt.Scanf("%s", &(passenger.MobileNum))
	fmt.Println("You typed: " + passenger.MobileNum)

	fmt.Scanf("%s", &(passenger.Email))
	fmt.Println(passenger.Email)
	fmt.Print("Enter updated Email: ")
	fmt.Scanf("%s", &(passenger.Email))
	fmt.Println("You typed: " + passenger.Email)

	postBody, _ := json.Marshal(passenger)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/passengers/"+passenger.PassengerID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("\nSuccess! - Passenger", passenger.PassengerID, "updated")
			} else if res.StatusCode == 404 {
				fmt.Println("\nOh No, Error! - Passenger", passenger.PassengerID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func listAllDrivers() {
	client := &http.Client{}

	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/drivers", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				var res Drivers
				json.Unmarshal(body, &res)
				println(body)

				for k, v := range res.Drivers {
					fmt.Println(v.DriverID, "(", k, ")")
					fmt.Println("First Name:", v.FirstName)
					fmt.Println("Last Name:", v.LastName)
					fmt.Println("Mobile Number:", v.MobileNum)
					fmt.Println("Email:", v.Email)
					fmt.Println("Identification Number (NRIC):", v.ID)
					fmt.Println("Email:", v.LicenseNum)
					fmt.Println()
				}
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func createDriver() {
	var driver Driver

	// fmt.Scanf("%s", &(driver.DriverID))
	// fmt.Println(driver.DriverID)
	// fmt.Print("Enter the ID of the Driver to be created: ")
	// fmt.Scanf("%s", &(driver.DriverID))
	// fmt.Println("You typed: " + driver.DriverID)

	fmt.Scanf("%s", &(driver.FirstName))
	fmt.Println(driver.FirstName)
	fmt.Print("Enter your First Name: ")
	fmt.Scanf("%s", &(driver.FirstName))
	fmt.Println("You typed: " + driver.FirstName)

	fmt.Scanf("%s", &(driver.LastName))
	fmt.Println(driver.LastName)
	fmt.Print("Enter your Last Name: ")
	fmt.Scanf("%s", &(driver.LastName))
	fmt.Println("You typed: " + driver.LastName)

	fmt.Scanf("%s", &(driver.MobileNum))
	fmt.Println(driver.MobileNum)
	fmt.Print("Enter your Mobile Number: ")
	fmt.Scanf("%s", &(driver.MobileNum))
	fmt.Println("You typed: " + driver.MobileNum)

	fmt.Scanf("%s", &(driver.Email))
	fmt.Println(driver.Email)
	fmt.Print("Enter your Email: ")
	fmt.Scanf("%s", &(driver.Email))
	fmt.Println("You typed: " + driver.Email)

	fmt.Scanf("%s", &(driver.ID))
	fmt.Println(driver.ID)
	fmt.Print("Enter your Identification Number (NRIC): ")
	fmt.Scanf("%s", &(driver.ID))
	fmt.Println("You typed: " + driver.ID)

	fmt.Scanf("%s", &(driver.LicenseNum))
	fmt.Println(driver.LicenseNum)
	fmt.Print("Enter your License Number: ")
	fmt.Scanf("%s", &(driver.LicenseNum))
	fmt.Println("You typed: " + driver.LicenseNum)

	postBody, _ := json.Marshal(driver)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/drivers/"+driver.DriverID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("\nSucess! - Driver", driver.DriverID, "created")
			} else if res.StatusCode == 409 {
				fmt.Println("\nOh No, Error! - Driver", driver.DriverID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

func updateDriver() {
	var driver Driver

	fmt.Scanf("%s", &(driver.DriverID))
	fmt.Println(driver.DriverID)
	fmt.Print("Enter the ID of the Driver to be updated: ")
	fmt.Scanf("%s", &(driver.DriverID))
	fmt.Println("You typed: " + driver.DriverID)

	fmt.Scanf("%s", &(driver.FirstName))
	fmt.Println(driver.FirstName)
	fmt.Print("Enter updated First Name: ")
	fmt.Scanf("%s", &(driver.FirstName))
	fmt.Println("You typed: " + driver.FirstName)

	fmt.Scanf("%s", &(driver.LastName))
	fmt.Println(driver.LastName)
	fmt.Print("Enter updated Last Name: ")
	fmt.Scanf("%s", &(driver.LastName))
	fmt.Println("You typed: " + driver.LastName)

	fmt.Scanf("%s", &(driver.MobileNum))
	fmt.Println(driver.MobileNum)
	fmt.Print("Enter updated Mobile Number: ")
	fmt.Scanf("%s", &(driver.MobileNum))
	fmt.Println("You typed: " + driver.MobileNum)

	fmt.Scanf("%s", &(driver.Email))
	fmt.Println(driver.Email)
	fmt.Print("Enter updated Email: ")
	fmt.Scanf("%s", &(driver.Email))
	fmt.Println("You typed: " + driver.Email)

	fmt.Scanf("%s", &(driver.LicenseNum))
	fmt.Println(driver.LicenseNum)
	fmt.Print("Enter new License Number: ")
	fmt.Scanf("%s", &(driver.LicenseNum))
	fmt.Println("You typed: " + driver.LicenseNum)

	postBody, _ := json.Marshal(driver)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/drivers/"+driver.DriverID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("\nSuccess! - Driver", driver.DriverID, "updated")
			} else if res.StatusCode == 404 {
				fmt.Println("\nOh No, Error! - Driver", driver.DriverID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}
