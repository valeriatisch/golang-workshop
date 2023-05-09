package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	health int
	mu     sync.RWMutex
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) damage(amount int) {
	p.mu.Lock()
	defer p.mu.Unlock()
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
