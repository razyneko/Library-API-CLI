# Library API

This is a simple RESTful API built using the [Gin Web Framework](https://github.com/gin-gonic/gin) in Go. The API manages a library system with features to view books, add new books, and check out or return books.

## Features

- View all books
- Get details of a specific book by ID
- Add a new book to the library
- Check out a book (reduce its quantity)
- Return a book (increase its quantity)

## Setup Instructions

Follow these steps to run the API on your local machine:

### Prerequisites

- [Go](https://go.dev/) installed (version 1.19 or higher)
- `git` installed

### Steps

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install Dependencies**
   The project uses the Gin framework. Run the following command to install dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Server**
   Start the API server by running:

   ```bash
   go run main.go
   ```

   The server will start on `localhost:8080`.

## API Endpoints

### 1. Get All Books

**Endpoint:** `GET /books`

**Description:** Fetch a list of all books in the library.

**Curl Command:**

```bash
curl localhost:8080/books
```

### 2. Get Book by ID

**Endpoint:** `GET /books/:id`

**Description:** Fetch details of a specific book by its ID.

**Curl Command:**

```bash
curl localhost:8080/books/<id>
```

Replace `<id>` with the book's ID (e.g., `1`, `2`).

### 3. Add a New Book

**Endpoint:** `POST /books`

**Description:** Add a new book to the library.

**Request Body Format:**

```json
{
  "id": "4",
  "title": "New Book Title",
  "author": "Author Name",
  "quantity": 10
}
```

**Curl Command:**

```bash
curl localhost:8080/books \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '@body.json'
```

### 4. Check Out a Book

**Endpoint:** `PATCH /checkout`

**Description:** Check out a book by reducing its quantity.

**Query Parameter:**

- `id` (required): The ID of the book to check out

**Curl Command:**

```bash
curl localhost:8080/checkout?id=<id> \
    --request "PATCH"
```

Replace `<id>` with the book's ID.

### 5. Return a Book

**Endpoint:** `PATCH /return`

**Description:** Return a book by increasing its quantity.

**Query Parameter:**

- `id` (required): The ID of the book to return

**Curl Command:**

```bash
curl localhost:8080/return?id=<id> \
    --request "PATCH"
```

Replace `<id>` with the book's ID.

## Project Structure

- `main.go`: Contains all the API routes and logic.
- `go.mod`: Manages the project dependencies.

## Example Usage

1. Start the server.
2. Use the curl commands to interact with the API.
3. Test the endpoints to manage your library system.

## Contributing

Feel free to open issues or submit pull requests if you find bugs or want to add new features.

##

