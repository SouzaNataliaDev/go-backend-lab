package main

import (
	"fmt"
	"sort"
)

func main() {
	// make inicializa um map pronto para escrita. var m map[string]int seria nil.
	stock := make(map[string]int)
	stock["PLA"] = 10
	stock["PETG"] = 0

	// O segundo retorno distingue chave ausente de um valor zero existente.
	quantity, exists := stock["PETG"]
	fmt.Println("PETG:", quantity, "existe:", exists)
	quantity, exists = stock["ABS"]
	fmt.Println("ABS:", quantity, "existe:", exists)

	delete(stock, "PETG") // Excluir chave ausente também é permitido.
	stock["TPU"] = 5
	// A ordem de range em map não é garantida. Ordene as chaves para exibir.
	keys := make([]string, 0, len(stock))
	for name := range stock {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		fmt.Printf("%s=%d\n", name, stock[name])
	}

	// Atribuir um map não copia seus elementos: ambos acessam o mesmo map.
	alias := stock
	alias["PLA"] = 9
	fmt.Println("Original após alias:", stock["PLA"])
	// USO: índices e contadores. Para escrita concorrente, use sincronização.
	// Ler map nil é permitido; escrever nele causa panic.
}
