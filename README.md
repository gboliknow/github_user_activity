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
