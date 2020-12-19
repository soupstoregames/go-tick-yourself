create table characters
(
	id bigserial not null,
	balance bigint default 0 not null,
	reputation smallint default 0 not null
);
