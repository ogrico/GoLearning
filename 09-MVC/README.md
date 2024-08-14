# <h1 align="center" > Modulo para trabajar con MVC </h1>
****
## Contenido

- [Base de datos](#base-de-datos)

- [Funcionalidades](#funcionalidades-expuestas)

- [Documentación extra](#documentacion-adicional)


****
## Base de datos

La siguiente es la estructura SQL para crear la tabla `cuentas_bancarias`:

```sql
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
```
Algunos datos base de pruebas para la la tabla `usuarios`:

```sql
INSERT INTO users (id, name, email) VALUES
(default, 'Alice Smith', 'alice.smith@example.com'),
(default, 'Bob Johnson', 'bob.johnson@example.com'),
(default, 'Charlie Brown', 'charlie.brown@example.com'),
(default, 'Diana Prince', 'diana.prince@example.com'),
(default, 'Edward Norton', 'edward.norton@example.com'),
(default, 'Fiona Apple', 'fiona.apple@example.com'),
(default, 'George Michael', 'george.michael@example.com'),
(default, 'Hannah Montana', 'hannah.montana@example.com'),
(default, 'Isaac Newton', 'isaac.newton@example.com'),
(default, 'Jane Austen', 'jane.austen@example.com');

```
****

## Funcionalidades expuestas

- `Listar Usuarios`: Permite consultar el listado de usuarios registrados
- `Crear usuarios`: Permite crear nuevos usuarios
- `Eliminar usuarios`: Permite eliminar nuevos usuarios


****

## Documentacion adicional

- https://pkg.go.dev/
- https://developer.mozilla.org/es/docs/Web/HTTP
- https://gorm.io/index.html
- https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/JSON