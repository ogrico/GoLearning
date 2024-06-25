# <h1 align="center" > Modulo para trabajar con procesamientos </h1>
****
## Contenido

- [Base de datos](#base-de-datos)

- [Funcionalidades](#funcionalidades-expuestas)

- [Documentación extra](#documentacion-adicional)


****
## Base de datos

La siguiente es la estructura SQL para crear la tabla `cuentas_bancarias`:

```sql
CREATE TABLE cuentasbancarias (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cedula VARCHAR(20) NOT NULL,
    nombre_completo VARCHAR(100) NOT NULL,
    numero_cuenta VARCHAR(20) NOT NULL,
    entidad_cuenta VARCHAR(50) NOT NULL,
    saldo_disponible DECIMAL(15, 2) NOT NULL,
    fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    estado_cuenta ENUM('activa', 'inactiva') NOT NULL
);
```
Algunos datos base de pruebas para la la tabla `cuentasbancarias`:

```sql
INSERT INTO cuentasbancarias (cedula, nombre_completo, numero_cuenta, entidad_cuenta, saldo_disponible, fecha_actualizacion, fecha_creacion, estado_cuenta)
VALUES
    ('10000001', 'Usuario 1', '0000000001', 'Entidad 1', 1000.00, NOW(), NOW(), 'activa'),
    ('10000002', 'Usuario 2', '0000000002', 'Entidad 2', 2000.00, NOW(), NOW(), 'activa'),
    ('10000003', 'Usuario 3', '0000000003', 'Entidad 3', 3000.00, NOW(), NOW(), 'activa'),
    ('10000004', 'Usuario 4', '0000000004', 'Entidad 4', 4000.00, NOW(), NOW(), 'activa'),
    ('10000005', 'Usuario 5', '0000000005', 'Entidad 5', 5000.00, NOW(), NOW(), 'activa'),
    ('10000006', 'Usuario 6', '0000000006', 'Entidad 6', 6000.00, NOW(), NOW(), 'activa'),
    ('10000007', 'Usuario 7', '0000000007', 'Entidad 7', 7000.00, NOW(), NOW(), 'activa'),
    ('10000008', 'Usuario 8', '0000000008', 'Entidad 8', 8000.00, NOW(), NOW(), 'activa'),
    ('10000009', 'Usuario 9', '0000000009', 'Entidad 9', 9000.00, NOW(), NOW(), 'activa'),
    ('10000010', 'Usuario 10', '0000000010', 'Entidad 10', 10000.00, NOW(), NOW(), 'activa'),
    ('10000011', 'Usuario 11', '0000000011', 'Entidad 11', 11000.00, NOW(), NOW(), 'activa'),
    ('10000012', 'Usuario 12', '0000000012', 'Entidad 12', 12000.00, NOW(), NOW(), 'activa'),
    ('10000013', 'Usuario 13', '0000000013', 'Entidad 13', 13000.00, NOW(), NOW(), 'activa'),
    ('10000014', 'Usuario 14', '0000000014', 'Entidad 14', 14000.00, NOW(), NOW(), 'activa'),
    ('10000015', 'Usuario 15', '0000000015', 'Entidad 15', 15000.00, NOW(), NOW(), 'activa'),
    ('10000016', 'Usuario 16', '0000000016', 'Entidad 16', 16000.00, NOW(), NOW(), 'activa'),
    ('10000017', 'Usuario 17', '0000000017', 'Entidad 17', 17000.00, NOW(), NOW(), 'activa'),
    ('10000018', 'Usuario 18', '0000000018', 'Entidad 18', 18000.00, NOW(), NOW(), 'activa'),
    ('10000019', 'Usuario 19', '0000000019', 'Entidad 19', 19000.00, NOW(), NOW(), 'activa'),
    ('10000020', 'Usuario 20', '0000000020', 'Entidad 20', 20000.00, NOW(), NOW(), 'activa')

```
****

## Funcionalidades expuestas

- `Listar Cuentas`: Permite consultar el listado de cuentas y saldos
- `Listar cuentas por parametros`: Permite consultar las cuentas por parametros
- `Procesar archivos planos`: Permite leer y procesar archivos planos (txt, json y yaml)

Archivo txt para simular la recepción y procesar la información

```
123456789,John Doe,0123456789,Banco A,1500.25,2024-06-20,2020-05-10,activo
987654321,Jane Smith,9876543210,Banco B,2500.75,2024-06-21,2019-11-15,inactivo
456789123,Alice Johnson,4567890123,Banco C,3500.50,2024-06-22,2021-04-30,activo
789123456,Michael Brown,7890123456,Banco D,4500.00,2024-06-23,2022-09-25,inactivo
321654987,Emily Davis,3216549870,Banco E,5500.80,2024-06-24,2023-02-18,activo
```

****

## Documentacion adicional

- https://pkg.go.dev/
- https://developer.mozilla.org/es/docs/Web/HTTP
- https://gorm.io/index.html
- https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/JSON