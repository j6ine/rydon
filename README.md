# Rydon: Ride Sharing Platform
Rydon is a Ride Sharing Platform using a microservice architecture. It is a console application that can be run from the command line.

Technologies used are Go for the microservices and MySQL for the persistent storage of information using database

## Architecture diagram


## Design consideration of microservices
	
	
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
