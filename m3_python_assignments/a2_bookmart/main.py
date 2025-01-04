from book_management import BookManager
from customer_management import CustomerManager
from sales_management import SalesManager

def get_valid_float(prompt: str) -> float:
    while True:
        try:
            value = float(input(prompt))
            if value <= 0:
                print("Please enter a positive number")
                continue
            return value
        except ValueError:
            print("Please enter a valid number")

def get_valid_int(prompt: str) -> int:
    while True:
        try:
            value = int(input(prompt))
            if value <= 0:
                print("Please enter a positive number")
                continue
            return value
        except ValueError:
            print("Please enter a valid number")

def book_menu(book_manager: BookManager):
    while True:
        print("\nBook Management")
        print("1. Add Book")
        print("2. View All Books")
        print("3. Search Book")
        print("4. Return to Main Menu")
        
        choice = input("Enter your choice: ")
        
        if choice == "1":
            title = input("Enter book title: ")
            author = input("Enter author name: ")
            price = get_valid_float("Enter price: $")
            quantity = get_valid_int("Enter quantity: ")
            book_manager.add_book(title, author, price, quantity)
            
        elif choice == "2":
            book_manager.view_all_books()
            
        elif choice == "3":
            search_term = input("Enter book title or author to search: ")
            book = book_manager.search_book(search_term)
            if book:
                print("\nBook found:")
                print(book.display_details())
            else:
                print("Book not found")
                
        elif choice == "4":
            break
            
        else:
            print("Invalid choice")

def customer_menu(customer_manager: CustomerManager):
    while True:
        print("\nCustomer Management")
        print("1. Add Customer")
        print("2. View All Customers")
        print("3. Return to Main Menu")
        
        choice = input("Enter your choice: ")
        
        if choice == "1":
            name = input("Enter customer name: ")
            email = input("Enter customer email: ")
            phone = input("Enter customer phone: ")
            customer_manager.add_customer(name, email, phone)
            
        elif choice == "2":
            customer_manager.view_all_customers()
            
        elif choice == "3":
            break
            
        else:
            print("Invalid choice")

def sales_menu(sales_manager: SalesManager):
    while True:
        print("\nSales Management")
        print("1. Sell Book")
        print("2. View All Sales")
        print("3. Return to Main Menu")
        
        choice = input("Enter your choice: ")
        
        if choice == "1":
            customer_email = input("Enter customer email: ")
            book_title = input("Enter book title: ")
            quantity = get_valid_int("Enter quantity: ")
            sales_manager.sell_book(customer_email, book_title, quantity)
            
        elif choice == "2":
            sales_manager.view_all_sales()
            
        elif choice == "3":
            break
            
        else:
            print("Invalid choice")

def main():
    book_manager = BookManager()
    customer_manager = CustomerManager()
    sales_manager = SalesManager(book_manager, customer_manager)
    
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        
        choice = input("Enter your choice: ")
        
        if choice == "1":
            book_menu(book_manager)
        elif choice == "2":
            customer_menu(customer_manager)
        elif choice == "3":
            sales_menu(sales_manager)
        elif choice == "4":
            print("Thank you for using BookMart!")
            break
        else:
            print("Invalid choice")

if __name__ == "__main__":
    main()