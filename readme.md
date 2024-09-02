---

# Music API - A Learning Tool for Frontend Developers

Welcome to the **Music API**! This project is designed to help frontend developers, especially beginners, to practice and enhance their skills in working with REST APIs. The API provides essential features such as user authentication, musics management, and search functionality.

## Features

- **User Authentication**: 
  - Register a new account.
  - Login to an existing account.

- **Music Management**:
  - Upload your musics.
  - Get your uploaded musics.
  - Edit your uploaded musics.
  - Delete your musics.

- **Search and Discovery**:
  - Search for other users by username.
  - View other users' musics.

## Getting Started

### Prerequisites

Make sure you have the following installed on your system:

- **Go** (latest version recommended)
- **Git**
- **PostgreSQL** (or any other database you prefer)
- **GORM** for database management
- **JWT** for token-based authentication

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/rrraf1/soundshare
   cd musics-api
   ```

2. **Create a `.env` file**:

   Create a `.env` file in the root directory of the project and add the following environment variables:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_database_user
   DB_PASSWORD=your_database_password
   DB_NAME=your_database_name
   JWT_SECRET=your_jwt_secret_key
   ```

3. **Install dependencies**:

   Use `go mod` to install the required dependencies:

   ```bash
   go mod tidy
   ```

4. **Setup the database**:

   Ensure your PostgreSQL server is running and create the required database.

   ```sql
   CREATE DATABASE your_database_name;
   ```

   Migrate the database using GORM in the Go project.

5. **Run the application**:

   Start the server by running:

   ```bash
   go run .
   ```

   The server will start on `http://localhost:3000`.

### API Endpoints

Here’s a quick overview of the available endpoints:

- **Authentication**:
  - `POST /register` - Register a new user.
  - `POST /login` - Login an existing user.

- **Music Management (Need a JWT token)**:
  - `POST /api/musics` - Upload a new music file.
  - `GET /api/musics` - Get all your uploaded music.
  - `PUT /api/musics/:id` - Edit a specific musics file.
  - `DELETE /api/musics/:id` - Delete a specific musics file.
  - `GET /api/musics/details` - Get a spesific musics.

- **Search and Discovery**:
  - `GET /api/users/` - View musics uploaded by a specific user.

### Usage

1. **Register**:
   
   Send a `POST` request to `/register` with the key your username and password.

2. **Login**:
   
   Send a `POST` request to `/login` to receive a JWT token, with the key your username and password.

3. **Upload Music**:

   Use the JWT token to authenticate and send a `POST` request to `/api/musics` with your musics file. The key is MusicName, artist, genre, and link

4. **Get Your Music**:

   Send a `GET` request to `/api/musics` to retrieve your uploaded musics.

5. **Delete musics**:

   Send a `POST` request to `/api/musics/:id` to delete the music that already posted.

6. **Update musics**:

   Send a `PUT` request to `/api/musics/:id` to update the music detail, the key is MusicName, artist, genre, and link

7. **Search Users**:

   Send a `GET` request to `/api/users` to find other users and their musics. The key is username.

8. **Search Musics**:

   Send a `GET` request to `/api/musics/details` to find the details of the music, the key is MusicName

### License

This project is licensed under the MIT License.

---
