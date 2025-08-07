# Go Address API

[![Go](https://github.com/iggym/goaddressapi/workflows/Go/badge.svg)](https://github.com/iggym/goaddressapi/actions)

An address API in Go. The API implements listing addresses, showing a single address, adding an address, and modifying an address. The API also enables exporting addresses to a CSV file and adding addresses by importing a CSV file.

## Project Overview

This project is a simple demonstration of a RESTful API built with Go and the Gorilla Mux router. It provides basic CRUD (Create, Read, Update, Delete) operations for a list of addresses. The addresses are stored in memory and are not persisted to a database.

### Features

*   List all addresses
*   Get a single address by ID
*   Create a new address
*   Update an existing address
*   Delete an address
*   Export addresses to a CSV file
*   Import addresses from a CSV file

### Data Model

The `Address` object has the following structure:

```json
{
  "id": "string",
  "firstname": "string",
  "lastname": "string",
  "emailaddress": "string",
  "phonenumber": "string"
}
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   Download and [install Go](https://golang.org/doc/install).
*   Make sure your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) is set.
*   Install [Gorilla Mux](https://github.com/gorilla/mux):
    ```sh
    go get -u github.com/gorilla/mux
    ```

### Installing

1.  Clone this repository:
    ```sh
    git clone https://github.com/iggym/goaddressapi.git
    ```
2.  Change into the project directory:
    ```sh
    cd goaddressapi
    ```
3.  Install the application:
    ```sh
    go install
    ```
4.  Run the server:
    ```sh
    goaddressapi
    ```
    The server will start on port `8001`.

## API Endpoints

The following endpoints are available:

| Method | Path                | Description                                |
| ------ | ------------------- | ------------------------------------------ |
| `GET`    | `/addresses`        | Get all addresses                          |
| `GET`    | `/addresses/{id}`   | Get a single address by ID                 |
| `POST`   | `/addresses/{id}`   | Create a new address                       |
| `PUT`    | `/addresses/{id}`   | Update an existing address                 |
| `DELETE` | `/addresses/{id}`   | Delete an address by ID                    |
| `GET`    | `/export`           | Export the list of addresses to a CSV file |
| `POST`   | `/import`           | Import addresses from a CSV file           |

### Examples

*   **Get all addresses:**
    ```sh
    curl localhost:8001/addresses
    ```
*   **Get a single address:**
    ```sh
    curl localhost:8001/addresses/2
    ```
*   **Create an address:**
    ```sh
    curl -X POST -d '{"firstname":"Bobby","lastname":"Axe","emailaddress":"axe@test.com","phonenumber":"214-545-5553"}' localhost:8001/addresses/31
    ```
*   **Update an address:**
    ```sh
    curl -X PUT -d '{"firstname":"Santa","lastname":"Claus","emailaddress":"santa@test.com","phonenumber":"214-545-5553"}' localhost:8001/addresses/4
    ```
*   **Delete an address:**
    ```sh
    curl -X DELETE localhost:8001/addresses/3
    ```
*   **Export to CSV:**
    ```sh
    curl localhost:8001/export --output list.csv
    ```
*   **Import from CSV:**
    ```sh
    curl -X POST -F "data=@tmp.csv" localhost:8001/import
    ```

## Running the Tests

To run the unit tests, use the following command:

```sh
go test -v
```

To view code coverage:

```sh
go test -cover
```

To generate an HTML coverage report:

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Potential Improvements

This project is a demonstration and has several areas for improvement:

*   **Code Organization:** The main `addressapi.go` file could be split into multiple files for better organization (e.g., `main.go`, `handlers.go`, `address.go`).
*   **Error Handling:** Error handling could be more consistent and robust.
*   **API Design:** The API could be improved to follow REST best practices more closely (e.g., using the request body for creating and updating resources).
*   **Persistence:** The data is currently stored in memory and is lost when the server restarts. A database or file-based persistence layer could be added.
*   **Testing:** The tests could be improved to provide more comprehensive coverage and better assertions.

## Built With

*   [Go](https://golang.org/)
*   [Gorilla Mux](https://github.com/gorilla/mux)

## Contributing

Contributions are welcome! Please submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
