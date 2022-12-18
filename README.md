# Rydon: Ride Sharing Platform
Rydon is a Ride Sharing Platform using a microservice architecture. It is a console application that can be run from the command line.

Technologies used are Go for the microservices and MySQL for the persistent storage of information using database.

## Rydon architecture
### Functionalities 
Rydon provides the following functionalities
* Passengers and drivers can set up their profile
* Passengers and drivers can update their personal details
* Passengers can use the Rydon console app to submit their trip request
* Rydon will match passengers who submitted a trip request with a driver and assign them to each valid trip

### Microservices 
To effectively provide functionalities above, the following microservices below enables Rydon Ride-sharing System workflow:
* Trip Service - Store trip information for retrieving purpose
* Passenger Profile Service - Store passengers profile information
* Driver Profile Service - Store driver profile information

## Design consideration of microservices
1. REpresentational State Transfer (REST) is used for the Application Programming Interface (API) 
	* Best practice for handling well-organized entities, data and resource, such as profile, trip 
	* Supports most CRUD-based operations such as GET, POST, PUT, DELETE, OPTIONS
	
## Setup
Step 1: In the directory you want to store the repo, clone Rydon using
```
git clone https://github.com/j6ine/rydon.git
```
Step 2: Change directory to where the backend folder is located. Run the Backend 
```golang
// add go mod file if running for the first time
go mod init backend

// install packages if running for the first time 
go get –u "github.com/go-sql-driver/mysql"
go get -u github.com/gorilla/mux
```
```golang
// go run [microservice name].go
go run passenger.go
```
Step 3: Change directory to where the backend folder is located. Run the Frontend 
```golang
// add go mod file if running for the first time
go mod init frontend
```
```golang
go run menu.go
```
