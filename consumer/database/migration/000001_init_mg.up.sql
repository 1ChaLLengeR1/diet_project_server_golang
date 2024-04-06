CREATE TABLE IF NOT EXISTS public.users (
    id UUID DEFAULT uuid_generate_v4() NOT NULL,
    "userName" VARCHAR COLLATE pg_catalog."default",
    "lastName" VARCHAR COLLATE pg_catalog."default",
    email VARCHAR COLLATE pg_catalog."default",
    role VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.post(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    day integer,
    weight double precision,
    fitatu integer,
    "createdUp" date,
    "updateUp" date,
    CONSTRAINT post_pkey PRIMARY KEY (id)
);

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

CREATE TABLE IF NOT EXISTS public.images (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    "postId" UUID,
    path VARCHAR COLLATE pg_catalog."default",
    url VARCHAR COLLATE pg_catalog."default",
    "createdUp" date,
    "updateUp" date,
    CONSTRAINT images_pkey PRIMARY KEY (id)
);