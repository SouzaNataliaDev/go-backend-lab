package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func order() {
	defer fmt.Println("defer registrado primeiro")
	defer fmt.Println("defer registrado depois")
	fmt.Println("corpo da função")
	// LIFO: o último defer registrado executa primeiro, na saída da função.
}
func arguments() {
	value := 1
	defer fmt.Println("argumento avaliado no registro:", value)
	defer func() { fmt.Println("closure lê na saída:", value) }()
	value = 2
}

// writeContent aceita um recurso já aberto e assume a responsabilidade de fechá-lo.
func writeContent(writer io.WriteCloser) (err error) {
	defer func() {
		closeErr := writer.Close()
		// Não esconda a falha original se a escrita já falhou.
		if err == nil {
			err = closeErr
		}
	}()
	_, err = io.Copy(writer, strings.NewReader("conteúdo"))
	return err
}

type demoWriter struct {
	bytes.Buffer
	closed bool
}

func (w *demoWriter) Close() error { w.closed = true; return nil }
func main() {
	order()
	arguments()
	writer := &demoWriter{}
	if err := writeContent(writer); err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Escrito:", writer.String(), "Fechado:", writer.closed)
	// USO: fechar arquivos, liberar mutex, executar cancel() e marcar wg.Done().
	// defer executa na saída da FUNÇÃO, não no final do bloco ou da iteração.
	// Para abrir muitos arquivos em um loop, extraia uma função por arquivo.
	// os.Exit/log.Fatal não executam os defers pendentes.
}
