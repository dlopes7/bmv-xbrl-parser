package main

import (
	"encoding/json"
	"fmt"
	"github.com/dlopes7/bmv-xbrl-parser/pkg"
	"io"
	"os"
)

func main() {
	var processed pkg.XBRL

	file, err := os.Open("fiduxbrl_1213276_3541-4_2022-01_1.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	contents, _ := io.ReadAll(file)
	if err := json.Unmarshal(contents, &processed); err != nil {
		panic(err)
	}
	for _, context := range processed.Contexts {
		fmt.Println(context.Period.Duration())
	}

}
