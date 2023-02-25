package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	EMPTY  = " "
	PLAYER = "◯"
	CPU    = "×"
	SIZE   = 5
)

var board [SIZE][SIZE]string

func initBoard() {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			board[i][j] = EMPTY
		}
	}
}

func drawBoard() {
	fmt.Print("\n")
	for i := 0; i < SIZE; i++ {
		fmt.Print("| ")
		for j := 0; j < SIZE; j++ {
			fmt.Printf("%s | ", board[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func getInput() (int, int) {
	var x, y int
	fmt.Print("Enter the line number : ")
	fmt.Scan(&x)
	fmt.Print("Enter the column number : ")
	fmt.Scan(&y)
	return x - 1, y - 1
}

func isValidMove(x, y int) bool {
	if x < 0 || x >= SIZE || y < 0 || y >= SIZE {
		return false
	}
	if board[x][y] != EMPTY {
		return false
	}
	return true
}

func isGameOver() bool {
	// Whether the win condition is met in the vertical column
	for i := 0; i < SIZE; i++ {
		if board[0][i] != EMPTY && board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] == board[3][i] && board[3][i] == board[4][i] {
			return true
		}
	}
	// Whether the win condition is satisfied in the horizontal column or not
	for i := 0; i < SIZE; i++ {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][2] == board[i][3] && board[i][3] == board[i][4] {
			return true
		}
	}
	// Whether the winning condition is satisfied diagonally or not
	if board[0][0] != EMPTY && board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] == board[3][3] && board[3][3] == board[4][4] {
		return true
	}
	if board[0][4] != EMPTY && board[0][4] == board[1][3] && board[1][3] == board[2][2] && board[2][2] == board[3][1] && board[3][1] == board[4][0] {
		return true
	}
	// If all squares are filled, the game is considered a tie
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == EMPTY {
				return false
			}
		}
	}
	return true
}

func playerMove() {
	var x, y int
	for {
		x, y = getInput()
		if isValidMove(x, y) {
			break
		}
		fmt.Println("Invalid square.")
	}
	board[x][y] = PLAYER
}

func cpuMove() {

	// Finding candidates for empty squares
	candidates := make([][2]int, 0, SIZE*SIZE)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == EMPTY {
				candidates = append(candidates, [2]int{i, j})
			}
		}
	}

	// Place in adjacent squares
	for _, p := range candidates {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				x, y := p[0]+i, p[1]+j
				if x < 0 || x >= SIZE || y < 0 || y >= SIZE {
					continue
				}
				if board[x][y] == CPU {
					board[p[0]][p[1]] = CPU
					return
				}
			}
		}
	}

	// If the CPU side has a winning move, select that move
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == EMPTY {
				board[i][j] = CPU
				if isGameOver() {
					return
				}
				board[i][j] = EMPTY
			}
		}
	}

	// Block the opposing side's winning move
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == EMPTY {
				board[i][j] = PLAYER
				if isGameOver() {
					board[i][j] = CPU
					return
				}
				board[i][j] = EMPTY
			}
		}
	}

	// If the middle is available, choose the middle.
	if board[SIZE/2][SIZE/2] == EMPTY {
		board[SIZE/2][SIZE/2] = CPU
		return
	}

	// Randomly select an empty square
	for {
		x := randInt(0, SIZE-1)
		y := randInt(0, SIZE-1)
		if board[x][y] == EMPTY {
			board[x][y] = CPU
			break
		}
	}

	// Randomly select an empty square
	time.Sleep(500 * time.Millisecond)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	initBoard()
	drawBoard()
	for {
		fmt.Println("It is your turn.")
		playerMove()
		drawBoard()
		if isGameOver() {
			fmt.Println("You win!")
			os.Exit(0)
		}
		fmt.Println("It is your opponent's turn.")
		// Wait 0.5 seconds before displaying the CPU's move
		time.Sleep(500 * time.Millisecond)
		cpuMove()
		drawBoard()
		if isGameOver() {
			fmt.Println("Your opponent wins!")
			os.Exit(0)
		}
	}
}
