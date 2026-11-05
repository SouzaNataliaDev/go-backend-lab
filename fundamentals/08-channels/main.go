package main

import "fmt"

// chan<- permite somente enviar; o produtor é dono do fechamento.
func produce(out chan<- int) {
	defer close(out)
	for _, id := range []int{1, 2, 3} {
		out <- id
	}
}

// <-chan permite somente receber.
func consume(in <-chan int) int {
	total := 0
	// range termina quando o canal é fechado e os valores já enviados são lidos.
	for id := range in {
		total += id
	}
	return total
}
func main() {
	jobs := make(chan int) // Sem buffer: envio aguarda um receptor disponível.
	go produce(jobs)
	fmt.Println("Soma recebida:", consume(jobs))

	queue := make(chan string, 2) // Até 2 valores pendentes sem receptor.
	queue <- "pedido A"
	queue <- "pedido B"
	// Um terceiro envio bloquearia aqui até alguém receber.
	close(queue)
	fmt.Println("Primeiro:", <-queue)
	fmt.Println("Segundo:", <-queue)
	value, ok := <-queue
	fmt.Printf("Fechado e vazio: valor=%q ok=%v\n", value, ok)
	// USO: pipelines e comunicação entre produtores e consumidores.
	// Enviar em canal fechado ou fechá-lo duas vezes causa panic.
	// Receber de canal fechado e vazio retorna zero imediatamente, com ok=false.
	// Enviar/receber em canal nil bloqueia; use context para interromper esperas reais.
}
