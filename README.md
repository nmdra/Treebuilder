# TreeBuilder

A fast CLI tool written in Go that reads a tree-formatted text file and instantly scaffolds the corresponding directory structure and empty files on your disk.

![Treebuilder png](https://github.com/user-attachments/assets/ffa6e4f2-5f7f-41f6-a876-71a827e5f404)

> [!NOTE]
> **Hi, Nimendra here!**   
> I built this tool out of personal frustration. Most of the time, AI tools like Gemini and ChatGPT give you a folder structure in a tree format. It's a massive headache to create those folders and files one by one. I built this tool so you can just take that text output and instantly convert it into the actual folders and files on your disk!

## Installation

Download the latest binary for your OS (macOS, Linux, Windows) from the [Releases page](https://github.com/nmdra/treebuilder/releases).

*Or build from source:*
```bash
make build

```

## Usage

Create a text file (e.g., `structure.txt`) defining your layout. **Note:** Directories must end with a trailing slash (`/`).

**structure.txt:**

```text
my-project/
├── cmd/
│   └── main.go
└── pkg/
    └── utils.go

```

**Run the CLI:**

```bash
treebuilder structure.txt

```

**Preview without creating files (Dry Run):**

```bash
treebuilder --dry-run structure.txt

```

---

Built with ❤️ by Nimendra — [blog.nimendra.xyz](https://blog.nimendra.xyz)
