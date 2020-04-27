# Go Address API

An address API in go.
The API implements listing addresses, showing a single address, adding an address and modifying an address.
The API also enables exporting addresses to a CSV file and adding addresses by importing a CSV file.

---

## Getting Started

These instructions will get you a copy of the project and set up your your local machine for development and testing of the project.

### Prerequisites
* Download and [install Go](https://golang.org/doc/install).
* Make sure your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) is set.
* Get [Gorilla Mux](https://github.com/gorilla/mux) (the command below assumes a [correctly configured](https://golang.org/doc/install#testing) Go toolchain)
```
go get -u github.com/gorilla/mux
```

### Installing (assumes terminal)

1. Clone this repository
```
git clone https://github.com/iggym/goaddressapi.git
```
2. change directory to **goaddressapi**
```
cd goaddressapi
```
3. run go install
```
go install
```
4. Run the server by typing the following command in your terminal. (the command below assumes a [correctly configured](https://golang.org/doc/install#testing) Go toolchain)
```
goaddressapi
```
5. Open another terminal window and test the server with the following [curl](https://curl.haxx.se/) command.
```
curl -v  GET localhost:8001/addresses
```
You should see the following:
```
[{"id":"1","firstname":"Bryan","lastname":"Connerty","emailaddress":"bryanc@test-email.com","phonenumber":"214-555-5551"},{"id":"2","firstname":"Kate","lastname":"Sacker","emailaddress":"kates@test-email.com","phonenumber":"214-555-5552"},{"id":"3","firstname":"Lonnie","lastname":"Watley","emailaddress":"lonniew@test-email.com","phonenumber":"214-555-5553"},{"id":"4","firstname":"Ira","lastname":"Schirmer","emailaddress":"iras@test-email","phonenumber":"214-555-5554"},{"id":"5","firstname":"Adam","lastname":"DeGiulio","emailaddress":"adamgs@test-email","phonenumber":"214-555-5554"}]
```

---

## Testing the API
You can test the API by sending HTTP requests to the running server using a HTTP Testing client like [Postman](https://www.getpostman.com/) or using a command line tool like [curl](https://curl.haxx.se/).

Start by running the server by typing the following command in your terminal. (see above section for Installing and starting the server.)
```
goaddressapi
```
* To get a single address (use the id for the parameter)
```
 curl -v GET localhost:8001/addresses/2
```
* To add an address
```
curl -v -X POST localhost:8001/addresses/31 --data '{"Firstname":"Bobby","lastname":"Axe","emailaddress":"axe@test.com","phonenumber":"214-545-5553"}'
```
* To Delete an address(use the id for the parameter)
```
curl -v -X DELETE localhost:8001/addresses/3
```
* To list all addresses
```
curl -v  GET localhost:8001/addresses
```
* To Modify an address
```
curl -v -X PUT localhost:8001/addresses/4 --data '{"ID":"4","Firstname":"Santa","lastname":"Claus","emailaddress":"santa@test.com","phonenumber":"214-545-5553"}'
```
* To Export the list of addresses to a CSV file
```
curl  localhost:8001/export --output list.csv
```
* To import addresses using a CSV file use the command below (assumes your file is a CSV file named tmp.csv)
```
curl -i -X POST -H "Content-Type: multipart/form-data" -F "data=@tmp.csv" localhost:8001/import
```
## Running the unit tests
* To run the unit tests enter the following command
```
go test -v
```
* To view code coverage
```
go test -cover
```
* To get a code coverage report use the below  
```
go tool cover -html=c.out -o coverage.html
```

## Deployment

This project should not be used in a production system as is. It is a  demonstration project.

## Built With

* [Gorilla Mux](https://github.com/gorilla/mux)
* [Go](https://golang.org/doc/install).

## Contributing

We would love your contributions. Please submit pull requests with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to the [Gorilla Mux](https://github.com/gorilla/mux) team.
