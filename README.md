# h2go

**A CLI tool and component library for building Go + HTMX applications faster.**  
Inspired by [shadcn/ui](https://github.com/shadcn/ui), `h2go` provides ready-to-use TailwindCSS-styled components and scaffolding to help Go developers build rich, interactive UIs with HTMX and Go templates.

---

## 🚀 Features

- CLI to scaffold UI components directly into your Go projects  
- Prebuilt components optimized for HTMX interactivity  
- TailwindCSS integration for styling flexibility  
- Easily extensible and customizable components  
- Open source and community driven  

---

## 📦 Installation

You can install the CLI tool by building from source (coming soon) or downloading binaries (planned).

```bash
# Example (once published)
go install github.com/your-org/h2go@latest
```

---

## 💡 Usage

### Initialize your project

```bash
h2go init
```

Sets up TailwindCSS config and creates the components directory structure.

### Add a component

```bash
h2go add button
```

Adds the `button` component templates to your project.

### List available components

```bash
h2go list
```

Shows all components you can add.

### Update components

```bash
h2go update
```

Updates your installed components to the latest versions.

---

## 📁 Project Structure

```plaintext
h2go/
├── cmd/                 # CLI command definitions
├── internal/            # Internal logic & scaffolding code
├── templates/           # Component template files
├── pkg/                 # Utility packages
├── examples/            # Example projects & demos
├── README.md
└── LICENSE
```

---

## 📅 Roadmap

- [x] CLI skeleton with basic commands (`init`, `add`, `list`, `update`)  
- [ ] Functional component scaffolding (buttons, inputs, modals)  
- [ ] TailwindCSS auto-configuration and theming  
- [ ] Example Go + HTMX demo app  
- [ ] Documentation website  
- [ ] Community contributions & plugin system  

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check [issues](https://github.com/your-org/h2go/issues) or submit a PR.

---

## 📞 Stay in Touch

Follow the project updates on X (Twitter): [@yourhandle](https://twitter.com/yourhandle)  
Join discussions and give feedback!

---

## 📄 License

[MIT License](./LICENSE)