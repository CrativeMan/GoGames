package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib - test window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		drawSpaceship()

		rl.EndDrawing()
	}
}

func drawSpaceship() {
	rl.DrawLine(0, 0, 20, 100, rl.White)
	rl.DrawLine(20, 100, 40, 0, rl.White)
}
