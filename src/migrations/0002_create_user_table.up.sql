CREATE TYPE IF NOT EXISTS AUTH_PROVIDERS AS ENUM ('google', 'apple');

CREATE TABLE IF NOT EXISTS public.user (
	pk: SERIAL PRIMARY KEY,
	user_id: INT4,
	provider: AUTH_PROVIDERS,
	UNIQUE (user_id, provider)
);


CREATE TABLE IF NOT EXISTS public.image (
	pk: SERIAL PRIMARY KEY,
	user: REFERENCES user.pk
	url: TEXT NOT NULL;
);

CREATE INDEX IF NOT EXISTS ON public.image (user);
