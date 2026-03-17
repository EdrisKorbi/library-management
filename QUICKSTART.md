# ⚡ Quick Start Guide

Get the Library Management System running in 5 minutes.

## 🚀 Prerequisites

- Go 1.21 or later
- Git

## 📥 Installation

\\\ash
git clone https://github.com/EdrisKorbi/library-management.git
cd library-management
go run main.go
\\\

That's it! The application will start.

## 🎮 First Run

When you start, you'll see a menu:

\\\
╔═══════════════════════════════════╗
║   Library Management System       ║
╚═══════════════════════════════════╝

--- Main Menu ---
1. Add Book
2. List Books
3. Add User
4. List Users
5. Borrow Book
6. Return Book
7. View Borrow History
8. View User's Borrowed Books
9. Exit

Enter choice (1-9):
\\\

## 📚 5-Minute Demo

Follow this to see all features:

### 1️⃣ Add a Book
\\\
Enter choice: 1
Enter book title: The Go Programming Language
Enter book author: Donovan & Kernighan
Enter book ISBN: 978-0134190440
✅ Book added successfully! ID: 1
\\\

### 2️⃣ Add a User
\\\
Enter choice: 3
Enter user's name: Alice Johnson
Enter user's email: alice@example.com
✅ User added successfully! ID: 1
\\\

### 3️⃣ Borrow a Book
\\\
Enter choice: 5
Enter user's ID: 1
Enter book's ID: 1
✅ Book borrowed successfully.
\\\

### 4️⃣ View History
\\\
Enter choice: 7
--- Borrow History ---
ID: 1 | UserID: 1 | BookID: 1 | BorrowDate: 2026-03-17 | ReturnDate: Not returned yet | Status: borrowed
Total records: 1.
\\\

### 5️⃣ Return the Book
\\\
Enter choice: 6
Enter user's ID: 1
Enter book's ID: 1
✅ Book returned successfully.
\\\

### 6️⃣ Exit
\\\
Enter choice: 9
Thank you for using Library Management System!
\\\

## ✅ Features You Can Try

- Add multiple books
- Add multiple users
- Borrow books and see status change
- Return books and see history update
- Try invalid IDs (error handling works)
- Try borrowing unavailable book (validation works)

## 🛠️ Common Commands

\\\ash
# Run the application
go run main.go

# Build executable
go build -o library-cli

# Run executable
./library-cli

# Check for errors
go vet ./...
\\\

## 📖 Learn More

- See **README.md** for full documentation
- See **PORTFOLIO.md** for interview talking points

---

**That's it!** You're ready to use the Library Management System. Enjoy! 🎉
