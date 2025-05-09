create table if not exists "order" (
	id uuid primary key default gen_random_uuid(),
	created_at timestamp default current_timestamp not null,
	created_by varchar(50) not null,
	updated_at timestamp null,
	updated_by varchar(50) null,
	deleted_at timestamp null,
  	deleted_by varchar(50) null,
	ref varchar(50) not null,
	customer_id uuid not null,
	total_quantity integer not null,
	total_price numeric(10,2) not null,

	constraint order_ref_unique unique (ref),
	constraint order_customer_id_fkey foreign key (customer_id) references customer (id) on update cascade on delete cascade
);

create table if not exists order_item (
	id uuid primary key default gen_random_uuid(),
	order_id uuid not null,
	product_id uuid not null,
	quantity integer not null,
	subtotal numeric(10,2) not null,

	constraint order_product_id_fkey foreign key (product_id) references product (id) on update cascade on delete cascade
);