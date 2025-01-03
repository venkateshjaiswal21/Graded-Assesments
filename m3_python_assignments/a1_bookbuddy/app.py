from flask import Flask, request, jsonify
import sqlite3
from sqlite3 import Error
from datetime import datetime

app = Flask(__name__)

VALID_GENRES = {'Fiction', 'Non-Fiction', 'Mystery', 'Science Fiction', 'Romance', 'Horror', 'Biography', 'History'}

def get_db_connection():
    try:
        conn = sqlite3.connect('bookbuddy.db')
        conn.row_factory = sqlite3.Row
        return conn
    except Error as e:
        return None

def validate_book_data(data, check_required_fields=True):
    errors = []
    
    required_fields = {'title', 'author', 'published_year', 'genre'}
    if check_required_fields:
        missing_fields = required_fields - set(data.keys())
        if missing_fields:
            errors.append(f"Missing required fields: {', '.join(missing_fields)}")
    
    if 'published_year' in data:
        try:
            year = int(data['published_year'])
            current_year = datetime.now().year
            if year < 1000 or year > current_year:
                errors.append(f"Published year must be between 1000 and {current_year}")
        except ValueError:
            errors.append("Published year must be a valid number")
    
    if 'genre' in data and data['genre'] not in VALID_GENRES:
        errors.append(f"Genre must be one of: {', '.join(VALID_GENRES)}")
    
    return errors

@app.route('/books', methods=['POST'])
def add_book():
    """Add a new book"""
    data = request.get_json()
    
    errors = validate_book_data(data)
    if errors:
        return jsonify({'error': 'Invalid input', 'messages': errors}), 400
    
    conn = get_db_connection()
    if conn is None:
        return jsonify({'error': 'Database connection failed'}), 500
    
    try:
        cursor = conn.cursor()
        cursor.execute('''
            INSERT INTO books (title, author, published_year, genre)
            VALUES (?, ?, ?, ?)
        ''', (data['title'], data['author'], data['published_year'], data['genre']))
        conn.commit()
        book_id = cursor.lastrowid
        return jsonify({'message': 'Book added successfully', 'book_id': book_id}), 201
    except Error as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500
    finally:
        conn.close()

@app.route('/books', methods=['GET'])
def get_books():
    """Get all books with optional filtering"""
    genre = request.args.get('genre')
    author = request.args.get('author')
    
    conn = get_db_connection()
    if conn is None:
        return jsonify({'error': 'Database connection failed'}), 500
    
    try:
        cursor = conn.cursor()
        query = "SELECT * FROM books"
        params = []
        
        if genre or author:
            conditions = []
            if genre:
                conditions.append("genre = ?")
                params.append(genre)
            if author:
                conditions.append("author = ?")
                params.append(author)
            query += " WHERE " + " AND ".join(conditions)
        
        cursor.execute(query, params)
        books = [dict(row) for row in cursor.fetchall()]
        return jsonify(books), 200
    except Error as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500
    finally:
        conn.close()

@app.route('/books/<int:id>', methods=['GET'])
def get_book(id):
    """Get a specific book by ID"""
    conn = get_db_connection()
    if conn is None:
        return jsonify({'error': 'Database connection failed'}), 500
    
    try:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (id,))
        book = cursor.fetchone()
        
        if book is None:
            return jsonify({
                'error': 'Book not found',
                'message': 'No book exists with the provided ID'
            }), 404
            
        return jsonify(dict(book)), 200
    except Error as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500
    finally:
        conn.close()

@app.route('/books/<int:id>', methods=['PUT'])
def update_book(id):
    data = request.get_json()
    
    errors = validate_book_data(data, check_required_fields=False)
    if errors:
        return jsonify({'error': 'Invalid input', 'messages': errors}), 400
    
    conn = get_db_connection()
    if conn is None:
        return jsonify({'error': 'Database connection failed'}), 500
    
    try:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (id,))
        if cursor.fetchone() is None:
            return jsonify({
                'error': 'Book not found',
                'message': 'No book exists with the provided ID'
            }), 404
        
        update_fields = []
        params = []
        for field in ['title', 'author', 'published_year', 'genre']:
            if field in data:
                update_fields.append(f"{field} = ?")
                params.append(data[field])
        
        if update_fields:
            query = f"UPDATE books SET {', '.join(update_fields)} WHERE id = ?"
            params.append(id)
            cursor.execute(query, params)
            conn.commit()
            
        return jsonify({'message': 'Book updated successfully'}), 200
    except Error as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500
    finally:
        conn.close()

@app.route('/books/<int:id>', methods=['DELETE'])
def delete_book(id):
    conn = get_db_connection()
    if conn is None:
        return jsonify({'error': 'Database connection failed'}), 500
    
    try:
        cursor = conn.cursor()
        cursor.execute('SELECT * FROM books WHERE id = ?', (id,))
        if cursor.fetchone() is None:
            return jsonify({
                'error': 'Book not found',
                'message': 'No book exists with the provided ID'
            }), 404
            
        cursor.execute('DELETE FROM books WHERE id = ?', (id,))
        conn.commit()
        return jsonify({'message': 'Book deleted successfully'}), 200
    except Error as e:
        return jsonify({'error': 'Database error', 'message': str(e)}), 500
    finally:
        conn.close()

if __name__ == '__main__':
    app.run(debug=True)
