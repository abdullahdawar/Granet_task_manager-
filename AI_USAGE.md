# AI Usage

AI was used as a development assistant throughout this project, mainly for guidance, implementation support, debugging, and reviewing decisions. I worked through the project step by step rather than asking AI to generate the entire application from scratch.

My approach was to understand what needed to be built, work on each part individually, use AI when I needed guidance or help with implementation, and then manually run and verify the result before moving forward.

## Key Prompts and Areas Where AI Helped

### 1. Understanding and Planning

"Explain each part of the assignment to me from the basics. What should be done first?"

I first used AI to break the assignment into manageable parts and understand how React, Go, and MySQL would work together.

From there, I decided to build the project in stages: database, backend architecture, CRUD API, frontend integration, testing, and finally UI improvements.

### 2. Backend Structure

"How should I structure the Go backend using handler, service, repository, model, and database layers?"

AI helped explain the purpose of each layer and provided guidance while I implemented the structure.

I worked through the layers individually and tested the backend as it developed instead of generating the whole backend in one step.

### 3. Database and Repository

"How should the Go application connect to MySQL and how should the repository layer handle the tasks?"

AI helped with SQL and Go syntax where needed.

I created and configured the MySQL database locally, set the environment variables, ran the SQL schema, started the server, and verified that data was actually being stored and retrieved.

### 4. REST API

"How should I implement the create, list, update and delete task endpoints?"

AI assisted with parts of the handler and service implementation and explained the request flow.

I manually tested each endpoint and verified the HTTP requests, responses, and database changes before continuing.

### 5. React and TypeScript

"How should I connect my React TypeScript frontend to the Go API?"

AI helped with the API service structure, TypeScript interfaces, and some React implementation details.

I then ran the frontend, tested it against the real backend, created tasks, edited them, deleted them, refreshed the page, and checked that the UI reflected the database correctly.

### 6. Debugging

"Why is my empty task list returning null instead of []?"

Testing revealed that an empty database returned `null` rather than an empty JSON array.

The repository implementation was changed to initialize an empty task slice. I restarted the backend and tested the endpoint again to verify that it now returned:

    []

This was one example where testing the actual behavior led to a change in the implementation.

### 7. UI Improvements

"Can we make the frontend more appealing with animations?"

"Take it up a notch and change the colours from white and purple to navy blue and golden."

Once the functionality was working, I used AI to help refine the CSS and animations.

I chose the final navy-and-gold direction, reviewed the result in the browser, tested the responsive layout and interactions, and kept the design once I was satisfied with it.

## My Contribution

My work on the project included:

- Breaking the assignment into implementation stages.
- Setting up the project and local development environment.
- Installing and configuring Go, Node.js, npm, MySQL, and Git.
- Creating and configuring the MySQL database.
- Working through the backend structure layer by layer.
- Configuring database environment variables rather than hard-coding credentials.
- Running and testing the Go server.
- Testing CRUD operations directly against the API.
- Connecting and testing the React frontend against the backend.
- Testing create, list, edit, and delete operations from the browser.
- Verifying that tasks persisted in MySQL.
- Running Go formatting and build/test checks.
- Running the React production build.
- Finding issues through testing and applying fixes.
- Choosing and refining the final frontend design.
- Reviewing the architecture so I could understand and explain the request flow.

## How AI Contributed

AI was particularly useful for:

- Explaining unfamiliar Go and React concepts.
- Suggesting implementation approaches.
- Providing code snippets for individual parts of the application.
- Helping debug errors and unexpected behavior.
- Reviewing the project structure.
- Suggesting improvements to validation and error handling.
- Assisting with CSS and animations.
- Explaining how the complete application works across its different layers.

## Development Approach

The project was developed iteratively.

For example, the backend was not generated as one finished application. I first established the database connection, verified it, then worked on the repository, service, and handler layers. After the API was working, I tested CRUD operations before connecting the React frontend.

Similarly, the frontend was first made functional and tested against the API. Visual improvements and animations were added only after the core application worked correctly.

This approach allowed me to understand each part and catch problems as they appeared.

## What I Learned

The project helped strengthen my understanding of:

- React state and form handling.
- TypeScript interfaces.
- REST API design.
- HTTP methods and status codes.
- Go HTTP handlers.
- Service and repository patterns.
- SQL and MySQL persistence.
- Environment variables.
- CORS.
- Frontend/backend integration.
- Debugging a full-stack request across multiple layers.

## Request Flow

The main application flow is:

React Frontend
→ HTTP Request
→ Go Handler
→ Service Layer
→ Repository Layer
→ MySQL

The result then travels back through:

MySQL
→ Repository
→ Service
→ Handler
→ JSON Response
→ React Frontend

AI helped me work through and understand individual parts of this process, while I remained involved in the setup, decisions, implementation process, testing, debugging, verification, and final presentation of the application.
