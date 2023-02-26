CREATE TABLE public.mission_type (
	id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	description varchar(255) NULL,
	CONSTRAINT mission_type_pk PRIMARY KEY (id)
);