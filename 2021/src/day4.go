package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	number string
	marked bool
}

type Board struct {
	cells  [5][5]Cell
	winner bool
}

type Game struct {
	input  []string
	boards []*Board
}

func (board *Board) isWinner() bool {
	i := 0

	for i < 5 {
		rowCount := 0
		columnCount := 0
		j := 0

		for j < 5 {
			row := board.cells[i][j]
			if row.marked {
				rowCount += 1
			}

			column := board.cells[j][i]
			if column.marked {
				columnCount += 1
			}

			j += 1
		}

		if rowCount == 5 || columnCount == 5 {
			return true
		}

		i += 1
	}

	return false
}

func (board *Board) mark(number string) {
	i := 0

	for i < 5 {
		j := 0

		for j < 5 {
			cell := board.cells[i][j]

			if cell.number == number {
				cell.marked = true
			}

			board.cells[i][j] = cell

			j += 1
		}

		i += 1
	}
}

func (board *Board) getScore() int {
	i := 0
	score := 0

	for i < 5 {
		j := 0

		for j < 5 {
			number := board.cells[i][j]

			if !number.marked {
				num, convError := strconv.Atoi(number.number)

				if convError != nil {
					log.Fatalf("Conversion error: %v", convError)
					continue
				}

				score += num
			}

			j += 1
		}

		i += 1
	}

	return score
}

func (game *Game) mark(number string) {
	i := 0

	for i < len(game.boards) {
		board := game.boards[i]

		if !board.winner {
			board.mark(number)
		}

		i += 1
	}
}

func (game *Game) getWinners() []int {
	i := 0
	indexes := []int{}

	for i < len(game.boards) {
		board := game.boards[i]

		if board.isWinner() && !board.winner {
			board.winner = true
			indexes = append(indexes, i)
		}

		game.boards[i] = board
		i += 1
	}

	return indexes
}

func (game *Game) allBoardsWon() bool {
	i := 0

	for i < len(game.boards) {
		board := game.boards[i]

		if !board.isWinner() {
			return false
		}

		i += 1
	}

	return true
}

func (game Game) start() {
	lastWinningScore := 0

	for {
		number, err := game.pop()

		if err != nil {
			log.Print("Input over")
			break
		}

		game.mark(*number)
		indices := game.getWinners()

		if len(indices) == 0 {
			continue
		}

		i := 0
		log.Printf("Indices %d", len(indices))

		for i < len(indices) {
			index := indices[i]
			board := game.boards[index]
			lastWinningScore = board.getTotalScore(*number)
			i += 1
		}

		if game.allBoardsWon() {
			log.Print("All Boards won")
			break
		}
	}

	log.Printf("Total Score %v", lastWinningScore)
}

func (board Board) getTotalScore(number string) int {
	board.print()
	log.Printf("Number %v", number)
	score := board.getScore()
	num, convError := strconv.Atoi(number)

	if convError != nil {
		log.Fatalf("Conversion error: %v", convError)
	}

	return score * num
}

func getEmptyBoard() [5][5]Cell {
	return [5][5]Cell{
		{Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}},
		{Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}},
		{Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}},
		{Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}},
		{Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}, Cell{number: "0", marked: false}},
	}
}

func createGame(entries []string) Game {
	input := strings.Split(entries[0], ",")
	i := 1
	boards := []*Board{}
	cells := getEmptyBoard()
	row := 0

	for i < len(entries) {
		column := 0
		columns := parseColumn(entries[i])

		for column < 5 {
			cells[row][column] = Cell{number: columns[column], marked: false}
			column += 1
		}

		if row == 4 {
			boards = append(boards, &Board{cells: cells, winner: false})
			cells = getEmptyBoard()
			row = 0
		} else {
			row += 1
		}

		i += 1
	}

	return Game{input: input, boards: boards}
}

func (game *Game) pop() (*string, error) {
	if len(game.input) == 0 {
		return nil, errors.New("empty_input")
	}

	value, input := game.input[0], game.input[1:]
	game.input = input
	return &value, nil
}

func (board Board) print() {
	log.Printf("Board : %v", board.cells)
}

func (game Game) print() {
	i := 0

	for i < len(game.boards) {
		log.Printf("Board #%d: %v", i, game.boards[i].cells)
		i += 1
	}

	log.Print("---")
}

func parseColumn(row string) []string {
	components := strings.Split(row, " ")
	columns := []string{}
	i := 0

	for i < len(components) {
		column := strings.TrimSpace(components[i])

		if column != "" {
			columns = append(columns, column)
		}

		i += 1
	}

	return columns
}

func Day4() {
	file, err := os.OpenFile("../input/day4.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []string{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		trimmed := strings.TrimSpace(line)

		if trimmed != "" {
			entries = append(entries, trimmed)
		}
	}

	game := createGame(entries)
	game.start()
}
