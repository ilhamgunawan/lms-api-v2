CREATE TABLE public.mission_user (
	id uuid NOT NULL,
	mission_id uuid NOT NULL,
	user_id uuid NOT NULL,
	status varchar(20) NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	duration int NULL,
	CONSTRAINT mission_user_pk PRIMARY KEY (id),
	CONSTRAINT mission_user_fk FOREIGN KEY (user_id) REFERENCES public.user_account(id),
	CONSTRAINT mission_user_fk_1 FOREIGN KEY (mission_id) REFERENCES public.mission(id)
);

-- Column comments

COMMENT ON COLUMN public.mission_user.duration IS 'Hour';