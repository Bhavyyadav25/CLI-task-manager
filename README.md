# TaskMgr

A cleanly-architected, Go-powered CLI task manager.  
Easily add, list, mark done, delete tasks — and export them to a polished PDF report.

---

## 🚀 Features

- **Add** new tasks with descriptions  
- **List** all tasks in your to-do list  
- **Mark** tasks as done  
- **Delete** tasks by ID  
- **Export** your task list to a styled PDF  
- **Clean Architecture**  
  - **Domain**: core `Task` entity & validation  
  - **Repo (Ports)**: CRUD interface  
  - **Store (Adapters)**: JSON-file implementation  
  - **Use Cases**: application logic + PDF exporter  
  - **CLI**: command-line interface layer  

---

## 📦 Project Layout

```
taskmgr/
├── cmd/
│   └── taskmgr/           # CLI entrypoint
├── internal/
│   ├── domain/            # Task entity & errors
│   ├── repo/              # Repository interfaces
│   ├── store/             # JSON file-based storage
│   └── usecase/           # Application services + PDF exporter
├── assets/
│   └── fonts/             # UTF-8 TTF fonts for PDF
│       └── DejaVuSans.ttf
├── go.mod
└── README.md
```

---

## ⚙️ Installation

1. **Clone the repo**  
   ```bash
   git clone https://github.com/you/taskmgr.git
   cd taskmgr
   ```

2. **Download dependencies & build**  
   ```bash
   go mod tidy
   go build -o taskmgr ./cmd/taskmgr
   ```

3. **Add your font**  
   Place `DejaVuSans.ttf` (or another UTF-8-capable TTF) at:
   ```
   assets/fonts/DejaVuSans.ttf
   ```

---

## 📝 Usage

```bash
# Add a new task
./taskmgr add "Buy groceries"

# List all tasks
./taskmgr list

# Mark a task done
./taskmgr done <task-id>

# Delete a task
./taskmgr del <task-id>

# Export tasks to PDF (default: tasks.pdf)
./taskmgr export [output.pdf]
```

- Tasks are persisted to `tasks.json` in the current working directory.  
- PDF export uses the font at `assets/fonts/DejaVuSans.ttf`.

---

## ✅ Testing

Run unit and integration tests:

```bash
go test ./...
```

- **Domain tests** for validation rules  
- **Use-case tests** with mocked repos  
- **Store tests** against temporary JSON files  

---

## 🤝 Contributing

1. Fork the repo  
2. Create your feature branch (`git checkout -b feat/my-feature`)  
3. Commit your changes (`git commit -m 'Add some feature'`)  
4. Push to the branch (`git push origin feat/my-feature`)  
5. Open a Pull Request  

Please follow Go formatting (`go fmt`) and include tests for new logic.

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for details.

---

## 🙏 Acknowledgments

- [GoFPDF](https://github.com/jung-kurt/gofpdf) for PDF generation  
- [DejaVu Sans](https://dejavu-fonts.github.io/) for UTF-8 font support  