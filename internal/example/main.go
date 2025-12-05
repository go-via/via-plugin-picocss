package main

import (
	"github.com/go-via/via"
	"github.com/go-via/via-plugin-picocss/picocss"
	. "github.com/go-via/via/h"
)

func main() {

	v := via.New()

	v.Config(via.Options{
		DocumentTitle: "PicoCSS Plugin for Via",
		Plugins: []via.Plugin{

			// Plugin is placed here.
			// Use picocss.Default to add the plugin with default options.
			picocss.WithOptions(picocss.Options{
				Theme:         picocss.ThemeAmber,
				IncludeColors: true,
			}),
		},
	})

	v.Page("/", func(c *via.Context) {

		c.View(func() H {
			return Main(Class("container"),
				Section(
					Nav(
						Ul(Li(Strong(Text("⚡Via")))),
						Ul(
							Li(A(Class("pico-color-slate-50"), Text("About"), Href("https://github.com/go-via/via"))),
							Li(A(Class("pico-color-slate-50"), Text("Resources"), Href("https://github.com/orgs/go-via/repositories"))),
							Li(A(Class("pico-color-slate-50"), Text("Comunity"), Href("http://github.com/go-via/via/discussions"))),
						),
					),
				),

				Section(
					H1(Text(" PicoCSS Plugin for Via")),
					P(Text("Minimal CSS styles for semantic HTML ― Right in your Via app! ✨")),
				),

				A(Button(Class("outline"), Text("PicoCSS docs")), Href("https://picocss.com/docs"), Style("margin-right:10px")),
				A(Button(Text("View on GitHub")), Href("https://github.com/go-via/via-plugin-picocss")),

				Section(
					Hr(),
					H5(Text("About the Plugin")),
					P(Text("The PicoCSS plugin integrates seamlessly with Via to provide elegant, responsive, and lightweight styling for your components without writing any CSS manually.")),

					Hr(),
					H5(Text("Available Themes")),
					Div(Class("grid"),
						Ul(
							Li(Mark(Class("pico-background-amber"), Text("Amber")), Text(" *")),
							Li(Mark(Class("pico-background-blue"), Text("Blue"))),
							Li(Mark(Class("pico-background-cyan"), Text("Cyan"))),
							Li(Mark(Class("pico-background-fuchsia"), Text("Fuchsia"))),
							Li(Mark(Class("pico-background-green"), Text("Green"))),
						),
						Ul(
							Li(Mark(Class("pico-background-grey"), Text("Grey"))),
							Li(Mark(Class("pico-background-indigo"), Text("Indigo"))),
							Li(Mark(Class("pico-background-jade"), Text("Jade"))),
							Li(Mark(Class("pico-background-lime"), Text("Lime"))),
							Li(Mark(Class("pico-background-orange"), Text("Orange"))),
						),
						Ul(
							Li(Mark(Class("pico-background-pink"), Text("Pink"))),
							Li(Mark(Class("pico-background-pumpkin"), Text("Pumpkin"))),
							Li(Mark(Class("pico-background-purple"), Text("Purple"))),
							Li(Mark(Class("pico-background-red"), Text("Red"))),
							Li(Mark(Class("pico-background-sand"), Text("Sand"))),
						),
						Ul(
							Li(Mark(Class("pico-background-slate"), Text("Slate"))),
							Li(Mark(Class("pico-background-violet"), Text("Violet"))),
							Li(Mark(Class("pico-background-yellow"), Text("Yellow"))),
							Li(Mark(Class("pico-background-zinc"), Text("Zinc"))),
						),
					),
					P(Small(Text("(*) Default Theme."))),
				))

		})
	})

	v.Start()
}
