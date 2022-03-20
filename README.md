# Bingo

## Tools

- docker / docker-compose
- golang-migrate

## Data

**Users**

| field    | type   | note   |
| -------- | ------ | ------ |
| user_id  | int    | serial |
| username | string |        |
| password | string |        |

**Game**

| field     | type | note                                                 |
| --------- | ---- | ---------------------------------------------------- |
| game_id   | int  | serial                                               |
| dimension | int  | boards are squares, we only need 1 dimensional value |

**Card**

| field   | type   | note       |
| ------- | ------ | ---------- |
| card_id | int    | serial     |
| game_id | int    | fk -> game |
| text    | string |            |

**Board**

| field    | type | note       |
| -------- | ---- | ---------- |
| board_id | int  | serial     |
| user_id  | int  | fk -> user |
| game_id  | int  | fk -> game |

**Tile**

A board is made of many tiles. Each tile has a card on top of it.

| field    | type    | note                                           |
| -------- | ------- | ---------------------------------------------- |
| tile_id  | int     | serial                                         |
| board_id | int     | fk -> board                                    |
| card_id  | int     | fk -> card                                     |
| row      | int     | row index of the card on the board's matrix    |
| column   | int     | column index of the card on the board's matrix |
| complete | boolean |                                                |
