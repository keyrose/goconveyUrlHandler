package main

import (
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	_args := os.Args
	var args1, args2 string = "", ""
	var _url, _line string = "", ""

	for v, i := range _args {
		fmt.Print("i ::::" + i)
		fmt.Println("::: v ::::" + strconv.Itoa(v))
	}

	if len(_args) >= 3 {
		args1 = _args[1]
		args2 = _args[2]
	}
	thisUrl, err := url.Parse(args2)

	if err == nil {
		values := thisUrl.Query()

		_url = strings.Replace(values.Get("url"), "file://", "", -1)
		_line = values.Get("line")

		args2 = _url
		if _line != "" {
			args2 = args2 + ":" + _line
		}

		cmd := exec.Command(args1, args2)
		cmd.Run()
	}
}
