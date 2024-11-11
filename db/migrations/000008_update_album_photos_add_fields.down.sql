alter table album_photos
drop column if exists user_id,
drop column if exists created_at;

alter table album_photos
drop constraint if exists album_photos_user_id_fkey;
