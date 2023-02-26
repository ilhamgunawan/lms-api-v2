CREATE TABLE public.course (
	id uuid NOT NULL,
	title varchar(255) NOT NULL,
	slug varchar(50) NULL,
	description varchar(255) NULL,
	status varchar(20) NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	deleted timestamptz NULL,
	author_id uuid NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY (id),
	CONSTRAINT course_fk FOREIGN KEY (author_id) REFERENCES public.user_account(id)
);