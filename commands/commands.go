package commands

import (
	"fmt"
	"net/url"
)

type Commands struct{}

func (c *Commands) Help() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

func (c *Commands) List() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

func (c *Commands) G(cmdArg string) string {
	return fmt.Sprintf("https://www.google.com/#q=%s", url.QueryEscape(cmdArg))
}

func (c *Commands) Author() string {
	return "https://www.markusdosch.com"
}

func (c *Commands) So(cmdArg string) string {
	return fmt.Sprintf("https://stackoverflow.com/search?q=%s", url.QueryEscape(cmdArg))
}
