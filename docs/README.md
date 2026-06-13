# ERP Interno Momentto Garden y Viajes Premium

Este proyecto es un ERP interno para administrar dos unidades de negocio dentro de una sola plataforma:

- Momentto Garden
- Viajes Premium

La idea no es separar dos sistemas distintos, sino tener un solo ERP con datos y permisos segmentados por negocio.

## Objetivo

El sistema debe permitir que:

- Un usuario pertenezca a uno o varios negocios.
- Cada negocio tenga sus propios leads, contactos, cotizaciones, clientes y operacion.
- Existan modulos compartidos entre negocios.
- Existan modulos especificos para cada negocio.

## Stack Backend

- Go
- Gin
- GORM
- PostgreSQL
- Arquitectura Vertical Slice

## Estructura conceptual

El sistema se organiza en tres niveles:

- `platform`: base del ERP
- `modules`: funcionalidades compartidas
- `business`: funcionalidades especificas por negocio

### Platform

Aqui viven las piezas globales del sistema:

- Users
- Auth
- Roles
- Permissions
- Business Units
- Departments
- Memberships

### Modules

Aqui van los modulos que pueden reutilizarse entre los negocios:

- CRM
- Sales
- Catalog
- Operations
- Finance
- Communication

### Business

Aqui van las funciones exclusivas de cada negocio:

- `business/momentto`
- `business/viajes`

## Regla principal

- Si algo existe igual en ambos negocios, va en `modules`.
- Si algo solo tiene sentido para un negocio, va en `business/{negocio}`.
- Si algo es base del sistema, va en `platform`.

## Separacion de datos

La separacion principal se hace por `business_unit_id`.

Esto significa que casi todos los registros operativos deben pertenecer a una unidad de negocio.

## Sobre departamentos

Los departamentos existen dentro de una unidad de negocio y sirven para organizar usuarios por area.

La idea es que el rol del usuario pueda depender del departamento donde este asignado.

Tanto los departamentos como los roles se definen por unidad de negocio, no de forma global.

Ejemplo:

- Business Unit: Momentto Garden
- Department: Ventas
- Role: Seller

Consulta el flujo completo en [backend_arquitecture.md](./backend_arquitecture.md).

## Modelo mental recomendado

No pensar el sistema como dos ERPs separados.

Pensarlo como:

- Un solo ERP interno
- Usuarios globales
- Modulos compartidos
- Registros separados por unidad de negocio
- Departamentos internos por negocio
