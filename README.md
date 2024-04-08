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

## Ejecutar el cliente

El cliente es una TUI (Text User Interface) escrito en [Go](https://go.dev/) en conjunto con las librerias [Bubbletea](https://github.com/charmbracelet/bubbletea), [Bubbles](https://github.com/charmbracelet/bubbles), [Lipgloss](https://github.com/charmbracelet/lipgloss) . Estos son los pasos para ejecutarlo

1. Iniciar la BD y desplegar el entorno de desarrollo
   
   ```bash
   nix develop --impure
   ```

2. Navegar a la carpeta con el codigo de cliente `/client`
   
   ```bash
   cd ./client
   ```

3. Por último ya solo basta compilar y ejecutar.
   
   ```bash
   make run
   ```
