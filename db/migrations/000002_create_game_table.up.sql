CREATE TABLE IF NOT EXISTS games(
  game_id serial PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  dimension INTEGER NOT NULL,
  user_id INTEGER REFERENCES users(user_id)
);
