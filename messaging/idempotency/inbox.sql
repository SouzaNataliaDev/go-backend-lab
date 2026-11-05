-- Laboratório PostgreSQL: deduplicação e efeito na MESMA transação.
-- Execute somente em banco de estudos. Não depende de cliente Go.
CREATE TABLE IF NOT EXISTS study_accounts (
    id TEXT PRIMARY KEY,
    balance_cents BIGINT NOT NULL DEFAULT 0
);
CREATE TABLE IF NOT EXISTS study_inbox (
    consumer_name TEXT NOT NULL,
    account_id TEXT NOT NULL REFERENCES study_accounts(id),
    event_id TEXT NOT NULL,
    cents BIGINT NOT NULL CHECK (cents > 0),
    processed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (consumer_name, account_id, event_id)
);
INSERT INTO study_accounts (id) VALUES ('company') ON CONFLICT DO NOTHING;

BEGIN;
-- Se outro consumidor já inseriu esta chave, accepted não retorna linhas.
-- O UPDATE só acontece quando o INSERT aceitou o evento.
WITH accepted AS (
    INSERT INTO study_inbox (consumer_name, account_id, event_id, cents)
    VALUES ('credit-consumer', 'company', 'evt-42', 5000)
    ON CONFLICT DO NOTHING
    RETURNING account_id, cents
)
UPDATE study_accounts AS account
SET balance_cents = account.balance_cents + accepted.cents
FROM accepted
WHERE account.id = accepted.account_id
RETURNING account.balance_cents;
COMMIT;
-- Repita o bloco BEGIN/COMMIT: o saldo não aumenta novamente.
-- A FK impede registro de evento para conta inexistente.
-- Em uma API real, ao detectar duplicata, compare o payload persistido com o
-- recebido e rejeite divergências. Este roteiro usa sempre o mesmo payload.
-- Confirme a mensagem no broker somente após o COMMIT bem-sucedido.
