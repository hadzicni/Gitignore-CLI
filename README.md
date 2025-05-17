# gitignore-cli

Generate `.gitignore` files by downloading templates directly from the official [github/gitignore](https://github.com/github/gitignore) repository.

## ✅ Usage

```bash
gitignore go
gitignore go,node
gitignore --list
gitignore -o myignore.txt python,java
```

## 🌐 Templates are fetched from:

- GitHub repo: https://github.com/github/gitignore
- Raw files: https://raw.githubusercontent.com/github/gitignore/main/

## 🧾 Flags

| Flag        | Description                              |
| ----------- | ---------------------------------------- |
| `--list`    | List all available templates (live)      |
| `-o <file>` | Output file name (default: `.gitignore`) |

## 🛠 Build

```bash
go mod tidy
go build -o gitignore.exe ./cmd/gitignore
```

## 💡 Notes

- Template names are case-insensitive
- Special cases like `c++` are supported as `c++`
- Multiple templates can be combined via comma

## 📄 License

MIT
