# 📄 Gitignore CLI

A handy command-line tool to quickly fetch and create `.gitignore` files for your project, directly from the terminal — using the official GitHub `.gitignore` templates.

![Go Version](https://img.shields.io/badge/Go-1.20+-blue?logo=go)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🔍 Search for available `.gitignore` templates
- 📄 Download `.gitignore` files for specific languages or frameworks
- ➕ Append or overwrite `.gitignore` content easily
- 🌍 Uses GitHub’s official [gitignore repository](https://github.com/github/gitignore)
- ⚙️ Minimal dependencies, blazing fast

---

## 📦 Installation

### Option 1: Go Install (recommended)

```bash
go install github.com/hadzicni/gitignore-cli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build

```bash
git clone https://github.com/hadzicni/gitignore-cli.git
cd gitignore-cli
go build -o gitignore ./cmd/gitignore
```

---

## 🚀 Usage

```bash
gitignore [flags] [template-name]
```

### Available Flags

| Flag               | Description                                        | Example                         |
|--------------------|----------------------------------------------------|----------------------------------|
| `--list`, `-l`     | List available gitignore templates                 | `gitignore -l`                  |
| `--append`, `-a`   | Append to existing .gitignore instead of overwrite | `gitignore -a go`               |
| `--output`, `-o`   | Specify output path (default: .gitignore)          | `gitignore -o .gitignore.dev`   |
| `--help`, `-h`     | Show help message                                  | `gitignore -h`                  |

---

## 🔧 Examples

Create a `.gitignore` for Go:

```bash
gitignore go
```

Append Python rules to existing file:

```bash
gitignore -a python
```

List all supported templates:

```bash
gitignore --list
```

Write to a custom file:

```bash
gitignore -o .gitignore.custom node
```

---

## 🧪 Development

Run the CLI locally:

```bash
go run ./cmd/gitignore
```

Run tests:

```bash
go test ./...
```

---

## 📁 Project Structure

```
.
├── cmd/
│   └── gitignore/       # CLI command logic
├── internal/
│   └── fetcher/         # Template fetching logic
├── go.mod               # Go module file
├── LICENSE              # License file
└── README.md            # Project readme
```

---

## 👨‍💻 Author

Made with ❤️ by **Nikola Hadzic**

- GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
