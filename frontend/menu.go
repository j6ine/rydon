package main

import (
	"fmt"
)

type Passenger struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	MobileNum string `json: "Mobile Number`
	Email     string `json: "Email Address"`
}

// type Passengers struct {
// 	Passengers map[string]Passenger `json:"Passengers"`
// }

func main() {
outer:
	for {
		fmt.Println("\n\n==============================")
		fmt.Println("Rydon: Ride Sharing Platform\n",
			"\nFor Passengers:\n",
			"1. Create Passenger Account\n",
			"2. Update Passenger Account\n",
			"3. Delete Passenger Account\n",
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
			// createPassenger()
		case 2:
			print("hi")
			// updatePassenger()
		case 3:
			// deletePassenger()
		case 4:
			// requestTrip()
		case 5:
			// listAllTrips()
		case 10:
			// listAllPassengers()
		case 12:
			break outer
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}
