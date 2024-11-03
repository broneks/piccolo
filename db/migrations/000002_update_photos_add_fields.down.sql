alter table photos
drop column if exists filename,
drop column if exists content_type,
drop column if exists file_size;
