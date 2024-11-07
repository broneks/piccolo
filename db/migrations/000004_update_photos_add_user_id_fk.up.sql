alter table photos
add column user_id uuid not null references users(id);
