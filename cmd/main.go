package main

import (
	"fmt"

	"github.com/jairoevaristo/assinatura-digital/cmd/internal"
	"github.com/jairoevaristo/assinatura-digital/cmd/internal/test"
)

func main() {
	internal.Init()

	fmt.Println("[TEMPO MEDIO]:")
	test.TestTime()

	fmt.Println("\n[TEMPO MEDIO EXECUTADO 10 VEZES]:")
	test.TestOverflow()
}
