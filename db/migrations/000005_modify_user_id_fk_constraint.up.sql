alter table photos
drop constraint photos_user_id_fkey;

alter table photos
add constraint photos_user_id_fkey
foreign key (user_id)
references users(id)
on update cascade
on delete cascade;
