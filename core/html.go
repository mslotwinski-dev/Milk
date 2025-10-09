package core

type Html struct {
	Title     string
	Body      Renderable
	GlobalCSS Sheet
	DevMode   bool
}

func Page(children ...interface{}) *Html {
	return &Html{
		Body: Element("body", children...),
	}
}

func HTML() *Html {
	return &Html{Body: Text("")}
}

func (h *Html) CSS(s Sheet) *Html {
	h.GlobalCSS = s
	return h
}

func (h *Html) Content(body Renderable) {
	h.Body = body
}

func (h *Html) SetTitle(title string) {
	h.Title = title
}

func (h *Html) Dev() *Html {
	h.DevMode = true
	return h
}

func (h *Html) Render() string {
	content := "<!DOCTYPE html>"
	content += "<html>"

	// Head
	{
		content += "<head>"

		if h.Title != "" {
			content += "<title>" + h.Title + "</title>"
		}

		if h.DevMode {
			content += `
			<script>
				console.log("Dev mode enabled");
			</script>`
		}

		if h.GlobalCSS != nil {
			content += "<style>"
			content += h.GlobalCSS.RenderGlobal()
			content += "</style>"
		}

		content += "</head>"
	}

	// Body
	content += "<body>" + h.Body.Render() + "</body>"

	content += "</html>"

	return content
}
