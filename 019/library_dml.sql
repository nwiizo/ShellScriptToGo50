-- データの挿入: authors
INSERT INTO authors (name, country) VALUES ('夏目 漱石', 'Japan');
INSERT INTO authors (name, country) VALUES ('芥川 龍之介', 'Japan');
INSERT INTO authors (name, country) VALUES ('宮沢 賢治', 'Japan');

-- データの挿入: books
INSERT INTO books (title, author_id, publication_year) VALUES ('吾輩は猫である', 1, 1905);
INSERT INTO books (title, author_id, publication_year) VALUES ('羅生門', 2, 1915);
INSERT INTO books (title, author_id, publication_year) VALUES ('銀河鉄道の夜', 3, 1934);

-- データの挿入: users
INSERT INTO users (name, email, birth_date) VALUES ('山田 太郎', 'taro.yamada@example.com', '2000-01-01');
INSERT INTO users (name, email, birth_date) VALUES ('鈴木 次郎', 'jiro.suzuki@example.com', '1995-02-14');
INSERT INTO users (name, email, birth_date) VALUES ('佐藤 三郎', 'saburo.sato@example.com', '1980-12-31');

-- データの挿入: loans
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (1, 1, '2022-01-01', '2022-01-15');
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (1, 2, '2022-01-01', '2022-01-15');
INSERT INTO loans (user_id, book_id, loan_date, due_date) VALUES (2, 3, '2022-01-05', '2022-01-19');

-- Authors
INSERT INTO authors (name) VALUES ('Haruki Murakami');
INSERT INTO authors (name) VALUES ('Banana Yoshimoto');
INSERT INTO authors (name) VALUES ('Ryu Murakami');
INSERT INTO authors (name) VALUES ('Yoko Ogawa');
INSERT INTO authors (name) VALUES ('Higashino Keigo');

-- Books
INSERT INTO books (title, author_id) VALUES ('Norwegian Wood', 1);
INSERT INTO books (title, author_id) VALUES ('Kafka on the Shore', 1);
INSERT INTO books (title, author_id) VALUES ('Kitchen', 2);
INSERT INTO books (title, author_id) VALUES ('In the Miso Soup', 3);
INSERT INTO books (title, author_id) VALUES ('The Housekeeper and the Professor', 4);
INSERT INTO books (title, author_id) VALUES ('The Devotion of Suspect X', 5);

-- Users
INSERT INTO users (name) VALUES ('Yuki Tanaka');
INSERT INTO users (name) VALUES ('Akiko Suzuki');
INSERT INTO users (name) VALUES ('Kenji Sato');
INSERT INTO users (name) VALUES ('Mika Yamamoto');

-- Loans
INSERT INTO loans (user_id, book_id, due_date) VALUES (1, 1, '2023-06-10');
INSERT INTO loans (user_id, book_id, due_date) VALUES (1, 2, '2023-06-10');
INSERT INTO loans (user_id, book_id, due_date) VALUES (2, 3, '2023-06-15');
INSERT INTO loans (user_id, book_id, due_date) VALUES (3, 4, '2023-05-30');
INSERT INTO loans (user_id, book_id, due_date) VALUES (4, 5, '2023-06-20');
