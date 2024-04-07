-- Inserts para la tabla 'area'
INSERT INTO area VALUES 
(default, 'El patio', true, true),
(default, 'Sala principal', false, false),
(default, 'Sala de niños', true, false),
(default, 'Bar', true, true),
(default, 'Sala de postres', true, false);


-- Inserts para la tabla 'puesto'
INSERT INTO puesto VALUES 
(default, 'Mesero'),
(default, 'Encargado'),
(default, 'Chef'),
(default, 'Bartender');

-- Inserts para la tabla 'empleado'
INSERT INTO empleado VALUES 
(default, 'Juan Lopez', 1, 1),
(default, 'Maria Melendez', 1, 1),
(default, 'Antonio Cruz', 1, 2),
(default, 'Luna Sancho', 1, 2),
(default, 'Jose Arquimedes', 1, 3),
(default, 'Maria Bartolome', 1, 3),
(default, 'Justo Guerra', 1, 4),
(default, 'Ramon de Leon', 1, 4),
(default, 'Michael Mejia', 1, 5),
(default, 'Alfonzo Monzon', 1, 5),
(default, 'Bartolome Chevez', 2, null),
(default, 'Loudes Montiel', 3, null),
(default, 'Jose Quinteros', 3, null),
(default, 'Gabriel Otero', 3, NULL),
(default, 'Diego Oliva', 3, null),
(default, 'Maria Camacho', 3, null),
(default, 'Pablo Rivas', 3, null),
(default, 'Iker Calderon', 4, null),
(default, 'Luis Domingo', 4, null),
(default, 'Gema Echeverria', 4, null),
(default, 'Alfonso Portillo', 4, null),
(default, 'Mariano Abad', 4, null),
(default, 'Emma Soler', 4, null);

-- Inserts para la tabla 'mesa'
INSERT INTO mesa VALUES 
(default, 1, 2),
(default, 1, 2),
(default, 1, 4),
(default, 1, 4),
(default, 1, 8),
(default, 1, 8),
(default, 1, 14),
(default, 2, 2),
(default, 2, 2),
(default, 2, 4),
(default, 2, 4),
(default, 2, 8),
(default, 2, 12),
(default, 3, 2),
(default, 3, 4),
(default, 3, 4),
(default, 3, 8),
(default, 3, 8),
(default, 3, 12),
(default, 3, 12),
(default, 4, 2),
(default, 4, 2),
(default, 4, 4),
(default, 4, 4),
(default, 4, 8),
(default, 4, 8),
(default, 4, 14),
(default, 4, 14),
(default, 5, 2),
(default, 5, 4),
(default, 5, 4),
(default, 5, 8),
(default, 5, 8),
(default, 5, 14),
(default, 5, 8),
(default, 5, 8);

-- Inserts para la tabla 'tipoUsuario'
INSERT INTO tipoUsuario VALUES 
(default, 'Mesero'),
(default, 'Chef'),
(default, 'Bartender'),
(default, 'Encargado');


-- Inserts para la tabla 'usuario'
INSERT INTO usuario VALUES 
('Juan Lopez', 'ratatouille1234', 1),
('Maria Melendez', 'ratatouille1234', 1),
('Gabriel Otero', 'ratatouille1234', 2),
('Jose Quinteros', 'ratatouille1234', 2),
('Alfonso Portillo', 'ratatouille1234', 3),
('Gema Echeverria', 'ratatouille1234', 3),
('Bartolome Chevez', 'ratatouille1234', 4);


-- Inserts para la tabla 'cuenta'
INSERT INTO cuenta VALUES 
(1, DEFAULT, false, 2),
(3, DEFAULT, false, 4),
(2, DEFAULT, true, 3),
(4, DEFAULT, true, 10),
(5, DEFAULT, false, 9),
(6, DEFAULT, false, 11),
(7, DEFAULT, false, 12),
(8, DEFAULT, true, 3),
(9, DEFAULT, false, 1),
(10, DEFAULT, false, 2),
(11, DEFAULT, true, 4),
(12, DEFAULT, true, 7),
(13, DEFAULT, true, 8),
(14, DEFAULT, true, 7),
(15, DEFAULT, false, 6);

-- Inserts para la tabla 'cliente'
INSERT INTO cliente VALUES 
(default, 'Carlos', 'Escuintla, Escuintla', '123456-7'),
(default, 'Ana', 'Ciudad de Guatemala', '987654-3'),
(default, 'Pedro', 'Chimaltenango', '456789-1'),
(default, 'Daniel', 'Ciudad de Guatemala', '451289-1'),
(default, 'Dylan', 'Ciudad de Guatemala', '123986-6'),
(default, 'Flavio', 'Escuintla, Escuintla', '829252-3'),
(default, 'Eddy', 'Chimaltenango', '123421-4'),
(default, 'Luis', 'Ciudad de Guatemala', '345623-2'),
(default, 'Erick', 'Escuintla, Escuintla', '912345-9'),
(default, 'Estuardo', 'Jalapa', '643134-3'),
(default, 'Brahian', 'Ciudad de Guatemala', '981256-7'),
(default, 'Matthew', 'Escuintla, Escuintla', '102342-8'),
(default, 'Douglas', 'Jutiapa', '114623-2'),
(default, 'Oscar', 'Escuintla, Escuintla', '102343-9'),
(default, 'Sebastian', 'Jutiapa', '125313-2'),
(default, 'Josue', 'Ciudad de Guatemala', '123635-5'),
(default, 'Bernardo', 'Escuintla, Escuintla', '431354-1'),
(default, 'Richie', 'Ciudad de Guatemala', '853251-1');


-- Inserts para la tabla 'estadosPedidos'
INSERT INTO estadosPedidos VALUES 
(default, 'En espera'),
(default, 'En preparación'),
(default, 'Cocinado'),
(default, 'Entregado'),
(default, 'Cancelado');


-- Inserts para la tabla 'tipoPago'
INSERT INTO tipoPago VALUES 
(default, 'Efectivo'),
(default, 'Tarjeta de Crédito'),
(default, 'Tarjeta de Débito'); 


-- Inserts para la tabla 'itemMenuCategoria'
INSERT INTO itemMenuCategoria VALUES 
(default, 'Platos'),
(default, 'Bebidas');


-- Inserts para la tabla 'itemMenu'
INSERT INTO itemMenu VALUES 
(default, 'Ensalada', 'Lechuga, cebolla, pollo, aderezo', 90.5, 1),
(default, 'Filete de Res', 'Filete de res a la parrilla con guarnicion', 60.5, 1),
(default, 'Fruit Punch', 'Naranja, jugo de manzana y sandia', 30, 2),
(default, 'Limonada', 'Limon, agua, azúcar y el toque especial de la casa.', 30, 2),
(default, 'Skinny Green', 'Espinaca, apio, piña, jugo de limón, pepino, jengibre y un toque de miel de abeja', 28, 2),
(default, 'Purple Fuel', 'Piña, banano, bluberries, manzana verde, avena y miel de abeja.', 30, 2),
(default, 'Té frío', 'Refrescante combinación de frutas y flores exoticas con un toque de naranja', 25, 2),
(default, 'Ceviche de Camarón', 'Ceviche de camarones acompañados de crackers de ajonjolí', 95, 1),
(default, 'Smoke tocino', 'Carne de res a la parrilla, con queso, tomates, lechuga y tocino', 68, 1);


-- Inserts para la tabla 'queja'
INSERT INTO queja VALUES 
(1, 3, 'Comida fría', now(), 12, 2),
(2, 5, 'Mal servicio', now(), 10, 1),
(3, 1, 'Plato equivocado', now(), 3, 3), --cambiar queja.empleado y .item a que puedan ser valores nulos 
(4, 1, 'Mesa en mal estado', now(), null, null),
(5, 3, 'Entrega de comida muy tardada', now(), null, null),
(6, 4, 'Me hablo mal el mesero', now(), 5, null),
(7, 2, 'El gerente no se preocupa por sus clientes', now(), 11, null),
(8, 3, 'Pedí bebida fría y me ha llegado caliente', now(), 19, 4);




-- Inserts para la tabla 'pedido'
INSERT INTO pedido VALUES 
(default, Now(), 1, 2, 1, 1),
(default, Now(), 4, 3, 2, 2),
(default, Now(), 3, 4, 3, 3),
(default, Now(), 4, 2, 4, 4),
(default, Now(), 4, 1, 4, 9),
(default, Now(), 4, 3, 8, 8),
(default, Now(), 4, 4, 11, 5),
(default, Now(), 4, 1, 12, 3),
(default, Now(), 4, 1, 13, 2),
(default, Now(), 2, 1, 6, 8),
(default, Now(), 4, 1, 12, 4),
(default, Now(), 4, 2, 14, 1);


-- Inserts para la tabla 'factura'
INSERT INTO factura VALUES 
(default, now(), 2, 1),
(default, now(), 4, 2),
(default, now(), 8, 3),
(default, now(), 11, 6),
(default, now(), 12, 4),
(default, now(), 13, 9),
(default, now(), 14, 2);


-- Inserts para la tabla 'pago'
INSERT INTO pago VALUES 
(1, 181.5, 1),
(2, 100, 2),
(3, 88, 2),
(1, 150, 3),
(2, 135, 3),
(1, 30, 4),
(2, 50, 4),
(1, 32, 4),
(3, 60, 5),
(3, 60.5, 6),
(2, 181, 7);

-- Inserts para la tabla 'encuesta'
INSERT INTO encuesta VALUES 
(9, 1, 4, 4, now()),
(10, 2, 3, 2, now()),
(3, 3, 4, 1, now()),
(5, 6, 1, 4, now()),
(2, 4, 5, 4, now()),
(1, 9, 5, 5, now()),
(11, 7, 1, 3, now());