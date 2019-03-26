package cmd

import (
	"encoding/json"
	"io"
)

func prettyJson(j interface{}, w io.Writer) {
	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		panic(err)
	}
	w.Write(b)
	io.WriteString(w, "\n")
}
