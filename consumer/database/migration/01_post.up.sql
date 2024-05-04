CREATE TABLE IF NOT EXISTS public.post(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    "projectId" UUID,
    day integer,
    weight double precision,
    fitatu integer,
    "createdUp" date,
    "updateUp" date,
    CONSTRAINT post_pkey PRIMARY KEY (id)
);