// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.

package main

import "fmt"

type player struct {
	name   string
	atBats int
	hits   int
}

func (p *player) average() float64 {
	if p.atBats == 0 {
		return 0.0
	}
	return float64(p.hits) / float64(p.atBats)
}

func main() {
	// create players
	players := []player{
		{"john", 10, 7},
		{"jim", 12, 6},
		{"alex", 6, 4},
	}

	for i := range players {
		fmt.Printf("%s: AVG[.%.f]\n", players[i].name, players[i].average()*1000)
	}

}
