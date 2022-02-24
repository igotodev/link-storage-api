CREATE TABLE "users"
(
    "id"       serial            NOT NULL,
    "username" character varying NOT NULL UNIQUE,
    "password" character varying NOT NULL,
    "email"    character varying NOT NULL UNIQUE,
    "active"   boolean           NOT NULL
) WITH (
      OIDS= FALSE
    );