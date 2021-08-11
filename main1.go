package main

import ("fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)

	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)

	fmt.Println("rooms :", rooms, " roomsOccupied :", roomsOccupied)
	// Example: 1
	fmt.Println("boolean expression : 'rooms > roomsOccupied' is equal to:")
	fmt.Println(rooms > roomsOccupied)

	// Example: 2
	fmt.Println("boolean expression : 'roomsOccupied > rooms' is equal to:")
	fmt.Println(roomsOccupied > rooms)

	// Example: 3
	fmt.Println("boolean expression : 'roomsOccupied == rooms' is equal to:")
	fmt.Println(roomsOccupied == rooms)
}
