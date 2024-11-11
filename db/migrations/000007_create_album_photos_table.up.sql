create table album_photos (
  album_id uuid not null,
  photo_id uuid not null,
  primary key (album_id, photo_id),
  foreign key (album_id) references albums(id) on delete cascade,
  foreign key (photo_id) references photos(id) on delete cascade
);
