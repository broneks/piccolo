create table if not exists photos(
  id uuid primary key default gen_random_uuid(),
  location text not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create trigger photos_update_updated_at
before update on photos
for each row
execute function update_updated_at();
