# sendtg

> 🛠️ Simple CLI tool to send files and folders to Telegram chats using bot tokens.

`sendtg` allows you to send files or folders to specific Telegram chats using Telegram Bot API. It supports interactive configuration of bots and chat IDs, file/folder zipping, and multiple bots.

---

## ✨ Features

- Send files or folders directly to Telegram chats  
- Auto-zip folders before sending  
- Manage multiple bots and chats via interactive menu  
- Choose bot per message or use default  
- Support for custom headings (in development)  
- Clean terminal feedback (status, bot/chat/file info)  

---

## 📦 Installation

### 🔧 Prerequisites

- Go 1.19+
- Telegram bot token ([create one here](https://t.me/botfather))
- Chat ID of the destination group or user

### 🔄 Clone & Build

```bash
git clone https://github.com/yourname/sendtg.git
cd sendtg
go build -o sendtg .
```

### 📂 Add to PATH (optional, recommended)

```bash
sudo mv sendtg /usr/local/bin/
# OR, for user-local bin
mkdir -p ~/bin
mv sendtg ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc
```

You can now run `sendtg` globally from any directory.

---

## 🧑‍💻 Usage

### 🧭 Open Interactive Menu

```bash
sendtg -menu
```

You’ll see:

```
--- sendtg Menu ---
1. Add Chat
2. Delete Chat
3. Add Bot
4. Delete Bot
5. List Config
0. Exit
```

You can:
- Add or delete chat names and their Telegram IDs
- Add or delete bot names and their tokens
- List current config for easy reference

### 📤 Send File

```bash
sendtg <filepath> <chatName>
```

Sends the file using the **first configured bot**.

### 📦 Send Folder (auto-zips)

```bash
sendtg <folderpath> <chatName>
```

Zips the folder into `/tmp/foldername.zip`, sends it, and deletes the zip.

### 🤖 Use Specific Bot

```bash
sendtg <filepath|folderpath> <chatName> <botName>
```

Send file using a specific bot.

### ✅ Example

```bash
sendtg ~/Documents/report.pdf devops
sendtg ~/Projects/codebase devops
sendtg ./demo.gif testers backup_bot
```

---

## 🛠️ Configuration Storage

Your bot and chat config is saved locally in:

```
~/.sendtg/config.json
```

Structure:

```json
{
  "Bots": {
    "default": "123456:ABCDEF...",
    "backup_bot": "987654:ZYXWVU..."
  },
  "Chats": {
    "group1": -103453454455,
    "testers": 123456789
  }
}
```

You can manually edit it if needed.

---

## 🧪 Development & Contributing

Clone the project, make changes in the `internal/` and `cmd/` directories following Go's clean architecture conventions. Then run:

```bash
go run main.go -menu
go run main.go <file> <chat>
```

Contributions welcome via PR. Please include tests for new features.

---

## 🛡️ Security

- Tokens are stored locally in plaintext (`~/.sendtg/config.json`) – avoid sharing access to your machine.  
- Files are sent over HTTPS via Telegram’s Bot API.  
- Future versions may support encrypted config storage and file signing.

---

<!-- ## 📃 License

MIT License. See [LICENSE](./LICENSE) for full details.

--- -->
