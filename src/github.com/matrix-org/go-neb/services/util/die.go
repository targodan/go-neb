package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type Dice []Die

func NewDice(dice ...Die) Dice {
	return append(Dice{}, dice...)
}

func (d Dice) Roll() []int {
	result := make([]int, d.Len())
	for i, die := range d {
		result[i] = die.Roll()
	}
	return result
}

func (d Dice) Len() int {
	return len(d)
}

func (d Dice) Join(other Dice) Dice {
	return append(d, other...)
}

func (d Dice) String() string {
	dice := make([]string, d.Len())
	for i, die := range d {
		dice[i] = die.String()
	}
	return strings.Join(dice, ", ")
}

func ItoaSlice(ints []int) []string {
	ret := make([]string, len(ints))
	for i := range ints {
		ret[i] = strconv.Itoa(ints[i])
	}
	return ret
}

type Die struct {
	Sides int
}

func (d Die) Roll() int {
	return rand.Intn(d.Sides) + 1
}

func (d Die) String() string {
	return "d" + strconv.Itoa(d.Sides)
}

// Note: w20 is the german version of d20. Feel free to add other letters as
// delimiters.
var diceRegex = regexp.MustCompile(`(?:\s|^)([0-9]*)(?:d|w)([0-9]+)`)

func ParseDice(text string) Dice {
	dice := make(Dice, 0)

	matches := diceRegex.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		fmt.Println(match)
		dieCount := match[1]
		dieSides := match[2]

		dieCountInt := 1
		if dieCount != "" {
			dieCountInt, _ = strconv.Atoi(dieCount)
		}
		dieSidesInt, _ := strconv.Atoi(dieSides)

		for i := 0; i < dieCountInt; i++ {
			dice = append(dice, Die{Sides: dieSidesInt})
		}
	}

	return dice
}
