# Task Management Application

A full-stack Task Management application built for the Granet Technologies AI-Assisted Full Stack Developer Assignment.

The application allows users to create, view, update, and delete tasks. Each task contains a title, description, due date, and status.

## Tech Stack

Frontend:
- React
- TypeScript
- Vite
- CSS

Backend:
- Go (Golang)
- REST API
- MySQL
- CORS middleware

## Architecture

The application uses a layered architecture:

React Frontend
→ REST API / Handler
→ Service Layer
→ Repository Layer
→ MySQL Database

### Handler Layer
Handles HTTP requests and responses.

### Service Layer
Contains application logic and validation.

### Repository Layer
Handles SQL queries and communication with MySQL.

### Database Layer
Creates and manages the MySQL connection using environment variables.

## Features

- Create tasks
- View tasks
- Update tasks
- Delete tasks
- Task status management
- Form validation
- Error handling
- Responsive interface
- Animated navy and gold UI
- MySQL persistence
- Modular backend structure

## API Endpoints

GET /api/tasks - Get all tasks

POST /api/tasks - Create a task

PUT /api/tasks/{id} - Update a task

DELETE /api/tasks/{id} - Delete a task

## Database Setup

Start MySQL and run:

    mysql -u root -p < backend/schema.sql

This creates the task_manager database and tasks table.

## Backend Setup

Move into the backend directory:

    cd backend

Set the database environment variables:

    export DB_USER=root
    export DB_PASSWORD="your_mysql_password"
    export DB_HOST=127.0.0.1
    export DB_PORT=3306
    export DB_NAME=task_manager

Start the backend:

    go run ./cmd/api

The API runs at http://localhost:8080.

## Frontend Setup

Open another terminal and move into the frontend directory:

    cd frontend

Install dependencies:

    npm install

Start the frontend:

    npm run dev

The application runs at http://localhost:5173.

## Testing

The following functionality was manually tested through the REST API and browser:

- Create task
- List tasks
- Update task
- Delete task
- MySQL persistence
- Frontend/backend communication
- Form validation

Check the Go backend with:

    go test ./...

Check the frontend production build with:

    npm run build

## Security

Database credentials are not stored directly in the source code. The backend reads database configuration from environment variables.

Local environment files, dependencies, and generated build files are excluded using .gitignore.

## AI-Assisted Development

AI was used for architecture planning, implementation assistance, debugging, UI improvements, and explanations.

Generated code was reviewed and tested before being included. Manual verification included compiling the backend, building the frontend, testing CRUD operations, checking MySQL persistence, and fixing issues discovered during testing.

See AI_USAGE.md for the key prompts and a detailed explanation of AI-assisted development.

## Future Improvements

- Authentication
- Task filtering and search
- Pagination
- Automated tests
- Deployment configuration
- Dedicated production database user

## Author

Abdullah Dawar
