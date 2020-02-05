CREATE USER wattx WITH PASSWORD 'rootroot';
CREATE DATABASE coindb OWNER wattx;

CREATE TABLE price_info (id SERIAL, coin_price numeric, coin_symbol varchar(10));
CREATE TABLE rank_info (id SERIAL, coin_symbol varchar(10), coin_rank integer);