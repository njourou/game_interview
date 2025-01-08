package main

import (
    "fmt"
)

type SpinResult struct {
    Grid      [][]string `json:"grid"`
    Winnings  float64    `json:"winnings"`
    WinLines  []string   `json:"winLines"`
    FreeSpins int        `json:"freeSpins"`
    BonusRound bool      `json:"bonusRound"`
}

func spin(betAmount float64) SpinResult {
    grid := generateGrid()

    winnings, winLines := calculateWinnings(grid, betAmount)
    freeSpins := calculateFreeSpins(winLines)
    bonusRound := calculateBonusRound(winLines)

    return SpinResult{
        Grid:      grid,
        Winnings:  winnings,
        WinLines:  winLines,
        FreeSpins: freeSpins,
        BonusRound: bonusRound,
    }
}

func generateGrid() [][]string {
    grid := make([][]string, 3)
    for i := range grid {
        grid[i] = make([]string, 5)
        for j := range grid[i] {
            grid[i][j] = getRandomSymbol()
        }
    }
    return grid
}

func calculateWinnings(grid [][]string, betAmount float64) (float64, []string) {
    var totalWinnings float64
    var winLines []string

    // Check all lines
    totalWinnings += checkLines(grid, &winLines, betAmount)

    // check on this !!!
    if len(winLines) == 0 {
        winLines = []string{}
    }

    return totalWinnings, winLines
}

func checkLines(grid [][]string, winLines *[]string, betAmount float64) float64 {
    var totalWinnings float64
    // Check horizontal lines
    for i := 0; i < 3; i++ {
        symbol := grid[i][0]
        count := 1
        for j := 1; j < 5; j++ {
            if grid[i][j] == symbol {
                count++
            } else {
                break
            }
        }
        if count >= 3 {
            winnings := calculateWinningsForLine(symbol, count, betAmount)
            totalWinnings += winnings
            *winLines = append(*winLines, fmt.Sprintf("Horizontal line %d: %d symbols matched", i+1, count))
        }
    }

    // Check vertical lines
    for j := 0; j < 5; j++ {
        symbol := grid[0][j]
        count := 1
        for i := 1; i < 3; i++ {
            if grid[i][j] == symbol {
                count++
            } else {
                break
            }
        }
        if count >= 3 {
            winnings := calculateWinningsForLine(symbol, count, betAmount)
            totalWinnings += winnings
            *winLines = append(*winLines, fmt.Sprintf("Vertical line %d: %d symbols matched", j+1, count))
        }
    }

    // Check diagonal lines
    if grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] {
        winnings := calculateWinningsForLine(grid[0][0], 3, betAmount)
        totalWinnings += winnings
        *winLines = append(*winLines, "Diagonal line 1: 3 symbols matched")
    }
    if grid[0][4] == grid[1][3] && grid[1][3] == grid[2][2] {
        winnings := calculateWinningsForLine(grid[0][4], 3, betAmount)
        totalWinnings += winnings
        *winLines = append(*winLines, "Diagonal line 2: 3 symbols matched")
    }

    return totalWinnings
}

func calculateWinningsForLine(symbol string, count int, betAmount float64) float64 {
    multiplier := getMultiplier(symbol)
    return float64(count) * multiplier * betAmount
}

func calculateFreeSpins(winLines []string) int {
    // award 1 free spin if 1 line is won
    return len(winLines)
}

func calculateBonusRound(winLines []string) bool {
    // added bonus round if there are more than 2 winning lines
    return len(winLines) > 2
}