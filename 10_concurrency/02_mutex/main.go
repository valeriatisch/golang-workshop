package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Mutex is a mutual exclusion lock.
// A mutex allows only one goroutine to access the shared resource at a time.
// RWMutex allows multiple readers to access the shared resource simultaneously but only one writer at a time.

type Player struct {
	health int
	// Add mutex

}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int {
	// Lock
	return p.health
}

func (p *Player) damage(amount int) {
	// Lock
	p.health -= amount
	fmt.Println("Player damaged with amount:", amount)
}

func startUI(p *Player) {
	for {
		fmt.Printf("Player health: %d\n", p.getHealth())
		time.Sleep(time.Millisecond * 800)
	}
}

func startGame(p *Player) {
	for {
		p.damage(rand.Intn(20))
		if p.getHealth() <= 0 {
			fmt.Println("Game over!")
			break
		}
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	player := NewPlayer()
	go startUI(player)
	startGame(player)
}
