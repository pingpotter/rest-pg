\connect rest_pg

SET TIME ZONE 'UTC';

DROP TABLE IF EXISTS public.account;
CREATE TABLE public.account (
    id character varying(100) NOT NULL,
    name character varying(250) NOT NULL,
    age numeric(3,0) NOT NULL,
    create_time timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.account OWNER TO postgres;