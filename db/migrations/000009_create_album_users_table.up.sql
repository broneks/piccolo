create type album_user_role_enum as enum ('viewer', 'editor');

create table album_users (
  album_id uuid not null,
  user_id uuid not null,
  role album_user_role_enum default 'viewer',
  created_at timestamptz not null default now(),
  primary key (album_id, user_id),
  foreign key (album_id) references albums(id) on delete cascade,
  foreign key (user_id) references users(id) on delete cascade
);
