-- Check that makes sure no payment exceeds the amount defined by account
CREATE OR REPLACE FUNCTION order_pay_overflow_check()
RETURNS trigger as
$BODY$
begin
	if (SELECT SUM(monto) +new.monto FROM pago WHERE factura = new.factura) 
		> (SELECT c.total 
			FROM pago p 
				INNER JOIN factura f ON p.factura = f.numFactura 
				INNER JOIN cuenta c ON f.cuenta = c.numCuenta
			WHERE p.factura = new.factura
			LIMIT 1) then
    raise exception 'El pago ha ingresar sobrepasa el monto por pagar de la cuenta.';
	end if;

	RETURN new;
end
$BODY$
LANGUAGE 'plpgsql';

CREATE OR REPLACE TRIGGER order_pay_overflow_check
BEFORE INSERT OR UPDATE
ON pago
FOR EACH ROW
EXECUTE PROCEDURE order_pay_overflow_check();

-- Check that makes sure no account can be marked as closed unless is payed fully
CREATE OR REPLACE FUNCTION account_must_be_fully_payed()
RETURNS TRIGGER AS
$BODY$
declare theres_more_to_pay boolean;
begin
		SELECT 
		(SELECT SUM(monto) 
			FROM pago p
				INNER JOIN factura f ON f.numFactura = p.factura AND f.cuenta = old.numCuenta)
		< old.total 
		INTO theres_more_to_pay;

		if new.estaCerrada AND theres_more_to_pay then
			raise exception 'La cuenta se debe pagar antes de poder cerrarla!';
		end if;

		RETURN new;
end
$BODY$
LANGUAGE 'plpgsql';

CREATE OR REPLACE TRIGGER account_must_be_fully_payed
BEFORE UPDATE
ON cuenta
FOR EACH ROW
EXECUTE PROCEDURE account_must_be_fully_payed();


-- Dejar un limite definido de la maxima cantidad de personas que puedan haber en una cuenta.
CREATE OR REPLACE FUNCTION assert_max_personas()
RETURNS TRIGGER AS
$BODY$
BEGIN
    IF NEW.numPersonas > 14 THEN
        RAISE EXCEPTION 'El número de personas no puede ser mayor de 14';
    END IF;
    RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;

-- Creación del trigger
CREATE TRIGGER assert_num_personas
BEFORE INSERT OR UPDATE
ON cuenta
FOR EACH ROW
EXECUTE FUNCTION assert_max_personas();



-- Crea la funciin para actualizar el total acumulado en la tabla cuenta
CREATE OR REPLACE FUNCTION actualizar_total_cuenta()
RETURNS TRIGGER AS
$BODY$
BEGIN
    -- Actualizar el campo total de la tabla cuenta
    UPDATE cuenta AS c
    SET total = (
        SELECT SUM(IM.precioUnitario * p.cantidad) AS cantidad_total
        FROM itemmenu IM
        INNER JOIN pedido p ON IM.id = p.item
        WHERE p.cuenta = NEW.cuenta
        GROUP BY p.cuenta
    )
    WHERE c.numCuenta = NEW.cuenta;
    
    RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;

-- Crear el trigger para despues de insertar o actualizar tabla pedido
CREATE TRIGGER trigger_actualizar_total_cuenta
AFTER INSERT OR UPDATE ON pedido
FOR EACH ROW
EXECUTE FUNCTION actualizar_total_cuenta();
