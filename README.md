# Web FOUND
Web Found is a powerful web scanning tool designed to help developers and security professionals identify potential vulnerabilities in web applications. The tool uses a word list to send HTTP requests to a target URL and prints the response status codes. With its simple and intuitive interface, Web Found makes it easy to scan web applications and identify potential security risks.
## Screenshot 📸
[![image](./cmd/assets//exaple.jpg)](#)

![Downloads](https://badgen.net/github/assets-dl/Ruwantha-OFFICIAL/Web-Found/?color=green)
![Release](https://badgen.net/github/release/Ruwantha-OFFICIAL/Web-Found/?color=green)
![go](https://badgen.net/badge/icon/golang?icon=go&label&color=green)

**Fast & Simple Web Content Discovery**


## Features ✨
- **Fast HEAD Requests** - Low bandwidth scanning
- **Status Code Detection** - 200, 301, 403, 404
- **Custom Wordlist** - Use any txt wordlist
- **Lightweight** - Single binary, no dependencies
- **Clean Output** - With banner and summary

----
## 📂 Project Structure
```markdown
web-found/
├── cmd/
│   ├── main.go
│   ├── check.go
├── internal/
│   ├── readfile/
│   │   ├── reader.go
│   ├── requests/
│   │   ├── request.go
├── go.mod
├── go.sum
```
---

## Installation 
Go to **[Release](https://google.com)** Page & Download Binary File 

### Downloads
| Architecture | Windows | Linux |
| :--- | :---: | :---: |
| **x86_64 / amd64** | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_windows_amd64.zip) | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_linux_amd64.tar.gz) |
| **aarch64 / arm64** | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_windows_arm64.zip) | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_linux_arm64.tar.gz) |
| **x86 / 386** | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_windows_386.zip) | [Download](https://github.com/Ruwantha-OFFICIAL/Web-Found/releases/download/1.0.0v/web-found_linux_386.tar.gz) |

### Linux Setup 
```bash
chmod +x web-found
./web-found
```

### Windows Setup

```powershell
.\web-found.exe
```

---

## Usage 🧾

```bash
./web-found.exe -url https://en.wikipedia.org -word webco.txt
```
### Options 🐾

- **url**: Rconsens Website Url/Domain Name
- **word**: You Provide Word List  `wordlist.txt/webco.txt`
- **r**: Redirect mode

---
## 📝 License
Web Found is licensed under the ISC License.

## 📬 Contact
For any questions or concerns, please contact us at [ruwanthalasith20@gmail.com](mailto:ruwanthalasith20@gmail.com).



Copyright 2026 Lasith Ruwantha Amarawansha
