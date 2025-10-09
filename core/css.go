package core

import (
	utility "milk/core/utility"
)

type Sheet []StyleGroup

type StyleGroup struct {
	Selector string
	Styles   []StyleByte
}

func CSS(selector string, styles ...StyleByte) StyleGroup {
	return StyleGroup{Selector: selector, Styles: styles}
}

func DefineCSS(styles ...StyleGroup) Sheet {
	return Sheet(styles)
}

func (c *Component) CSS(sheet Sheet) *Component {
	id := "milk-scope-" + utility.NewID(4)
	c.Class(id)

	scoped := make([]StyleGroup, 0, len(sheet))
	for _, s := range sheet {
		scoped = append(scoped, StyleGroup{
			Selector: "." + id + " " + s.Selector,
			Styles:   s.Styles,
		})
	}

	scopedSheet := DefineCSS(scoped...)
	c.LocalCSS = &scopedSheet

	return c
}

func (s Sheet) RenderGlobal() string {
	var result string
	for _, style := range s {
		result += style.Selector + " {"

		for _, st := range style.Styles {
			result += st.Render()
		}

		result += "}"
	}

	return result
}
