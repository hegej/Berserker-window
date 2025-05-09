package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1000)
	screenHight := int32(1000)
	var shakeOffsetX, shakeOffsetY float32
	var message string
	var textColor rl.Color

	rl.InitWindow(screenWidth, screenHight, "Berserker Windows")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	berserker := false

	for !rl.WindowShouldClose() {
		berserker = rl.IsMouseButtonDown(rl.MouseLeftButton)

		rl.BeginDrawing()

		if berserker {
			shakeOffsetX = float32(rand.Intn(11) - 5)
			shakeOffsetY = float32(rand.Intn(11) - 5)
			message = "The world is at WAR - BERSERKER MODE ON!!!"
			textColor = rl.Black
		} else {
			shakeOffsetX = 0
			shakeOffsetY = 0
			message = "The world is at peace..."
			textColor = rl.White
		}

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(shakeOffsetX, shakeOffsetY), 0, 1.0))
		rl.DrawText(message, 100, 250, 30, textColor)
		rl.EndMode2D()

		if berserker {
			rl.DrawRectangle(0, 0, screenWidth, screenHight, rl.NewColor(255, 0, 0, 80))
		}

		rl.EndDrawing()
	}
}
