CREATE TABLE IF NOT EXISTS public.users (
    id UUID DEFAULT uuid_generate_v4() NOT NULL,
    "userName" VARCHAR COLLATE pg_catalog."default",
    "lastName" VARCHAR COLLATE pg_catalog."default",
    email VARCHAR COLLATE pg_catalog."default",
    role VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);