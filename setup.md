## Setup
## Step 1: Download
Go to  [golang.org/dl/](https://golang.org/dl/)  and download the installer for your operating system. Open the downloaded file and follow the instructions. <br>On **Linux** extract the downloaded `.tar.gz` file.
```
tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz
```
## Step 2: Path
Set the `PATH` environment variable. <br>
To setup the `PATH` on **Windows** please follow this [tutorial](https://medium.com/@bhanotvardana/setting-up-golang-environment-on-windows-3d50c2dbffe7). <br>
On **macOS** and **Linux**:
1. Navigate to your shell configuration file in your home directory:
    - For Bash: `vim ~/.bash_profile` or `vim ~/.bashrc`
    - For Zsh: `vim ~/.zshrc`
2. Add the following line at the end of the file: 
   ```
   export PATH=$PATH:/usr/local/go/bin
   ```
3. If you followed the workspace setup instructions, also add the `$GOPATH/bin` directory to the `PATH` by adding the following line: 
   ```
   export PATH=$PATH:$GOPATH/bin
   ```
4. Save the file and reload the configuration.
   ```
   source ~/.your_shell_config
   ```
## Step 3: Verification
Open a new terminal session and verify the installation.
```
go version
```
## Step 4: IDE
Download [VSCode](https://code.visualstudio.com/download) or another appropriate IDE like [GoLand](https://www.jetbrains.com/go/). <br>
In **VSCode**, go to the Extensions Marketplace, search for the Go plugin and install it.
<img src=go-plugin.png width=500>
Also install the necessary Go tools that comes with it.
Open the Command Palette by pressing `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (macOS), type "Go: Install/Update Tools" and press `Enter`.
