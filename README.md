# Employee Dashboard Project

This project is an employee dashboard developed using the Go programming language for the REST API and MongoDB for database connectivity. The front-end interface is built using React and NextJS.

## Table of Contents

- [Technologies](#technologies)
- [Setup](#setup)
- [API Routes](#api-routes)
- [Contributing](#contributing)
- [License](#license)

## Technologies

The following technologies were used in this project:

- Go
- MongoDB
- React
- NextJS

## Setup

### Prerequisites

- Go
- MongoDB
- Node.js and npm

### Installation

1.  **Clone the repository**

    ```sh
    git clone https://github.com/alexdang1993374/employee-performance.git
    ```

2.  **Setup Go and MongoDB**

    Make sure you have Go installed on your machine and running properly.

    3. **Setup the backend**

    Navigate to the backend directory, install the dependencies and start the server.

    ```sh
    cd backend
    go get
    ```

3.  **Setup environment variables**

    In the backend directory, create a .env file with the MongoDB URI:

    ```sh
    cd backend
    touch .env
    ```

    Open the .env file and add the following line:

    ```sh
    MONGODB_URI=your_mongodb_uri
    ```

    Replace your_mongodb_uri with your actual MongoDB URI.

    Start the server

    ```sh
    go run main.go
    ```

    The server should be running on [http://localhost:5001](http://localhost:5001)

    ```

    ```

4.  **Setup the frontend**

    Open a new terminal and navigate to the frontend directory, install the dependencies and start the server.

    ```sh
    cd frontend
    npm install
    npm run dev
    ```

    The frontend should be running on [http://localhost:3000](http://localhost:3000)

## API Routes

The Go API supports the following routes:

- **GET** `/employees`: Retrieve a list of all employees
- **POST** `/employees`: Add a new employee
