// PicoCSSPlugin contains extentions plugins to configure and add PicoCSS to a Via application.
//
// See https://picocss.com for documentation and examples.
package picocss

import (
	_ "embed"
	"net/http"

	"github.com/go-via/via"
	"github.com/go-via/via/h"
)

//go:embed pico.amber.min.css
var picoAmber []byte

//go:embed pico.blue.min.css
var picoBlue []byte

//
//go:embed pico.cyan.min.css
var picoCyan []byte

//go:embed pico.fuchsia.min.css
var picoFuchsia []byte

//go:embed pico.green.min.css
var picoGreen []byte

//go:embed pico.grey.min.css
var picoGrey []byte

//go:embed pico.indigo.min.css
var picoIndigo []byte

//go:embed pico.jade.min.css
var picoJade []byte

//go:embed pico.lime.min.css
var picoLime []byte

//go:embed pico.orange.min.css
var picoOrange []byte

//go:embed pico.pink.min.css
var picoPink []byte

//go:embed pico.pumpkin.min.css
var picoPumpkin []byte

//go:embed pico.purple.min.css
var picoPurple []byte

//go:embed pico.red.min.css
var picoRed []byte

//go:embed pico.sand.min.css
var picoSand []byte

//go:embed pico.slate.min.css
var picoSlate []byte

//go:embed pico.violet.min.css
var picoViolet []byte

//go:embed pico.yellow.min.css
var picoYellow []byte

//go:embed pico.zinc.min.css
var picoZinc []byte

//go:embed pico.colors.min.css
var picoColors []byte

// Default adds the PicoCSS amber variant stylesheet to the base html document. To enable
// the plugin, add it to the *via.v runtime configuration in via.option.plugins.
func Default(v *via.V) {
	v.HandleFunc("GET /_plugins/picocss/assets/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		_, _ = w.Write(picoAmber)
	})
	v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/style.css")))
}

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

// WithOptions adds the PicoCSS with the provided Options to the base html document.
func WithOptions(options Options) via.Plugin {
	var cssBytes []byte
	switch options.Theme {
	case ThemeBlue:
		cssBytes = picoBlue
	case ThemeCyan:
		cssBytes = picoCyan
	case ThemeFuchia:
		cssBytes = picoFuchsia
	case ThemeGreen:
		cssBytes = picoGreen
	case ThemeGrey:
		cssBytes = picoGrey
	case ThemeIndigo:
		cssBytes = picoIndigo
	case ThemeJade:
		cssBytes = picoJade
	case ThemeLime:
		cssBytes = picoLime
	case ThemeOrange:
		cssBytes = picoOrange
	case ThemePink:
		cssBytes = picoPink
	case ThemePumpkin:
		cssBytes = picoPumpkin
	case ThemePurple:
		cssBytes = picoPurple
	case ThemeRed:
		cssBytes = picoRed
	case ThemeSand:
		cssBytes = picoSand
	case ThemeSlate:
		cssBytes = picoSlate
	case ThemeViolet:
		cssBytes = picoViolet
	case ThemeYellow:
		cssBytes = picoYellow
	case ThemeZinc:
		cssBytes = picoZinc
	default:
		cssBytes = picoAmber
	}

	return func(v *via.V) {
		v.HandleFunc("GET /_plugins/picocss/assets/pico.css", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css")
			_, _ = w.Write(cssBytes)
		})
		v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/pico.css")))

		if options.IncludeColors {
			v.HandleFunc("GET /_plugins/picocss/assets/pico-colors.css", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/css")
				_, _ = w.Write(picoColors)
			})
			v.AppendToHead(h.Link(h.Rel("stylesheet"), h.Href("/_plugins/picocss/assets/pico-colors.css")))
		}
	}
}
