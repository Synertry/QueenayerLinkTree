# QueenayerLinkTree

A web application backend that serves a customizable link tree for queenayer.me, allowing users to access various resources through a single page.

## Features

- **Web Server**: Serves a static website with a link tree
- **Self-Update**: Ability to update the application to the latest version
- **Version Information**: Display version information of the application

## Installation

### Prerequisites

- Go 1.24.2 or higher
- Node.js and npm (for web development)

### Building from Source

1. Clone the repository:
   ```
   git clone https://github.com/Synertry/QueenayerLinkTree.git
   cd QueenayerLinkTree
   ```

2. Build the application:
   ```
   make build
   ```

   This will compile the Go application and build the web assets.

3. Run the application:
   ```
   ./bin/windows-amd64/QueenayerLinkTree serve
   ```

### Using Docker

1. Build and run using Docker Compose:
   ```
   docker-compose up -d
   ```

## Usage

### Serving the Website

```
QueenayerLinkTree serve [--port|-p PORT]
```

Options:
- `--port, -p`: Port to serve the website on (default: 8080)

### Updating the Application

```
QueenayerLinkTree update PATH_TO_NEW_BINARY [--patch|-p]
```

Arguments:
- `PATH_TO_NEW_BINARY`: Path to the new binary file

Options:
- `--patch, -p`: Use patch file to update binary

### Displaying Version Information

```
QueenayerLinkTree version
```

## Development

### Project Structure

- `cmd/`: Command-line interface definitions
- `internal/app/`: Application logic
  - `SelfUpdate/`: Self-update functionality
  - `Website/`: Web server functionality
- `internal/pkg/`: Shared packages
- `web/`: Web assets and frontend code
- `build/`: Build-related files
- `scripts/`: Build and utility scripts

### Building the Web Assets

```
cd web
npm install
npm run build
```

## License

Distributed under the Boost Software License, Version 1.0. See `LICENSE` for more information.