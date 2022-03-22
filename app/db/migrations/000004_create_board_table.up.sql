CREATE TABLE IF NOT EXISTS boards(
  board_id serial PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  game_id INT REFERENCES games(game_id)
);
