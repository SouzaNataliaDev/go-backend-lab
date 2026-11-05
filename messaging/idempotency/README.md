# Consumidor idempotente

Idempotência significa que repetir uma operação com a mesma identidade não repete seu efeito. A mensagem ainda pode ser entregue e lida várias vezes.

```sh
go run ./messaging/idempotency
go test -race ./messaging/idempotency
```

Saída:

```text
Entrega 1: aplicado=true erro=<nil>
Entrega 2: aplicado=false erro=<nil>
Saldo em centavos: 5000
```

O [código comentado](main.go) usa mutex para atualizar saldo e registro de eventos juntos. Também detecta mesmo ID com valor diferente. O teste envia o mesmo evento por 50 goroutines e espera um único efeito.

## Como levar a ideia para banco e Kafka

O exemplo em memória só funciona no processo atual. Em várias réplicas, use armazenamento compartilhado com chave única. O [inbox.sql](inbox.sql) mostra uma transação PostgreSQL: inserir ID da mensagem e aplicar o efeito no mesmo commit. A deduplicação tem escopo por consumidor e conta; preserve esse escopo ao projetar a chave.

Fluxo: ler mensagem → transação de deduplicação + efeito → commit no banco → confirmar processamento no broker. No Kafka, o offset confirmado normalmente aponta para a próxima mensagem a ler naquela partição. Uma queda após o commit do banco e antes do commit do offset causa reentrega; a chave persistente evita repetir o efeito.

Em processamento concorrente de uma partição, confirmar um offset maior exige que todos os offsets anteriores relevantes estejam concluídos. Não avance sobre uma mensagem ainda em processamento.

Se o efeito for uma chamada a um provedor externo, use a idempotência oferecida por ele e coordene reprocessamento. Se também precisar publicar um evento, considere transactional outbox. Um simples `SELECT` antes do `INSERT` não impede duas réplicas de processarem juntas.

O SQL é material de estudo e não é executado pelos testes Go. A tabela de eventos cresce; retenção precisa considerar a janela de possíveis reentregas. Não prometa exactly-once global para banco + broker + serviço externo sem definir as garantias de cada etapa.
