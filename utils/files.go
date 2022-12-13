package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Check(err error, message string) {
	if err != nil {
		formatted := fmt.Sprintf("%s\n%s", message, err)
		log.Fatal(formatted)
	}
}

func WritePrettyRspToFile(filename string, rsp []byte) {
	var prettyRsp bytes.Buffer
	err := json.Indent(&prettyRsp, rsp, "", "    ")
	Check(err, "Couldn't format JSON response body!")

	os.WriteFile(filename, prettyRsp.Bytes(), 0644)
}
