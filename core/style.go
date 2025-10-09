package core

type StyleByte struct {
	Property string
	Value    string
}

func GetStyle(property string, value string) StyleByte {
	return StyleByte{Property: property, Value: value}
}

func (c *Component) Style(styles ...StyleByte) *Component {
	if c.Styles == nil {
		c.Styles = make(map[string]string)
	}

	for _, s := range styles {
		c.Styles[s.Property] = s.Value
	}

	return c
}

func (c *Component) Class(class ...string) *Component {
	if c.Classes == nil {
		c.Classes = make([]string, 0)
	}

	c.Classes = append(c.Classes, class...)

	return c
}

func (c *Component) ID(id string) *Component {
	c.Id = id
	return c
}

// ---- Position ----
func Width(value string) StyleByte     { return GetStyle("width", value) }
func MinWidth(value string) StyleByte  { return GetStyle("min-width", value) }
func MaxWidth(value string) StyleByte  { return GetStyle("max-width", value) }
func Height(value string) StyleByte    { return GetStyle("height", value) }
func MinHeight(value string) StyleByte { return GetStyle("min-height", value) }
func MaxHeight(value string) StyleByte { return GetStyle("max-height", value) }
func Top(value string) StyleByte       { return GetStyle("top", value) }
func Right(value string) StyleByte     { return GetStyle("right", value) }
func Bottom(value string) StyleByte    { return GetStyle("bottom", value) }
func Left(value string) StyleByte      { return GetStyle("left", value) }
func Position(value string) StyleByte  { return GetStyle("position", value) }
func ZIndex(value string) StyleByte    { return GetStyle("z-index", value) }

// ---- Margins ----
func Margin(value string) StyleByte       { return GetStyle("margin", value) }
func MarginTop(value string) StyleByte    { return GetStyle("margin-top", value) }
func MarginRight(value string) StyleByte  { return GetStyle("margin-right", value) }
func MarginBottom(value string) StyleByte { return GetStyle("margin-bottom", value) }
func MarginLeft(value string) StyleByte   { return GetStyle("margin-left", value) }

func Padding(value string) StyleByte       { return GetStyle("padding", value) }
func PaddingTop(value string) StyleByte    { return GetStyle("padding-top", value) }
func PaddingRight(value string) StyleByte  { return GetStyle("padding-right", value) }
func PaddingBottom(value string) StyleByte { return GetStyle("padding-bottom", value) }
func PaddingLeft(value string) StyleByte   { return GetStyle("padding-left", value) }

// ---- Background ----
func Background(value string) StyleByte         { return GetStyle("background", value) }
func BackgroundColor(value string) StyleByte    { return GetStyle("background-color", value) }
func BackgroundImage(value string) StyleByte    { return GetStyle("background-image", value) }
func BackgroundPosition(value string) StyleByte { return GetStyle("background-position", value) }
func BackgroundSize(value string) StyleByte     { return GetStyle("background-size", value) }
func BackgroundRepeat(value string) StyleByte   { return GetStyle("background-repeat", value) }
func Color(value string) StyleByte              { return GetStyle("color", value) }
func Opacity(value string) StyleByte            { return GetStyle("opacity", value) }

// ---- Text and Font ----
func FontSize(value string) StyleByte       { return GetStyle("font-size", value) }
func FontWeight(value string) StyleByte     { return GetStyle("font-weight", value) }
func FontFamily(value string) StyleByte     { return GetStyle("font-family", value) }
func FontStyle(value string) StyleByte      { return GetStyle("font-style", value) }
func TextAlign(value string) StyleByte      { return GetStyle("text-align", value) }
func TextTransform(value string) StyleByte  { return GetStyle("text-transform", value) }
func TextDecoration(value string) StyleByte { return GetStyle("text-decoration", value) }
func LineHeight(value string) StyleByte     { return GetStyle("line-height", value) }
func LetterSpacing(value string) StyleByte  { return GetStyle("letter-spacing", value) }
func WordSpacing(value string) StyleByte    { return GetStyle("word-spacing", value) }
func WhiteSpace(value string) StyleByte     { return GetStyle("white-space", value) }

// ---- Border ----
func Border(value string) StyleByte       { return GetStyle("border", value) }
func BorderTop(value string) StyleByte    { return GetStyle("border-top", value) }
func BorderRight(value string) StyleByte  { return GetStyle("border-right", value) }
func BorderBottom(value string) StyleByte { return GetStyle("border-bottom", value) }
func BorderLeft(value string) StyleByte   { return GetStyle("border-left", value) }
func BorderColor(value string) StyleByte  { return GetStyle("border-color", value) }
func BorderWidth(value string) StyleByte  { return GetStyle("border-width", value) }
func BorderStyle(value string) StyleByte  { return GetStyle("border-style", value) }
func BorderRadius(value string) StyleByte { return GetStyle("border-radius", value) }

// ---- Flexbox ----
func Display(value string) StyleByte        { return GetStyle("display", value) }
func Flex(value string) StyleByte           { return GetStyle("flex", value) }
func FlexDirection(value string) StyleByte  { return GetStyle("flex-direction", value) }
func FlexWrap(value string) StyleByte       { return GetStyle("flex-wrap", value) }
func JustifyContent(value string) StyleByte { return GetStyle("justify-content", value) }
func AlignItems(value string) StyleByte     { return GetStyle("align-items", value) }
func AlignContent(value string) StyleByte   { return GetStyle("align-content", value) }
func Gap(value string) StyleByte            { return GetStyle("gap", value) }

// ---- Grid ----
func GridTemplateColumns(value string) StyleByte { return GetStyle("grid-template-columns", value) }
func GridTemplateRows(value string) StyleByte    { return GetStyle("grid-template-rows", value) }
func GridColumn(value string) StyleByte          { return GetStyle("grid-column", value) }
func GridRow(value string) StyleByte             { return GetStyle("grid-row", value) }
func GridGap(value string) StyleByte             { return GetStyle("grid-gap", value) }

// ---- Transformations and Animations ----
func Transform(value string) StyleByte  { return GetStyle("transform", value) }
func Transition(value string) StyleByte { return GetStyle("transition", value) }
func Animation(value string) StyleByte  { return GetStyle("animation", value) }
func Cursor(value string) StyleByte     { return GetStyle("cursor", value) }

// ---- Other ----
func Overflow(value string) StyleByte   { return GetStyle("overflow", value) }
func OverflowX(value string) StyleByte  { return GetStyle("overflow-x", value) }
func OverflowY(value string) StyleByte  { return GetStyle("overflow-y", value) }
func ClipPath(value string) StyleByte   { return GetStyle("clip-path", value) }
func Filter(value string) StyleByte     { return GetStyle("filter", value) }
func Visibility(value string) StyleByte { return GetStyle("visibility", value) }
