psql postgresql://postgres:postgresSuperPassword@postgres/postgres << EOF
  CREATE TABLE "public"."senders" (
    "id" bigserial NOT NULL,
    "name" text NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE("id")
  );

  CREATE TABLE "public"."messages" (
    id bigserial NOT NULL,
    timestamp timestamptz NOT NULL DEFAULT now(),
    message text NOT NULL,
    ip inet,
    priority integer,
    sender_id integer,
    PRIMARY KEY ("id"),
    FOREIGN KEY("sender_id") REFERENCES "public"."senders" ("id") ON UPDATE RESTRICT ON DELETE CASCADE,
    UNIQUE("id")
  );

EOF
