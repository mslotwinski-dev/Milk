package examples

import (
	c "milk/core"
	"milk/core/web"
)

var counterCSS = c.DefineCSS(
	c.CSS("body", c.FontSize("16px"), c.FontFamily("Arial, sans-serif"), c.BackgroundColor("#e3e3e3"), c.Height("80vh")),
	c.CSS("main", c.Display("flex"), c.AlignItems("center"), c.JustifyContent("center"), c.FlexDirection("column"), c.Height("100%")),
	c.CSS(".header", c.FontSize("24px"), c.Color("#00647d"), c.BackgroundColor("white"), c.BorderRadius("5px"), c.Padding("10px")),
	c.CSS(".button-container", c.Color("white"), c.BackgroundColor("#333"), c.Padding("10px"), c.FontSize("18px"), c.BorderRadius("5px"), c.Display("flex"), c.AlignItems("center"), c.Gap("20px"), c.Margin("20px")),
	c.CSS("button", c.Padding("5px 20px"), c.FontSize("16px"), c.Border("none"), c.BorderRadius("5px"), c.BackgroundColor("#cc2222"), c.Color("white"), c.Cursor("pointer")),
)

var counterState = c.State{
	"count": 0,
	"name":  "Mateusz",
}

func MySuperCounterHeader() *c.Component {
	return c.H1(c.Div("Welcome to Milk, ", c.Span().GetState(counterState, "name"), "!").Class("header"))
}

func MySushiImage() *c.Component {
	return c.Img().Src("assets/sushi.png").Alt("Sushi").Width(200).Height(200)
}

func Counter() {

	page := c.Page(
		c.Main(
			MySuperCounterHeader(),

			c.Div(
				"How many sushi have you eaten?",
				c.Div().GetState(counterState, "count"),
				c.Button("Eat").Click(counterState.Mutate("count+", 1)),
			).Class("button-container"),

			MySushiImage(),
		).State(counterState),
	).CSS(counterCSS)

	page.SetTitle("Who doesn't love counting sushi?")
	page.SetFavicon("assets/sushi.png")
	page.SetAuthor("Giovani Giorgio (but everyone calls me Giorgio)")

	web.Run(page, 7250)
}
