package core

type StyleByte struct {
	property string
	value    string
}

func GetStyle(property string, value string) StyleByte {
	return StyleByte{property: property, value: value}
}

func (c *Component) Style(styles ...StyleByte) *Component {
	if c.Styles == nil {
		c.Styles = make(map[string]string)
	}

	for _, s := range styles {
		c.Styles[s.property] = s.value
	}

	return c
}

func Width(value string) StyleByte           { return GetStyle("width", value) }
func Height(value string) StyleByte          { return GetStyle("height", value) }
func BackgroundColor(value string) StyleByte { return GetStyle("background-color", value) }
func Color(value string) StyleByte           { return GetStyle("color", value) }
func FontSize(value string) StyleByte        { return GetStyle("font-size", value) }
