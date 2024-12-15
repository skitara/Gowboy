package horse

import "fmt"

type Horse struct {
	Name       string
	Health     int
	Speed      int
	Endurance  int
	Level      int
	Experience int
}

func (h *Horse) DisplayStats() {
	fmt.Printf("Horse: %s\n", h.Name)
	fmt.Printf("Health:     %d\n", h.Health)
	fmt.Printf("Speed:      %d\n", h.Speed)
	fmt.Printf("Endurance:  %d\n", h.Endurance)
	fmt.Printf("Level:      %d\n", h.Level)
	fmt.Printf("Experience: %d\n", h.Experience)
}

func (h *Horse) Heal(amount int) {
	h.Health += amount
	if h.Health > 100 {
		h.Health = 100
	}
	fmt.Printf("%s's health restored to %d!\n", h.Name, h.Health)
}

func (h *Horse) Train(attribute string) {
	switch attribute {
	case "speed":
		h.Speed += 1
		fmt.Printf("%s's speed increased to %d!\n", h.Name, h.Speed)
	case "endurance":
		h.Endurance += 1
		fmt.Printf("%s's endurance increased to %d!\n", h.Name, h.Endurance)
	default:
		fmt.Println("Invalid attribute to train!")
	}
}

func (h *Horse) GainExperience(amount int) {
	h.Experience += amount
	fmt.Printf("%s gained %d experience!\n", h.Name, amount)

	if h.Experience >= h.Level*10 {
		h.Experience = 0
		h.Level++
		fmt.Printf("%s leveled up! Now at level %d!\n", h.Name, h.Level)

		h.Speed += 2
		h.Endurance += 2
		h.Health = 100
	}
}
