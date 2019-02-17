package printer

import (
	"astrid/board"
	"astrid/wordcolumn"
	"fmt"
	"sort"
	"sync"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

type printColumn struct {
	idx            uint16
	words          []string
	wordCount      uint16
	longestWordLen uint16
}

const spaceBetweenColumns uint16 = 2

var printColumns []printColumn
var longestWordLen uint16

func makePrintColumns(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	printColumns = make([]printColumn, board.Size)

	var wg sync.WaitGroup

	wg.Add(int(board.Size))

	for i, wc := range wordColumns {
		go func(pc *printColumn, wc wordcolumn.WordColumn) {
			defer wg.Done()

			pc.words = make([]string, len(wc.Words))
			pc.longestWordLen = wc.LongestWordLen
			pc.idx = wc.RootIndex
			pc.wordCount = wc.WordCount
			j := 0

			for k := range wc.Words {
				pc.words[j] = k
				j++
			}
			sort.Sort(byLength(pc.words))
		}(&printColumns[i], wc)
	}
	wg.Wait()
}

func findLongestWord() {
	for _, pc := range printColumns {
		if pc.longestWordLen > longestWordLen {
			longestWordLen = pc.longestWordLen
		}
	}
}

func pad(length uint16) uint16 {
	return (longestWordLen - length) + spaceBetweenColumns
}

func printColumnHeaders(start int, end int) {
	for i := start; i < end; i++ {
		str := fmt.Sprintf("[%d]", i+1)
		fmt.Printf("%s%*s", str, pad(uint16(len(str))), "")
	}
	fmt.Println()
}

func getLongestColumn(start int, end int) int {
	var count uint16

	for i := start; i < end; i++ {
		if printColumns[i].wordCount > count {
			count = printColumns[i].wordCount
		}
	}

	return int(count)
}

func printWord(word string, endColumn bool) {
	var padding uint16
	if endColumn {
		padding = 0
	} else {
		padding = pad(uint16(len(word)))
	}
	fmt.Printf("%s%*s", word, padding, "")
}

//PrintWords ...
func PrintWords(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	makePrintColumns(board, wordColumns)
	findLongestWord()

	//////////////////
	//values to read from configuration
	colsPerRow := 16
	maxWordsPerRow := 15
	/////////////////

	colHeaderStart := 0
	colHeaderEnd := colHeaderStart + colsPerRow

	var i uint16
	for i = 0; i < board.Size; i += uint16(colsPerRow) {
		printColumnHeaders(colHeaderStart, colHeaderEnd)

		longestColumn := getLongestColumn(colHeaderStart, colHeaderEnd)
		numPrintedRows := 0

		for j := 0; j < longestColumn; j++ {
			if numPrintedRows == maxWordsPerRow {
				numPrintedRows = 0
				fmt.Println()
				printColumnHeaders(colHeaderStart, colHeaderEnd)
			}
			numPrintedRows++

			for k := colHeaderStart; k < colHeaderEnd; k++ {
				if int(printColumns[k].wordCount) > j {
					printWord(printColumns[k].words[j], k == (colHeaderEnd-1))
				} else {
					fmt.Printf("%*s", pad(0), "")
				}
			}
			fmt.Println()
		}
		fmt.Println()

		colHeaderStart += colsPerRow
		colHeaderEnd = colHeaderStart + colsPerRow
		if colHeaderEnd >= int(board.Size) {
			colHeaderEnd = int(board.Size)
		}
	}

	for i = 0; i < (longestWordLen+spaceBetweenColumns)*uint16(colsPerRow); i++ {
		fmt.Print("+")
	}
	fmt.Printf("\n\n")
}
