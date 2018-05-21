package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/mdlayher/waveform"
)

func main() {

	if len(os.Args) > 1 {
		file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0755)

		if err != nil {
			panic(err)
		}

		w, err := waveform.New(file)

		if err != nil {
			panic(err)
		}

		values, err := w.Compute()

		if err != nil {
			panic(err)
		}

		encoded, err := json.Marshal(values)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(encoded))
	} else {
		fmt.Println("Please prvode file")
	}
}
