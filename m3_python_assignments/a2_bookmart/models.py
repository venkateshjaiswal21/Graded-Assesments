from datetime import datetime

class Book:
    def __init__(self, title: str, author: str, price: float, quantity: int):
        if price <= 0:
            raise ValueError("Price must be a positive number")
        if quantity < 0:
            raise ValueError("Quantity cannot be negative")
            
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity
    
    def display_details(self) -> str:
        return f"Title: {self.title}\nAuthor: {self.author}\nPrice: ${self.price:.2f}\nQuantity: {self.quantity}"
    
    def update_quantity(self, amount: int) -> None:
        new_quantity = self.quantity + amount
        if new_quantity < 0:
            raise ValueError("Cannot reduce quantity below 0")
        self.quantity = new_quantity

class Customer:
    def __init__(self, name: str, email: str, phone: str):
        if not all([name, email, phone]):
            raise ValueError("All customer details are required")
        if not '@' in email:
            raise ValueError("Invalid email format")
            
        self.name = name
        self.email = email
        self.phone = phone
    
    def display_details(self) -> str:
        return f"Name: {self.name}\nEmail: {self.email}\nPhone: {self.phone}"

class Transaction(Customer):
    def __init__(self, customer: Customer, book_title: str, quantity_sold: int, total_price: float):
        super().__init__(customer.name, customer.email, customer.phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold
        self.total_price = total_price
        self.transaction_date = datetime.now()
    
    def display_details(self) -> str:
        return (f"Transaction Details:\n"
                f"Date: {self.transaction_date.strftime('%Y-%m-%d %H:%M:%S')}\n"
                f"Customer: {self.name}\n"
                f"Book: {self.book_title}\n"
                f"Quantity: {self.quantity_sold}\n"
                f"Total Price: ${self.total_price:.2f}")