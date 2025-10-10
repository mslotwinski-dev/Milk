package core

type Head struct {
	Title         string
	Keywords      []string
	Description   string
	Author        string
	Refresh       int
	NotResponsive bool

	OgTitle       string
	OgDescription string
	OgImage       string
	OgType        string
	OgUrl         string
	Robots        string

	Favicon   string
	AppleIcon string
	Manifest  string

	Scripts []Script
}

func (h *Html) SetTitle(title string) {
	h.Head.Title = title
}

func (h *Html) SetKeywords(keywords ...string) {
	h.Head.Keywords = keywords
}

func (h *Html) SetDescription(description string) {
	h.Head.Description = description
}

func (h *Html) SetAuthor(author string) {
	h.Head.Author = author
}

func (h *Html) SetRefresh(seconds int) {
	h.Head.Refresh = seconds
}

func (h *Html) SetNotResponsive() {
	h.Head.NotResponsive = true
}

func (h *Html) SetOgTitle(title string) {
	h.Head.OgTitle = title
}

func (h *Html) SetOgDescription(description string) {
	h.Head.OgDescription = description
}

func (h *Html) SetOgImage(url string) {
	h.Head.OgImage = url
}

func (h *Html) SetOgType(t string) {
	h.Head.OgType = t
}

func (h *Html) SetOgUrl(url string) {
	h.Head.OgUrl = url
}

func (h *Html) SetRobots(content string) {
	h.Head.Robots = content
}

func (h *Html) SetFavicon(url string) {
	h.Head.Favicon = url
}

func (h *Html) SetAppleIcon(url string) {
	h.Head.AppleIcon = url
}

func (h *Html) SetManifest(url string) {
	h.Head.Manifest = url
}
