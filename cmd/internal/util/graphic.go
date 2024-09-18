package util

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GenerateGraphicTime(value plotter.XYs, filename string) {
	p := plot.New()

	p.Title.Text = "Tempos de Execução do Algoritmo"
	p.X.Label.Text = "Execução"
	p.Y.Label.Text = "Tempo (ms)"

	line, err := plotter.NewLine(value)
	if err != nil {
		panic(err)
	}

	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}
