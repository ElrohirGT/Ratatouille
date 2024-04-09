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
