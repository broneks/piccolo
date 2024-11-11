alter table album_photos
add column user_id uuid not null,
add column created_at timestamptz not null default now();

alter table album_photos
add constraint album_photos_user_id_fkey
foreign key (user_id)
references users(id)
on update cascade
on delete cascade;
