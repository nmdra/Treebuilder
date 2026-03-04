# TreeBuilder

A fast CLI tool written in Go that reads a tree-formatted text file and instantly scaffolds the corresponding directory structure and empty files on your disk.

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
