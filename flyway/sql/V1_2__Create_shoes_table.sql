CREATE TABLE if not exists stockx.shoes
(
    productid bigint NOT NULL,
    productname text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT shoes_pkey PRIMARY KEY (productid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE stockx.shoes
    OWNER to postgres;