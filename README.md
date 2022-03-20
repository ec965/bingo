# Bingo

## Data

**Users**

| field    | type   | note   |
| -------- | ------ | ------ |
| id       | int    | serial |
| username | string |        |
| password | string |        |

**Game**

| field     | type | note                                                 |
| --------- | ---- | ---------------------------------------------------- |
| id        | int  | serial                                               |
| dimension | int  | boards are squares, we only need 1 dimensional value |

**Card**

| field   | type   | note       |
| ------- | ------ | ---------- |
| id      | int    | serial     |
| game_id | int    | fk -> game |
| text    | string |            |

**Board**

| field   | type | note       |
| ------- | ---- | ---------- |
| id      | int  | serial     |
| user_id | int  | fk -> user |

**Card Board Junction**

A board is made of many cards.

| field    | type    | note                                           |
| -------- | ------- | ---------------------------------------------- |
| id       | int     | serial                                         |
| board_id | int     | fk -> board                                    |
| card_id  | int     | fk -> card                                     |
| row      | int     | row index of the card on the board's matrix    |
| column   | int     | column index of the card on the board's matrix |
| complete | boolean |                                                |
