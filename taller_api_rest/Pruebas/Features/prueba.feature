Feature: Gestionar usuarios haciendo uso de los métodos proporcionados por la API

    Scenario: Creación de un nuevo usuario
    Given Que se proporcionen los siguientes datos sobre el nuevo usuario:
             | username | contraseña | email            |
             | user1    | password1  | user1@example.com |
    And se proporciona un token JWT válido
    When Se realiza una solicitud de tipo POST a /users con los datos del nuevo usuario y el token JWT válido
    Then La respuesta contiene un mensaje de confirmación y un codigo de estado 201

    Scenario: Obtener todos los usuarios registrados 
    Given Que se proporcione un token JWT válido de un usuario autenticado en el sistema
    When se realiza la solicitud tipo GET a /users, adicionando el token JWT a la cabecera de autenticación de la solicitud
    Then La respuesta contiene un listado de los usuarios y un codigo de estado 200

    Scenario: Obtener un usuario de acuerdo a su ID
    Given Que se proporcione un ID de un usuario registrado en el sistema
    And se proporciona un token JWT válido 
    When Se realiza una solicitud tipo GET a /users/{ID}
    Then la respuesta contiene la información del usuario solicitado y un codigo de estado 200

    Scenario:Actualizar los datos de un usuario de acuerdo a su ID
    Given Que se proporcionen algunos de los siguientes datos sobre el usuario que se quiere actualizar:
             | username | contraseña | email            |
             | user2    | password2  | user2@example.com |
    