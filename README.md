# gRPC Chess

## Descripción
**gRPC Chess** es un proyecto que permite realizar consultas sobre partidas de ajedrez en línea utilizando la tecnología gRPC. Este sistema ofrece un backend eficiente desarrollado en Go, proporcionando acceso a datos detallados de partidas jugadas en plataformas online. Utilizando Microsoft Azure para la infraestructura, Docker para la contenedorización y Postman para las pruebas, el sistema ofrece un análisis rápido y preciso de las partidas de ajedrez.

## Funcionalidad
El proyecto gRPC Chess permite:
- Consultar el nombre del usuario.
- Consultar el tipo de victoria.
- Consultar el color que usó el jugador.

## Tecnologías Utilizadas
- **Lenguaje de Programación**: Go
- **Infraestructura**: Microsoft Azure
- **Contenedorización**: Docker
- **Pruebas**: Postman

## Cómo Ejecutarlo Localmente

### Prerrequisitos
- Docker instalado en tu máquina.
- Go instalado en tu máquina.

### Instrucciones
1. Clona el repositorio:
    ```sh
    git clone https://github.com/DiegoMejiaa/Proyecto.git
    cd Proyecto
    ```
2. Construye y ejecuta el contenedor Docker:
    ```sh
    docker build -t grpc-chess .
    docker run -p 8080:8080 grpc-chess
    ```
3. Accede al servicio en `http://localhost:8080` y utiliza los endpoints gRPC.

## Cómo Interactuar con el Servicio en Azure
El servicio está desplegado en Microsoft Azure y se puede acceder mediante la siguiente URL:
- [Servicio en Azure](https://grcpchess.azurewebsites.net/)

Para interactuar con el servicio:
1. Abre Postman o cualquier cliente gRPC de tu elección.
2. Configura una nueva solicitud gRPC apuntando a `https://grcpchess.azurewebsites.net/`.
3. Utiliza los métodos proporcionados para realizar las consultas necesarias.

## Entrada en Tu Portafolio

### Enlaces Importantes
- **Repositorio del Proyecto**: [https://github.com/DiegoMejiaa/Proyecto](https://github.com/DiegoMejiaa/Proyecto)
- **Servicio en Azure**: [https://grcpchess.azurewebsites.net/](https://grcpchess.azurewebsites.net/)

### Enlace para Descargar el Archivo .proto
- **Descargar archivo .proto**: [https://grcpchess.azurewebsites.net/proto/juegos.proto](https://grcpchess.azurewebsites.net/proto/juegos.proto)



