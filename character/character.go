package character

import "fmt"

type Stats struct {
	Strength  int
	Dexterity int
	Endurance int
	Charisma  int
	Fortune   int
}

func (s *Stats) getStatValue(stat int) int {
	switch stat {
	case 1:
		return s.Strength
	case 2:
		return s.Dexterity
	case 3:
		return s.Endurance
	case 4:
		return s.Charisma
	case 5:
		return s.Fortune
	default:
		return 0
	}
}

func (s *Stats) incrementStat(stat int) {
	switch stat {
	case 1:
		if s.Strength < 10 {
			s.Strength++
		}
	case 2:
		if s.Dexterity < 10 {
			s.Dexterity++
		}
	case 3:
		if s.Endurance < 10 {
			s.Endurance++
		}
	case 4:
		if s.Charisma < 10 {
			s.Charisma++
		}
	case 5:
		if s.Fortune < 10 {
			s.Fortune++
		}
	}
}

func (s *Stats) decrementStat(stat int) {
	switch stat {
	case 1:
		if s.Strength > 1 {
			s.Strength--
		}
	case 2:
		if s.Dexterity > 1 {
			s.Dexterity--
		}
	case 3:
		if s.Endurance > 1 {
			s.Endurance--
		}
	case 4:
		if s.Charisma > 1 {
			s.Charisma--
		}
	case 5:
		if s.Fortune > 1 {
			s.Fortune--
		}
	}
}

type Skills struct {
	Persuasion   int
	Haggling     int
	Stealth      int
	Safecracking int
	Gambling     int
	Revolvers    int
	Rifles       int
	Brawling     int
}

type Cowboy struct {
	Name      string
	Health    int
	Stats     Stats
	Skills    Skills
	Inventory []string
	Dollars   int
}

func (c *Cowboy) DistributePoints() {
	points := 5 // Total points to redistribute

	for points > 0 {
		fmt.Println("You have", points, "points left to distribute.")
		fmt.Println("Choose a stat to increase:")
		fmt.Println("1-Strength, 2-Dexterity, 3-Endurance, 4-Charisma, 5-Fortune")

		var choice int
		fmt.Scanln(&choice)

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice. Please select a number between 1 and 5.")
			continue
		}

		// Increment the chosen stat
		if c.Stats.getStatValue(choice) < 10 {
			c.Stats.incrementStat(choice)
			points--
		} else {
			fmt.Println("This stat is already at the maximum value of 10.")
		}
	}
}
