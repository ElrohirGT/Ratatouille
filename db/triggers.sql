-- Check that makes sure no payment exceeds the amount defined by account
CREATE OR REPLACE FUNCTION order_pay_overflow_check()
RETURNS trigger as
$BODY$
declare pays_exceeds_total boolean;
begin
		SELECT 
		((SELECT SUM(monto) FROM pago WHERE factura = new.factura) + new.monto) 
		> (SELECT c.total 
			FROM pago p 
				INNER JOIN factura f ON p.factura = f.numFactura 
				INNER JOIN cuenta c ON f.cuenta = c.numCuenta
			WHERE p.factura = new.factura) 
		INTO pays_exceeds_total;

    if pays_exceeds_total then
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

-- Actualizacion de la suma total del pedido 
-- un pedido se ponga en una cuenta se agarre el precio unitario de la tabla itemMenu y 
--lo multiplique por la cantidad del pedido y eso es un total pero ese total se tiene que sumar por 
--todos los pedidos ingresados a la misma cuenta
--PENDIENTE DE REVISAR.
CREATE OR REPLACE FUNCTION actualizar_total_cuenta()
RETURNS TRIGGER AS
$BODY$
DECLARE
    total_pedido money;
    total_cuenta money;
BEGIN
    -- Calcular el total del pedido
    total_pedido := NEW.cantidad * (
        SELECT precioUnitario FROM itemMenu WHERE id = NEW.item
    );
    
    -- Actualizar el total de la cuenta sumando el total del pedido
    UPDATE cuenta
    SET total = total + total_pedido
    WHERE numCuenta = NEW.cuenta;
    
    RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;

-- Creación del trigger
CREATE TRIGGER calcular_total_pedido
AFTER INSERT OR UPDATE
ON pedido
FOR EACH ROW
EXECUTE FUNCTION actualizar_total_cuenta();

