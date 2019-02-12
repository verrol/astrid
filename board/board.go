package board

import (
	"astrid/dimension"
	"astrid/tile"
	"fmt"
)

//Board holds all tiles
type Board struct {
	Tiles     []tile.Tile
	dimension dimension.Dimension
}

//MakeBoard makes the word puzzle board from tiles
func (board *Board) MakeBoard(tiles []tile.Tile, height uint16, width uint16) {
	board.dimension.Height = height
	board.dimension.Width = width
	board.Tiles = tiles

	var i uint16
	var j uint16

	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			idx := (i * width) + j
			board.Tiles[idx].Coordinate.X = j
			board.Tiles[idx].Coordinate.Y = i
			board.Tiles[idx].SetPaths(board.dimension, board.Tiles)
		}
	}
}

//PlaceLetters places letters onto the board
func (board *Board) PlaceLetters(letters string) {
	var i uint16

	for i = 0; i < board.GetBoardSize(); i++ {
		board.Tiles[i].Letter = rune(letters[i])
	}
}

//GetBoardSize returns the size of the board
func (board Board) GetBoardSize() uint16 {
	return board.dimension.Height * board.dimension.Width
}

//PrintBoard ...
func (board *Board) PrintBoard() {
	fmt.Print("Printing board\n")
	var i uint16

	for i = 0; i < board.GetBoardSize(); i++ {
		board.Tiles[i].PrintTile()
		fmt.Print("\n")
	}
}
