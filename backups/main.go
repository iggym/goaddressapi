package main

import (
	"fmt"
	"log"
	"net/http"
)

var addresses []Address

func main() {
	// add intial data
	addresses = append(addresses, Address{ID: "1", Firstname: "Bryan", Lastname: "Connerty", EmailAddress: "bryanc@test-email.com", PhoneNumber: "214-555-5551"})
	addresses = append(addresses, Address{ID: "2", Firstname: "Kate", Lastname: "Sacker", EmailAddress: "kates@test-email.com", PhoneNumber: "214-555-5552"})
	addresses = append(addresses, Address{ID: "3", Firstname: "Lonnie", Lastname: "Watley", EmailAddress: "lonniew@test-email.com", PhoneNumber: "214-555-5553"})
	addresses = append(addresses, Address{ID: "4", Firstname: "Ira", Lastname: "Schirmer", EmailAddress: "iras@test-email", PhoneNumber: "214-555-5554"})
	addresses = append(addresses, Address{ID: "5", Firstname: "Adam", Lastname: "DeGiulio", EmailAddress: "adamgs@test-email", PhoneNumber: "214-555-5554"})
	fmt.Println("Server starting (serving port: 8001)")
	log.Fatal(http.ListenAndServe(":8001", handlers()))
}
