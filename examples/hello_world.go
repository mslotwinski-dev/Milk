package examples

import (
	c "milk/core"
	"milk/core/web"
)

func HelloWorld() {

	var css = c.DefineCSS(
		c.CSS("body", c.FontSize("16px"), c.FontFamily("Arial, sans-serif"), c.BackgroundColor("#e3e3e3"), c.Display("flex"), c.AlignItems("center"), c.JustifyContent("center"), c.FlexDirection("column"), c.Height("80vh")),
		c.CSS(".header", c.FontSize("24px"), c.Color("#00647d")),
	)

	var css2 = c.DefineCSS(c.CSS("", c.BackgroundColor("white"), c.BorderRadius("5px"), c.Padding("10px")))

	page := c.Page(
		c.H1(
			c.Div("Witaj w Milk!").CSS(css2).Class("header"),
		),
		c.Div("Abcdefghijklmnopqrstuvwxyz").Style(c.Color("red"), c.FontSize("20px")),
		c.P(
			"To jest akapit tekstu wygenerowany w Go przy użyciu Milk .",
		).Style(c.Color("white"), c.BackgroundColor("#333"), c.Padding("10px"), c.FontSize("18px"), c.BorderRadius("5px")),
		c.Img().Src("assets/sushi.png").Alt("Sushi").Width(200).Height(200),
	).CSS(css)

	page.SetTitle("Moja Pierwsza Strona w Milk")

	page.RawScript(`console.log("Hello from Milk!!!!")`, c.WithDefer())

	web.Run(page, 7250)

}
