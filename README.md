# gopherlol
> [bunnylol](https://www.quora.com/What-is-Facebooks-bunnylol) / [bunny1](http://www.bunny1.org) -like smart bookmarking tool, written in Go

```bash
go run . 
# then add `http://localhost:8080/?q=%s` as a search engine to your browser (where `%s` will be replaced with the search term by the browser)
```

## How to set it as default search engine
- in Chrome: [Set your default search engine](https://support.google.com/chrome/answer/95426)

## FAQ
- Where are the commands currently supported by gopherlol? => See `commands/commands.go`
- Why would a company run such a service internally? => Read about [how Facebook uses it internally](http://www.ccheever.com/blog/?p=74)