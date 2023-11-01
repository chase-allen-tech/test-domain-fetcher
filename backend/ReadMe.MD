# Domain Status Fetcher in Golang #
## Overview ##
This application is designed to provide status code with the provided domain in request path.

## Getting Started ##
### Prerequisites ###
Go 1.15 or newer

### Running the Application ###
1. Clone this repository to your local machine.
2. Navigate to the root directory of the application in your terminal.
3. Run `go mod tidy` to download the necessary Go modules.
4. Run `go run main.go` to start the application.

The server will start and listen on localhost:8000. You can change the port by modifying the PORT constant in main.go.

### Using the Application ###
Make a `GET` request to `localhost:8000/v1` with a proper path

The application will respond with a JSON object containing status information

### Sample Response ###
{
	"url": "https://www.google.com",
	"statusCode": 200,
	"duration": 154,
	"date": 1606663220
}

### Notes ###
Troubleshooting
If you encounter any errors, the application will log relevant error messages to help you troubleshoot.

### Future Improvements ###
