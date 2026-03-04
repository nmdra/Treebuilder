# TreeBuilder

TreeBuilder is a fast, lightweight Command Line Interface (CLI) tool written in Go. It reads a standard tree-formatted text file and automatically generates the corresponding directory structure and empty files on your local disk. 

Whether you are scaffolding a new project, setting up a monorepo, or generating boilerplate, TreeBuilder saves you from running endless `mkdir` and `touch` commands.

## Features
* **Fast & Lightweight:** Built with Go and the Cobra CLI framework.
* **Smart Parsing:** Automatically distinguishes between directories (ending in `/`) and files.
* **Robust:** Safely handles UTF-8 characters and malformed indentation.
* **Dry-Run Mode:** Preview the folder structure in your terminal without writing anything to the disk.
