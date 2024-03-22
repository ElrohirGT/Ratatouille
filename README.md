# Ratatouille
Dependencias del proyecto no manejadas dentro de Nix:
- Docker

Una vez se tengan instaladas las dependencias del proyecto se puede utilizar:
```bash
nix develop
```

Para entrar a una shell con las demás dependencias del proyecto.

## Iniciar la Base de datos
La base de datos del proyecto la tenemos en un docker, para iniciar el docker se utiliza:
```bash
nix run .#restartDBDocker
```
