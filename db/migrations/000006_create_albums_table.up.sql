create table if not exists albums(
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null,
  name text not null,
  description text,
  cover_photo_id uuid,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  foreign key (user_id)
    references users(id)
    on update cascade
    on delete cascade,
  foreign key (cover_photo_id)
    references photos(id)
    on update cascade
    on delete set null
);

create trigger albums_update_updated_at
before update on albums
for each row
execute function update_updated_at();
