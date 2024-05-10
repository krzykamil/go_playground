package main

import (
    "fmt";
    "math/rand"
)

type Player interface {
    KickBall()
}

type CR7 struct {
    stamina int
    power int
    SUI int
}

func (f CR7) KickBall() {
    shot := f.stamina + f.power * f.SUI
    fmt.Println("CR7 kicking", shot)
}

type FootballPlayer struct {
    stamina int
    power int
}

func (f FootballPlayer) KickBall() {
    shot := f.stamina + f.power
    fmt.Println("Kicking", shot)
}

func main() {
    team := make([]Player, 11)
    for i := 0; i < len(team)-1; i++ {
        team[i] = FootballPlayer{rand.Intn(10), rand.Intn(10)}
    }
    team[len(team) - 1] = CR7{10, 10, 10}
    for i :=0; i < len(team); i++ {
        team[i].KickBall()
    }
}
