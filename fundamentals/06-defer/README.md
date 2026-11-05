# defer

## Quando usar

Agendar limpeza e entender ordem de execução e captura de argumentos.

## Executar

Na raiz do repositório:

```sh
go run ./fundamentals/06-defer
```

Abra [main.go](main.go) para ler as explicações `//` junto de cada operação.

## Saída de referência

```text
corpo da função
defer registrado depois
defer registrado primeiro
closure lê na saída: 2
argumento avaliado no registro: 1
Escrito: conteúdo Fechado: true
```

A saída corresponde à simulação apresentada no código.

## Erro comum

Defer não executa ao sair de um simples bloco. Adiá-lo dentro de um loop grande pode manter muitos recursos abertos.

## Pratique

Altere o writer para falhar no Close e confira o erro retornado.
