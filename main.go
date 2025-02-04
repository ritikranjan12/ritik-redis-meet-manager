package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

// Version file
const (
	versionFile      = "version.txt"
	passwordFilePath = ".ubuntu_password" // Hidden file for security
)

// Google Meet Links
const (
	misikiMeetLink     = "https://meet.google.com/ssfa-tdxv-ayz"
	sparkinityMeetLink = "https://meet.google.com/sfa-tdxv-ayz"
)

// VS Code Projects List
var projects = map[string]struct {
	Path    string
	Command string
}{
	"Project 1": {Path: "E:\\Project1", Command: "bun dev"},
	"Project 2": {Path: "E:\\Project2", Command: "bun dev"},
	"Project 3": {Path: "E:\\Project3", Command: "bun dev"},
	"Project 4": {Path: "E:\\Project4", Command: "bun dev"},
	"Project 5": {Path: "E:\\Project5", Command: "bun dev"},
}

func readSavedPassword() string {
	data, err := os.ReadFile(passwordFilePath)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

// Save the Ubuntu password securely
func savePassword(password string) {
	os.WriteFile(passwordFilePath, []byte(password), 0600)
}

// Get Ubuntu password (saved or new)
func getUbuntuPassword() string {
	password := readSavedPassword()
	if password != "" {
		return password
	}

	fmt.Print(Bold + Green + "🔑 Enter your Ubuntu password (saved securely): " + Reset)
	fmt.Scanln(&password)
	savePassword(password)
	return password
}

// Read version as a string
func readVersionString() string {
	data, err := os.ReadFile(versionFile)
	if err != nil {
		return "1.0.0"
	}
	return strings.TrimSpace(string(data))
}

func openVSCodeProject() {
	fmt.Println(Bold + Cyan + "\n📂 Available VS Code Projects:" + Reset)
	projectNames := []string{}
	for name := range projects {
		projectNames = append(projectNames, name)
	}

	// Display project list
	for i, name := range projectNames {
		fmt.Printf("%s[%d] %s%s\n", Green, i+1, name, Reset)
	}

	fmt.Print(Bold + "\n👉 Choose a project to open: " + Reset)
	var choice int
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(projectNames) {
		fmt.Println(Red + "❌ Invalid choice." + Reset)
		return
	}

	selectedProject := projects[projectNames[choice-1]]

	// Open VS Code
	fmt.Println(Bold + Yellow + "🚀 Opening VS Code: " + projectNames[choice-1] + Reset)
	cmd := exec.Command("cmd", "/c", "code", selectedProject.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()

	fmt.Print(Bold + "\n🔄 Do you want to wait for VS Code exit? (y/n): " + Reset)
	var waitChoice string
	fmt.Scanln(&waitChoice)

	if strings.ToLower(waitChoice) == "y" {
		cmd.Wait()
		fmt.Println(Green + "✅ VS Code closed." + Reset)
		return
	}

	fmt.Print(Bold + "💻 Do you want to run a command in VS Code terminal? (y/n): " + Reset)
	var terminalChoice string
	fmt.Scanln(&terminalChoice)

	if strings.ToLower(terminalChoice) == "y" {
		// Open VS Code terminal and run the command
		fmt.Println(Bold + Blue + "Running command: " + selectedProject.Command + Reset)

		// Open the VS Code integrated terminal and send the command
		terminalCmd := exec.Command("cmd", "/c", "code", selectedProject.Path, "--new-window", "--command", "workbench.action.terminal.new")
		terminalCmd.Stdout = os.Stdout
		terminalCmd.Stderr = os.Stderr
		terminalCmd.Start()

		// Wait a bit for VS Code to open the terminal
		time.Sleep(3 * time.Second)

		// Now execute the command in the VS Code terminal manually (This will work with a known terminal extension in VS Code)
		execTerminalCmd := exec.Command("cmd", "/c", "code", selectedProject.Path, "--command", "workbench.action.terminal.sendText", "--args", selectedProject.Command)
		execTerminalCmd.Stdout = os.Stdout
		execTerminalCmd.Stderr = os.Stderr
		execTerminalCmd.Run()
	}
}

// Build EXE and prompt for password
func buildExecutable(version string) {
	fmt.Print(Bold + Green + "🔑 Enter your Ubuntu password for EXE (saved securely): " + Reset)
	var password string
	fmt.Scanln(&password)
	savePassword(password)

	fmt.Println(Yellow + "🔧 Building EXE..." + Reset)
	exeName := fmt.Sprintf("ritik_redis_manager_v%s.exe", version)
	cmd := exec.Command("go", "build", "-o", exeName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(Red + "❌ Build failed: " + err.Error() + Reset)
	} else {
		fmt.Println(Green + "✅ Build successful! EXE saved as " + exeName + Reset)
	}
}

func startRedis() error {
	password := getUbuntuPassword()

	fmt.Println(Bold + Yellow + "🚀 Starting Redis server..." + Reset)
	startCmd := fmt.Sprintf("echo %s | sudo -S redis-server --daemonize yes", password)
	cmd := exec.Command("wsl", "-e", "bash", "-c", startCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows clear screen
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayTitle(version string) {
	clearScreen()
	title := fmt.Sprintf(`
██████╗ ██╗████████╗██╗██╗  ██╗ ██████╗      
██╔══██╗██║╚══██╔══╝██║██║  ██║██╔════╝      
██████╔╝██║   ██║   ██║███████║██║  ███╗     
██╔═══╝ ██║   ██║   ██║██╔══██║██║   ██║     
██║     ██║   ██║   ██║██║  ██║╚██████╔╝     
╚═╝     ╚═╝   ╚═╝   ╚═╝╚═╝  ╚═╝ ╚═════╝      
🚀 Ritik Ranjan's Redis & Meet Manager v%s 🚀  
`, version)

	fmt.Println(Yellow + title + Reset)
}

// Stop Redis Server
func stopRedis() error {
	password := getUbuntuPassword()

	fmt.Println(Bold + Yellow + "🛑 Stopping Redis server..." + Reset)
	stopCmd := fmt.Sprintf("echo %s | sudo -S systemctl stop redis", password)
	cmd := exec.Command("wsl", "-e", "bash", "-c", stopCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func changeNodejsVersion() error {
	nodejsVersion := "23.7.0"
	fmt.Print(Bold + Green + "🔑 Enter your Node js Version: " + Reset)
	fmt.Scanln(&nodejsVersion)
	cmd := exec.Command("cmd", "/c", "code", nodejsVersion)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

// Main menu
func main() {
	version := readVersionString()

	for {
		displayTitle(version)
		fmt.Println(Bold + "\n📌 Main Menu:" + Reset)
		options := []string{
			"💻 Change Node js Version",
			"🔴 Start Redis Server",
			"🟢 Stop Redis Server",
			"📞 Join Google Meet 1",
			"📞 Join Google Meet 2",
			"📲 Open Telegram",
			"💬 Open a Specific Telegram Chat",
			"💻 Open VS Code Project",
			"🚀 Increment Version (Major, Minor, Patch)",
			"🛠️ Build EXE (Save Ubuntu Password)",
			"🔑 Change Saved Ubuntu Password",
			"🚪 Exit",
		}

		for i, option := range options {
			fmt.Printf("%s[%d] %s%s\n", Cyan, i+1, option, Reset)
		}

		fmt.Print(Bold + "\n👉 Choose an option: " + Reset)
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			fmt.Println(Green + "✅ Changing Node js Version..." + Reset)
			changeNodejsVersion()
		case 1:
			fmt.Println(Green + "✅ Starting Redis..." + Reset)
			startRedis()
		case 2:
			fmt.Println(Red + "❌ Stopping Redis..." + Reset)
			stopRedis()
		case 3:
			fmt.Println(Blue + "📞 Opening Google Meet 1..." + Reset)
			exec.Command("cmd", "/c", "start", misikiMeetLink).Run()
		case 4:
			fmt.Println(Blue + "📞 Opening Google Meet 2..." + Reset)
			exec.Command("cmd", "/c", "start", sparkinityMeetLink).Run()
		case 5:
			fmt.Println(Blue + "📲 Opening Telegram..." + Reset)
			exec.Command("cmd", "/c", "start", "tg://").Run()
		case 6:
			fmt.Print(Bold + "💬 Enter Telegram Username or Phone: " + Reset)
			var userInput string
			fmt.Scanln(&userInput)
			link := fmt.Sprintf("tg://resolve?domain=%s", userInput)
			exec.Command("cmd", "/c", "start", link).Run()
		case 7:
			openVSCodeProject()
		case 8:
			fmt.Print("📌 Increment version (major/minor/patch): ")
			var vType string
			fmt.Scanln(&vType)
			newVersion := updateVersion(vType)
			fmt.Println(Green + "✅ Version updated to " + newVersion + Reset)
		case 9:
			buildExecutable(version)
		case 10:
			fmt.Print(Bold + Green + "🔑 Enter new Ubuntu password: " + Reset)
			var newPassword string
			fmt.Scanln(&newPassword)
			savePassword(newPassword)
			fmt.Println(Green + "✅ Password updated successfully!" + Reset)
		case 11:
			fmt.Println(Yellow + "👋 Exiting... Have a great day, Ritik!" + Reset)
			os.Exit(0)
		default:
			fmt.Println(Red + "❌ Invalid choice, try again." + Reset)
		}
	}
}

// Increment version based on user input
func updateVersion(changeType string) string {
	major, minor, patch := 1, 0, 0

	data, err := os.ReadFile(versionFile)
	if err == nil {
		parts := strings.Split(strings.TrimSpace(string(data)), ".")
		if len(parts) == 3 {
			major, _ = strconv.Atoi(parts[0])
			minor, _ = strconv.Atoi(parts[1])
			patch, _ = strconv.Atoi(parts[2])
		}
	}

	switch changeType {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		patch++
	}

	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	os.WriteFile(versionFile, []byte(newVersion), 0644)
	return newVersion
}
