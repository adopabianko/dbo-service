create table if not exists "order" (
	id uuid primary key default gen_random_uuid(),
	created_at timestamp default current_timestamp not null,
  	created_by varchar(50) not null,
	ref varchar(50) not null,
	customer_id uuid not null,
	product_id uuid not null,
	quantity integer not null,
	total_price numeric(10,2) not null,

	constraint order_ref_unique unique (ref),
	constraint order_customer_id_fkey foreign key (customer_id) references customer (id) on update cascade on delete cascade,
	constraint order_product_id_fkey foreign key (product_id) references product (id) on update cascade on delete cascade
);