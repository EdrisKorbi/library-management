package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/EdrisKorbi/library-management/library"
)

type CLI struct {
	lib     *library.Library
	reader  *bufio.Reader
	running bool
}

func NewCLI() *CLI {
	return &CLI{
		lib:     library.NewLibrary(),
		reader:  bufio.NewReader(os.Stdin),
		running: true,
	}
}

func (cli *CLI) Run() {
	fmt.Println("╔═══════════════════════════════════╗")
	fmt.Println("║   Library Management System       ║")
	fmt.Println("╚═══════════════════════════════════╝")

	for cli.running {
		cli.showMenu()
		cli.processCommand()
	}

	fmt.Println("\nThank you for using Library Management System!")

}

func (cli *CLI) processCommand() {
	input, _ := cli.reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		cli.handleAddBook()
	case "2":
		cli.handleListBooks()
	case "3":
		cli.handleAddUser()
	case "4":
		cli.handleListUsers()
	case "5":
		cli.handleBorrowBook()
	case "6":
		cli.handleReturnBook()
	case "7":
		cli.handleViewBorrowedHistory()
	case "8":
		cli.handleViewUserBorrowedHistory()
	case "9":
		cli.running = false
	default:
		fmt.Println("❌ Invalid choice. Please try again.")

	}
}

func (cli *CLI) showMenu() {
	fmt.Println()
	fmt.Println("\n--- Main Menu ---")
	fmt.Println("1. Add Book")
	fmt.Println("2. List Books")
	fmt.Println("3. Add User")
	fmt.Println("4. List Users")
	fmt.Println("5. Borrow Book")
	fmt.Println("6. Return Book")
	fmt.Println("7. View Borrow History")
	fmt.Println("8. View User's Borrowed Books")
	fmt.Println("9. Exit")
	fmt.Print("\nEnter choice (1-9): ")
}

func (cli *CLI) handleAddBook() {

	title := cli.readInput("Enter book title: ")
	author := cli.readInput("Enter book author: ")
	isbn := cli.readInput("Enter book ISBN: ")

	if title == "" || author == "" || isbn == "" {
		fmt.Println("❌ All fields are required.")
		return
	}

	book := cli.lib.AddBook(title, author, isbn)
	fmt.Printf("Book added successfully! ID: %d.\n", book.ID)

}

func (cli *CLI) handleListBooks() {
	list := cli.lib.ListBooks()

	if len(list) == 0 {
		fmt.Println("📚 No books in library.")
		return
	}

	fmt.Println("\n--- All Books ---")

	for _, book := range list {
		fmt.Println(book.DisplayInfo())
	}
	fmt.Printf("Total books: %d.\n", len(list))
}

func (cli *CLI) handleAddUser() {
	name := cli.readInput("Enter user's name: ")
	email := cli.readInput("Enter user's email: ")

	if name == "" || email == "" {
		fmt.Println("❌ All fields are required.")
		return
	}

	user := cli.lib.AddUser(name, email)
	fmt.Printf("User added successfully! ID: %d.\n", user.ID)

}

func (cli *CLI) handleListUsers() {
	list := cli.lib.ListUsers()

	if len(list) == 0 {
		fmt.Println("📚 No books in library.")
		return
	}

	fmt.Println("\n--- All Users ---")

	for _, user := range list {
		fmt.Println(user.DisplayInfo())
	}
	fmt.Printf("Total users: %d.\n", len(list))
}

func (cli *CLI) handleBorrowBook() {
	userID, err1 := cli.readInt("Enter user's ID: ")

	bookID, err2 := cli.readInt("Enter books's ID: ")

	if err1 != nil || err2 != nil {
		fmt.Println("❌ Invalid input. IDs must be numbers.")
		return
	}

	if err := cli.lib.BorrowBook(userID, bookID); err != nil {
		fmt.Printf("❌ Error: %v.\n", err)
		return
	}
	fmt.Println("✅ Book borrowed successfully.")
}

func (cli *CLI) handleReturnBook() {
	userID, err1 := cli.readInt("Enter user's ID: ")

	bookID, err2 := cli.readInt("Enter books's ID: ")

	if err1 != nil || err2 != nil {
		fmt.Println("❌ Invalid input. IDs must be numbers.")
		return
	}

	if err := cli.lib.ReturnBook(userID, bookID); err != nil {
		fmt.Printf("❌ Error: %v.\n", err)
		return
	}
	fmt.Println("✅ Book returned successfully.")

}

func (cli *CLI) handleViewBorrowedHistory() {
	list := cli.lib.ListBorrowHistory()

	if len(list) == 0 {
		fmt.Println("📋 No borrow history.")
		return
	}

	fmt.Println("\n--- Borrow History ---")
	for _, value := range list {
		fmt.Println(value.DisplayInfo())
	}

	fmt.Printf("Total records: %d.\n", len(list))
}

func (cli *CLI) handleViewUserBorrowedHistory() {

	userID, err := cli.readInt("Enter user's ID: ")
	if err != nil {
		fmt.Println("❌ All fields are required.")
		return
	}

	list := cli.lib.GetUserBorrowedRecords(userID)

	if len(list) == 0 {
		fmt.Println("📚 No books in histroy.")
		return
	}

	fmt.Printf("\n--- All Books borrowed for user: %d ---\n", userID)

	for _, value := range list {
		fmt.Println(value.DisplayInfo())
	}

}

func (cli *CLI) readInput(prompt string) string {
	fmt.Print(prompt)

	input, _ := cli.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (cli *CLI) readInt(prompt string) (int, error) {
	input := cli.readInput(prompt)
	return strconv.Atoi(input)
}
