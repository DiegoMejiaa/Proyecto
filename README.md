# Chess GRPC Service

Este proyecto implementa un servicio gRPC para gestionar información de juegos de ajedrez. Proporciona operaciones como consultar información de juegos, listar todos los juegos, añadir nuevos juegos y buscar juegos por estado de victoria.

## Funcionalidad

El servicio incluye las siguientes funcionalidades:

- **GetJuegosInfo**: Obtiene información de un juego específico basado en el nombre.
- **GetJuegosList**: Devuelve un flujo con todos los juegos almacenados en la base de datos.
- **AddJuegos**: Permite añadir nuevos juegos enviando un flujo de datos.
- **GetJuegosByWin**: Devuelve un flujo de juegos filtrados por estado de victoria.

---

## Cómo ejecutar el proyecto localmente

### Prerrequisitos

1. **Instalar Docker **.
2. Crear un archivo `.env` en la raíz del proyecto con las siguientes variables configuradas:
   ```env
   DB_SERVER=<direccion_del_servidor>
   DB_USER=<usuario>
   DB_PASSWORD=<contraseña>
   DB_NAME=<nombre_de_base_de_datos>
   DB_PORT=<puerto_de_base_de_datos>
   ```
3. Configurar una base de datos Microsoft SQL Server con la tabla `tablero`:
   ```sql
   CREATE TABLE chess.tablero (
       Name NVARCHAR(50),
       Win NVARCHAR(50),
       Color NVARCHAR(50)
   );
   ```

### Pasos

1. **Clonar el repositorio**:
   ```bash
   git clone https://github.com/DiegoMejiaa/Proyecto
   cd Proyecto
   ```

2. **Construir y ejecutar el contenedor Docker**:
   ```bash
   docker build -t chess_grpc .
   docker run -p 50051:50051 -p 8080:8080 --env-file .env chess_grpc
   ```

3. **Verificar que el servicio esté funcionando**:
   - Endpoint de salud: [http://localhost:8080/health](http://localhost:8080/health)

---

## Interacción con el servicio en Azure

1. **Despliegue en Azure**:
   - Usa un servicio como Azure Kubernetes Service (AKS) o Azure Container Instances para alojar el contenedor.
   - Configura las variables de entorno necesarias en el servicio.

2. **Acceso al servicio**:
   - La API gRPC estará disponible en el puerto `50051`.
   - Usa Postman para probar las operaciones con un cliente gRPC:
     - Configura la URL del servidor gRPC en Postman.
     - Usa el archivo `.proto` para generar las definiciones de los métodos.

3. **Verificación del servicio de salud**:
   - Endpoint: `http://<azure_endpoint>:8080/health`

---

## Entrada en tu portafolio

- **Enlace al repositorio**: [GitHub](https://github.com/DiegoMejiaa/Proyecto)
- **Enlace al portafolio**: [Portafolio](https://diegomejia.pages.dev)

---

## Tecnologías utilizadas

- **Go**: Lenguaje principal para el servidor gRPC.
- **gRPC**: Framework para comunicación de alto rendimiento.
- **Docker**: Contenerización de la aplicación.
- **Microsoft SQL Server**: Base de datos para almacenar los juegos.
- **Postman**: Herramienta para pruebas de las APIs.
- **Azure**: Despliegue del servicio.

---



