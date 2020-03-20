CREATE TABLE price_info (id SERIAL, coin_price numeric, coin_symbol varchar(10));
ALTER TABLE price_info ADD constraint price_symbol UNIQUE (coin_symbol);

CREATE TABLE rank_info (id SERIAL, coin_symbol varchar(10), coin_rank integer);
ALTER TABLE rank_info ADD constraint rank_symbol UNIQUE (coin_symbol);