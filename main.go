package main

import (
	"fmt"
	"gowboy/character"
	"gowboy/horse"
)

func main() {
	var name string

	// Name validation
	fmt.Println("What's your name, cowboy?")
	fmt.Scanln(&name)
	for name == "" {
		fmt.Println("Name cannot be empty. Please enter a valid name.")
		fmt.Scanln(&name)
	}

	// Create a Gunslinger with default stats and skills
	gunslinger := character.Cowboy{
		Name:   name,
		Health: 100,
		Stats: character.Stats{
			Strength:  5,
			Dexterity: 5,
			Endurance: 5,
			Charisma:  5,
			Fortune:   5,
		},
		Skills: character.Skills{
			Persuasion:   0,
			Haggling:     0,
			Stealth:      0,
			Safecracking: 0,
			Gambling:     0,
			Revolvers:    0,
			Rifles:       0,
			Brawling:     0,
		},
		Inventory: []string{},
		Dollars:   20,
		Horse: horse.Horse{
			Name:   "Swift",
			Health: 50,
			Speed:  10,
		},
	}

	// Display initial stats
	fmt.Println("\nAttention, cowboy! You'll need to redistribute some points for survival.")
	fmt.Println("Current stats:")
	displayStats(gunslinger.Stats)

	// Redistribute points
	gunslinger.DistributePoints()

	// Display final stats
	fmt.Println("\nYour final stats:")
	displayStats(gunslinger.Stats)
}

// Helper function to display stats
func displayStats(stats character.Stats) {
	fmt.Printf("Strength:  %d\n", stats.Strength)
	fmt.Printf("Dexterity: %d\n", stats.Dexterity)
	fmt.Printf("Endurance: %d\n", stats.Endurance)
	fmt.Printf("Charisma:  %d\n", stats.Charisma)
	fmt.Printf("Fortune:   %d\n", stats.Fortune)
}
