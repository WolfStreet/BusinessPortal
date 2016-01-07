package main

import (
	"fmt"
	"time"
)

func trial() {
	TimeNow := time.Now()
	var clock Clock
	var MidNight Clock
	MidNight.Hour = 0
	MidNight.Min = 0
	MidNight.Second = 0
	clock.Hour, clock.Min, clock.Second = TimeNow.Clock()
	if clock == MidNight {
		fmt.Println("Its Mid Night Now")
	} else {
		fmt.Println("Mid Night is still due")
	}
}
