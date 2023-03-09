create table business (
	id bigint auto_increment unique not null primary key,
    alias varchar(255),
    name varchar(255),
	image_url varchar(255),
    is_closed tinyint unsigned default 0,
    url varchar(255),
    review_count int,
    rating float,
    coordinate_id bigint unsigned null,
    transactions text,
    location_id bigint unsigned null,
    price varchar(255),
    phone varchar(255),
    display_phone varchar(255),
    distance float
);

drop table business;

create table coordinates (
	id bigint auto_increment unique not null primary key,
    latitude float,
    logitude float
);

create table categories (
	id bigint auto_increment unique not null primary key,
    alias varchar(255),
    title varchar(255),
    business_id bigint unsigned null
);

create table locations (
	id bigint auto_increment unique not null primary key,
    address1 varchar(255),
    address2 varchar(255),
    address3 varchar(255),
    city varchar(255),
    zip_code varchar(255),
    country varchar(255),
    state varchar(255),
    display_address text
);