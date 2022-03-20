CREATE TABLE IF NOT EXISTS cards(
  card_id serial PRIMARY KEY,
  game_id INTEGER REFERENCES games(game_id) ON DELETE CASCADE,
  -- as long as a tweet
  text VARCHAR(280) NOT NULL
)
