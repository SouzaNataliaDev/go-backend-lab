package main

import (
	"errors"
	"testing"
)

func TestErrorChain(t *testing.T) {
	if _, err := loadOrder("99"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("causa perdida: %v", err)
	}
	var validation *ValidationError
	if _, err := loadOrder(""); !errors.As(err, &validation) || validation.Field != "id" {
		t.Fatalf("erro tipado perdido: %v", err)
	}
}
