-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS consolidations (
    slot_number INT NOT NULL,
    slot_index INT NOT NULL,
    slot_root bytea NOT NULL,
    orphaned bool NOT NULL DEFAULT FALSE,
    source_index BIGINT NOT NULL,
    target_index BIGINT NOT NULL,
    epoch BIGINT NOT NULL,
    CONSTRAINT consolidation_pkey PRIMARY KEY (slot_index, slot_root)
);

CREATE INDEX IF NOT EXISTS "consolidations_source_idx"
    ON public."consolidations"
    ("source_index" ASC NULLS FIRST);

CREATE INDEX IF NOT EXISTS "consolidations_target_idx"
    ON public."consolidations"
    ("target_index" ASC NULLS FIRST);

CREATE INDEX IF NOT EXISTS "consolidations_epoch_idx"
    ON public."consolidations"
    ("epoch" ASC NULLS FIRST);

CREATE INDEX IF NOT EXISTS "consolidations_slot_idx"
    ON public."consolidations"
    ("slot_number" ASC NULLS FIRST);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'NOT SUPPORTED';
-- +goose StatementEnd
