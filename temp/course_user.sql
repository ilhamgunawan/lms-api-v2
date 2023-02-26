CREATE TABLE public.course_user (
	id uuid NOT NULL,
	course_id uuid NOT NULL,
	user_id uuid NOT NULL,
	status varchar(20) NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	duration int NULL,
	CONSTRAINT course_user_pk PRIMARY KEY (id),
	CONSTRAINT course_user_fk FOREIGN KEY (course_id) REFERENCES public.course(id),
	CONSTRAINT course_user_fk_1 FOREIGN KEY (user_id) REFERENCES public.user_account(id)
);

-- Column comments

COMMENT ON COLUMN public.course_user.duration IS 'Hour';