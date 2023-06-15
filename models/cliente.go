package models

type Cliente struct {
	Rg             string
	Nome           string
	ValorParaPagar float64
}

var ExemploDeEntradaParaClientes = []Cliente{
	{"12345", "Sofia", 20.0},
	{"54321", "Sami", 40.0},
	{"12345", "Blade", 20.0},
	{"12345", "Sofia", 60.0},
}
