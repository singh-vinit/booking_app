# Booking App CLI

A simple command-line application written in Go for booking tickets to the Go Conference. This app demonstrates basic Go concepts and uses goroutines for concurrent execution.

## Features

- Book tickets for a conference via the command line
- Collects user details (first name, last name, email, number of tickets)
- Displays remaining tickets and all bookings
- Uses goroutines to simulate concurrent ticket processing

## Usage

1. **Build the app:**
   ```sh
   go build .
   ```
2. **Run the app:**
   ```sh
   go run .
   ```
3. ** Run the executable file ./file_name.exe **

4. **Follow the prompts** to enter your name, email, and number of tickets to book.

## Requirements

- Go 1.18 or higher
- Windows, macOS, or Linux

## Concurrency

This app uses goroutines to handle ticket booking and email confirmation concurrently, providing a basic example of Go's concurrency model.

## License

MIT
