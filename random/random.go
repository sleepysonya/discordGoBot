package randomCommand

import (
	"math/rand"
	"strconv"
	"strings"
)

func Random(rolls int64, sides int64) (int, string) {
	var total int64 = 0
	eachRolls := make([]string, rolls)
	var j int64 = 0
	for ; j < rolls; j++ {
		var oneRoll int64 = rand.Int63n(sides) + 1
		total += oneRoll
		eachRolls[j] = "#**" + strconv.Itoa(int(j+1)) + "**: " + strconv.Itoa(int(oneRoll))

	}
	var eachRollsString string = strings.Join(eachRolls, "\n")
	return int(total), eachRollsString
}
