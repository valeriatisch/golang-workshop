
# Setup
### Step 1: Download
Go to  [golang.org/dl/](https://golang.org/dl/) and download the installer for your operating system. Open the downloaded file and follow the instructions. <br>
On **Linux** extract the downloaded `.tar.gz` file.
```bash
tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```
### Step 2: Path
Set the `PATH` environment variable. <br>
To setup the `PATH` on **Windows** please follow this [tutorial](https://medium.com/@bhanotvardana/setting-up-golang-environment-on-windows-3d50c2dbffe7). <br>
On **macOS** and **Linux**:
1. Navigate to your shell configuration file in your home directory:
    - For Bash: `vim ~/.bash_profile` or `vim ~/.bashrc`
    - For Zsh: `vim ~/.zshrc`
2. Add the following line at the end of the file: 
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
3. If you followed the workspace setup instructions, also add the `$GOPATH/bin` directory to the `PATH` by adding the following line: 
   ```bash
   export PATH=$PATH:$GOPATH/bin
   ```
4. Save the file and reload the configuration.
   ```bash
   source ~/.your_shell_config
   ```
### Step 3: Verification
Open a new terminal session and verify the installation.
```bash
go version
```
### Step 4: IDE
Download [VSCode](https://code.visualstudio.com/download) or another appropriate IDE like [GoLand](https://www.jetbrains.com/go/).
In **VSCode**, go to the Extensions Marketplace, search for the Go plugin and install it. <br>
Also install the necessary Go tools that comes with it.
Open the Command Palette by pressing `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (macOS), type *"Go: Install/Update Tools"* and press `Enter`.
### Step 5: Project
1. Create a new directory for your project. This directory will be the workspace for your Go code, dependencies, and other project-related files.
   ```bash
   mkdir your_go_project
   cd your_go_project
   ```
2. Go uses modules to manage dependencies. Initialise a new Go module with your desired module path (a unique identifier for your Go module)
   ```bash
   go mod init module_path/your_go_project
   ```
   A `go.mod` file will be created in your project directory, which will be used to manage your project's dependencies.
3. Create a new file named `main.go` in your project directory and write a simple Hello World program.
   ```go
   package main
   
   import "fmt"
   
   func main() {
	   fmt.Println("Hello, world!")
   }
   ```
4. Compile your program.
   ```bash
   go build
   ```
   This will create an executable binary named `main` (or `main.exe` on Windows).
   To run your program, simply execute the binary:
   ```bash
   ./main
   ```
   Alternatively, you can use compile and run your program in a single step with:
   ```bash
   go run main.go
   ```
