## Ecommerce Backend in Go (WIP)
This repository contains a simple example of a backend for an Ecommerce Application with User Authentication with JWT in Go.

### Installation

Before running the project, ensure that you have the following tools installed on your machine:

- **Migrate**: For managing database migrations.

### Running the Project

Firstly, ensure you have a MySQL database running on your machine. You can swap it for any storage you prefer under `/db`.

Then, create a database with the desired name (default is `ecom`) and run the migrations:

```bash
make migrate-up
```

### TODOS
1. [ ] Update Dockerfile and setup.
2. [ ] Update Postman collection.
3. [ ] Update Tests.
4. [ ] Implement Refresh Token in JWT
5. [ ] Implement ATOMIC design of Types
6. [ ] Add User Order History
7. [ ] Add User Address feature
8. [ ] Implement Cancel Order for Items that are not shipped
9. [ ] Perform optimizations.
