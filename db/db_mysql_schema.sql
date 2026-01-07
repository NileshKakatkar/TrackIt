CREATE TABLE user_info (
    uniq_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    email VARCHAR(150),
    age INT
);
CREATE TABLE billing_info (
    uniq_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    total_balance INT DEFAULT 0,
    month_income INT DEFAULT 0,
    expenses INT DEFAULT 0,
    upcoming_bills INT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE transactions (
    uniq_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    category VARCHAR(100),
    expense_type VARCHAR(50),
    amount INT NOT NULL,
    trans_date DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE budgets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    category VARCHAR(50) NOT NULL,
    budget_amount INT NOT NULL,
    month TINYINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uniq_user_category_month (username, category, month)
);
