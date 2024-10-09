package models

import (
	"time"
)

type Eleicao struct {
	ID         string    `json:"id"`
	Nome       string    `json:"nome"`
	Descricao  string    `json:"descricao"`
	DataInicio time.Time `json:"dataInicio"`
	DataFim    time.Time `json:"dataFim"`
}
