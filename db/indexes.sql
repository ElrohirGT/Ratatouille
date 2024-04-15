CREATE INDEX usuario_username_idx ON usuario (nombre);
CREATE INDEX usuario_password_idx ON usuario (contraseña);

CREATE INDEX pedido_date_idx ON pedido (fecha);
CREATE INDEX item_nombre_idx ON itemMenu (nombre);
CREATE INDEX item_descripcion_idx ON itemMenu (descripcion);

CREATE INDEX pedido_hour_from_date_idx ON pedido (EXTRACT(hour from fecha));

CREATE INDEX queja_date_idx ON queja (fecha);
