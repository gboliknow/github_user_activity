
# GitHub User Activity CLI

A feature-rich command-line interface tool to fetch and display GitHub user information, including profiles, activities, and repositories.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Basic Commands](#basic-commands)
  - [Interactive Mode](#interactive-mode)
  - [Output Formats](#output-formats)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- 🔍 Fetch user profiles with detailed information
- 📊 Display recent GitHub activities with filtering
- 📚 List and explore user repositories
- � Interactive mode for easy navigation
- 🎨 Colored output for better readability
- 💾 Smart caching for faster responses
- ⚡ Rate limit handling for GitHub API
- ⚙️ Configurable settings and preferences
- 🌟 View user issues and starred repositories

## Installation

1. Ensure you have Go installed on your system
2. Clone the repository:
```bash
git clone https://github.com/gboliknow/github_user_activity.git
cd github_user_activity
```

3. Build the project:
```bash
# Using make
make build

# Or using go
go build
```

## Configuration

Set up your preferences using the config command:

```bash
# Set GitHub token
go run main.go config set --token "your_github_token"

# Set default user
go run main.go config set --user "your_username"

# Set output format
go run main.go config set --format "json"

# View current settings
go run main.go config view
```

## Usage

### Basic Commands

1. **Profile Command**:
```bash
# View user profile
go run main.go profile gboliknow

# With different output format
go run main.go profile gboliknow -o json
```

2. **Activity Command**:
```bash
# View recent activities
go run main.go activity gboliknow

# Filter by activity type
go run main.go activity gboliknow -t PushEvent

# Limit number of results
go run main.go activity gboliknow -l 5

# Use GitHub token
go run main.go activity gboliknow -k "your_token"
```

3. **Repository Command**:
```bash
# List repositories
go run main.go repo gboliknow
```

### Interactive Mode

Launch the interactive interface:
```bash
go run main.go interactive
```

Features available in interactive mode:
- View profile information
- Check recent activities
- List repositories
- View issues
- Check starred repositories
- Easy navigation with arrow keys

### Output Formats

Supported output formats:
- **Table** (default): Colored, well-formatted tables
- **JSON**: Machine-readable JSON format

Specify format using the -o flag:
```bash
go run main.go profile gboliknow -o json
```

## Development

### Project Structure
```
github_user_activity/
├── cmd/
│   ├── activity.go       # Activity command
│   ├── cache_activity.go # Caching system
│   ├── colors.go        # Color formatting
│   ├── config.go        # Configuration
│   ├── config_cmd.go    # Config command
│   ├── interactive.go   # Interactive mode
│   ├── interactive_cmd.go # Interactive command
│   ├── output.go        # Output formatting
│   ├── profile.go       # Profile command
│   ├── rate_limit.go    # Rate limiting
│   ├── repo.go          # Repository command
│   ├── root.go          # Root command
│   ├── stars_issues.go  # Stars and issues
│   └── types.go         # Data structures
├── main.go              # Entry point
├── Makefile            # Build commands
└── README.md           # Documentation
```

### Available Make Commands
```bash
make build          # Build the binary
make test          # Run tests
make clean         # Clean build files
make test-coverage # Run tests with coverage
make deps          # Update dependencies
make install       # Install the binary
make run          # Run the application
make help         # Show all commands
```

### Running Tests
```bash
# Run all tests
go test ./... -v

# Run specific tests
go test -v ./cmd -run TestFetchUserActivity

# Run tests with coverage
make test-coverage
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) for the CLI framework
- [Go](https://golang.org/) for the programming language
- [promptui](https://github.com/manifoldco/promptui) for interactive prompts
- [fatih/color](https://github.com/fatih/color) for colored output
- Project inspired by [roadmap.sh projects](https://roadmap.sh/projects/github-user-activity)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

