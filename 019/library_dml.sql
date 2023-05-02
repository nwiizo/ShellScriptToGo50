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
