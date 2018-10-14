CREATE TABLE if not exists stockx.truesize
(
    id bigint NOT NULL,
    productid integer,
    fitvalue integer NOT NULL,
    CONSTRAINT truesize_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE stockx.truesize
    OWNER to postgres;