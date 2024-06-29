package main

import "fmt"

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func SortByPrice(flights []Flight) []Flight {
	// implement
	return nil
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	var flights []Flight

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
