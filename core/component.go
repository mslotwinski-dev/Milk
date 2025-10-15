package core

import "fmt"

type Renderable interface {
	Render() string
}

type Component struct {
	Tag      string
	Children []Renderable
	Styles   []StyleByte
	Classes  []string
	Id       string
	LocalCSS *Sheet
	Alpine   AlpineJS

	Atributes Atributes
}

func (c Component) Render() string {
	var result string

	result += "<" + c.Tag

	if len(c.Classes) > 0 {
		result += " class=\""
		for i, class := range c.Classes {
			result += class
			if i < len(c.Classes)-1 {
				result += " "
			}
		}
		result += "\""
	}

	if c.Id != "" {
		result += " id=\"" + c.Id + "\""
	}

	result += c.Alpine.Render()

	if c.Tag == "img" {
		if c.Atributes.ImageSrc != "" {
			result += " src=\"" + c.Atributes.ImageSrc + "\""
		}

		if c.Atributes.ImageAlt != "" {
			result += " alt=\"" + c.Atributes.ImageAlt + "\""
		}

		if c.Atributes.ImageWidth > 0 {
			result += " width=\"" + fmt.Sprint(c.Atributes.ImageWidth) + "\""
		}

		if c.Atributes.ImageHeight > 0 {
			result += " height=\"" + fmt.Sprint(c.Atributes.ImageHeight) + "\""
		}
	}

	if len(c.Styles) > 0 {
		result += " style=\""
		for _, style := range c.Styles {
			result += style.Render()
		}
		result += "\""
	}

	result += ">"

	for _, child := range c.Children {
		result += child.Render()
	}

	result += "</" + c.Tag + ">"

	if c.LocalCSS != nil {
		result += "<style>" + c.LocalCSS.RenderGlobal() + "</style>"
	}

	return result
}

type Text string

func (t Text) Render() string {
	return string(t)
}
