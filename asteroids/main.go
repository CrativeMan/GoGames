package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	CamerOffsetX int32
	CamerOffsetY int32
)

func main() {
	rl.InitWindow(800, 450, "Asteroids")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		inputManagment()
		drawSpaceship(CamerOffsetX, CamerOffsetY)

		rl.EndDrawing()
	}
}

func drawSpaceship(xOffset int32, yOffset int32) {
	rl.DrawLine(0+xOffset, 0+yOffset, 20+xOffset, 100+yOffset, rl.White)
	rl.DrawLine(20+xOffset, 100+yOffset, 40+xOffset, 0+yOffset, rl.White)
}

func inputManagment() {
	if rl.IsKeyPressed(rl.KeyW) {
		CamerOffsetY -= 20
	} else if rl.IsKeyPressed(rl.KeyS) {
		CamerOffsetY += 20
	} else if rl.IsKeyPressed(rl.KeyA) {
		CamerOffsetX -= 20
	} else if rl.IsKeyPressed(rl.KeyD) {
		CamerOffsetX += 20
	}
}
