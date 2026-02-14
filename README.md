# 🤖 AI Programming Mentor Bot

[![Go Report Card](https://goreportcard.com/badge/github.com/Endropr/ai-programming-mentor)](https://goreportcard.com/report/github.com/Endropr/ai-programming-mentor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An advanced Telegram bot built with **Go** that serves as a personal coding mentor. The system leverages **OpenAI's LLMs** to provide expert-level guidance, featuring persistent conversation history and a modular architecture.

---

## Key Features

* **Interactive Specialization Selection**: Users can choose their learning track (Go, Python, JS, C++, PHP, HTML/CSS) via a sleek, multi-row inline keyboard.
* **Context-Aware Mentorship**: The bot maintains the state of the selected programming language for each user, ensuring personalized guidance.
* **Clean Architecture**: Strictly follows a modular design (Domain, Repository, Application layers), making the codebase easy to scale and maintain.
* **LLM Agnostic Design**: Decoupled service layer allows seamless switching between AI providers (OpenAI, DeepSeek, Anthropic) without rewriting core logic.
* **Persistent Storage**: Full audit logs and user preferences are stored in a **PostgreSQL** database with high-performance interaction via the `pgx` driver.
* **Rich UI/UX**: Utilizes HTML-formatted responses and real-time callback feedback for a professional look and feel.

---

## Project Structure

Following the **Standard Go Project Layout**, the code is organized into logical layers:

```text
ai-programming-mentor/
├── cmd/bot/           # Entry point: Bot initialization and update loop
├── internal/
│   ├── domain/        # Core entities: Business models (Message, User)
│   └── repository/    # Data Layer: PostgreSQL implementation (Save/Load logic)
├── migrations/        # SQL scripts: Schema definition and table structures
├── .env.example       # Environment configuration template
└── README.md          # Project documentation
```
---

## 🛠 Tech Stack

• Language: Go (Golang) 1.21+

• Database: PostgreSQL

• AI Engine: OpenAI GPT-4o-mini

• APIs: Telegram Bot API v5

• Libraries: pgx/v5, godotenv, go-openai

---

## ⚙️ Installation & Setup
### 1. Prerequisites

• Installed Go and PostgreSQL.

• API keys from OpenAI and Telegram (@BotFather).

### 2. Configuration

Clone the repository and create a .env file:

```
git clone [https://github.com/Endropr/ai-programming-mentor.git](https://github.com/Endropr/ai-programming-mentor.git)
```

```
cd ai-programming-mentor
```

```
cp .env.example .env
```

#### Fill in your credentials:

```
TELEGRAM_APITOKEN=your_token
OPENAI_API_KEY=your_key
DB_URL=postgres://user:password@localhost:5432/dbname
```

## 3. Database Initialization

Run the following SQL command to set up the storage:

```
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    role VARCHAR(20),
    content TEXT,
    selected_language VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 4. Running the Bot

```
go run cmd/bot/main.go
```

---

## 📸 Project Demo & Database
<details>
  <summary>Click to view bot interface and database logs</summary>
  <br>
  
  ### Bot Interface Preview
  <p align="left">
    <img src="assets/bot.png?v=2" alt="Telegram Bot Interface" width="500"/>
  </p>
  <p align="left"><i>Figure 1: The interactive menu with 3x2 inline keyboard.</i></p>

  <br>

  ### Database Logs (PostgreSQL)
  <p align="left">
    <img src="assets/db_full_preview.png" alt="Database Logs" width="900"/>
  </p>
  <p align="left"><i>Figure 2: The <code>messages</code> table tracking language selections (Python, HTML, C++, PHP).</i></p>
</details>
