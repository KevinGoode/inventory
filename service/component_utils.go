package service

import (
	"encoding/json"
	"io"
)

//Write writes any entity in json format to writer (eg stdout)
func Write(entity interface{}, w io.Writer) {
	b, _ := json.Marshal(entity)
	w.Write(b)
	w.Write([]byte("\n"))
}
