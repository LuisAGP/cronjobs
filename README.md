# Cronjobs - Cron as a Service

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Cronjobs** es una plataforma moderna que permite a usuarios programar y gestionar tareas HTTP recurrentes con facilidad. Imagínalo como un sistema de cron distribuido en la nube con una interfaz intuitiva, donde puedes configurar peticiones periódicas a cualquier endpoint HTTP sin necesidad de acceder a servidores o configurar complejos sistemas.

### 🚀 Características Principales

* Programación flexible: Ejecuta tareas cada X segundos (mínimo 60s)
* HTTP Requests: Soporta métodos GET, POST, PUT, DELETE
* Payload personalizado: Envía datos JSON personalizados en cada petición
* Gestión de headers: Configura headers de autenticación y contenido
* Multi-usuario: Sistema de autenticación seguro con JWT
* Panel intuitivo: Interfaz web para gestionar todas tus tareas programadas
* Motor confiable: Ejecución puntual con manejo de errores y reintentos

### ⚙️ Arquitectura Técnica

graph TD
    U[Usuario] -->|Programa tarea| W[Web UI]
    W --> A[API Go]
    A --> D[(MySQL)]
    S[Scheduler] -->|Consulta tareas| D
    S -->|Ejecuta| E[Endpoint Externo]
    E -->|Respuesta| L[Logs]

### 🌐 Endpoints de la API

POST   /api/register     # Registro usuario
POST   /api/login        # Autenticación
GET    /api/tasks        # Listar tareas
POST   /api/tasks        # Crear nueva tarea
PUT    /api/tasks/{id}   # Actualizar tarea
DELETE /api/tasks/{id}   # Eliminar tarea

### 📄 Licencia

Este proyecto está licenciado bajo la licencia MIT - ver el archivo [LICENSE](LICENSE.txt) para más detalles.
Cronjobs es software open-source bajo la licencia MIT.
