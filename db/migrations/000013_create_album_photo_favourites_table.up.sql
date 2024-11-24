create table album_photo_favourites (
  album_id uuid not null,
  photo_id uuid not null,
  user_id uuid not null,
  created_at timestamptz not null default now(),
  primary key (album_id, photo_id, user_id),
  foreign key (album_id) references albums(id) on delete cascade,
  foreign key (photo_id) references photos(id) on delete cascade,
  foreign key (user_id) references users(id) on delete cascade
);

create index idx_album_photo_favourites_album_id_photo_id on album_photo_favourites (album_id, photo_id);

