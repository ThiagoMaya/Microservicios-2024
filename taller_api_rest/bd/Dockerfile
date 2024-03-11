# Usa la imagen oficial de MySQL
# Usa la imagen oficial de MySQL
FROM mysql:latest

# Copia el script de inicialización SQL al directorio /docker-entrypoint-initdb.d/
# Este script se ejecutará automáticamente al iniciar el contenedor MySQL
COPY init.sql /docker-entrypoint-initdb.d/
# Variables de entorno para configurar la base de datos, el usuario y la contraseña
ENV MYSQL_DATABASE=database \
    MYSQL_USER=santiago \
    MYSQL_PASSWORD=12345 \
    MYSQL_ROOT_PASSWORD=12345