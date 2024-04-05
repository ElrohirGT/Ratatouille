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
