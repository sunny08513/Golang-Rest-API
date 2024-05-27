# SQL 

1. Basic Query: Write a query to retrieve all authors from the authors table.
```
SELECT * FROM authors;
```

2. Filtering: Write a query to retrieve books published after '2000-01-01' from the books table.

```
SELECT * FROM books WHERE published_date > '2000-01-01';
```

3. Join: Write a query to retrieve the title of each book along with its author's name.
```
SELECT b.title, a.name as author_name
FROM books b
JOIN authors a ON b.author_id = a.author_id;
```

4. Aggregate Function: Write a query to calculate the total price of all orders.

```
SELECT SUM(total_amount) as total_price FROM orders;
```

4. Group By: Write a query to calculate the total price of each order.

```
SELECT order_id, SUM(price) as total_price
FROM order_items
GROUP BY order_id;
```

5. Subquery: Write a query to find the customers who have placed more than one order.
```
SELECT customer_id
FROM orders
GROUP BY customer_id
HAVING COUNT(order_id) > 1;
```

6. Indexing: Explain the importance of indexing in a database and how you would decide which columns to index.

```
Indexing is important because it improves the speed of data retrieval operations on a database table. It works by creating a sorted list of specific columns, allowing the database to quickly locate and retrieve data.

Columns that are frequently used in WHERE clauses, JOIN operations, and ORDER BY clauses are good candidates for indexing. However, indexing too many columns can lead to decreased performance in insert, update, and delete operations, so it's important to strike a balance.
```

# Sample query
**Query Example**

We will retrieve the names of customers, the number of orders they have placed, and the total amount spent on those orders. We will filter for customers who have placed more than one order, group the results by customer, and order the results by the total amount spent in descending order.

```
SELECT 
    c.name AS customer_name,
    COUNT(o.order_id) AS total_orders,
    SUM(o.total_amount) AS total_spent
FROM 
    customers c
JOIN 
    orders o ON c.customer_id = o.customer_id
WHERE 
    o.order_date >= '2024-01-01'
GROUP BY 
    c.customer_id
HAVING 
    COUNT(o.order_id) > 1
ORDER BY 
    total_spent DESC;

```