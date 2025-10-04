package core

type Html struct {
	Title string
	Body  Renderable
}

func Page(children ...interface{}) Html {
	return Html{
		Body: Element("body", children...),
	}
}

func HTML() *Html {
	return &Html{Body: Text("")}
}

func (h *Html) Content(body Renderable) {
	h.Body = body
}

func (h *Html) SetTitle(title string) {
	h.Title = title
}

func (h Html) Render() string {
	content := "<!DOCTYPE html>"
	content += "<html>"

	// Head
	{
		content += "<head>"

		if h.Title != "" {
			content += "<title>" + h.Title + "</title>"
		}

		content += "</head>"
	}

	// Body
	content += "<body>" + h.Body.Render() + "</body>"

	content += "</html>"

	return content
}
