# GMP
 The Go/Golang package manager

# How to install GPM
- For linux / Mac based systems
```bash
curl -fsSL https://raw.githubusercontent.com/Nadhila-
```
- For windows, go to the builds folder and download the .exe

## What is gpm?
Gpm or Go package manager is wrapper around the already implemented go modules system, it makes life easier by replacing go get github.com/my_go_package_stuffies/todo-app to gpm add todo. It's suppose to feel like pip or npm but for go.

## Why should anyone use gpm?
Gpm makes life simpler, You can have your own registar or use the public registar at go.dev.pkg.lat. It automatically manages the installing while being beginner friendly, or if your a hardcore Gopher, you can simply use GPM to remove like 13 characters from your go get commands. 

## Language and Structure 
Gpm is written in go, unlike other packages managers that are fast (bun 👀) we can maintain our manager with the language it's ment for. 

## Modify GPM?
- Later i'll add docs on this. 

## Commands on GPM
- gpm add/install/i all do the same thing, you can install a go module. 
- gpm refersh to get "cached" packages from the source.
- gpm help to get a simple help screen
- gpm list to list installed packages
- gpm remove to remove packages
- gpm setsource to set a new source. <-- More on this source thing later...

