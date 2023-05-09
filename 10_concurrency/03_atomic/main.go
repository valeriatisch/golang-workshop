package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) damage(amount int) {
	// atomic.StoreInt32(&p.health, int32(int(p.health)-amount))
	atomic.AddInt32(&p.health, -int32(amount))
	fmt.Println("Player damaged with amount:", amount)
}

func startUI(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 800)
	for {
		fmt.Printf("Player health: %d\n", p.getHealth())
		<-ticker.C
	}
}

func startGame(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 700)
	for {
		p.damage(rand.Intn(50))
		if p.getHealth() <= 0 {
			fmt.Println("Game over!")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()
	go startUI(player)
	startGame(player)
}
