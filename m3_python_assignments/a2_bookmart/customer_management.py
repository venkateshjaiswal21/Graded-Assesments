from typing import List, Optional
from models import Customer

class CustomerManager:
    def __init__(self):
        self.customers: List[Customer] = []
    
    def add_customer(self, name: str, email: str, phone: str) -> Optional[Customer]:
        try:
            # Check if customer already exists
            for customer in self.customers:
                if customer.email == email:
                    print("Customer with this email already exists")
                    return customer
                    
            new_customer = Customer(name, email, phone)
            self.customers.append(new_customer)
            print("Customer added successfully!")
            return new_customer
        except ValueError as e:
            print(f"Error adding customer: {str(e)}")
            return None
    
    def view_all_customers(self) -> None:
        if not self.customers:
            print("No customers registered")
            return
            
        print("\nRegistered Customers:")
        for customer in self.customers:
            print("\n" + "="*40)
            print(customer.display_details())
    
    def find_customer(self, email: str) -> Optional[Customer]:
        for customer in self.customers:
            if customer.email.lower() == email.lower():
                return customer
        return None