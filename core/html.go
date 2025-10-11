package core

import (
	"fmt"
	"strings"
)

type Html struct {
	Head      Head
	Body      Renderable
	GlobalCSS Sheet
	DevMode   bool
}

func Page(children ...interface{}) *Html {
	return &Html{
		Body: Element("body", children...),
	}
}

func (h *Html) CSS(s Sheet) *Html {
	h.GlobalCSS = s
	return h
}

func (h *Html) Dev() *Html {
	h.DevMode = true
	return h
}

func (h *Html) Render() string {

	var b strings.Builder

	b.WriteString("<!DOCTYPE html>")
	b.WriteString("<html>")

	// Head
	{
		b.WriteString("<head>")

		b.WriteString(h.Head.Render())

		if h.GlobalCSS != nil {
			b.WriteString("<style>")
			b.WriteString(h.GlobalCSS.RenderGlobal())
			b.WriteString("</style>")
		}

		if h.DevMode {
			b.WriteString(`
			<script>
				console.log("Dev mode enabled");
			</script>`)
		}

		b.WriteString("</head>")

	}

	fmt.Fprintf(&b, "<body>%s</body>", h.Body.Render())

	fmt.Fprintf(&b, "</html>")

	return b.String()
}
