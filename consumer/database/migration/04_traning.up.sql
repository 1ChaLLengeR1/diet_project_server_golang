CREATE TABLE IF NOT EXISTS public.traning (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    type VARCHAR COLLATE pg_catalog."default",
    "time" daterange,
    kcal double precision,
    "createdUp" date,
    "updateUp" date,
    CONSTRAINT traning_pkey PRIMARY KEY (id)
);