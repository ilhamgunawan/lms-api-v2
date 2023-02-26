CREATE TABLE public.mission (
	id uuid NOT NULL,
	topic_id uuid NOT NULL,
	type_id uuid NOT NULL,
	status varchar(20) NOT NULL,
	content_url text NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	deleted timestamptz NULL,
	author_id uuid NOT NULL,
	CONSTRAINT mission_pk PRIMARY KEY (id),
	CONSTRAINT mission_fk FOREIGN KEY (topic_id) REFERENCES public.topic(id),
	CONSTRAINT mission_fk_1 FOREIGN KEY (type_id) REFERENCES public.mission_type(id),
	CONSTRAINT mission_fk_2 FOREIGN KEY (author_id) REFERENCES public.user_account(id)
);