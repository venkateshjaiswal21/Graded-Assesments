class BorrowLimitExceededException(Exception):
    pass

class BookAlreadyBorrowedException(Exception):
    pass

class BookNotBorrowedException(Exception):
    pass


class Book:
    def __init__(self, book_id, title, author, category):
        if not title.strip() or not author.strip():
            raise ValueError("Book title and author cannot be empty.")
        
        self.book_id = book_id
        self.title = title
        self.author = author
        self.category = category
        self.is_borrowed = False

    def borrow(self):
        if self.is_borrowed:
            raise BookAlreadyBorrowedException(f"'{self.title}' is already borrowed.")
        self.is_borrowed = True

    def return_book(self):
        if not self.is_borrowed:
            raise BookNotBorrowedException(f"'{self.title}' was not borrowed.")
        self.is_borrowed = False


class Member:
    def __init__(self, member_id, name, member_type):
        if not name.strip():
            raise ValueError("Member name cannot be empty.")
        
        self.member_id = member_id
        self.name = name
        self.max_books = 10 if member_type == 'subscribed' else 7
        self.borrowed_books = set()

    def can_borrow(self):
        return len(self.borrowed_books) < self.max_books

    def borrow_book(self, book):
        if len(self.borrowed_books) >= self.max_books:
            raise BorrowLimitExceededException(f"{self.name} cannot borrow more than {self.max_books} books.")
        self.borrowed_books.add(book.book_id)

    def return_book(self, book):
        self.borrowed_books.remove(book.book_id)


class Library:
    def __init__(self):
        self.books = {}
        self.members = {}

    def add_book(self, book):
        if book.book_id in self.books:
            raise ValueError(f"Book ID {book.book_id} already exists in the library.")
        self.books[book.book_id] = book
        print(f"Book '{book.title}' added to the library.")

    def register_member(self, member):
        if member.member_id in self.members:
            raise ValueError(f"Member ID {member.member_id} already exists.")
        self.members[member.member_id] = member
        print(f"Member '{member.name}' registered.")

    def lend_book(self, member_id, book_id):
        member = self._get_member(member_id)
        book = self._get_book(book_id)

        if not member.can_borrow():
            raise BorrowLimitExceededException(f"{member.name} has reached their borrowing limit.")

        book.borrow()
        member.borrow_book(book)
        print(f"{member.name} borrowed '{book.title}'.")

    def receive_book(self, member_id, book_id):
        member = self._get_member(member_id)
        book = self._get_book(book_id)

        book.return_book()
        member.return_book(book)
        print(f"{member.name} returned '{book.title}'.")

    def _get_member(self, member_id):
        if member_id not in self.members:
            raise ValueError(f"Member with ID {member_id} does not exist.")
        return self.members[member_id]

    def _get_book(self, book_id):
        if book_id not in self.books:
            raise ValueError(f"Book with ID {book_id} does not exist.")
        return self.books[book_id]


def main():
    library = Library()

    # Add books to the library
    library.add_book(Book(10, "The Power of Habit", "Charles Duhigg", "Self-help"))
    library.add_book(Book(11, "Sapiens", "Yuval Noah Harari", "History"))
    library.add_book(Book(12, "Educated", "Tara Westover", "Memoir"))
    library.add_book(Book(13, "Becoming", "Michelle Obama", "Biography"))

    # Register members
    library.register_member(Member(201, "Alice Johnson", "regular"))
    library.register_member(Member(202, "Bob Smith", "subscribed"))

    # normal lending and returning process
    library.lend_book(201, 10)  # Alice borrows "The Power of Habit"
    library.lend_book(201, 11)  # Alice borrows "Sapiens"
    library.receive_book(201, 10)  # Alice returns "The Power of Habit"
    library.receive_book(201, 11)  # Alice returns "Sapiens"

    # borrowing beyond the limit for a "regular" member
    try:
        library.lend_book(201, 10)
        library.lend_book(201, 11)
        library.lend_book(201, 12)
        library.lend_book(201, 13)  # Should raise BorrowLimitExceededException
    except BorrowLimitExceededException as e:
        print(e)

    # borrowing and returning for a "subscribed" member
    library.lend_book(202, 13)  # Bob borrows "Becoming"
    library.receive_book(202, 13)  # Bob returns "Becoming"

    # returning a book that was not borrowed
    try:
        library.receive_book(202, 10)  # Bob tries to return a book he never borrowed
    except BookNotBorrowedException as e:
        print(e)

    # trying to borrow a book that is already borrowed
    library.lend_book(201, 12)  # Alice borrows "Educated"
    try:
        library.lend_book(202, 12)  # Bob tries to borrow "Educated" that Alice has already borrowed
    except BookAlreadyBorrowedException as e:
        print(e)

    # adding a book that already exists
    try:
        library.add_book(Book(10, "The Power of Habit", "Charles Duhigg", "Self-help"))  # Duplicate book ID
    except ValueError as e:
        print(e)

    # registering a member that already exists
    try:
        library.register_member(Member(201, "Alice Johnson", "regular"))  # Duplicate member ID
    except ValueError as e:
        print(e)

    # non-existent member and book cases
    try:
        library.lend_book(999, 10)  # Non-existent member tries to borrow
    except ValueError as e:
        print(e)

    try:
        library.lend_book(201, 999)  # Alice tries to borrow a non-existent book
    except ValueError as e:
        print(e)

    # Test that the limit for subscribed member is higher
    library.lend_book(202, 10)  # Bob borrows "The Power of Habit"
    library.lend_book(202, 11)  # Bob borrows "Sapiens"
    library.lend_book(202, 12)  # Bob borrows "Educated"
    library.lend_book(202, 13)  # Bob borrows "Becoming"
    # Bob is allowed to borrow more books since he's a subscribed member (limit is 10)
    print(f"Bob currently borrowed {len(library.members[202].borrowed_books)} books.")

if __name__ == "__main__":
    main()
