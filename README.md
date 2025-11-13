#  ✨ PicoCSS Plugin
The PicoCSS plugin integrates seamlessly with Via to provide elegant, responsive, and lightweight styling for your components without writing any CSS manually.

## Example
```go
package main

import (
	"github.com/go-via/via"
	"github.com/go-via/via-plugin-picocss/picocss"
	"github.com/go-via/via/h"
)

type Counter struct{ Count int }

func main() {
	v := via.New()

	v.Config(via.Options{
		DocumentTitle: "Via Counter",
		Plugins:       []via.Plugin{picocss.Default},
	})

	v.Page("/", func(c *via.Context) {
		data := Counter{Count: 0}
		step := c.Signal(1)

		increment := c.Action(func() {
			data.Count += step.Int()
			c.Sync()
		})

		c.View(func() h.H {
			return h.Main(h.Class("container"), h.Br(),
				h.H1(h.Text("⚡ Via Counter w/ PicoCSS Plugin")),
				h.Hr(),
				h.Div(
					h.H2(h.Strong(h.Text("Count - ")), h.Textf("%d", data.Count)),
					h.H5(h.Strong(h.Text("Step - ")), h.Span(step.Text())),
					h.Div(h.Role("group"),
						h.Input(h.Type("number"), step.Bind()),
						h.Button(h.Text("Increment"), increment.OnClick()),
					),
				),
			)
		})
	})

	v.Start()
}
```

## Docs
See PicoCSS [docs](https://picocss.com/docs) to learn how to use it.
