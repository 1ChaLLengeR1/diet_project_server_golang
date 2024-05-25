CREATE TABLE IF NOT EXISTS public.dictionary(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    key VARCHAR COLLATE pg_catalog."default",
    translation VARCHAR COLLATE pg_catalog."default",
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT dictionary_pkey PRIMARY KEY (id)
);


INSERT INTO public.dictionary (id, "key", translation, "createdUp", "updateUp") VALUES
('e09bd685-aaf8-4d65-bcdd-aadca85670bc', 'PL', 'polski', '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.dictionary (id, "key", translation, "createdUp", "updateUp") VALUES
('df88c32f-f71d-41bc-84ac-7cc36e37305f', 'EN', 'english', '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.dictionary (id, "key", translation, "createdUp", "updateUp") VALUES
('f63ce14b-af7a-4fc9-abc5-c68bc3254e2c', 'GER', 'deutsch', '2024-05-24 12:00:00', '2024-05-24 12:00:00');




