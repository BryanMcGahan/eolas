package gapbuffer

import (
	"fmt"
	"strings"
)

type GapBuffer struct {
	Buffer    []string
	GapStart  int
	GapLength int
}

func Init(content string) GapBuffer {
	return GapBuffer{
		Buffer:    append([]string{"󰂭", "󰂭", "󰂭", "󰂭", "󰂭", "󰂭", "󰂭", "󰂭", "󰂭", "󰂭"}, strings.Split(content, "")...),
		GapStart:  0,
		GapLength: 10,
	}
}

func (gb *GapBuffer) Insert(char string, x int) {
}

func (gb *GapBuffer) MoveGap(x int) {

	if x > len(gb.Buffer)-1 || x < 0 {
		fmt.Printf("Cursor is outside of buffer\n")
		return
	}

	// Gap Start is before cursor position
	// Need to move chars after gap to the front of gap
	if gb.GapStart < x {
		for x > gb.GapStart {
			var indexToMove int = gb.GapStart + gb.GapLength
			if indexToMove < len(gb.Buffer) {
				gb.Buffer[gb.GapStart] = gb.Buffer[indexToMove]
				gb.Buffer[indexToMove] = "󰂭"
			}
			gb.GapStart++
		}
	}

	// Gap Start is after cursor position
	// Need to move chars before gap to the end of the gap
	if gb.GapStart > x {
		for x < gb.GapStart {
			var indexToMove int = gb.GapStart - 1
			if indexToMove >= 0 {
				gb.Buffer[gb.GapStart+gb.GapLength-1] = gb.Buffer[indexToMove]
				gb.Buffer[indexToMove] = "󰂭"
			}

			gb.GapStart--
		}
	}
}

func (gb *GapBuffer) Print() {
	fmt.Println(gb.Buffer)
}
