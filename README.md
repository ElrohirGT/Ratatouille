# Ratatouille
Este proyecto utiliza [Nix Flakes](https://nixos.wiki/wiki/Flakes) y [DevEnv](https://devenv.sh/) para sus ambientes de desarrollo. Si quieres tener la mejor experiencia te recomendamos instalarlos.

Para ingresar a una consola de desarrollador con todas las dependencias del proyecto preinstaladas se utiliza:
```bash
nix develop --impure
```

## Iniciar la Base de datos
Para iniciar la base de datos y cualquier otro servicio que el proyecto necesite tener corriendo usar:
```bash
devenv up
```

Recuerda que debes estar dentro de la consola de desarrollador descrita anteriormente.

Si alguna vez los scripts dentro de `db` cambian, debes borrar la carpeta `.devenv/postgres` y volver a hacer `devenv up`.

Para automatizar este proceso se creó el comando:
```bash
nix run .#restartServices
```

Por favor recuerda correr este comando estando en la carpeta root del directorio, no importa si ya corriste `nix develop --impure` antes puesto que este comando lo corre por tí!

Se necesitan algunos queries de análisis de la data dentro de la aplicación, estos queries se pueden ver en el directorio `db`.
