CREATE TABLE users(
   id UUID NOT NULL PRIMARY KEY,
   email VARCHAR(255) UNIQUE NOT NULL,
   password VARCHAR(255) NOT NULL
);