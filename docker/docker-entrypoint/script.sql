CREATE DATABASE books_db;
USE books_db;

create table authors(
id int primary key auto_increment,
first_name varchar(60),
last_name varchar(60)
);

create table books(
id int primary key auto_increment,
book_title varchar(60),
book_description varchar(60)
);

create table authors_books(
author_id int,
book_id int,
primary key (author_id,book_id), 
constraint fk_authorid foreign key(author_id) references authors(id),
constraint fk_bookid foreign key(book_id) references books(id)
);
INSERT INTO  authors(first_name,last_name)
VALUES 
    ('Nikolai','Tukachev'),
    ('Alexadnr','Pushkin'),
    ('Petr','Petrov'),
    ('Anton','Golovanov'),
    ('Andrey','Nikitin'),
    ('German','Varis'),
    ('Lucifer','Bedro');


INSERT INTO books(book_title,book_description)
VALUES 
    ('Computer science for beginners','a lot of practice'),
    ('English from A1 to B2','easy to start'),
    ('Golang middle level','create apis from scratch'),
    ('Python for web developers','lear how to use django'),
    ('Photoshop for begginers','create your first meme in 5 minutes');

INSERT INTO authors_books(author_id,book_id)
VALUES 
    (1,1),
    (2,3),
    (3,1),
    (4,2),
    (5,4),
    (6,4),
    (7,2),
    (4,1);

CREATE VIEW view_authorsbooks
AS
SELECT a.id AS author_id ,a.first_name, a.last_name, b.book_title,b.book_description
  FROM authors a
LEFT OUTER JOIN authors_books ab
  ON a.id = ab.author_id
LEFT OUTER JOIN books b
  ON ab.book_id = b.id