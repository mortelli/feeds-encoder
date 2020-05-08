package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read input
	input := os.Args[1:] // first element is path to the program
	if len(input) != 1 {
		panic("exactly one argument must be provided")
	}

	// split comma-separated values
	prefixes := strings.Split(input[0], ",")

	var targets [][]byte
	// for each comma-separate prefix, decode hex to bytes
	for _, prefix := range prefixes {
		target, err := hex.DecodeString(prefix)
		if err != nil {
			panic(err)
		}
		// append result to target list
		targets = append(targets, target)
	}

	// encode targets
	encodedTargets, err := json.Marshal(targets)
	if err != nil {
		panic(err)
	}

	// return hex string of encoded targets
	strEncoding := hex.EncodeToString(encodedTargets)
	fmt.Println(strEncoding)
}
