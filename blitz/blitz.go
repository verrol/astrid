package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/printer"
	"astrid/tile"
	"astrid/wordcolumn"
	"fmt"
)

//Blitz ...
type Blitz struct {
	WordColumn []wordcolumn.WordColumn
	Board      *board.Board
}

//New ...
func New(height uint16, width uint16) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	wc := make([]wordcolumn.WordColumn, board.Size)

	for i := 0; i < int(board.Size); i++ {
		wc[i].Words = make(map[string]struct{})
	}

	return &Blitz{wc, board}
}

//PrintWords ...
func (blitz *Blitz) PrintWords() {
	for _, wc := range blitz.WordColumn {
		for k := range wc.Words {
			fmt.Printf("%s:[%d]\n", k, wc.RootIndex)
		}
	}
}

//Start ...
func (blitz Blitz) Start() {
	blitz.Board.PlaceLetters("aseeaiaengseitse")
	//blitz.Board.PrintBoard()
	finder.FindWords(blitz.Board, blitz.WordColumn)
	//blitz.PrintWords()
	printer.PrintWords(blitz.Board, blitz.WordColumn)
}
