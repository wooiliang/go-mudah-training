// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import (
	"fmt"
)

// Add imports.

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.
type Player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (player *Player) average() float64 {
	if player.atBats == 0 {
		return 0.0
	}
	return float64(player.hits) / float64(player.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []Player{
		{"A", 188, 99},
		{"B", 177, 88},
		{"C", 166, 77},
	}

	// Display the batting average for each player in the slice.
	// dont use this method!
	// for _, player := range players {
	// 	fmt.Println(player.name, player.average())
	// }

	for i := range players {
		fmt.Println(players[i].name, players[i].average())
	}
}
