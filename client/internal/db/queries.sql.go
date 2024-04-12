// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addPayment = `-- name: AddPayment :exec
INSERT INTO pago (tipo, monto, factura) VALUES ($1, $2, $3)
`

type AddPaymentParams struct {
	Tipo    int32
	Monto   float64
	Factura int32
}

func (q *Queries) AddPayment(ctx context.Context, arg AddPaymentParams) error {
	_, err := q.db.ExecContext(ctx, addPayment, arg.Tipo, arg.Monto, arg.Factura)
	return err
}

const closeAccount = `-- name: CloseAccount :exec
UPDATE cuenta
SET estaCerrada = true
WHERE numCuenta = $1
`

func (q *Queries) CloseAccount(ctx context.Context, numcuenta int32) error {
	_, err := q.db.ExecContext(ctx, closeAccount, numcuenta)
	return err
}

const createClient = `-- name: CreateClient :one
INSERT INTO cliente (nombre, direccion, nit) VALUES ($1, $2, $3) RETURNING id, nombre, direccion, nit
`

type CreateClientParams struct {
	Nombre    string
	Direccion sql.NullString
	Nit       string
}

// MESERO
func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, createClient, arg.Nombre, arg.Direccion, arg.Nit)
	var i Cliente
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Direccion,
		&i.Nit,
	)
	return i, err
}

const generateBill = `-- name: GenerateBill :one
INSERT INTO factura (fecha, cuenta, cliente, total) VALUES (NOW(), $1, $2, (select total from cuenta where numCuenta = $1))
RETURNING numfactura, fecha, cuenta, cliente
`

type GenerateBillParams struct {
	Cuenta  int32
	Cliente int32
}

func (q *Queries) GenerateBill(ctx context.Context, arg GenerateBillParams) (Factura, error) {
	row := q.db.QueryRowContext(ctx, generateBill, arg.Cuenta, arg.Cliente)
	var i Factura
	err := row.Scan(
		&i.Numfactura,
		&i.Fecha,
		&i.Cuenta,
		&i.Cliente,
	)
	return i, err
}

const getActiveAccounts = `-- name: GetActiveAccounts :many
SELECT mesa, numcuenta, estacerrada, numpersonas, total
FROM cuenta
WHERE estaCerrada = false
`

func (q *Queries) GetActiveAccounts(ctx context.Context) ([]Cuentum, error) {
	rows, err := q.db.QueryContext(ctx, getActiveAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cuentum
	for rows.Next() {
		var i Cuentum
		if err := rows.Scan(
			&i.Mesa,
			&i.Numcuenta,
			&i.Estacerrada,
			&i.Numpersonas,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAverageTimeToEatPerClientQuantity = `-- name: GetAverageTimeToEatPerClientQuantity :many

SELECT
	c.numPersonas,
	AVERAGE(MAX(p.fecha) OVER (PARTITION BY c.numCuenta) - MIN(p.fecha) OVER (PARTITION by c.numCuenta)) as timeToEat
FROM cuenta c
	LEFT JOIN pedido p ON p.numCuenta = c.numCuenta
WHERE c.estaCerrada AND p.fecha BETWEEN $1 AND $2
GROUP BY numPersonas
`

type GetAverageTimeToEatPerClientQuantityParams struct {
	Fecha   time.Time
	Fecha_2 time.Time
}

type GetAverageTimeToEatPerClientQuantityRow struct {
	Numpersonas int32
	Timetoeat   interface{}
}

// 3. Promedio de tiempo en que se tardan los clientes en comer, agrupando la cantidad de
// personas comiendo, por ejemplo: 2 personas: 1 hora 10 minutos, 3 personas: 1 hora 15
// minutos, etc. entre un rango de fechas solicitadas al usuario.
func (q *Queries) GetAverageTimeToEatPerClientQuantity(ctx context.Context, arg GetAverageTimeToEatPerClientQuantityParams) ([]GetAverageTimeToEatPerClientQuantityRow, error) {
	rows, err := q.db.QueryContext(ctx, getAverageTimeToEatPerClientQuantity, arg.Fecha, arg.Fecha_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAverageTimeToEatPerClientQuantityRow
	for rows.Next() {
		var i GetAverageTimeToEatPerClientQuantityRow
		if err := rows.Scan(&i.Numpersonas, &i.Timetoeat); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getClient = `-- name: GetClient :one
SELECT id, nombre, direccion, nit FROM cliente WHERE id=$1
`

func (q *Queries) GetClient(ctx context.Context, id int32) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, getClient, id)
	var i Cliente
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Direccion,
		&i.Nit,
	)
	return i, err
}

const getClients = `-- name: GetClients :many
SELECT id, nombre, direccion, nit FROM cliente
`

func (q *Queries) GetClients(ctx context.Context) ([]Cliente, error) {
	rows, err := q.db.QueryContext(ctx, getClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cliente
	for rows.Next() {
		var i Cliente
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Direccion,
			&i.Nit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getComplaintsForDishBetween = `-- name: GetComplaintsForDishBetween :many

SELECT cliente, gravedad, motivo, fecha, empleado, item
FROM queja
WHERE item = $1 AND fecha $2 AND $3
`

type GetComplaintsForDishBetweenParams struct {
	Item    sql.NullInt32
	Column2 interface{}
	Column3 interface{}
}

// 5. Reporte de las quejas agrupadas por plato para un rango de fechas solicitadas al usuario.
func (q *Queries) GetComplaintsForDishBetween(ctx context.Context, arg GetComplaintsForDishBetweenParams) ([]Queja, error) {
	rows, err := q.db.QueryContext(ctx, getComplaintsForDishBetween, arg.Item, arg.Column2, arg.Column3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Queja
	for rows.Next() {
		var i Queja
		if err := rows.Scan(
			&i.Cliente,
			&i.Gravedad,
			&i.Motivo,
			&i.Fecha,
			&i.Empleado,
			&i.Item,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getComplaintsForEmployeeBetween = `-- name: GetComplaintsForEmployeeBetween :many

SELECT cliente, gravedad, motivo, fecha, empleado, item
FROM queja 
WHERE empleado = $1 AND fecha BETWEEN $2 AND $3
`

type GetComplaintsForEmployeeBetweenParams struct {
	Empleado sql.NullInt32
	Fecha    time.Time
	Fecha_2  time.Time
}

// 4. Reporte de las quejas agrupadas por persona para un rango de fechas solicitadas al usuario.
func (q *Queries) GetComplaintsForEmployeeBetween(ctx context.Context, arg GetComplaintsForEmployeeBetweenParams) ([]Queja, error) {
	rows, err := q.db.QueryContext(ctx, getComplaintsForEmployeeBetween, arg.Empleado, arg.Fecha, arg.Fecha_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Queja
	for rows.Next() {
		var i Queja
		if err := rows.Scan(
			&i.Cliente,
			&i.Gravedad,
			&i.Motivo,
			&i.Fecha,
			&i.Empleado,
			&i.Item,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEfficiencyReport = `-- name: GetEfficiencyReport :many


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
GROUP BY en.empleado, mes
`

type GetEfficiencyReportRow struct {
	Empleado int32
	Mes      string
	Avg      float64
	Avg_2    float64
}

// 6. Reporte de eficiencia de meseros mostrando los resultados de las encuestas, agrupado
// por personas y por mes para los últimos 6 meses.
// | Persona | Mes | avgAmabilidad | avgExactirud |
func (q *Queries) GetEfficiencyReport(ctx context.Context) ([]GetEfficiencyReportRow, error) {
	rows, err := q.db.QueryContext(ctx, getEfficiencyReport)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEfficiencyReportRow
	for rows.Next() {
		var i GetEfficiencyReportRow
		if err := rows.Scan(
			&i.Empleado,
			&i.Mes,
			&i.Avg,
			&i.Avg_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmployees = `-- name: GetEmployees :many
SELECT id, nombre, puesto, area FROM empleado
`

func (q *Queries) GetEmployees(ctx context.Context) ([]Empleado, error) {
	rows, err := q.db.QueryContext(ctx, getEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Empleado
	for rows.Next() {
		var i Empleado
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Puesto,
			&i.Area,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMenuItems = `-- name: GetMenuItems :many
SELECT im.id, im.nombre, im.descripcion, im.preciounitario, im.categoria, imc.nombre as NombreCategoria
FROM itemMenu im
	INNER JOIN itemmenucategoria imc ON im.categoria = imc.id
`

type GetMenuItemsRow struct {
	ID              int32
	Nombre          string
	Descripcion     string
	Preciounitario  string
	Categoria       int32
	Nombrecategoria string
}

func (q *Queries) GetMenuItems(ctx context.Context) ([]GetMenuItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMenuItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMenuItemsRow
	for rows.Next() {
		var i GetMenuItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Descripcion,
			&i.Preciounitario,
			&i.Categoria,
			&i.Nombrecategoria,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMostFamousDishesBetween = `-- name: GetMostFamousDishesBetween :many

SELECT p.item, i.nombre, i.descripcion, COUNT(p) as count
FROM pedido p
	RIGHT JOIN itemMenu i on p.item = i.id
WHERE p.fecha BETWEEN $1 AND $2 
GROUP BY p.item, i.nombre, i.descripcion
ORDER BY count DESC
`

type GetMostFamousDishesBetweenParams struct {
	Fecha   time.Time
	Fecha_2 time.Time
}

type GetMostFamousDishesBetweenRow struct {
	Item        sql.NullInt32
	Nombre      string
	Descripcion string
	Count       int64
}

// 1. Reporte de los platos más pedidos por los clientes en un rango de fechas solicitadas al usuario.
func (q *Queries) GetMostFamousDishesBetween(ctx context.Context, arg GetMostFamousDishesBetweenParams) ([]GetMostFamousDishesBetweenRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostFamousDishesBetween, arg.Fecha, arg.Fecha_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMostFamousDishesBetweenRow
	for rows.Next() {
		var i GetMostFamousDishesBetweenRow
		if err := rows.Scan(
			&i.Item,
			&i.Nombre,
			&i.Descripcion,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPendingDishes = `-- name: GetPendingDishes :many
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
where (EP.nombre = 'En espera' or EP.nombre = 'Cocinado') and IM.categoria = 1
`

type GetPendingDishesRow struct {
	ID       int32
	Fecha    time.Time
	Nombre   string
	Cantidad int32
	Nombre_2 string
}

// CHEF --Get Pending Dishes
// Para platillos CHEF
func (q *Queries) GetPendingDishes(ctx context.Context) ([]GetPendingDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPendingDishes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPendingDishesRow
	for rows.Next() {
		var i GetPendingDishesRow
		if err := rows.Scan(
			&i.ID,
			&i.Fecha,
			&i.Nombre,
			&i.Cantidad,
			&i.Nombre_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPendingDrinks = `-- name: GetPendingDrinks :many
SELECT p.id, fecha, estado, cantidad, cuenta, item, e.id, e.nombre, im.id, im.nombre, descripcion, preciounitario, categoria, imc.id, imc.nombre
FROM pedido p
	INNER JOIN estadosPedidos e ON p.estado = e.id
	INNER JOIN itemMenu im ON p.item = im.id
	INNER JOIN itemMenuCategoria imc ON im.categoria = imc.id
WHERE 
	imc.nombre = 'Bebidas' 
	AND e.nombre = 'Pedido'
ORDER BY p.fecha DESC
`

type GetPendingDrinksRow struct {
	ID             int32
	Fecha          time.Time
	Estado         int32
	Cantidad       int32
	Cuenta         int32
	Item           int32
	ID_2           int32
	Nombre         string
	ID_3           int32
	Nombre_2       string
	Descripcion    string
	Preciounitario string
	Categoria      int32
	ID_4           int32
	Nombre_3       string
}

// BARTENDER
func (q *Queries) GetPendingDrinks(ctx context.Context) ([]GetPendingDrinksRow, error) {
	rows, err := q.db.QueryContext(ctx, getPendingDrinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPendingDrinksRow
	for rows.Next() {
		var i GetPendingDrinksRow
		if err := rows.Scan(
			&i.ID,
			&i.Fecha,
			&i.Estado,
			&i.Cantidad,
			&i.Cuenta,
			&i.Item,
			&i.ID_2,
			&i.Nombre,
			&i.ID_3,
			&i.Nombre_2,
			&i.Descripcion,
			&i.Preciounitario,
			&i.Categoria,
			&i.ID_4,
			&i.Nombre_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoleOfUser = `-- name: GetRoleOfUser :many

SELECT 
	tu.nombre AS tipo_usuario
FROM 
	usuario u
JOIN 
	tipoUsuario tu ON u.tipo = tu.id
WHERE 
	u.nombre = $1 AND u.contraseña = $2
`

type GetRoleOfUserParams struct {
	Nombre     string
	Contraseña string
}

// 0. Saber el rol de la persona a partir del usuario y contraseña ingresados por el usuario.
func (q *Queries) GetRoleOfUser(ctx context.Context, arg GetRoleOfUserParams) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getRoleOfUser, arg.Nombre, arg.Contraseña)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var tipo_usuario string
		if err := rows.Scan(&tipo_usuario); err != nil {
			return nil, err
		}
		items = append(items, tipo_usuario)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRushHourBetween = `-- name: GetRushHourBetween :one

SELECT EXTRACT(hour from p.fecha) as horario, COUNT(*) as count
FROM pedido p
WHERE p.fecha BETWEEN $1 AND $2
GROUP BY horario
ORDER BY count DESC
LIMIT 1
`

type GetRushHourBetweenParams struct {
	Fecha   time.Time
	Fecha_2 time.Time
}

type GetRushHourBetweenRow struct {
	Horario string
	Count   int64
}

// 2. Horario en el que se ingresan más pedidos entre un rango de fechas solicitadas al usuario.
func (q *Queries) GetRushHourBetween(ctx context.Context, arg GetRushHourBetweenParams) (GetRushHourBetweenRow, error) {
	row := q.db.QueryRowContext(ctx, getRushHourBetween, arg.Fecha, arg.Fecha_2)
	var i GetRushHourBetweenRow
	err := row.Scan(&i.Horario, &i.Count)
	return i, err
}

const logIn = `-- name: LogIn :one
SELECT t.nombre as TipoUsuario, e.id as IdEmpleado
FROM usuario u
	INNER JOIN tipoUsuario t ON u.tipo = t.id
	INNER JOIN empleado e ON u.empleado = e.id
WHERE u.nombre = $1 AND u.contraseña = $2
LIMIT 1
`

type LogInParams struct {
	Nombre     string
	Contraseña string
}

type LogInRow struct {
	Tipousuario string
	Idempleado  int32
}

func (q *Queries) LogIn(ctx context.Context, arg LogInParams) (LogInRow, error) {
	row := q.db.QueryRowContext(ctx, logIn, arg.Nombre, arg.Contraseña)
	var i LogInRow
	err := row.Scan(&i.Tipousuario, &i.Idempleado)
	return i, err
}

const openAccount = `-- name: OpenAccount :one
INSERT INTO cuenta (mesa, estaCerrada, numPersonas) VALUES ($1, false, $2) RETURNING mesa, numcuenta, estacerrada, numpersonas, total
`

type OpenAccountParams struct {
	Mesa        int32
	Numpersonas int32
}

func (q *Queries) OpenAccount(ctx context.Context, arg OpenAccountParams) (Cuentum, error) {
	row := q.db.QueryRowContext(ctx, openAccount, arg.Mesa, arg.Numpersonas)
	var i Cuentum
	err := row.Scan(
		&i.Mesa,
		&i.Numcuenta,
		&i.Estacerrada,
		&i.Numpersonas,
		&i.Total,
	)
	return i, err
}

const registerComplaint = `-- name: RegisterComplaint :exec
INSERT INTO queja (cliente, gravedad, motivo, fecha, empleado, item) VALUES ($1, $2, $3, NOW(), $4, $5)
`

type RegisterComplaintParams struct {
	Cliente  int32
	Gravedad int32
	Motivo   string
	Empleado sql.NullInt32
	Item     sql.NullInt32
}

func (q *Queries) RegisterComplaint(ctx context.Context, arg RegisterComplaintParams) error {
	_, err := q.db.ExecContext(ctx, registerComplaint,
		arg.Cliente,
		arg.Gravedad,
		arg.Motivo,
		arg.Empleado,
		arg.Item,
	)
	return err
}

const setOrderDelivered = `-- name: SetOrderDelivered :exec
UPDATE pedido p
SET p.estado = (SELECT ep.id FROM estadosPedidos ep WHERE nombre = 'Entregado')
WHERE p.id = $1
`

func (q *Queries) SetOrderDelivered(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, setOrderDelivered, id)
	return err
}

const setOrderPreparing = `-- name: SetOrderPreparing :exec
UPDATE pedido p
SET p.estado = (SELECT ep.id FROM estadosPedidos ep WHERE nombre = 'En preparación')
WHERE p.id = $1
`

func (q *Queries) SetOrderPreparing(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, setOrderPreparing, id)
	return err
}

const signIn = `-- name: SignIn :exec
INSERT INTO 
	usuario (nombre, contraseña, tipo, empleado) 
VALUES ($1, $2, $3, $4)
`

type SignInParams struct {
	Nombre     string
	Contraseña string
	Tipo       int32
	Empleado   int32
}

func (q *Queries) SignIn(ctx context.Context, arg SignInParams) error {
	_, err := q.db.ExecContext(ctx, signIn,
		arg.Nombre,
		arg.Contraseña,
		arg.Tipo,
		arg.Empleado,
	)
	return err
}

const takeOrder = `-- name: TakeOrder :one
INSERT INTO pedido (fecha, estado, cantidad, cuenta, item) VALUES 
(NOW(), (SELECT id FROM estadosPedidos WHERE nombre = 'Pedido'), $1, $2, $3)
RETURNING id, fecha, estado, cantidad, cuenta, item
`

type TakeOrderParams struct {
	Cantidad int32
	Cuenta   int32
	Item     int32
}

func (q *Queries) TakeOrder(ctx context.Context, arg TakeOrderParams) (Pedido, error) {
	row := q.db.QueryRowContext(ctx, takeOrder, arg.Cantidad, arg.Cuenta, arg.Item)
	var i Pedido
	err := row.Scan(
		&i.ID,
		&i.Fecha,
		&i.Estado,
		&i.Cantidad,
		&i.Cuenta,
		&i.Item,
	)
	return i, err
}

const takeSurvey = `-- name: TakeSurvey :exec
INSERT INTO encuesta (empleado, cliente, gradoAmabilidad, gradoExactitud, fecha) VALUES ($1, $2, $3, $4, NOW())
`

type TakeSurveyParams struct {
	Empleado        int32
	Cliente         int32
	Gradoamabilidad int32
	Gradoexactitud  int32
}

func (q *Queries) TakeSurvey(ctx context.Context, arg TakeSurveyParams) error {
	_, err := q.db.ExecContext(ctx, takeSurvey,
		arg.Empleado,
		arg.Cliente,
		arg.Gradoamabilidad,
		arg.Gradoexactitud,
	)
	return err
}
