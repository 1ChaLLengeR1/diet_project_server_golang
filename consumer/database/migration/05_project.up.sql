CREATE TABLE IF NOT EXISTS public.project(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    title VARCHAR COLLATE pg_catalog."default",
    description VARCHAR COLLATE pg_catalog."default",
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT project_pkey PRIMARY KEY (id)
);