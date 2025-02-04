# Ritik's Redis & Meet Manager

![GitHub stars](https://img.shields.io/github/stars/ritikranjan12/ritik-redis-meet-manager)
![GitHub views](https://komarev.com/ghpvc/?username=ritikranjan12&repo=ritik-redis-meet-manager)

Welcome to **Ritik's Redis & Meet Manager**! This project is a command-line tool designed to help you manage Redis servers, join Google Meet sessions, open VS Code projects, and more, all from a single interface.

## Features

- **Change Node.js Version**: Easily switch between different Node.js versions.
- **Start/Stop Redis Server**: Manage your Redis server with simple commands.
- **Join Google Meet**: Quickly join predefined Google Meet sessions.
- **Open Telegram**: Open Telegram or a specific chat directly from the command line.
- **Open VS Code Projects**: Launch your VS Code projects with predefined commands.
- **Increment Version**: Automatically increment the version of your project (major, minor, patch).
- **Build EXE**: Build an executable file and save your Ubuntu password securely.
- **Change Saved Ubuntu Password**: Update your saved Ubuntu password.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/ritik-redis-meet-manager.git
    cd ritik-redis-meet-manager
    ```

2. **Install dependencies**:
    ```sh
    go mod tidy
    ```

3. **Build the executable**:
    ```sh
    go build -o ritik_redis_manager_v3.0.0.exe
    ```

## Usage

Run the executable to start the command-line interface:
```sh
./ritik_redis_manager_v3.0.0.exe
```

Follow the on-screen instructions to navigate through the menu and use the various features.

### Configuration

- **VS Code Projects**: Update the `projects` variable in `main.go` with your project paths and commands.
- **Google Meet Links**: Update the `misikiMeetLink` and `sparkinityMeetLink` constants with your Google Meet URLs.

### Example

Here's an example of how to use the tool:

1. Start the tool:
    ```sh
    ./ritik_redis_manager_v3.0.0.exe
    ```

2. Choose an option from the menu:

    📌 **Main Menu**:
    ```
    [1] 💻 Change Node.js Version
    [2] 🔴 Start Redis Server
    [3] 🟢 Stop Redis Server
    [4] 📞 Join Google Meet 1
    [5] 📞 Join Google Meet 2
    [6] 📲 Open Telegram
    [7] 💬 Open a Specific Telegram Chat
    [8] 💻 Open VS Code Project
    [9] 🚀 Increment Version (Major, Minor, Patch)
    [10] 🛠️ Build EXE (Save Ubuntu Password)
    [11] 🔑 Change Saved Ubuntu Password
    [12] 🚪 Exit
    ```

Follow the prompts to complete your task.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Contact

For any questions or suggestions, please contact Ritik Ranjan at `ritik123453@gmail.com`.