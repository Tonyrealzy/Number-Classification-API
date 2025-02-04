# Number Classification API

#### https://number-classification-api-twv8.onrender.com/api/classify-number?number=<your-number>

This is a simple API built with Go (Golang) that classifies a given number and provides various properties, including:

- Whether the number is prime
- Whether the number is perfect
- Whether the number is an Armstrong number
- Whether the number is even or odd
- The sum of its digits
- A fun fact fetched from an external API

## Features

- Provides multiple classifications for a number.
- Fetches a fun fact about the number from the Numbers API.
- Returns data in JSON format.
- Handles errors gracefully (e.g., invalid input).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the API](#running-the-api)
- [Example Requests and Responses](#example-requests-and-responses)

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher.
- An internet connection to access the Numbers API for fun facts.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/number-classifier-api.git
   ```

2. Navigate into the project directory:

```bash
cd number-classifier-api

```

3. Initialize Go modules (if you haven't already):

```bash
go mod init number-classifier-api

```

4. Install dependencies:

```bash
go mod tidy

```

5. Running the API
   Run the server using the following command:

```bash
go run main.go

```

The server will start and listen on port 8080 by default.

You can access the API by sending requests to:

```bash
http://localhost:8080/api/classify-number?number=<your-number>

```

#### Example

For example, to classify the number 371:

```bash "http://localhost:8080/api/classify-number?number=371"

```

#### Example Request

```bash
GET /api/classify-number?number=371

```

#### Example Response (200 OK)

{
"number": 371,
"is_prime": false,
"is_perfect": false,
"properties": ["armstrong", "odd"],
"digit_sum": 11,
"fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}

#### Example Response (400 Bad Request)

If an invalid input is provided (e.g., a string instead of a number):
{
"number": "alphabet",
"error": true
}
