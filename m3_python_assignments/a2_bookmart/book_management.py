from typing import List, Optional
from models import Book

class BookManager:
    def __init__(self):
        self.books: List[Book] = []
    
    def add_book(self, title: str, author: str, price: float, quantity: int) -> None:
        try:
            existing_book = self.search_book(title)
            if existing_book:
                existing_book.quantity += quantity
                print(f"Updated quantity for existing book: {title}")
                return
                
            new_book = Book(title, author, price, quantity)
            self.books.append(new_book)
            print("Book added successfully!")
        except ValueError as e:
            print(f"Error adding book: {str(e)}")
    
    def view_all_books(self) -> None:
        if not self.books:
            print("No books in inventory")
            return
            
        print("\nCurrent Inventory:")
        for book in self.books:
            print("\n" + "="*40)
            print(book.display_details())
    
    def search_book(self, search_term: str) -> Optional[Book]:
        search_term = search_term.lower()
        for book in self.books:
            if search_term in book.title.lower() or search_term in book.author.lower():
                return book
        return None
    
    def update_book_quantity(self, title: str, change: int) -> None:
        book = self.search_book(title)
        if not book:
            raise ValueError(f"Book '{title}' not found")
        book.update_quantity(change)