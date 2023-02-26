CREATE TABLE public.topic (
	id uuid NOT NULL,
	course_id uuid NOT NULL,
	title varchar(255) NOT NULL,
	status varchar(20) NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	deleted timestamptz NULL,
	author_id uuid NOT NULL,
	CONSTRAINT topic_pk PRIMARY KEY (id),
	CONSTRAINT topic_fk FOREIGN KEY (course_id) REFERENCES public.course(id),
	CONSTRAINT topic_fk_1 FOREIGN KEY (author_id) REFERENCES public.user_account(id)
);