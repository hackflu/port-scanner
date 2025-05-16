# port-scanner

A fast and simple port scanner written in Go to scan open ports on a target IP or hostname.

## 📦 Features

- Scan a range of TCP ports on any IP address or hostname
- Detect open ports efficiently using concurrent scanning
- Simple and easy-to-use command-line interface
- Supports customizable port ranges for flexible scanning
- Displays open ports clearly in the output
- Lightweight and fast with minimal dependencies

## ⚙️ Requirements

- Go 1.18+

## 🛠️ Libraries Used

- **net** — For TCP connections to test if ports are open.
- **sync** — For managing concurrency to speed up scans.
- **flag** — For parsing command-line arguments.

## 🚀 Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/hackflu/port-scanner.git
   cd port-scanner
