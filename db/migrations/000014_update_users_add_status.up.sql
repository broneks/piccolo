create type user_status_enum as enum ('pending', 'active', 'suspended');

alter table users
add column status user_status_enum default 'pending';
