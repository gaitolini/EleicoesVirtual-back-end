package models

import (
	"time"
)

// Função auxiliar para tratar parsing customizado da data
type Eleicao struct {
	ID         string     `json:"id"`
	Nome       string     `json:"nome"`
	Descricao  string     `json:"descricao"`
	DataInicio CustomTime `json:"dataInicio"`
	DataFim    CustomTime `json:"dataFim"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Remover as aspas da string JSON
	s := string(b)
	s = s[1 : len(s)-1]

	// Primeiro, tentar o formato ISO completo
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if err != nil {
		// Caso falhe, tentar apenas a data
		t, err = time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
	}

	ct.Time = t
	return nil
}
