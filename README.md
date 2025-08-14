
# GitHub User Activity CLI

A command-line interface tool to fetch and display GitHub user information, including profiles, activities, and repositories.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Make Commands](#make-commands)
- [Usage](#usage)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

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

3. Build the project (choose one method):

```bash
# Using go build
go build

# Using make
make build
```

## Make Commands

The project includes a Makefile with several useful commands:

```bash
make build           # Build the binary (output in bin/)
make clean          # Clean build files
make test           # Run tests
make test-coverage  # Run tests with coverage report
make deps           # Update dependencies
make install        # Install the binary
make run            # Run the application
make help           # Show all available commands
```

Example usage with make:
```bash
# Build and run
make build
./bin/github_user_activity profile gboliknow

# Run directly (for development)
make run profile gboliknow
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
=======
github_user_activity CLI

github_user_activity is a CLI tool built in Go for fetching and displaying GitHub user data. It supports retrieving user activity, profile information, and repositories, and includes caching functionality for improved performance.
Features

    Fetch User Activity: Get recent activity for a specified GitHub user.
    Fetch User Profile: Retrieve and display user profile information.
    Fetch User Repositories: List all repositories for a specified GitHub user.
    Caching: Save and retrieve data using file-based caching to reduce API calls and improve performance.

Installation

    Clone the repository:

    bash

git clone https://github.com/gboliknow/github_user_activity.git

Navigate to the project directory:

bash

cd github_user_activity

Build the CLI tool:

bash

    make build

    The executable will be located in the bin directory.

Usage
Fetch User Activity

bash

./bin/github_user_activity activity [username] --type [activityType]

    [username]: GitHub username.
    --type [activityType]: Optional filter for activity type (e.g., "PushEvent").

Fetch User Profile

bash

./bin/github_user_activity profile [username]

    [username]: GitHub username.

Fetch User Repositories

bash

./bin/github_user_activity repo [username]

    [username]: GitHub username.

Commands

    activity [username]: Fetches user activity.
    profile [username]: Fetches user profile information.
    repo [username]: Fetches user repositories.

Caching

The application uses file-based caching to store and retrieve data. Cached data will be saved in the cache directory and will be used to improve performance for subsequent requests.
Development

To contribute to the project:

    Fork the repository.
    Create a new branch (git checkout -b feature/YourFeature).
    Commit your changes (git commit -am 'Add new feature').
    Push to the branch (git push origin feature/YourFeature).
    Create a new Pull Request.

License

This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

    Cobra for the CLI framework.
    Go for the programming language.
https://roadmap.sh/projects/github-user-activity****

