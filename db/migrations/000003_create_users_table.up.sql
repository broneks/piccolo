create table if not exists users(
  id uuid primary key default gen_random_uuid(),
  username text unique,
  email text not null unique,
  hash text not null,
  hashed_at timestamptz not null default now(),
  last_login_at timestamptz not null default now(),
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create trigger users_update_updated_at
before update on users
for each row
execute function update_updated_at();
