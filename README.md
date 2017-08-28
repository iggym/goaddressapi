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
4. Open another terminal window and test the server with the following [curl](https://curl.haxx.se/) command.
```
curl -v  GET localhost:8001/addresses
```
You should see the following:
```
[{"id":"1","firstname":"Bryan","lastname":"Connerty","emailaddress":"bryanc@test-email.com","phonenumber":"214-555-5551"},{"id":"2","firstname":"Kate","lastname":"Sacker","emailaddress":"kates@test-email.com","phonenumber":"214-555-5552"},{"id":"3","firstname":"Lonnie","lastname":"Watley","emailaddress":"lonniew@test-email.com","phonenumber":"214-555-5553"},{"id":"4","firstname":"Ira","lastname":"Schirmer","emailaddress":"iras@test-email","phonenumber":"214-555-5554"},{"id":"5","firstname":"Adam","lastname":"DeGiulio","emailaddress":"adamgs@test-email","phonenumber":"214-555-5554"}]
```

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
