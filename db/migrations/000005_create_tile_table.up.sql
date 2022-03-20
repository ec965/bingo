CREATE TABLE IF NOT EXISTS tiles(
  tile_id serial PRIMARY KEY,
  board_id INT REFERENCES boards(board_id) ON DELETE CASCADE,
  card_id INT REFERENCES cards(card_id) ON DELETE CASCADE,
  row INT NOT NULL CHECK (row >= 0),
  col INT NOT NULL CHECK (col >= 0),
  complete BOOLEAN DEFAULT 'f'
)
