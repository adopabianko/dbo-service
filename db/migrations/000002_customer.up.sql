do $$
begin
    if not exists (select 1 from pg_type where typname = 'gender') then
        create type gender as enum ('male', 'female');
    end if;
end
$$;

create table if not exists customer (
	id uuid primary key default gen_random_uuid(),
	created_at timestamp default current_timestamp not null,
  	created_by varchar(50) not null,
  	updated_at timestamp null,
  	updated_by varchar(50) null,
	deleted_at timestamp null,
  	deleted_by varchar(50) null,
	"name" varchar(150) not null,
	phone varchar(16) not null,
	email varchar(150) not null,
	gender gender not null,
	address text not null,

	constraint customer_email_unique unique (email),
    constraint customer_phone_unique unique (phone)
);