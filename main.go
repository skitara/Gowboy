package main

import (
	"fmt"
	"gowboy/character"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 40
	height = 20
)

type Tile rune

const (
	Empty  Tile = '.'
	Player Tile = '@'
	Cow    Tile = '🐄'
	Cactus Tile = '🌵'
	House1 Tile = '⌂'
	House2 Tile = '╔'
	House3 Tile = '╗'
	House4 Tile = '╝'
	House5 Tile = '╚'
)

type Position struct {
	x, y int
}

type Game struct {
	grid    [height][width]Tile
	player  Position
	running bool
}

func clearScreen() {
	cmd := exec.Command("clear") // Linux/Mac
	// cmd := exec.Command("cmd", "/c", "cls") // Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Game) init() {
	// Fill the field with empty cells
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			g.grid[y][x] = Empty
		}
	}

	// Place player at the center
	g.player = Position{width / 2, height / 2}
	g.grid[g.player.y][g.player.x] = Player

	// Place random objects
	rand.Seed(time.Now().UnixNano())
	g.placeRandomObjects(Cow, 5)
	g.placeRandomObjects(Cactus, 10)
	g.placeHouses(3)
}

func (g *Game) placeRandomObjects(obj Tile, count int) {
	for i := 0; i < count; i++ {
		x, y := rand.Intn(width), rand.Intn(height)
		if g.grid[y][x] == Empty {
			g.grid[y][x] = obj
		}
	}
}

func (g *Game) placeHouses(count int) {
	for i := 0; i < count; i++ {
		x, y := rand.Intn(width-3), rand.Intn(height-3)

		canPlace := true
		for dy := 0; dy < 3; dy++ {
			for dx := 0; dx < 3; dx++ {
				if g.grid[y+dy][x+dx] != Empty {
					canPlace = false
					break
				}
			}
			if !canPlace {
				break
			}
		}

		if canPlace {
			g.grid[y][x] = House2
			g.grid[y][x+1] = House3
			g.grid[y+1][x] = House5
			g.grid[y+1][x+1] = House4
			g.grid[y+2][x] = House1
			g.grid[y+2][x+1] = House1
		}
	}
}

func (g *Game) render() {
	// Move cursor to top-left instead of clearing
	fmt.Print("\033[H")

	// Set background to dusty brown
	fmt.Print("\033[48;2;139;69;19m")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch g.grid[y][x] {
			case Player:
				fmt.Print("\033[38;2;178;34;34m@\033[0m")
			case Cow:
				fmt.Print("\033[38;2;220;220;220m🐄\033[0m")
			case Cactus:
				fmt.Print("\033[38;2;34;139;34m🌵\033[0m")
			case House1, House2, House3, House4, House5:
				fmt.Print("\033[38;2;160;82;45m", string(g.grid[y][x]), "\033[0m")
			default:
				fmt.Print("\033[38;2;160;82;45m.\033[0m")
			}
		}
		fmt.Println()
	}

	// UI in sand color
	fmt.Println("\033[38;2;244;164;96mWASD - movement, I - inventory, H - horse, Q - quit\033[0m")
}

func (g *Game) movePlayer(dx, dy int, cowboy *character.Cowboy) {
	newX, newY := g.player.x+dx, g.player.y+dy

	if newX < 0 || newX >= width || newY < 0 || newY >= height {
		return
	}

	switch g.grid[newY][newX] {
	case Empty:
		g.grid[g.player.y][g.player.x] = Empty
		g.player.x, g.player.y = newX, newY
		g.grid[newY][newX] = Player
	case Cactus:
		cowboy.Health -= 5
		fmt.Printf("\aCactus has injured you (-5 HP)! %d HP remaining\n", cowboy.Health)
	case Cow:
		fmt.Println("\aYou scared the cow! She ran away.")
		g.grid[newY][newX] = Empty
	}
}

// func displayStats(stats character.Stats) {
// 	fmt.Printf("Strength:  %d\n", stats.Strength)
// 	fmt.Printf("Dexterity: %d\n", stats.Dexterity)
// 	fmt.Printf("Endurance: %d\n", stats.Endurance)
// 	fmt.Printf("Charisma:  %d\n", stats.Charisma)
// 	fmt.Printf("Fortune:   %d\n", stats.Fortune)
// }

func main() {
	// Console setup
	cmd := exec.Command("cmd", "/c", "chcp 65001 > nul")
	cmd.Run()
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	// Player creation
	var name string
	fmt.Println("\nWhat's your name, cowboy?")
	for {
		fmt.Scanln(&name)
		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty. Try again:")
	}

	gunslinger := character.NewCowboy(name)
	gunslinger.DistributePoints()

	// Game init
	game := Game{running: true}
	game.init()
	game.render()

	// Main loop
	for game.running {
		var input string
		fmt.Scanln(&input)

		if len(input) == 0 {
			continue
		}

		switch input[0] {
		case 'w', 'W':
			game.movePlayer(0, -1, gunslinger)
		case 'a', 'A':
			game.movePlayer(-1, 0, gunslinger)
		case 's', 'S':
			game.movePlayer(0, 1, gunslinger)
		case 'd', 'D':
			game.movePlayer(1, 0, gunslinger)
		case 'i', 'I':
			gunslinger.ShowInventory()
		case 'h', 'H':
			gunslinger.CheckHorse()
		case 'q', 'Q':
			game.running = false
		}

		game.render()

		if gunslinger.Health <= 0 {
			fmt.Println("\nYOU DIED")
			game.running = false
		}
	}
}
