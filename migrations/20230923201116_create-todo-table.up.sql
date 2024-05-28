CREATE TABLE IF NOT EXISTS todos (
   id serial PRIMARY KEY,
   user_id INTEGER NOT NULL,
   title VARCHAR(255) NOT NULL,
   content VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);