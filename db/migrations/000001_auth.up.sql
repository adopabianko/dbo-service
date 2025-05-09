create table if not exists auth (
	id uuid primary key default gen_random_uuid(),
	created_at timestamp default current_timestamp not null,
	created_by varchar(50) not null,
	updated_at timestamp null,
	updated_by varchar(50) null,
	email varchar(50) not null,
	password varchar(255) not null,

	constraint auth_email_unique unique (email)
);



-- email = 'admin@dbo.id'
-- password = 'dboadmin'
insert into auth (created_at, created_by, email, password) values (now(), 'system', 'admin@dbo.id', '$2a$10$UM6IEWjIUHgs1gcR3xDT0eOXttTdXPA0Om7tu12bc3loToGbfjyZe');