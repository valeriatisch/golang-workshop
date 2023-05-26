1. Retrieve all users (GET):
```bash
curl -X GET http://localhost:8080/users
```

2. Retrieve a user by ID (GET):
```bash
curl -X GET http://localhost:8080/users/1
```

3. Create a new user (POST):
```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": 4, "username": "user4", "email": "user4@example.com", "age": 35}' http://localhost:8080/users
```

4. Delete a user by ID (DELETE):
```bash
curl -X DELETE http://localhost:8080/users/4
```
