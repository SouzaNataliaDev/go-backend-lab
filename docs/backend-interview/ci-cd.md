# CI e CD

## “O que é CI?”

> Integração contínua valida mudanças automaticamente ao integrá-las ao repositório. Configuro verificações de formatação, análise estática, testes e build para detectar problemas cedo. Os resultados ajudam a revisar a mudança; a política de merge define se uma falha bloqueia a integração.

O projeto já possui [GitHub Actions](../../.github/workflows/ci.yml), acionado em push e pull request:

| Etapa | Comando ou ação | O que verifica |
| --- | --- | --- |
| Checkout | actions/checkout | Obtém o código da revisão |
| Toolchain | actions/setup-go | Seleciona Go pelo go.mod |
| Formatação | `test -z "$(gofmt -l .)"` | Falha se houver arquivos fora do formato |
| Análise estática | `go vet ./...` | Detecta construções suspeitas |
| Testes | `go test -race -cover ./...` | Executa testes, detecta races observadas e mede cobertura |
| Build | `go build ./...` | Compila os pacotes |

Os fundamentos novos entram automaticamente nesses comandos por usarem `./...`. `make fundamentals` também executa cada exemplo, mas não é uma etapa separada da CI atual.

A CI reporta falhas. Exigir checks verdes para merge depende de configuração de regras da branch, que não foi alterada neste projeto. Cobertura não prova qualidade, e a CI não garante que uma alteração nunca falhará em produção.

## “E CD?”

- **Continuous delivery:** deixa uma versão pronta para entrega; produção pode depender de aprovação.
- **Continuous deployment:** versões aprovadas pelas verificações são implantadas automaticamente.

Em um serviço real, acrescentaria artefato versionado, promoção entre ambientes, testes de contrato/integração, gestão de segredos, estratégia gradual e verificação após deploy. O rollback deve considerar migrações e compatibilidade dos dados.

Este laboratório tem CI; não possui deploy automatizado para produção.

Fonte: [GitHub — Building and testing Go](https://docs.github.com/en/actions/tutorials/build-and-test-code/go).
