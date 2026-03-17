# 📚 Library Management System

A CLI-based library management system built with Go, demonstrating core Go concepts like structs, methods, pointers, interfaces, and error handling.

## ✨ Features

- ➕ Add books to the library
- 📖 List all books with availability status
- 👤 Register library users
- 👥 List all registered users
- 📤 Borrow books (with validation)
- 📥 Return borrowed books
- 📋 View complete borrowing history
- 📚 View books borrowed by a specific user

## 🏗️ Architecture
`
library-management/
├── main.go                    # Entry point
├── cli/cli.go                 # User interface
├── library/library.go         # Business logic
├── models/                    # Data structures
│   ├── person.go
│   ├── book.go
│   ├── user.go
│   └── borrow_record.go
└── interfaces/displayable.go  # Interface
`

## 🎓 Go Concepts Demonstrated

- **Structs**: Person, Book, User, BorrowRecord
- **Methods**: Value and pointer receivers
- **Embedding**: User embeds Person
- **Interfaces**: Displayable interface
- **Error Handling**: Idiomatic Go patterns
- **Slices**: Managing collections
- **Packages**: Clean organization
- **Pointers**: For mutations

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or later

### Installation

\\\ash
git clone https://github.com/EdrisKorbi/library-management.git
cd library-management
go mod tidy
go run main.go
\\\

## 💻 Usage

The application provides a menu-driven interface:

\\\
1. Add Book
2. List Books
3. Add User
4. List Users
5. Borrow Book
6. Return Book
7. View Borrow History
8. View User's Borrowed Books
9. Exit
\\\

### Example Workflow

1. **Add a book**: Choose option 1, enter title, author, ISBN
2. **Register a user**: Choose option 3, enter name and email
3. **Borrow a book**: Choose option 5, enter user ID and book ID
4. **View history**: Choose option 7 to see all transactions

## 📊 Features in Detail

### Add Book
- Requires: title, author, ISBN
- Auto-generates unique ID
- Books start as available

### Add User
- Requires: name, email
- Auto-generates unique ID
- Records join date

### Borrow Book
- Validates user exists
- Validates book exists
- Checks book is available
- Creates borrow record
- Marks book as unavailable

### Return Book
- Validates user exists
- Validates book exists
- Checks book is borrowed
- Records return date
- Marks book as available

## 🧪 Testing

Test the application by:
1. Running \go run main.go\
2. Add multiple books and users
3. Borrow books and verify status changes
4. Return books and check history
5. Try edge cases (invalid IDs, unavailable books, etc.)

## 📈 Code Quality

- ✅ Follows Go conventions
- ✅ Proper error handling
- ✅ Clean architecture
- ✅ Well-organized packages
- ✅ Input validation

## 🔮 Future Enhancements

- Add persistence (save to JSON/database)
- Add search functionality
- Add user authentication
- Add book ratings
- Add REST API
- Add web interface

## 📝 Project Structure

| File | Lines | Purpose |
|------|-------|---------|
| main.go | ~7 | Entry point |
| cli/cli.go | ~200 | User interface |
| library/library.go | ~175 | Business logic |
| models/*.go | ~150 | Data structures |
| interfaces/displayable.go | ~3 | Interface |

**Total**: ~600 lines of production code

## 👨‍💻 Author

Korbi Edris

## 📄 License

This project is open source and available under the MIT License.

---

**Ready to use!** Clone the repository and run \go run main.go\ to get started.
