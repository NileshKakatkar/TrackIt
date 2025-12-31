CREATE TABLE user_info (
    uniq_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    password TEXT,
    name TEXT,
    email TEXT,
    age INTEGER
);

CREATE TABLE billing_info(
    uniq_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    total_balance INTEGER,
    month_income INTEGER,
    expenses INTEGER,
    upcoming_bills INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions(
    uniq_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    description TEXT,
    category TEXT,
    expense_type TEXT,
    amount INTEGER,
    trans_date DATETIME DEFAULT CURRENT_TIMESTAMP
);
