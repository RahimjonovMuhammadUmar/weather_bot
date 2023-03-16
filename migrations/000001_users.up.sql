CREATE TABLE IF NOT EXISTS "users"(
    "tg_id" BIGINT PRIMARY KEY,
    "first_name" VARCHAR(100),
    "last_name" VARCHAR(100),
    "created_at " TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);