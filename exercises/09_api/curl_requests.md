**APIs** (Application Programming Interfaces) are part of the backend and act as intermediaries between databases storing data and the frontend which interacts with users.
This communication is done via standard HTTP methods like GET, POST, and DELETE.

- **GET**: This is used to retrieve data. The information is retrieved in JSON (JavaScript Object Notation) format, a common data format that's human and machine-readable.
- **POST**: This method is used to send data to the server. The request body contains JSON-formatted.
- **DELETE**: This method is used to remove data from the server.

**HTTP status codes** are included in the response to indicate whether an HTTP request has been successful.
Common ones are 200 (OK), 201 (Created), 204 (No Content), 400 (Bad Request), 404 (Not Found), and 500 (Internal Server Error).

1. **Retrieving all users (GET)**:

```bash
curl -X GET http://localhost:8080/users
```

2. **Retrieving a specific user by their ID (GET)**:

```bash
curl -X GET http://localhost:8080/users/1
```

3. **Creating a new user (POST)**: By sending JSON formatted data as part of the request body. Header is used to indicate the JSON content type.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 4, "username": "user4", "email": "user4@example.com", "age": 35}' http://localhost:8080/users
```

4. **Deleting a user by their ID (DELETE)**:

```bash
curl -X DELETE http://localhost:8080/users/4
```
