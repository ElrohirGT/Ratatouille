CREATE ROLE postgres WITH LOGIN SUPERUSER;

CREATE DATABASE ratatouille;
\c ratatouille postgres;

CREATE USER root WITH PASSWORD 'root';
GRANT ALL PRIVILEGES ON DATABASE ratatouille TO root;

CREATE USER backend WITH PASSWORD 'backend';
REVOKE ALL ON DATABASE ratatouille FROM backend;
GRANT CONNECT ON DATABASE ratatouille TO backend;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO backend;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO backend;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, UPDATE ON SEQUENCES TO backend;

create table area(
	id serial primary key,
	nombre varchar(30) not null,
	admiteMoverMesas boolean not null,
	admiteFumadores boolean not null
);

create table puesto(
	id serial primary key,
	nombrePuesto varchar (30)
);

create table empleado (
	id serial primary key,
	nombre varchar(30),
	puesto integer references puesto (id) on delete cascade, 
	area integer references area(id) on delete cascade
);

create table mesa (
	id serial primary key,
	area int references area(id) on delete cascade,
	capacidad int
);

create table tipoUsuario (
	id serial primary key,
	nombre varchar(30)
);

create table usuario (
	nombre varchar (30) not null,
	contraseña varchar (30) not null,
	tipo int not null, 
	foreign key (tipo) references tipoUsuario(id) on delete cascade
);

create table cuenta (
	mesa int references mesa(id) on delete cascade not null,
	numCuenta serial primary key,
	estaCerrada boolean, 
	numPersonas int
);

create table cliente (
	id SERIAL primary key,
	nombre varchar (30),
	direccion varchar (30),
	nit varchar (30),

	unique(nombre, nit)
);

create table queja (
	cliente int references cliente(id) on delete cascade,
	gravedad int not null,
	motivo text not null,
	fecha timestamp not null,
	empleado integer references empleado (id) on delete cascade
);

create table estadosPedidos (
	id SERIAL primary key,
	nombre varchar(30)
);

create table tipoPago (
	id SERIAL primary key,
	nombre varchar (30)
);

create table itemMenuCategoria (
	id serial primary key,
	nombre varchar(30)
);

create table itemMenu (
	id SERIAL primary key,
	nombre varchar (30),
	descripcion text,
	precioUnitario money,
	categoria int references itemMenuCategoria(id) on delete cascade
);

create table pedidos (
	id serial primary key,
	fecha timestamp,
	estado int references estadosPedidos(id) on delete cascade,
	cantidad int,
	numCuenta int references cuenta(numCuenta) on delete cascade,
	item int references itemMenu(id) on delete cascade 
);

create table factura(
	numFactura serial primary key,
	fecha timestamp,
	cuenta int references cuenta(id) on delete cascade,
	cliente int references cliente(id) on delete cascade
);

create table pago(
	tipo int references tipoPago(id) on delete cascade,
	monto float,
	numFactura int references factura(numFactura) on delete cascade
);

create table encuesta (
	empleado int references empleado(id) on delete cascade,
	cliente int references cliente(id) on delete cascade,
	gradoAmabilidad int,
	gradoExactitud int,
	fecha timestamp
);

