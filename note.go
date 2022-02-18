// Print out the note for guitar string/fret pairs, or vice versa.
// SvM 17-JUN-2021
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// We could also just calculate it.
var notes = [][]string{
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
	{"B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G"},
	{"D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D"},
	{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A"},
	{"E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E"},
}

// Translate something like "6/1" to "F"
func pos2note(s string) {
	ss := strings.Split(s, "/")
	string, _ := strconv.Atoi(ss[0])
	fret, _ := strconv.Atoi(ss[1])

	fmt.Println(string, "/", fret, "=", notes[string-1][fret])
}

// Translate something like "F" to "6/1" etc.
func note2pos(s string) {
	s = strings.ToUpper(s)
	for str := 0; str < 6; str++ {
		for not := 0; not < 12; not++ {
			if notes[str][not] == s {
				fmt.Println(s, "=", str+1, "/", not)

			}

		}
	}
}

func feed(s string) {
	if strings.Contains(s, "/") {
		pos2note(s)
	} else {
		note2pos(s)
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			feed(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			feed(s)
		}
	}

}
