# Idempotência, DLQ e offset

## “Como garantir que uma mensagem não seja consumida duas vezes?”

> Em sistemas com reentrega, não assumo que a mensagem será lida uma única vez. Torno o processamento idempotente: uso uma identidade estável, deduplicação persistente com chave única e transação para registrar a mensagem junto do efeito no banco. Confirmo o processamento depois do commit. Assim, uma reentrega não repete o efeito já aplicado.

Veja o [exemplo Go executável com 50 entregas concorrentes no teste](../../messaging/idempotency/README.md) e o [roteiro SQL](../../messaging/idempotency/inbox.sql).

Idempotência controla o efeito, não a quantidade de vezes que o broker entrega a mensagem. Se há um serviço externo envolvido, precisamos das garantias de idempotência dele e de coordenação do fluxo. As garantias de transação Kafka não tornam qualquer efeito em banco ou HTTP externo automaticamente exactly-once.

## “Para onde vão mensagens problemáticas?”

> Após classificar a falha e aplicar a política de tentativas, encaminho falhas permanentes ou esgotadas para uma DLQ, com contexto de diagnóstico, monitoramento e procedimento de reprocessamento. Corrijo a causa e reprocesso de modo controlado e idempotente.

DLQ é uma fila ou tópico separado. Guarde identidade, origem, quantidade de tentativas e erro útil, com cuidado para não vazar dados sensíveis. Uma DLQ sem alerta, responsável e plano de replay só acumula problemas. Falhas de infraestrutura não precisam ir imediatamente para DLQ; dependem da política do fluxo. Mover uma mensagem pode afetar a ordem de processamento de eventos do mesmo agregado.

DLQ é conteúdo conceitual deste guia; não há integração com broker neste projeto.

## “O que é offset?”

| Contexto | Significado | Exemplo |
| --- | --- | --- |
| Kafka | Posição de uma mensagem dentro de uma partição | Partição 2, offset 123 |
| Commit do consumidor Kafka | Posição persistida de onde retomar, normalmente a próxima mensagem | Após concluir 123, confirmar 124 |
| SQL/paginação | Quantas linhas pular antes de retornar o resultado | `LIMIT 20 OFFSET 40` |

No Kafka não existe um offset global para todo o tópico. Cada partição tem sua sequência; grupos mantêm seu progresso por partição. Com concorrência dentro da partição, não confirme além de mensagens anteriores ainda não concluídas.

Na paginação SQL, use `ORDER BY` determinístico. Offset alto pode ficar caro e mudanças nos dados entre páginas podem produzir lacunas ou repetição; paginação por cursor pode ser mais adequada.

Fonte: [Apache Kafka — design e garantias de entrega](https://kafka.apache.org/41/design/design/).
