import sqlite3
from sqlite3 import Error

def create_connection():
    try:
        conn = sqlite3.connect('bookbuddy.db')
        return conn
    except Error as e:
        print(f"Error connecting to database: {e}")
        return None

def create_table(conn):
    try:
        cursor = conn.cursor()
        cursor.execute('''
            CREATE TABLE IF NOT EXISTS books (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT NOT NULL,
                author TEXT NOT NULL,
                published_year INTEGER NOT NULL,
                genre TEXT NOT NULL
            )
        ''')
        conn.commit()
    except Error as e:
        print(f"Error creating table: {e}")

def insert_sample_data(conn):
    sample_books = [
        ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction'),
        ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction'),
        ('1984', 'George Orwell', 1949, 'Science Fiction'),
        ('Pride and Prejudice', 'Jane Austen', 1813, 'Romance')
    ]
    
    try:
        cursor = conn.cursor()
        cursor.executemany('''
            INSERT INTO books (title, author, published_year, genre)
            VALUES (?, ?, ?, ?)
        ''', sample_books)
        conn.commit()
    except Error as e:
        print(f"Error inserting sample data: {e}")

def initialize_database():
    """Initialize the database with the table and sample data"""
    conn = create_connection()
    if conn is not None:
        create_table(conn)
        insert_sample_data(conn)
        conn.close()
        print("Database initialized successfully!")
    else:
        print("Error: Could not establish database connection")

if __name__ == '__main__':
    initialize_database()
