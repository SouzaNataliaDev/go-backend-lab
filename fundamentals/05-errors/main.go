package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("pedido não encontrado")

// Erro tipado permite transportar informação estruturada.
type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "campo inválido: " + e.Field }

func findOrder(id string) (string, error) {
	if id == "" {
		return "", &ValidationError{Field: "id"}
	}
	if id != "42" {
		return "", ErrNotFound
	}
	return "pedido 42", nil
}
func loadOrder(id string) (string, error) {
	order, err := findOrder(id)
	if err != nil {
		return "", fmt.Errorf("carregar pedido %q: %w", id, err)
	}
	return order, nil
}
func main() {
	for _, id := range []string{"42", "99", ""} {
		order, err := loadOrder(id)
		if err == nil {
			fmt.Println("OK:", order)
			continue
		}
		var validation *ValidationError
		switch {
		case errors.Is(err, ErrNotFound): // percorre a cadeia preservada por %w.
			fmt.Println("Resposta 404: pedido não encontrado")
		case errors.As(err, &validation): // encontra um erro pelo tipo.
			fmt.Println("Resposta 400: corrigir", validation.Field)
		default:
			fmt.Println("Resposta 500: erro interno")
		}
	}
	// USO: transformar erros de domínio em respostas de API.
	// Registre detalhes internos na camada apropriada; evite expor esses detalhes.
	// Erros esperados são retornados. panic não substitui esse fluxo.
}
