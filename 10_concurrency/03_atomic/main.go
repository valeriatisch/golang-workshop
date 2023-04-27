package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Atomic: low-level atomic memory operations on primitive data types, which are safe for concurrent access
// Atomic is much faster than mutexes.

type Player struct {
	health int32
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int {
	// Atomic read
	return 1
} 

func (p *Player) damage(amount int) {
	// atomic write

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
		p.damage(rand.Intn(50))
		if p.getHealth() <= 0 {
			fmt.Println("Game over!")
			break
		}
		time.Sleep(time.Millisecond * 700)
	}
}

func main() {
	player := NewPlayer()
	go startUI(player)
	startGame(player)
}

// 