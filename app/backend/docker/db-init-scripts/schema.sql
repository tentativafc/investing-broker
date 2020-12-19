CREATE TABLE IF NOT EXISTS public.user (
  id VARCHAR NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  user_password VARCHAR NOT NULL,
  user_sso_provider VARCHAR NOT NULL,
  CONSTRAINT user PRIMARY KEY (id)
);