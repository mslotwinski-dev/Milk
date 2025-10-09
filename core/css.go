package core

type Sheet []StyleGroup

type StyleGroup struct {
	Selector string
	Styles   map[string]string
}

func CSS(selector string, styles ...StyleByte) StyleGroup {
	css := StyleGroup{Selector: selector, Styles: make(map[string]string)}
	for _, style := range styles {
		css.Styles[style.Property] = style.Value
	}
	return css
}

func DefineCSS(styles ...StyleGroup) Sheet {
	return Sheet(styles)
}

// func (s Sheet) RenderGlobal() string {
// 	var result string
// 	for _, style := range s {
// 		result += style.Render()
// 	}
// 	return result
// }
