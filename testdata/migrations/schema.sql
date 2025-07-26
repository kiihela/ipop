CREATE TABLE IF NOT EXISTS "schema_migration"
(
    "version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         TEXT PRIMARY KEY,
    "name"       TEXT     NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "teams"
(
    "id"         TEXT PRIMARY KEY,
    "name"       TEXT     NOT NULL,
    "leader" user NOT NULL,
    "members" [] user NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL
);
