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

	if cmdName == "List" || cmdName == "Help" {
		commandsType := reflect.TypeOf(&commands.Commands{})

		var html strings.Builder
		html.WriteString("<h1>gopherlol command list</h1>")
		html.WriteString("<ul>")
		for i := 0; i < commandsType.NumMethod(); i++ {
			method := commandsType.Method(i)

			takesArgs := ""
			if method.Type.NumIn() == 1 {
				takesArgs = ", takes args"
			}
			html.WriteString(fmt.Sprintf(
				"<li><strong>%s</strong>%s</li>",
				strings.ToLower(method.Name),
				takesArgs,
			))
		}
		html.WriteString("</ul>")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, html.String())
		return
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
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
