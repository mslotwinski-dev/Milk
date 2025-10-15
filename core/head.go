package core

import (
	"fmt"
	"strings"
)

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
}

func (h *Head) Render() string {
	var b strings.Builder

	b.WriteString(`<meta charset="UTF-8">`)

	if h.Title != "" {
		fmt.Fprintf(&b, "<title>%s</title>", h.Title)
	}

	if h.Description != "" {
		fmt.Fprintf(&b, `<meta name="description" content="%s">`, h.Description)
	}

	if len(h.Keywords) > 0 {
		fmt.Fprintf(&b, `<meta name="keywords" content="%s">`, strings.Join(h.Keywords, ", "))
	}

	if h.Author != "" {
		fmt.Fprintf(&b, `<meta name="author" content="%s">`, h.Author)
	}

	if h.Robots != "" {
		fmt.Fprintf(&b, `<meta name="robots" content="%s">`, h.Robots)
	}

	if h.Refresh > 0 {
		fmt.Fprintf(&b, `<meta http-equiv="refresh" content="%d">`, h.Refresh)
	}

	if !h.NotResponsive {
		b.WriteString(`<meta name="viewport" content="width=device-width, initial-scale=1.0">`)
	}

	if h.OgTitle != "" {
		fmt.Fprintf(&b, `<meta property="og:title" content="%s">`, h.OgTitle)
	}
	if h.OgDescription != "" {
		fmt.Fprintf(&b, `<meta property="og:description" content="%s">`, h.OgDescription)
	}
	if h.OgImage != "" {
		fmt.Fprintf(&b, `<meta property="og:image" content="%s">`, h.OgImage)
	}
	if h.OgType != "" {
		fmt.Fprintf(&b, `<meta property="og:type" content="%s">`, h.OgType)
	}
	if h.OgUrl != "" {
		fmt.Fprintf(&b, `<meta property="og:url" content="%s">`, h.OgUrl)
	}

	if h.Favicon != "" {
		fmt.Fprintf(&b, `<link rel="icon" href="%s">`, h.Favicon)
	}
	if h.AppleIcon != "" {
		fmt.Fprintf(&b, `<link rel="apple-touch-icon" href="%s">`, h.AppleIcon)
	}
	if h.Manifest != "" {
		fmt.Fprintf(&b, `<link rel="manifest" href="%s">`, h.Manifest)
	}

	return b.String()
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
