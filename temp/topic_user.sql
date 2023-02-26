CREATE TABLE public.topic_user (
	id uuid NOT NULL,
	topic_id uuid NOT NULL,
	user_id uuid NOT NULL,
	status varchar(20) NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NULL,
	duration int NULL,
	CONSTRAINT topic_user_pk PRIMARY KEY (id),
	CONSTRAINT topic_user_fk FOREIGN KEY (user_id) REFERENCES public.user_account(id),
	CONSTRAINT topic_user_fk_1 FOREIGN KEY (topic_id) REFERENCES public.topic(id)
);

-- Column comments

COMMENT ON COLUMN public.topic_user.duration IS 'Hour';