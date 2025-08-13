# 🎯 Number Guessing Game

A modular, test-driven, command-line **Number Guessing Game** written in Go.  
Guess the secret number within a limited number of tries, with support for **difficulty levels**, **timer tracking**, and a clean, extensible codebase.

---

##  About this Project

This project is part of the **Number Guessing Game** challenge from [roadmap.sh](https://roadmap.sh/projects/number-guessing-game), designed to help learners build skills in conditional logic, loops, user input handling, and CLI interaction.  

---

## ✨ Features
- 🎮 Interactive CLI gameplay
- 📊 Multiple difficulty levels (easy, medium, hard)
- ⏱ Timer to track gameplay duration
- 🔄 Random number generation
- 🧪 Unit tests for all core components
- 🗂 Modular project structure

---

## 📦 Project Structure
```
.
├── difficulty
│   ├── difficulty.go
│   └── difficulty_test.go
├── game
│   ├── game.go
│   └── game_test.go
├── random
│   ├── random.go
│   └── random_test.go
├── timer
│   ├── timer.go
│   └── timer_test.go
├── ui
│   └── ui.go
├── README.md
├── go.mod
└── main.go
```


---

## 🚀 Getting Started

### Installation & Run
```bash
# Clone the repo
git clone https://github.com/amiralitaherkhany/number-guessing-game.git
cd number-guessing-game

# Install dependencies
go mod tidy

# Run the game
go run main.go
```
--- 

## 🕹 How to Play
 - Select difficulty (which sets the number of attempts).
 - Guess a number between 1 and 100.
 - Receive feedback، too high or too low.
 - Try to guess correctly within limited attempts.
 - Timer shows how long you took **and how many guesses you made**.
 - choose to play again or exit

## ✅ Running Tests

```bash
go test ./...
```
## License

This project is licensed under the **MIT License**.

Read the full license in the [LICENSE](./LICENSE) file.