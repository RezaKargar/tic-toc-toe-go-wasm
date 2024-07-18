# Go WebAssembly Tic-Tac-Toe Game

This project is a simple Tic-Tac-Toe game written in Go and compiled to WebAssembly (Wasm). The game demonstrates basic DOM manipulation and event handling using the `syscall/js` package in Go.

## Project Structure

- `main.go`: The Go source code that compiles to WebAssembly.
- `index.html`: The HTML file that loads the WebAssembly module and includes the game UI.
- `wasm_exec.js`: The JavaScript support file required to run Go WebAssembly.
  - To get this file, you can run this command:
    ```shell
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    ```

## Getting Started

### Step 0: Prerequisites

- [Go](https://golang.org/doc/install) 1.11+ (WebAssembly support)


### Step 1: Clone the Repository

```sh
git clone https://github.com/RezaKargar/tic-toc-toe-go-wasm.git
cd tic-toc-toe-go-wasm
```

### Step 2: Compile the source code

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```