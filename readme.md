# GMP
 The Go/Golang package manager
 The npm of Go ?

# How to install GPM
- For linux / Mac based systems
```bash
curl -fsSL https://raw.githubusercontent.com/Nadhila-dot/GPM/refs/heads/main/install.sh | bash
```
- For windows, go to the builds folder and download the .exe

## What is gpm?
Gpm or Go package manager is wrapper around the already implemented go modules system, it makes life easier by replacing go get github.com/my_go_package_stuffies/todo-app to gpm add todo. It's suppose to feel like pip or npm but for go.

## Why should anyone use gpm?
Gpm makes life simpler, You can have your own registar or use the public registar at go.dev.pkg.lat. It automatically manages the installing while being beginner friendly, or if your a hardcore Gopher, you can simply use GPM to remove like 13 characters from your go get commands. 

## Language and Structure 
Gpm is written in go, unlike other packages managers that are fast (bun 👀) we can maintain our manager with the language it's ment for. 

### The gpm.toml file
The gpm.toml file holds all the information on your packages and source. 

```toml
[Configuration]
source = "http://localhost:3000"
OS = "MacOS"
metadata = ""

[packages]
Fiber = "github.com/gofiber/fiber/v2"


```
### Sources
A source is basically a registar you call to get package details, you can call a registar it sends data like
```json
"GORM": {
      "versions": {
        "default": "gorm.io/gorm",
        "latest": "gorm.io/gorm",
        "major": {
          "1": "github.com/jinzhu/gorm",
          "2": "gorm.io/gorm"
        },
        "patch": []
      }
    },
```
Or if it's an 3rd party enabled registar.
```json
"GoGitDumper": {
      "party": {
        "score": "8",
        "source": "go.pkg.dev",
        "type": "3rd"
      },
      "versions": {
        "default": "github.com/pirateinformatique1337/allah/GoGitDumper",
        "latest": "github.com/pirateinformatique1337/allah/GoGitDumper",
        "major": {
          "1": "github.com/pirateinformatique1337/allah/GoGitDumper"
        },
        "patch": []
      }
    },
```
## Modify GPM?
- Coming soon... 

## Commands on GPM
- gpm add/i all do the same thing, you can install a go module. 
- gpm refersh to get "cached" packages from the source.
- gpm help to get a simple help screen
- gpm list to list installed packages
- gpm remove to remove packages
- gpm setsource to set a new source. <-- More on this source thing later...

# What happens under the hood?
When you do gpm add Fiber, gpm calls the offical go.pkg.lat registar or the source you set to get the module path for go get, it will automatically run go get and install the module. Afterwards gpm will run go mod verify and go mod tidy to sync and compile everything for you. 
- An example output of Gpm.
```bash
nadhi@Mac testing % ./gpm add Fiber                      
 
  ____ ____  __  __  
 / ___|  _ \|  \/  | 
| |  _| |_) | |\/| | Simplify your Go modules
| |_| |  __/| |  | | Go Package Manager
 \____|_|   |_|  |_| 
A package manager for Go Lang
Made for MacOS with go1.25.1

[✓] Found gpm.toml file..
[✓] Using package source: http://localhost:3000
[✓] Fetching packages from source..
[✓] Installing package ..
[✓] Refetching packages..
[✓] Checking a version for package..
[✓] Installing package via "go get"..
[✓] Running command: go get github.com/gofiber/fiber/v2
go: writing stat cache: open /Users/nadhi/go/pkg/mod/cache/download/github.com/valyala/fasthttp/@v/v1.66.0.info2419137.tmp: permission denied
go: writing go.mod cache: open /Users/nadhi/go/pkg/mod/cache/download/github.com/valyala/fasthttp/@v/v1.66.0.mod790127875.tmp: permission denied
go: added github.com/gofiber/fiber/v2 v2.52.9

[Success] Downloaded package Fiber..

[↻] Starting verification of installed packages..
[↻] Verifying packages..
[↻] Running 'go mod verify' to check package integrity..
[↻] Analyzing verification results..

[Success] All modules safe and verified

```

Gpm cuts down all the time wasted on doing extra steps like this. Ofcourse all you can do with GPM can be done by a human but why not automate it properly? 


## Todo
- Fix return statements.
- Make sure a go.mod file is present, if not do not proceed
- Sometimes even a fail return a sucess which is horrible.
