from typing import List
from models import Customer, Transaction
from book_management import BookManager
from customer_management import CustomerManager

class SalesManager:
    def __init__(self, book_manager: BookManager, customer_manager: CustomerManager):
        self.transactions: List[Transaction] = []
        self.book_manager = book_manager
        self.customer_manager = customer_manager
    
    def sell_book(self, customer_email: str, book_title: str, quantity: int) -> None:
        try:
            # Find customer
            customer = self.customer_manager.find_customer(customer_email)
            if not customer:
                raise ValueError("Customer not found")
            
            # Find book
            book = self.book_manager.search_book(book_title)
            if not book:
                raise ValueError("Book not found")
            
            # Check quantity
            if book.quantity < quantity:
                raise ValueError(f"Insufficient stock. Only {book.quantity} copies available")
            
            # Process sale
            total_price = book.price * quantity
            self.book_manager.update_book_quantity(book_title, -quantity)
            
            # Record transaction
            transaction = Transaction(customer, book_title, quantity, total_price)
            self.transactions.append(transaction)
            
            print(f"Sale successful! Total price: ${total_price:.2f}")
            
        except ValueError as e:
            print(f"Error processing sale: {str(e)}")
    
    def view_all_sales(self) -> None:
        if not self.transactions:
            print("No sales records found")
            return
            
        print("\nSales Records:")
        for transaction in self.transactions:
            print("\n" + "="*50)
            print(transaction.display_details())