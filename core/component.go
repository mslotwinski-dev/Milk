package core

type Renderable interface {
	Render() string
}

type Component struct {
	Tag      string
	Children []Renderable
	Styles   map[string]string
	Classes  []string
	Id       string
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

	if len(c.Styles) > 0 {
		result += " style=\""
		for k, v := range c.Styles {
			result += k + ":" + v + ";"
		}
		result += "\""
	}

	result += ">"

	for _, child := range c.Children {
		result += child.Render()
	}

	result += "</" + c.Tag + ">"
	return result
}

type Text string

func (t Text) Render() string {
	return string(t)
}
