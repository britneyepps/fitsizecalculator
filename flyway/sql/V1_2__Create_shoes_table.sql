CREATE TABLE if not exists stockx.shoes
(
    productid serial,
    productname text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT shoes_pkey PRIMARY KEY (productid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE stockx.shoes
    OWNER to postgres;