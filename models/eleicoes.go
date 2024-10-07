// models/eleicao.go
package models

type Eleicao struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}
