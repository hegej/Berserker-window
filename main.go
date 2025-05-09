package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHight := int32(1000)

	rl.InitWindow(screenWidth, screenHight, "Berserker Windows")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rand.seed(time.Now().UnixNano())
}
