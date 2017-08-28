# Go Address API

An address API in go.
The API implements listing addresses, showing a single address, adding an address and modifying an address.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

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

## Testing the API
You can test the API by sending HTTP requests to the running server using HTTP Testing client like [Postman](https://www.getpostman.com/) or using a command line tool like [curl](https://curl.haxx.se/).
Start by running the server by typing the following command in your terminal. (see above section in Installing and starting the server.)
```
goaddressapi
```
* To add an address
```
curl -v -X POST localhost:8001/addresses/31 --data '{"Firstname":"Bobby","lastname":"Axe","emailaddress":"axe@test.com","phonenumber":"214-545-5553"}'
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

* Hat tip to [Gorilla Mux](https://github.com/gorilla/mux) team.
