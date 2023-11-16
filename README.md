# Go Multi-Platform Builder

## Introduction
This tool is designed to automate the process of building Go programs for multiple platforms and architectures. It reads a configuration file to determine the target platforms and builds your Go script accordingly. This tool is especially useful for developers looking to streamline their build process for cross-platform Go applications.

## Features
- Builds a Go script for multiple target OSs and architectures.
- Uses a configuration file to specify target platforms.
- Option to specify a specific Go file for building.
- Automatically names the output binaries based on the script name and target platform.

## Getting Started

### Prerequisites
- Go programming language installed.
- Basic understanding of Go project structure.

### Installation
1. Clone the repository to your local machine:
   ```
   git clone https://github.com/shadowraven65/gobuilder.git
   ```
2. Or download the binary from the latest release (Only Linux amd64 as of now)
3. Place the binary of the builder script in a directory included in your PATH (e.g., `/usr/local/bin`) for easy access.

### Configuration
Create a `buildconfig.json` file in the root of your Go project with the following structure:
```
{
    "targets": [
        {"os": "linux", "arch": "amd64"},
        {"os": "windows", "arch": "amd64"},
        {"os": "darwin", "arch": "arm64"}
        // Add more targets as needed
    ]
}
```

### Usage
Run the builder script from the root of your Go project:
```
gobuilder
```

To build a specific Go file if there is more than 1:
```
gobuilder -file yourfile.go
```
Keep in mind you will still build for all in the json  
## Contributing
Contributions to improve this tool are welcome. Please feel free to fork the repository, make your changes, and submit a pull request.

## License
MIT

