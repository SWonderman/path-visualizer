package main

import (
	"slices"
	"sw/visualizer/algo"
	"sw/visualizer/graph"
	"sw/visualizer/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func win() {
	const WINDOW_WIDTH int32 = 450
	const WINDOW_HEIGHT int32 = 450

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Path Visualizer")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	const ROWS int32 = 9
	const BLOCK_SIZE int32 = 50

	matrix := [][]byte{
		{'S', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', 'x', '-', '-'},
		{'x', '-', '-', '-', 'x', '-', '-', '-', 'x'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', 'x'},
		{'-', '-', '-', '-', '-', 'x', '-', '-', '-'},
		{'x', '-', 'x', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', '-'},
		{'-', '-', '-', '-', '-', '-', '-', '-', 'F'},
	}

	start := graph.GridNode{Position: graph.Vector2{X: 0, Y: 0}}
	end := graph.GridNode{Position: graph.Vector2{X: 8, Y: 8}}

	obstacles := []byte{'x'}

	result := algo.RunUcs(&matrix, &start, &end, &obstacles)

	visitedIndex := 0
	pathIndex := 0
	fillInterval := float32(0.05)
	fillPathInterval := float32(0.3)

	intervalAccumulator := float32(0.0)
	pathIntervalAccumulator := float32(0.0)

	readyToDrawVisitedNodes := []*graph.GridNode{}
	readyToDrawPathEdges := []*graph.Edge{}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		intervalAccumulator += dt
		pathIntervalAccumulator += dt

		// Push nodes that should be drawn on interval basis
		if visitedIndex < len(result.Visited) && intervalAccumulator >= fillInterval {
			readyToDrawVisitedNodes = append(readyToDrawVisitedNodes, result.Visited[visitedIndex])
			visitedIndex += 1
			intervalAccumulator = 0.0
		}

		if result.CompletePath != nil && pathIndex < result.CompletePath.Len() && pathIntervalAccumulator >= fillPathInterval {
			readyToDrawPathEdges = append(readyToDrawPathEdges, utils.FindInLinkedListByIndex(result.CompletePath, pathIndex).(*graph.Edge))
			pathIndex += 1
			pathIntervalAccumulator = 0.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// TODO: optimize those nested loops
		for i := range ROWS {
			for j := range ROWS {
				// Draw start and end nodes
				if i == int32(start.Position.X) && j == int32(start.Position.Y) {
					rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Orange)
				}

				if i == int32(end.Position.X) && j == int32(end.Position.Y) {
					rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Green)
				}

				// Draw obstacles
				if slices.Contains(obstacles, matrix[i][j]) {
					rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Black)
				}

				// Draw visited
				for _, v := range readyToDrawVisitedNodes {
					// Do not draw on top of the start node
					if i == int32(start.Position.X) && j == int32(start.Position.Y) {
						continue
					}

					if i == int32(v.Position.X) && j == int32(v.Position.Y) {
						rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.LightGray)
					}
				}

				// Draw path
				for _, e := range readyToDrawPathEdges {
					if i == int32(e.To.Position.X) && j == int32(e.To.Position.Y) {
						rl.DrawRectangle(i*BLOCK_SIZE+2, j*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Beige)
					}
				}

				// Draw grid
				rl.DrawLine(i*BLOCK_SIZE, 0, i*BLOCK_SIZE, WINDOW_HEIGHT, rl.Black)
				rl.DrawLine(0, j*BLOCK_SIZE, WINDOW_WIDTH, j*BLOCK_SIZE, rl.Black)
			}
		}

		rl.EndDrawing()
	}
}

func main() {
	win()
}
