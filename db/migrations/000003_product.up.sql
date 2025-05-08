create table if not exists product (
	id uuid primary key default gen_random_uuid(),
	created_at timestamp default current_timestamp not null,
	created_by varchar(50) not null,
	updated_at timestamp null,
	updated_by varchar(50) null,
	sku varchar(50) not null,
	name varchar(150) not null,
	description text not null,
	price numeric(10,2) not null,

	constraint product_sku_unique unique (sku)
);

insert into product (created_by, sku, name, description, price) values
	('system', 'P001', 'Product 1', 'Description for Product 1', 10000),
	('system', 'P002', 'Product 2', 'Description for Product 2', 20000),
	('system', 'P003', 'Product 3', 'Description for Product 3', 30000),
	('system', 'P004', 'Product 4', 'Description for Product 4', 40000),
	('system', 'P005', 'Product 5', 'Description for Product 5', 50000);