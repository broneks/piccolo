alter table photos
add column filename text not null,
add column content_type varchar(50),
add column file_size int;
