package main

import (
	"fmt"
	"os"

	codec_json "github.com/muranoya/structured-editor/src/codec/json"
	"github.com/muranoya/structured-editor/src/view/tui"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	j := codec_json.NewCodecJSON()
	dataObj, err := j.Decode(file)
	if err != nil {
		fmt.Printf("json decode error: %v", err)
		os.Exit(1)
	}

	tui := tui.NewTUIView(dataObj)
	if err := tui.Start(); err != nil {
		fmt.Println(err)
	}
}
