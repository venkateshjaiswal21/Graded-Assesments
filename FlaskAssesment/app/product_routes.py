from flask import Blueprint, render_template, request, redirect

product_routes = Blueprint('product_routes', __name__)

# In Memory List of Products
products = []

# Routes for Products
# To get the list of products


@product_routes.route('/products')
def index():
    return render_template('products/index.html', products=products)

# To create a new product


@product_routes.route('/products/create', methods=['GET', 'POST'])
def create():
    if request.method == 'POST':
        new_product = {
            'id': request.form['id'], 'name': request.form['name'], 'price': request.form['price']}
        if new_product:
            products.append(new_product)
            print(products)
        return redirect('/products')
    return render_template('products/create.html')

# To Update a product


@product_routes.route('/products/update/<int:product_id>', methods=['GET', 'POST'])
def update(product_id):
    if request.method == 'POST':
        updated_product = {
            'id': request.form['id'], 'name': request.form['name'], 'price': request.form['price']}
        if updated_product:
            products[product_id-1] = updated_product
            print(products)
        return redirect('/products')
    return render_template('products/update.html', product=products[product_id-1], product_id=product_id)

# To delete a product


@product_routes.route('/products/delete/<int:product_id>')
def delete(product_id):
    # products.pop(product_id-1)
    del products[product_id-1]
    return redirect('/products')
