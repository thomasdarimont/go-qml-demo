package main

import (
	"github.com/go-qamel/qamel"
	"math/rand"
	"time"
)

// BackEnd is the bridge for communicating between QML and Go
type BackEnd struct {
	qamel.QmlObject
	_ func() int `slot:"getRandomNumber"`
	_ func() int `slot:"getRandomNumber2"`
}

func (b *BackEnd) getRandomNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(9999)
}

func (b *BackEnd) getRandomNumber2() int {
	return 2
}
