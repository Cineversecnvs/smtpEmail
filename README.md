# Go Email Utility

A simple command-line utility for sending emails using SMTP in Go. This tool allows you to send emails directly from the command line with customizable parameters.

## Project Structure

```
go-email-utility
├── main.go               # Complete email utility implementation
└── README.md             # Project documentation
|__ go.mod                # Minimum GO version
|__ smtpEmailMac          # Executable file for mac
|__ smtpEmailWindows.exe  # Executable file for windows
|__ smtpEmailLinux        # Executable file for linux 
```

## Features

- Send emails via SMTP with TLS encryption
- Command-line interface with flags
- Support for custom SMTP servers
- Configurable email content (subject, body)
- Error handling and validation

## Getting Started

### Prerequisites

- Go 1.16 or later
- An SMTP server (e.g., Gmail, Outlook, SendGrid, etc.)

### Installation

1. Clone or download the project:

   ```bash
   git clone <repository-url>
   cd go-email-utility
   ```

2. Initialize the Go module (if not already done):

   ```bash
   go mod init go-email-utility
   ```

## Usage

Run the utility with command-line flags:

```bash
go run main.go -server <smtp-server> -username <email> -password <password> -from <sender> -to <recipient> [options]
```

### Required Parameters

- `-host`: SMTP server address (e.g., smtp.gmail.com)
- `-username`: SMTP username (usually your email address)
- `-password`: SMTP password or app password
- `-from`: Sender email address
- `-to`: Recipient email address

### Optional Parameters

- `-port`: SMTP port (default: 587)
- `-subject`: Email subject (default: "Test Email")
- `-body`: Email body content (default: "This is a test email.")

### Examples

```bash
smtpEmail \
  -host smtp.gmail.com \
  -port 587 \
  -username your@gmail.com \
  -password your-app-password \
  -from your@gmail.com \
  -to recipient@example.com \
  -subject "Hello from Go!" \
  -body "This email was sent using Go!"
```
### OS Specific

Windows:

```bash
- Download the smtpEmailWindows.exe file
- run the below command
.\smtpEmailWindows.exe 
  -host smtp.gmail.com 
  -port 587 
  -username your@gmail.com 
  -password your-app-password 
  -from your@gmail.com 
  -to recipient@example.com 
  -subject "Hello from Go!" 
  -body "This email was sent using Go!"
- You will get this message "Email sent successfully!"
- Now you can check the receipent mailbox
```

Linux:
<<<<<<< HEAD

```bash
- Download the smtpEmailLinux file
=======
- Install the smtpEmailLinux file 
>>>>>>> 2d9ba7636b8a90d667821036262a432e611b866d
- run the below command
./smtpEmailLinux \
  -host smtp.gmail.com \
  -port 587 \
  -username your@gmail.com \
  -password your-app-password \
  -from your@gmail.com \
  -to recipient@example.com \
  -subject "Hello from Go!" \
  -body "This email was sent using Go!"
- You will get this message "Email sent successfully!"
- Now you can check the receipent mailbox
```

Mac:

```bash
- Download the smtpEmailMac file
- run the below command
./smtpEmailMac \
  -host smtp.gmail.com \
  -port 587 \
  -username your@gmail.com \
  -password your-app-password \
  -from your@gmail.com \
  -to recipient@example.com \
  -subject "Hello from Go!" \
  -body "This email was sent using Go!"
- You will get this message "Email sent successfully!"
- Now you can check the receipent mailbox
```

