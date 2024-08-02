package main

import (
	"sw/visualizer/algo"
	"sw/visualizer/graph"
	"sw/visualizer/matrix"
	"sw/visualizer/stack"
	"sw/visualizer/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const COLUMNS int32 = 18
const ROWS int32 = 9
const BLOCK_SIZE int32 = 50

func convertMousePositionToGrid(mousePosition rl.Vector2) *graph.GridNode {
	row := int(mousePosition.Y / float32(BLOCK_SIZE))
	column := int(mousePosition.X / float32(BLOCK_SIZE))

	return &graph.GridNode{
		Row:    row,
		Column: column,
	}
}

func main() {
	matrix := matrix.GenerateEmptyMatrix(ROWS, COLUMNS)

	const windowWidth int32 = COLUMNS * BLOCK_SIZE
	const windowHeight int32 = ROWS * BLOCK_SIZE

	rl.InitWindow(windowWidth, windowHeight, "Path Visualizer")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	obstaclePositions := make(map[graph.GridNode]bool)

	startEndStack := stack.NewStack()
	startEndStack.Push("end")
	startEndStack.Push("start")

	// index 0 = start, index 1 = end
	startEndPositions := [2]*graph.GridNode{nil, nil}

	var result *algo.SearchResult
	var color rl.Color
	runAlgorithm := false
	wasAlgorithmRun := false

	visitedIndex := 0
	pathIndex := 0
	fillInterval := float32(0.02)
	fillPathInterval := float32(0.3)

	intervalAccumulator := float32(0.0)
	pathIntervalAccumulator := float32(0.0)

	readyToDrawVisitedNodes := []*graph.GridNode{}
	readyToDrawPathEdges := []*graph.Edge{}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		intervalAccumulator += dt
		pathIntervalAccumulator += dt

		if rl.IsKeyDown(rl.KeySpace) {
			if wasAlgorithmRun {
				visitedIndex = 0
				pathIndex = 0
				intervalAccumulator = 0.0
				pathIntervalAccumulator = 0.0
				readyToDrawVisitedNodes = nil
				readyToDrawPathEdges = nil
				result = nil
				wasAlgorithmRun = false
			}
		}
		if rl.IsKeyDown(rl.KeyEnter) {
			runAlgorithm = true
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			newObstaclePosition := convertMousePositionToGrid(rl.GetMousePosition())
			_, exists := obstaclePositions[*newObstaclePosition]
			if exists {
				delete(obstaclePositions, *newObstaclePosition)
			} else {
				obstaclePositions[*newObstaclePosition] = true
			}
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			selectedGridPosition := convertMousePositionToGrid(rl.GetMousePosition())
			if startEndPositions[0] != nil && selectedGridPosition.Row == startEndPositions[0].Row && selectedGridPosition.Column == startEndPositions[0].Column {
				startEndPositions[0] = nil
				startEndStack.Push("start")
			} else if startEndPositions[1] != nil && selectedGridPosition.Row == startEndPositions[1].Row && selectedGridPosition.Column == startEndPositions[1].Column {
				startEndPositions[1] = nil
				startEndStack.Push("end")
			} else {
				value := startEndStack.Pop()
				if value == "start" {
					startEndPositions[0] = selectedGridPosition
				} else if value == "end" {
					startEndPositions[1] = selectedGridPosition
				}
			}
		}

		if runAlgorithm {
			result = algo.RunUcs(matrix, startEndPositions[0], startEndPositions[1], obstaclePositions)
			runAlgorithm = false
			wasAlgorithmRun = true
		}

		if result != nil {
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
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for i := range ROWS {
			for j := range COLUMNS {

				if startEndPositions[0] != nil && (int32(startEndPositions[0].Row) == i && int32(startEndPositions[0].Column) == j) {
					color = rl.Orange
				} else if startEndPositions[1] != nil && (int32(startEndPositions[1].Row) == i && int32(startEndPositions[1].Column) == j) {
					color = rl.Green
				} else if obstaclePositions[graph.GridNode{Row: int(i), Column: int(j)}] {
					color = rl.Black
				} else {
					color = rl.White
				}

				rl.DrawRectangle(j*BLOCK_SIZE+2, i*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, color)

				if result != nil {
					// Draw visited
					for _, v := range readyToDrawVisitedNodes {
						// Do not draw on top of the start node
						if i == int32(startEndPositions[0].Row) && j == int32(startEndPositions[0].Column) {
							continue
						}

						if i == int32(v.Row) && j == int32(v.Column) {
							rl.DrawRectangle(j*BLOCK_SIZE+2, i*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.LightGray)
						}
					}

					// Draw path
					for _, e := range readyToDrawPathEdges {
						if i == int32(e.To.Row) && j == int32(e.To.Column) {
							rl.DrawRectangle(j*BLOCK_SIZE+2, i*BLOCK_SIZE+2, BLOCK_SIZE-4, BLOCK_SIZE-4, rl.Beige)
						}
					}
				}

				// Draw grid
				rl.DrawLine(j*BLOCK_SIZE, 0, j*BLOCK_SIZE, windowHeight, rl.Black)
				rl.DrawLine(0, i*BLOCK_SIZE, windowWidth, i*BLOCK_SIZE, rl.Black)
			}
		}

		rl.EndDrawing()
	}
}
