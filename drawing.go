package main

import (
	"image/color"

	"github.com/go-p5/p5"
)

func setup() {
	p5.Canvas(400, 400)
	p5.Background(color.Gray{Y: 255})
}

func drawTruthTable() {

}

func draw() {

	if len(vhdl_vars) == 0 {
		p5.Text("Please re-load the file to compile!", 30, float64(40)+40)
	}

	for i := 0; i < len(vhdl_vars); i++ {

		p5.TextSize(24)
		p5.Text(vhdl_vars[i].name, 10, float64(40*i)+40)
		p5.Text(vhdl_vars[i].value, 30, float64(40*i)+40)

		if vhdl_vars[i].value == "0" {

		}

		switch vhdl_vars[i].value {
		case "0":
			{
				p5.Stroke(color.RGBA{R: 255, A: 208})
				p5.Line(60, float64(40*i)+35, 290, float64(40*i)+35)

			}
			break

		case "1":
			{
				p5.Stroke(color.RGBA{G: 200, A: 208})
				p5.Line(60, float64(40*i)+25, 290, float64(40*i)+25)
			}
			break

		default:
			{
				p5.Stroke(color.RGBA{R: 255, G: 255, A: 208})
				p5.Line(60, float64(40*i)+25, 290, float64(40*i)+25)
			}
		}

		i := len(vhdl_vars)
		p5.Stroke(color.RGBA{B: 255, A: 208})
		p5.Line(60, float64(40*i)+25, 290, float64(40*i)+25)
		p5.Text("100 ns", 60, float64(40*i)+60)

	}

	// p5.StrokeWidth(2)
	// p5.Fill(color.RGBA{R: 255, A: 208})
	// p5.Ellipse(50, 50, 80, 80)

	// p5.Fill(color.RGBA{B: 255, A: 208})
	// p5.Quad(50, 50, 80, 50, 80, 120, 60, 120)

	// p5.Fill(color.RGBA{G: 255, A: 208})
	// p5.Rect(200, 200, 50, 100)

	// p5.Fill(color.RGBA{G: 255, A: 208})
	// p5.Triangle(100, 100, 120, 120, 80, 120)

	// p5.TextSize(24)
	// p5.Text("Hello, World!", 10, 300)

	// p5.Stroke(color.Black)
	// p5.StrokeWidth(5)
	// p5.Arc(300, 100, 80, 20, 0, 1.5*math.Pi)
}
