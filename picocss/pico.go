// PicoCSSPlugin contains extentions plugins to configure and add PicoCSS to a Via application.
//
// See https://picocss.com for documentation and examples.
package picocss

import (
	"embed"
	"net/http"

	"github.com/go-via/via"
	"github.com/go-via/via/h"
)

//go:embed assets/*
var assets embed.FS

// Options can be used with WithOptions to configure the picocss plugin.
type Options struct {
	// Configures the plugin to add a color variant of the PicoCSS stylesheet.
	// See https://picocss.com/docs/version-picker.
	Theme Theme

	// If true, configures the plugin to add the PicoCSS colors stylesheet.
	// See https://picocss.com/docs/colors.
	IncludeColors bool
}

type Theme int

const (
	undefined Theme = iota
	colors
	ThemeAmber
	ThemeBlue
	ThemeCyan
	ThemeFuchia
	ThemeGreen
	ThemeGrey
	ThemeIndigo
	ThemeJade
	ThemeLime
	ThemeOrange
	ThemePink
	ThemePumpkin
	ThemePurple
	ThemeRed
	ThemeSand
	ThemeSlate
	ThemeViolet
	ThemeYellow
	ThemeZinc
)

var themes = map[Theme]string{
	undefined:    "assets/pico.amber.min.css",
	colors:       "assets/pico.colors.min.css",
	ThemeAmber:   "assets/pico.amber.min.css",
	ThemeBlue:    "assets/pico.blue.min.css",
	ThemeCyan:    "assets/pico.cyan.min.css",
	ThemeFuchia:  "assets/pico.fuchia.min.css",
	ThemeGreen:   "assets/pico.green.min.css",
	ThemeGrey:    "assets/pico.grey.min.css",
	ThemeIndigo:  "assets/pico.indigo.min.css",
	ThemeJade:    "assets/pico.jade.min.css",
	ThemeLime:    "assets/pico.lime.min.css",
	ThemeOrange:  "assets/pico.orange.min.css",
	ThemePink:    "assets/pico.pink.min.css",
	ThemePumpkin: "assets/pico.pumpkin.min.css",
	ThemePurple:  "assets/pico.purple.min.css",
	ThemeRed:     "assets/pico.red.min.css",
	ThemeSand:    "assets/pico.sand.min.css",
	ThemeSlate:   "assets/pico.slate.min.css",
	ThemeViolet:  "assets/pico.violet.min.css",
	ThemeYellow:  "assets/pico.yellow.min.css",
	ThemeZinc:    "assets/pico.zinc.min.css",
}

// Default adds the PicoCSS amber variant stylesheet to the base html document. To enable
// the plugin, add it to the *via.v runtime configuration in via.option.plugins.
var Default via.Plugin = func(v *via.V) {

	v.HandleFunc("GET /_plugins/picocss/assets/style.css", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/css")
		pico, err := assets.ReadFile(themes[undefined])
		if err != nil {
			http.Error(w, "failed to open stylesheet file", http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(pico)
	})
	v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/style.css")))
}

// WithOptions adds the PicoCSS with the provided Options to the base html document.
func WithOptions(options Options) via.Plugin {
	theme, ok := themes[options.Theme]
	if !ok {
		theme = themes[undefined]
	}

	return func(v *via.V) {

		v.HandleFunc("GET /_plugins/picocss/assets/style.css", func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "text/css")
			pico, err := assets.ReadFile(theme)
			if err != nil {
				http.Error(w, "failed to open stylesheet file", http.StatusInternalServerError)
				return
			}
			_, _ = w.Write(pico)
		})
		v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/style.css")))

		if options.IncludeColors {

			v.HandleFunc("GET /_plugins/picocss/assets/style.colors.css", func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-Type", "text/css")
				picoColors, err := assets.ReadFile(themes[colors])
				if err != nil {
					http.Error(w, "failed to open stylesheet file", http.StatusInternalServerError)
					return
				}
				_, _ = w.Write(picoColors)
			})
			v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/style.colors.css")))
		}
	}
}
