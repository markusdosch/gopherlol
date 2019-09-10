package main

import (
	"fmt"
	"github.com/markusdosch/gopherlol/commands"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var commandsObject = new(commands.Commands)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	arr := strings.SplitN(q, " ", 2)
	cmdName := strings.Title(arr[0])
	cmdArg := ""
	if len(arr) > 1 {
		cmdArg = arr[1]
	}

	cmdMethod := reflect.ValueOf(commandsObject).MethodByName(cmdName)
	if cmdMethod == reflect.ValueOf(nil) {
		// cmdMethod not found => fall back to google
		url := fmt.Sprintf("https://www.google.com/#q=%s", url.QueryEscape(q))
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	url := ""
	cmdMethodNumIn := cmdMethod.Type().NumIn()
	if cmdMethodNumIn == 0 {
		res := cmdMethod.Call([]reflect.Value{})
		url = res[0].String()
	} else if cmdMethodNumIn == 1 {
		in := []reflect.Value{reflect.ValueOf(cmdArg)}
		res := cmdMethod.Call(in)
		url = res[0].String()
	} else {
		// cmdMethod was wrongly defined.
		// We currently only support cmdMethods with 0 or 1 parameters
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
