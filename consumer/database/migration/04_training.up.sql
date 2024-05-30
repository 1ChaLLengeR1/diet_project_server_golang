CREATE TABLE IF NOT EXISTS public.training (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "postId" UUID,
    type VARCHAR COLLATE pg_catalog."default",
    time TIME,
    kcal double precision,
   "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT training_pkey PRIMARY KEY (id)
);