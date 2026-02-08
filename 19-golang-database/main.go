package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)
	opts := Options{}
}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("error", err)
	}

	employees := []User{
		{"john", "23", "9487670923", "ABC Com", Address{"Chennai", "TN", "India", "638001"}},
		{"Paul", "25", "6546676", "RRR Com", Address{"GMAK", "TN", "India", "638001"}},
		{"Quert", "28", "2262836", "BPTAE Com", Address{"GOA", "TN", "India", "638001"}},
		{"kumar", "30", "9487670923", "XYZ Com", Address{"bichapur", "TN", "India", "638001"}},
		{"Rmau", "29", "9487670923", "POI Com", Address{"hydrabad", "TELU", "India", "638001"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println(allusers)

	// if err := db.Delete("user", "john"); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
