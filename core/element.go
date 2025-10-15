package core

import "fmt"

func Element(tag string, children ...interface{}) *Component {
	renderables := make([]Renderable, 0, len(children))

	for _, child := range children {
		switch v := child.(type) {
		case string:
			renderables = append(renderables, Text(v))
		case int, float64, bool:
			renderables = append(renderables, Text(fmt.Sprint(v)))
		case Renderable:
			renderables = append(renderables, v)
		}
	}

	return &Component{Tag: tag, Children: renderables}
}

func A(children ...interface{}) *Component          { return Element("a", children...) }
func Abbr(children ...interface{}) *Component       { return Element("abbr", children...) }
func Address(children ...interface{}) *Component    { return Element("address", children...) }
func Area(children ...interface{}) *Component       { return Element("area", children...) }
func Article(children ...interface{}) *Component    { return Element("article", children...) }
func Aside(children ...interface{}) *Component      { return Element("aside", children...) }
func Audio(children ...interface{}) *Component      { return Element("audio", children...) }
func B(children ...interface{}) *Component          { return Element("b", children...) }
func Base(children ...interface{}) *Component       { return Element("base", children...) }
func Bdi(children ...interface{}) *Component        { return Element("bdi", children...) }
func Bdo(children ...interface{}) *Component        { return Element("bdo", children...) }
func Blockquote(children ...interface{}) *Component { return Element("blockquote", children...) }
func Br(children ...interface{}) *Component         { return Element("br", children...) }
func Button(children ...interface{}) *Component     { return Element("button", children...) }
func Canvas(children ...interface{}) *Component     { return Element("canvas", children...) }
func Caption(children ...interface{}) *Component    { return Element("caption", children...) }
func Cite(children ...interface{}) *Component       { return Element("cite", children...) }
func Code(children ...interface{}) *Component       { return Element("code", children...) }
func Col(children ...interface{}) *Component        { return Element("col", children...) }
func Colgroup(children ...interface{}) *Component   { return Element("colgroup", children...) }
func Datalist(children ...interface{}) *Component   { return Element("datalist", children...) }
func Dd(children ...interface{}) *Component         { return Element("dd", children...) }
func Del(children ...interface{}) *Component        { return Element("del", children...) }
func Details(children ...interface{}) *Component    { return Element("details", children...) }
func Dfn(children ...interface{}) *Component        { return Element("dfn", children...) }
func Dialog(children ...interface{}) *Component     { return Element("dialog", children...) }
func Div(children ...interface{}) *Component        { return Element("div", children...) }
func Dl(children ...interface{}) *Component         { return Element("dl", children...) }
func Dt(children ...interface{}) *Component         { return Element("dt", children...) }
func Em(children ...interface{}) *Component         { return Element("em", children...) }
func Embed(children ...interface{}) *Component      { return Element("embed", children...) }
func Fieldset(children ...interface{}) *Component   { return Element("fieldset", children...) }
func Figcaption(children ...interface{}) *Component { return Element("figcaption", children...) }
func Figure(children ...interface{}) *Component     { return Element("figure", children...) }
func Footer(children ...interface{}) *Component     { return Element("footer", children...) }
func Form(children ...interface{}) *Component       { return Element("form", children...) }
func H1(children ...interface{}) *Component         { return Element("h1", children...) }
func H2(children ...interface{}) *Component         { return Element("h2", children...) }
func H3(children ...interface{}) *Component         { return Element("h3", children...) }
func H4(children ...interface{}) *Component         { return Element("h4", children...) }
func H5(children ...interface{}) *Component         { return Element("h5", children...) }
func H6(children ...interface{}) *Component         { return Element("h6", children...) }
func Header(children ...interface{}) *Component     { return Element("header", children...) }
func Hr(children ...interface{}) *Component         { return Element("hr", children...) }
func I(children ...interface{}) *Component          { return Element("i", children...) }
func Iframe(children ...interface{}) *Component     { return Element("iframe", children...) }
func Img(children ...interface{}) *Component        { return Element("img", children...) }
func Input(children ...interface{}) *Component      { return Element("input", children...) }
func Ins(children ...interface{}) *Component        { return Element("ins", children...) }
func Kbd(children ...interface{}) *Component        { return Element("kbd", children...) }
func Label(children ...interface{}) *Component      { return Element("label", children...) }
func Legend(children ...interface{}) *Component     { return Element("legend", children...) }
func Li(children ...interface{}) *Component         { return Element("li", children...) }
func Link(children ...interface{}) *Component       { return Element("link", children...) }
func Main(children ...interface{}) *Component       { return Element("main", children...) }
func Map(children ...interface{}) *Component        { return Element("map", children...) }
func Mark(children ...interface{}) *Component       { return Element("mark", children...) }
func Meta(children ...interface{}) *Component       { return Element("meta", children...) }
func Meter(children ...interface{}) *Component      { return Element("meter", children...) }
func Nav(children ...interface{}) *Component        { return Element("nav", children...) }
func Noscript(children ...interface{}) *Component   { return Element("noscript", children...) }
func Object(children ...interface{}) *Component     { return Element("object", children...) }
func Ol(children ...interface{}) *Component         { return Element("ol", children...) }
func Optgroup(children ...interface{}) *Component   { return Element("optgroup", children...) }
func Option(children ...interface{}) *Component     { return Element("option", children...) }
func Output(children ...interface{}) *Component     { return Element("output", children...) }
func P(children ...interface{}) *Component          { return Element("p", children...) }
func Param(children ...interface{}) *Component      { return Element("param", children...) }
func Picture(children ...interface{}) *Component    { return Element("picture", children...) }
func Pre(children ...interface{}) *Component        { return Element("pre", children...) }
func Progress(children ...interface{}) *Component   { return Element("progress", children...) }
func Q(children ...interface{}) *Component          { return Element("q", children...) }
func Rp(children ...interface{}) *Component         { return Element("rp", children...) }
func Rt(children ...interface{}) *Component         { return Element("rt", children...) }
func Ruby(children ...interface{}) *Component       { return Element("ruby", children...) }
func S(children ...interface{}) *Component          { return Element("s", children...) }
func Samp(children ...interface{}) *Component       { return Element("samp", children...) }
func Section(children ...interface{}) *Component    { return Element("section", children...) }
func Select(children ...interface{}) *Component     { return Element("select", children...) }
func Small(children ...interface{}) *Component      { return Element("small", children...) }
func Source(children ...interface{}) *Component     { return Element("source", children...) }
func Span(children ...interface{}) *Component       { return Element("span", children...) }
func Strong(children ...interface{}) *Component     { return Element("strong", children...) }
func Sub(children ...interface{}) *Component        { return Element("sub", children...) }
func Summary(children ...interface{}) *Component    { return Element("summary", children...) }
func Sup(children ...interface{}) *Component        { return Element("sup", children...) }
func Svg(children ...interface{}) *Component        { return Element("svg", children...) }
func Table(children ...interface{}) *Component      { return Element("table", children...) }
func Tbody(children ...interface{}) *Component      { return Element("tbody", children...) }
func Td(children ...interface{}) *Component         { return Element("td", children...) }
func Template(children ...interface{}) *Component   { return Element("template", children...) }
func Textarea(children ...interface{}) *Component   { return Element("textarea", children...) }
func Tfoot(children ...interface{}) *Component      { return Element("tfoot", children...) }
func Th(children ...interface{}) *Component         { return Element("th", children...) }
func Thead(children ...interface{}) *Component      { return Element("thead", children...) }
func Time(children ...interface{}) *Component       { return Element("time", children...) }
func Title(children ...interface{}) *Component      { return Element("title", children...) }
func Tr(children ...interface{}) *Component         { return Element("tr", children...) }
func Track(children ...interface{}) *Component      { return Element("track", children...) }
func U(children ...interface{}) *Component          { return Element("u", children...) }
func Ul(children ...interface{}) *Component         { return Element("ul", children...) }
func Var(children ...interface{}) *Component        { return Element("var", children...) }
func Video(children ...interface{}) *Component      { return Element("video", children...) }
func Wbr(children ...interface{}) *Component        { return Element("wbr", children...) }
