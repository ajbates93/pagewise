CREATE TABLE IF NOT EXISTS waitlist (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  email citext UNIQUE NOT NULL
)