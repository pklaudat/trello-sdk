# Trello SDK for Golang

A lightweight Go SDK for Trello, designed to simplify the management of boards and cards using either a CLI interface or a REST API. This project wraps the [Trello public API](https://developer.atlassian.com/cloud/trello/rest/api-group-actions/) to offer an intuitive developer experience via terminal.

---

## 🚀 Features

- Manage Trello boards and cards from your terminal
- Interact with Trello via a simple REST API (built with Gin)
- Modular SDK design for future extensibility
- Built entirely in Go — no external runtime dependencies

---

## 🛠 Project Structure

```bash
.
├── cmd/              # entrypoint - both local CLI and REST API
├── internal/         # Trello SDK logic and client implementation
├── go.mod
└── main.go
