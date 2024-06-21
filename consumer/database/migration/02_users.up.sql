CREATE TABLE IF NOT EXISTS public.users (
    id UUID DEFAULT uuid_generate_v4() NOT NULL,
    "userName" VARCHAR COLLATE pg_catalog."default",
    "lastName" VARCHAR COLLATE pg_catalog."default",
    "nickName" VARCHAR COLLATE pg_catalog."default",
    email VARCHAR COLLATE pg_catalog."default",
    role VARCHAR COLLATE pg_catalog."default",
    sub VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);


INSERT INTO users (id,"userName", "lastName", "nickName", email, role, sub) 
VALUES('21c2b0d3-e045-48f9-98d9-39d8ffbf7597','','','test','test@gmail.com','user','1234567890');