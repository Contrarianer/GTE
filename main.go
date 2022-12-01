package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/robertkrimen/otto"
)

// JsParser is a load js parse function, you can use it to call js get result
func JsParser(filePath string, functionName string, args ...interface{}) (result string, err error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	vm := otto.New()
	if _, err := vm.Run(string(bytes)); err != nil {
		return "", err
	}

	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		return "", err
	}
	return value.String(), err
}

func init() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	http.HandleFunc("/bs", func(writer http.ResponseWriter, request *http.Request) {
		jsFilePath := `web/example.js`
		result := strconv.FormatInt(time.Now().Unix(), 10)
		if rt, err := JsParser(jsFilePath, "encodeInp", result); err != nil {
			fmt.Println(`Error Js Parse call encodeInp`)
		} else {
			fmt.Fprintln(writer, rt)
		}
	})
}

func main() {
	if err := http.ListenAndServe(`:9090`, nil); err != nil {
		fmt.Printf("http server failed, err:%+v\n", err)
	}
}
