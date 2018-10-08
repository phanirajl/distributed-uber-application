package main

type Driver struct {
	name     string
	lastName string
}

type Car struct {
	licensePlate int
	model        string
	owner        Driver
}
