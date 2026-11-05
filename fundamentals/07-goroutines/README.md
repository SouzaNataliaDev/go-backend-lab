# Goroutines

## Quando usar

Executar tarefas independentes e aguardar sua conclusão com WaitGroup.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/07-goroutines
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
Consultando: frete
Consultando: estoque
Consultando: preço
Todas as consultas concluídas
```

A ordem das consultas varia; a última linha aparece depois de todas elas.

## Erro comum

A função main pode terminar antes das goroutines. WaitGroup aguarda; time.Sleep não garante a conclusão.

## Pratique

Adicione uma quarta consulta e observe a variação na ordem de conclusão.
