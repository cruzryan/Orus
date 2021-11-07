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
}
