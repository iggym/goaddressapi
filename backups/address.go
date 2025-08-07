package main

// Address represents a single address record.
type Address struct {
	ID           string `json:"id,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	EmailAddress string `json:"emailaddress,omitempty"`
	PhoneNumber  string `json:"phonenumber,omitempty"`
}

const csvfilepath = "addresses.csv"
const tmpdir = "tmp"
