CREATE TABLE IF NOT EXISTS stations(
    id INTEGER PRIMARY KEY,
    console_id INTEGER NOT NULL);

CREATE TABLE IF NOT EXISTS consoles(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    image TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS console_controllers(
    id INTEGER PRIMARY KEY,
    console_id INTEGER NOT NULL,
    controller_id INTEGER NOT NULL);

CREATE TABLE IF NOT EXISTS controllers(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    image TEXT,
    count INTEGER NOT NULL);

CREATE TABLE IF NOT EXISTS games(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    count INTEGER NOT NULL);

CREATE TABLE IF NOT EXISTS barcodes(
    id INTEGER PRIMARY KEY,
    game_id INTEGER NOT NULL,
    barcode TEXT NOT NULL);

CREATE TABLE IF NOT EXISTS rentals(
    id INTEGER PRIMARY KEY,
    badge_id INTEGER NOT NULL,
    controller_id INTEGER,
    game_id INTEGER);

CREATE TABLE IF NOT EXISTS transactions(
    id INTEGER PRIMARY KEY,
    type TEXT NOT NULL,
    badge_id INTEGER,
    station_id INTEGER,
    game_id INTEGER,
    controller_id INTEGER,
    created TEXT DEFAULT now());
