CREATE TABLE IF NOT EXISTS public.project_multi_language(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "idProject" UUID,
    "idLanguage" UUID,
    title VARCHAR COLLATE pg_catalog."default",
    description VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT project_multi_language_pkey PRIMARY KEY (id)
);