package examples

import (
	"fmt"
	c "milk/core"
	"milk/core/web"
)

func HelloWorld() {

	css := c.DefineCSS(
		c.CSS("body", c.FontSize("16px"), c.FontFamily("Arial, sans-serif"), c.BackgroundColor("red")),
		c.CSS(".header", c.FontSize("24px"), c.Color("#00647d")))

	fmt.Println(css)

	page := c.Page(
		c.H1(
			c.Div("Witaj w Milk!").Class("header"),
		),
		c.P(
			"To jest akapit tekstu wygenerowany w Go przy użyciu Milk.",
		).Style(c.Color("white"), c.BackgroundColor("black"), c.FontSize("18px")),
	).CSS(css)

	page.SetTitle("Moja Pierwsza Strona w Milk")

	web.Run(page, 7250)

}
