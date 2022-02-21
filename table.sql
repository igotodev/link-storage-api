CREATE TABLE "links"
(
    "id"       serial            NOT NULL,
    "name"     character varying NOT NULL UNIQUE,
    "category" character varying NOT NULL,
    "url"      character varying NOT NULL UNIQUE,
    "date"     timestamp         NOT NULL
) WITH (
      OIDS= FALSE
    );