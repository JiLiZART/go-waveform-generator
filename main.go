package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/mdlayher/waveform"
)

func main() {

	if len(os.Args) > 1 {
		file, err := os.OpenFile(os.Args[1], os.O_RDWR|os.O_CREATE, 0755)

		checkError(err)

		w, err := waveform.New(file)

		checkError(err)

		values, err := w.Compute()

		encoded, _ := json.Marshal(values)
		fmt.Println(string(encoded))
	} else {
		fmt.Println("Please prvode file")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
