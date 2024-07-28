package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"sw/visualizer/algo"
)

func win() {
	const WINDOW_WIDTH int32 = 450
	const WINDOW_HEIGHT int32 = 450

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Hello Window!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	const ROWS int32 = 9
	const BLOCK_SIZE int32 = 50

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := range ROWS {
			rl.DrawLine(i*BLOCK_SIZE, 0, i*BLOCK_SIZE, WINDOW_HEIGHT, rl.Black)
			rl.DrawLine(0, i*BLOCK_SIZE, WINDOW_WIDTH, i*BLOCK_SIZE, rl.Black)
		}

		rl.EndDrawing()
	}
}

func main() {
	searchResult := algo.Ucs()
	searchResult.ShowCompletePath()
}
