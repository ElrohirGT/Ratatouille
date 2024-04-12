-- 0. Saber el rol de la persona a partir del usuario y contraseña ingresados por el usuario.

-- name: GetRoleOfUser :many
SELECT 
	tu.nombre AS tipo_usuario
FROM 
	usuario u
JOIN 
	tipoUsuario tu ON u.tipo = tu.id
WHERE 
	u.nombre = $1 AND u.contraseña = $2;


-- 1. Reporte de los platos más pedidos por los clientes en un rango de fechas solicitadas al usuario.

-- name: GetMostFamousDishesBetween :many
SELECT p.item, i.nombre, i.descripcion, COUNT(p) as count
FROM pedido p
	RIGHT JOIN itemMenu i on p.item = i.id
WHERE p.fecha BETWEEN $1 AND $2 
GROUP BY p.item, i.nombre, i.descripcion
ORDER BY count DESC;

-- 2. Horario en el que se ingresan más pedidos entre un rango de fechas solicitadas al usuario.

-- name: GetRushHourBetween :one
SELECT EXTRACT(hour from p.fecha) as horario, COUNT(*) as count
FROM pedido p
WHERE p.fecha BETWEEN $1 AND $2
GROUP BY horario
ORDER BY count DESC
LIMIT 1;

-- 3. Promedio de tiempo en que se tardan los clientes en comer, agrupando la cantidad de
-- personas comiendo, por ejemplo: 2 personas: 1 hora 10 minutos, 3 personas: 1 hora 15
-- minutos, etc. entre un rango de fechas solicitadas al usuario.

-- name: GetAverageTimeToEatPerClientQuantity :many
SELECT
	c.numPersonas,
	AVERAGE(MAX(p.fecha) OVER (PARTITION BY c.numCuenta) - MIN(p.fecha) OVER (PARTITION by c.numCuenta)) as timeToEat
FROM cuenta c
	LEFT JOIN pedido p ON p.numCuenta = c.numCuenta
WHERE c.estaCerrada AND p.fecha BETWEEN $1 AND $2
GROUP BY numPersonas;

-- 4. Reporte de las quejas agrupadas por persona para un rango de fechas solicitadas al usuario.

-- name: GetComplaintsForEmployeeBetween :many
SELECT *
FROM queja 
WHERE empleado = $1 AND fecha BETWEEN $2 AND $3;

-- 5. Reporte de las quejas agrupadas por plato para un rango de fechas solicitadas al usuario.

-- name: GetComplaintsForDishBetween :many
SELECT *
FROM queja
WHERE item = $1 AND fecha $2 AND $3;


-- 6. Reporte de eficiencia de meseros mostrando los resultados de las encuestas, agrupado
-- por personas y por mes para los últimos 6 meses.

-- | Persona | Mes | avgAmabilidad | avgExactirud |

-- name: GetEfficiencyReport :many
SELECT 
	en.empleado,
	EXTRACT(MONTH from en.fecha) as mes,
	AVG(e.gradoAmabilidad),
	AVG(e.gradoExactitud) 
FROM encuesta en
	INNER JOIN empleado em ON em.id = en.empleado
	INNER JOIN puesto p ON em.puesto = p.id
WHERE 
	em.puesto = 'Mesero' AND
	en.fecha BETWEEN NOW() AND NOW() - interval '6 months'
GROUP BY en.empleado, mes;

-- name: SignIn :exec
INSERT INTO 
	usuario (nombre, contraseña, tipo, empleado) 
VALUES ($1, $2, $3, $4);

-- name: LogIn :one
SELECT t.nombre as TipoUsuario, e.id as IdEmpleado
FROM usuario u
	INNER JOIN tipoUsuario t ON u.tipo = t.id
	INNER JOIN empleado e ON u.empleado = e.id
WHERE u.nombre = $1 AND u.contraseña = $2
LIMIT 1;

-- BARTENDER
-- name: GetPendingDrinks :many
SELECT *
FROM pedido p
	INNER JOIN estadosPedidos e ON p.estado = e.id
	INNER JOIN itemMenu im ON p.item = im.id
	INNER JOIN itemMenuCategoria imc ON im.categoria = imc.id
WHERE 
	imc.nombre = 'Bebidas' 
	AND e.nombre = 'Pedido'
ORDER BY p.fecha DESC;

-- name: SetOrderPreparing :exec
UPDATE pedido p
SET p.estado = (SELECT ep.id FROM estadosPedidos ep WHERE nombre = 'En preparación')
WHERE p.id = $1;

-- name: SetOrderDelivered :exec
UPDATE pedido p
SET p.estado = (SELECT ep.id FROM estadosPedidos ep WHERE nombre = 'Entregado')
WHERE p.id = $1;

-- MESERO
-- name: CreateClient :one
INSERT INTO cliente (nombre, direccion, nit) VALUES ($1, $2, $3) RETURNING *;

-- name: OpenAccount :one
INSERT INTO cuenta (mesa, estaCerrada, numPersonas) VALUES ($1, false, $2) RETURNING *;

-- name: TakeOrder :one
INSERT INTO pedido (fecha, estado, cantidad, cuenta, item) VALUES 
(NOW(), (SELECT id FROM estadosPedidos WHERE nombre = 'Pedido'), $1, $2, $3)
RETURNING *;

-- name: CloseAccount :exec
UPDATE cuenta
SET estaCerrada = true
WHERE numCuenta = $1;

-- name: GetActiveAccounts :many
SELECT *
FROM cuenta
WHERE estaCerrada = false;

-- name: GetClients :many
SELECT * FROM cliente;

-- name: GetEmployees :many
SELECT * FROM empleado;

-- name: GenerateBill :one
INSERT INTO factura (fecha, cuenta, cliente) VALUES (NOW(), $1, $2)
RETURNING *;

-- name: GetClient :one
SELECT * FROM cliente WHERE id=$1;

-- name: AddPayment :exec
INSERT INTO pago (tipo, monto, factura) VALUES ($1, $2, $3);

-- name: TakeSurvey :exec
INSERT INTO encuesta (empleado, cliente, gradoAmabilidad, gradoExactitud, fecha) VALUES ($1, $2, $3, $4, NOW());

--CHEF --Get Pending Dishes
--Para platillos CHEF
-- name: GetPendingDishes :many
SELECT 
	IM.id,
	P.fecha,
	IM.nombre, 
	P.cantidad,
	EP.nombre
FROM pedido P
	INNER JOIN estadosPedidos EP on P.estado = EP.id
	INNER JOIN itemMenu IM on P.item = IM.id
	inner join itemmenucategoria IMC on IM.categoria = IMC.id
where (EP.nombre = 'En espera' or EP.nombre = 'Cocinado') and IM.categoria = 1 ;

-- name: GetMenuItems :many
SELECT im.*, imc.nombre as NombreCategoria
FROM itemMenu im
	INNER JOIN itemmenucategoria imc ON im.categoria = imc.id;
