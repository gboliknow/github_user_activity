# GitHub User Activity CLI

A command-line interface tool to fetch and display GitHub user information, including profiles, activities, and repositories.

## Features

- 🔍 Fetch user profiles
- 📊 Display recent GitHub activities
- 📚 List user repositories
- 💾 Activity caching for faster subsequent queries
- 📋 Clean, tabulated output format

## Installation

1. Make sure you have Go installed on your system
2. Clone the repository:
```bash
git clone https://github.com/gboliknow/github_user_activity.git
cd github_user_activity
```

3. Build the project:
```bash
go build
```

## Usage

### Profile Command
Fetch and display a GitHub user's profile information:

```bash
go run main.go profile <username>
```

Example:
```bash
go run main.go profile gboliknow
```

Output includes:
- Username
- Name
- Bio
- Company
- Location
- Email
- Website
- Public repository count
- Follower and following counts
- Account creation date

### Activity Command
View a user's recent GitHub activities:

```bash
go run main.go activity <username>
```

Example:
```bash
go run main.go activity gboliknow
```

The output shows:
- Activity type (Push, Pull Request, etc.)
- Creation date
- Repository name
- Actor
- Public/Private status

### Repository Command
List all repositories for a user:

```bash
go run main.go repo <username>
```

Example:
```bash
go run main.go repo gboliknow
```

The output includes:
- Repository name
- Full name
- Description
- Fork status

### Help Commands
Get help for any command:

```bash
# General help
go run main.go --help

# Command-specific help
go run main.go activity --help
go run main.go profile --help
go run main.go repo --help
```

## Caching

The activity command implements caching to improve performance:
- Activity data is cached locally
- Cache duration: 1 minute
- Cache location: System temporary directory
- Cache is automatically invalidated after expiration

## Development

### Running Tests

To run all tests:
```bash
go test ./... -v
```

The test suite includes:
- Activity formatting tests
- Cache operation tests
- API fetching tests

### Project Structure
```
github_user_activity/
├── cmd/
│   ├── activity.go       # Activity command implementation
│   ├── cache_activity.go # Caching functionality
│   ├── profile.go        # Profile command implementation
│   ├── repo.go          # Repository command implementation
│   ├── root.go          # Root command configuration
│   └── types.go         # Data structures
├── cache/               # Cache storage directory
├── main.go             # Entry point
└── README.md           # Documentation
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
