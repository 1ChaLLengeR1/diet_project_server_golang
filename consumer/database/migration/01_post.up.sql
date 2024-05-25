CREATE TABLE IF NOT EXISTS public.post(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    "projectId" UUID,
    day integer,
    weight double precision,
    kcal integer,
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT post_pkey PRIMARY KEY (id)
);