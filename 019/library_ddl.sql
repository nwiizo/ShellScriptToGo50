-- テーブルの作成: authors
CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL
);

-- テーブルの作成: books
CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author_id INTEGER REFERENCES authors(id),
  publication_year INTEGER,
  available BOOLEAN DEFAULT TRUE
);

-- テーブルの作成: users
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  birth_date DATE NOT NULL
);

-- テーブルの作成: loans
CREATE TABLE loans (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  loan_date DATE NOT NULL,
  due_date DATE NOT NULL,
  returned BOOLEAN DEFAULT FALSE
);
