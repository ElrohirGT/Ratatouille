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
	nombre varchar(30) not null,
	puesto integer references puesto (id) on delete cascade not null, 
	area integer references area(id) on delete cascade --Puede ser null, ya que si no es mesero no està asignado a ningun area.
);

create table mesa (
	id serial primary key,
	area int references area(id) on delete cascade not null,
	capacidad int not null
);

create table tipoUsuario (
	id serial primary key,
	nombre varchar(30) not null
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
	estaCerrada boolean not null, 
	numPersonas int not null
);

create table cliente (
	id SERIAL primary key,
	nombre varchar (30) not null, 
	direccion varchar (30),
	nit varchar (30) not null,

	unique(nombre, nit)
);


create table estadosPedidos (
	id SERIAL primary key,
	nombre varchar(30) not null
);

create table tipoPago (
	id SERIAL primary key,
	nombre varchar (30) not null
);

create table itemMenuCategoria (
	id serial primary key,
	nombre varchar(30) not null
);

create table itemMenu (
	id SERIAL primary key,
	nombre varchar (30) not null,
	descripcion text not null,
	precioUnitario money not null,
	categoria int references itemMenuCategoria(id) on delete cascade not null
);

create table queja (
	cliente int references cliente(id) on delete cascade not null,
	gravedad int not null,
	motivo text not null,
	fecha timestamp not null,
	empleado integer references empleado (id) on delete cascade not null,
	item integer references itemMenu (id) on delete cascade not null
);

create table pedido (
	id serial primary key,
	fecha timestamp not null,
	estado int references estadosPedidos(id) on delete cascade not null,
	cantidad int not null,
	cuenta int references cuenta(numCuenta) on delete cascade not null,
	item int references itemMenu(id) on delete cascade not null 
);

create table factura(
	numFactura serial primary key,
	fecha timestamp not null,
	cuenta int references cuenta(numCuenta) on delete cascade not null,
	cliente int references cliente(id) on delete cascade not null
);

create table pago(
	tipo int references tipoPago(id) on delete cascade not null,
	monto float not null,
	numFactura int references factura(numFactura) on delete cascade not null
);

create table encuesta (
	empleado int references empleado(id) on delete cascade not null,
	cliente int references cliente(id) on delete cascade not null,
	gradoAmabilidad int not null,
	gradoExactitud int not null,
	fecha timestamp not null
);

