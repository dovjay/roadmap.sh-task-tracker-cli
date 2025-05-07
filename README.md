# 📝 Task Tracker CLI in Go

A lightweight CLI tool to manage tasks using a JSON file.  
Supports adding, updating, deleting, and marking tasks with statuses like **todo**, **in-progress**, and **done**.

---

## 📦 Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer
- A terminal or command prompt (PowerShell, Git Bash, etc.)

---

## 🚀 Installation

1. **Clone this repo or copy the code into a directory:**

   ```bash
   git clone https://github.com/dovjay/roadmap.sh-task-tracker-cli
   cd taskcli
   ```

2. **Build the executable:**

   ```bash
   go build -o taskcli
   ```

3. **Run it from your terminal:**

   ```bash
   ./taskcli
   ```

---

## ⚙️ Usage

### ✅ Add a Task

```bash
./taskcli add Buy groceries
```

---

### 🔄 Update a Task Title

```bash
./taskcli update <id> New task title
# Example:
./taskcli update 1 Buy almond milk
```

---

### 🗑 Delete a Task

```bash
./taskcli delete <id>
# Example:
./taskcli delete 1
```

---

### 🚧 Mark Task Status

```bash
./taskcli mark <id> <todo|in-progress|done>
# Example:
./taskcli mark 2 done
```

---

### 📋 List Tasks

#### All Tasks

```bash
./taskcli list
```

#### Filtered Tasks

```bash
./taskcli list done
./taskcli list todo
./taskcli list in-progress
```

---

## 📌 Output Example

```plaintext
[2] Buy almond milk (done)
  Created: 2025-05-07T13:00:00Z
  Updated: 2025-05-07T14:22:00Z
```

---

## 🧠 Notes

- Tasks are saved in `tasks.json` in the current directory.
- You can safely edit or copy the JSON file if needed.
- All timestamps are stored in RFC3339 format.
