package models

import (
	"fmt"
	"time"
)

type Eleicao struct {
	ID         string     `json:"id" firestore:"id,omitempty"`
	Nome       string     `json:"nome" firestore:"nome"`
	Descricao  string     `json:"descricao" firestore:"descricao"`
	DataInicio CustomTime `json:"dataInicio" firestore:"dataInicio"`
	DataFim    CustomTime `json:"dataFim" firestore:"dataFim"`
}

type CustomTime struct {
	time.Time
}

// UnmarshalJSON desserializa a data do JSON em formatos suportados
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

// MarshalJSON serializa a data em JSON no formato ISO
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Format("2006-01-02T15:04:05Z07:00") + `"`), nil
}

// ToValue converte o CustomTime em um valor que pode ser armazenado no Firestore
func (ct CustomTime) ToValue() (interface{}, error) {
	return ct.Time, nil
}

// FromValue lê um valor do Firestore para CustomTime
func (ct *CustomTime) FromValue(val interface{}) error {
	switch v := val.(type) {
	case time.Time:
		ct.Time = v
	default:
		return fmt.Errorf("erro ao converter valor para CustomTime: tipo inesperado %T", val)
	}
	return nil
}
