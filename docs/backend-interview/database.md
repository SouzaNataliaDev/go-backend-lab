# Índices e EXPLAIN

## “O que é um índice?”

> É uma estrutura auxiliar que ajuda o banco a localizar dados sem percorrer necessariamente toda a tabela. Escolho índices a partir das consultas, filtros, joins e ordenações mais relevantes. Eles têm custo de espaço e manutenção em escritas, então verifico o plano e meço antes e depois.

## “Para que serve EXPLAIN?”

> Mostra o plano escolhido pelo otimizador: como acessa tabelas, aplica filtros e faz joins/ordenação. No PostgreSQL, EXPLAIN ANALYZE executa a consulta e acrescenta medições reais. Comparo estimativas com linhas reais, loops, tempo e buffers, considerando o volume de dados.

Exemplo PostgreSQL para um banco de estudos com tabela `orders`:

```sql
-- Consulta que queremos investigar.
EXPLAIN SELECT id, created_at
FROM orders
WHERE customer_id = 42
ORDER BY created_at DESC, id DESC
LIMIT 20;

-- Hipótese de índice alinhado ao filtro e à ordem pedida.
CREATE INDEX orders_customer_created_id_idx
ON orders (customer_id, created_at DESC, id DESC);

-- Depois, medir a consulta de leitura com dados representativos.
EXPLAIN (ANALYZE, BUFFERS)
SELECT id, created_at
FROM orders
WHERE customer_id = 42
ORDER BY created_at DESC, id DESC
LIMIT 20;
```

Esse SQL é um trecho didático: não cria a tabela nem é executado pela CI. Não há plano real ou ganho de desempenho medido neste material.

| O que observar | Interpretação |
| --- | --- |
| Seq Scan | Leitura sequencial; pode ser boa em tabela pequena ou consulta que retorna muito |
| Index Scan / Bitmap Scan | Uso de índice; ainda pode ter alto custo dependendo das linhas acessadas |
| estimated rows × actual rows | Grande diferença pode sugerir estatísticas pouco representativas |
| loops | Multiplica trabalho; ajuda a entender custo de nós repetidos |
| Buffers | Ajuda a distinguir acesso em cache e leitura de blocos |
| Sort | Verificar volume e eventual uso de disco |

Custo estimado do plano não é tempo em milissegundos. Criar um índice não garante que o banco o usará. Funções sobre colunas, seletividade, estatísticas e ordem das colunas influenciam a escolha.

`EXPLAIN ANALYZE` executa a instrução. Em comandos de escrita isso pode alterar dados; estude esses casos em ambiente apropriado. Criação de índice em produção também precisa de planejamento operacional.

Fonte: [PostgreSQL — Using EXPLAIN](https://www.postgresql.org/docs/current/using-explain.html).
