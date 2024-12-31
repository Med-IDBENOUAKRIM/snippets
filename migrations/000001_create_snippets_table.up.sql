CREATE TABLE IF NOT EXISTS snippets (
  id bigserial PRIMARY KEY,
  title text NOT NULL,
  content text NOT NULL,
  created timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  expires timestamp(0) with time zone NOT NULL
);