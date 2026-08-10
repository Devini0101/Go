Go Projects Portfolio
Welcome to my repository containing my learning journey and projects developed using Go (Golang). Here you will find applications ranging from CLI tools and background service monitors to full-stack web applications following software architecture patterns.

-- Projects Included
1. Website Logger (logger)
A robust command-line monitoring system built in Go designed to track the uptime and availability of target websites and log this information.

Key Features:
 - Interactive CLI menu supporting live monitoring loops and log generation.
 - File-based configuration (sites.txt) to dynamically load URLs to monitor.
 - Automated periodic testing with configurable retry limits and delays (time.Sleep).
 - Real-time HTTP status checking (net/http) and persistent audit logging (log.txt) with timestamps.
 - Log viewer functionality reading directly from the generated system history.


2. E-Commerce CRUD (loja)
A full-stack web application built using the MVC (Model-View-Controller) architecture pattern. This project serves as a complete product management system (CRUD) interacting with a relational database.

Key Features:
 - Clean architectural separation (Controllers, Models, and Views).
 - Server-side HTML rendering using Go's html/template package.
 - HTTP routing and request handling.
 - Database persistence layer integrated with PostgreSQL (database/sql and lib/pq).

 -- Technologies & Concepts Applied
 - Language: Go (Golang)

Web & Routing: net/http, HTML Templates

Database & Drivers: PostgreSQL, SQL queries, connection pooling

Architecture: MVC (Model-View-Controller)

Concurrency: Goroutines and Channels

 -- Getting Started
To run or explore any of the projects in this repository:

 - Clone the repository:

-- bash code: git clone https://github.com/Devini0101/Go.git

 - Navigate to the desired project folder:

-- bash code: cd project-folder-name

Run the application:
-- bash code: go run main.go
-- if dont have the needed dependencies, like go on your environment, please download and set it up before trying it out
