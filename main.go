package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHight := int32(600)
	var shakeOffsetX, shakeOffsetY float32
	centerX := float32(screenWidth) / 2
	centerY := float32(screenHight) / 2
	var zoom float32 = 1.0
	var rotation float32 = 0
	var message string
	var textColor rl.Color

	rl.InitWindow(screenWidth, screenHight, "Berserker Window")
	defer rl.CloseWindow()

	berserkerTexture := rl.LoadTexture("Viking_berserker_character.png")
	defer rl.UnloadTexture(berserkerTexture)

	peaceTexture := rl.LoadTexture("Viking_peace.png")
	defer rl.UnloadTexture(peaceTexture)

	rl.SetTargetFPS(60)

	berserker := false

	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(screenWidth)/2, float32(screenHight)/2)
	camera.Offset = rl.NewVector2(float32(screenWidth)/2+shakeOffsetX, float32(screenHight)/2+shakeOffsetY)
	camera.Zoom = 1.0

	for !rl.WindowShouldClose() {
		berserker = rl.IsMouseButtonDown(rl.MouseLeftButton)

		if berserker {
			shakeOffsetX = float32(rand.Intn(11) - 5)
			shakeOffsetY = float32(rand.Intn(11) - 5)
			zoom = 1.03 + rand.Float32()*0.02
			rotation = float32(rand.Intn(7) - 3)
			message = "The world is at WAR - BERSERKER MODE ON!!!"
			textColor = rl.DarkBrown
		} else {
			shakeOffsetX = 0
			shakeOffsetY = 0
			zoom = 1.0
			rotation = 0
			message = "The world is at peace..."
			textColor = rl.White
		}

		camera.Offset = rl.NewVector2(centerX+shakeOffsetX, centerY+shakeOffsetY)
		camera.Target = rl.NewVector2(centerX, centerY)
		camera.Zoom = zoom
		camera.Rotation = rotation

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(camera)

		iconSize := float32(128)
		var texture rl.Texture2D
		var tint rl.Color = rl.White
		var scale float32 = iconSize / float32(berserkerTexture.Width)
		var spriteY float32

		if berserker {
			texture = berserkerTexture
			scale *= 1.0 + 0.05*float32(rand.Intn(3))
			tint = rl.Red
			spriteY = centerY - 140
		} else {
			texture = peaceTexture
			spriteY = centerY - 100
		}

		spritePos := rl.NewVector2(
			centerX-float32(texture.Width)*scale/2,
			spriteY,
		)

		shadowWidth := float32(texture.Width) * scale * 0.6
		shadowHeight := float32(texture.Height) * scale * 0.2
		shadowX := centerX - shadowWidth/2
		shadowY := spriteY + float32(texture.Height)*scale - shadowHeight/2

		rl.DrawEllipse(
			int32(shadowX+shadowWidth/2),
			int32(shadowY+shadowHeight/2),
			shadowWidth/2,
			shadowHeight/2,
			rl.Fade(rl.Black, 0.4),
		)

		rl.DrawTextureEx(texture, spritePos, 0, scale, tint)

		textWidth := rl.MeasureText(message, 30)
		rl.DrawText(message, screenWidth/2-textWidth/2, screenHight/2-15, 30, textColor)
		rl.EndMode2D()

		if berserker {
			rl.DrawRectangle(0, 0, screenWidth, screenHight, rl.NewColor(255, 0, 0, 70))
		}

		rl.EndDrawing()
	}
}
