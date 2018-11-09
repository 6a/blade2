package db

import (
	"log"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

const databasePath = "../../db/"

// Note: https://www.tutorialspoint.com/sqlite/sqlite_indexed_by.htm
var schemas = [3]string{
	`CREATE TABLE IF NOT EXISTS player (
	player_id varchar NOT NULL,
	player_name varchar NOT NULL UNIQUE,
	profile_picture varchar NOT NULL,
	bio varchar NOT NULL,
	email varchar NOT NULL UNIQUE,
	password_hash varchar NOT NULL,
	token varchar NOT NULL,
	token_expiry INT NOT NULL DEFAULT '0',
	PRIMARY KEY (player_id)
);
CREATE INDEX player_name_idx ON player (player_name);
CREATE INDEX email_idx ON player (email);`,

	`CREATE TABLE IF NOT EXISTS rating (
	player_id varchar NOT NULL,
	elo INT NOT NULL DEFAULT '1200',
	wins INT NOT NULL DEFAULT '0',
	losses INT NOT NULL DEFAULT '0',
	draws INT NOT NULL DEFAULT '0',
	PRIMARY KEY (player_id)
);`,

	`CREATE TABLE IF NOT EXISTS game (
	game_id INT NOT NULL AUTO_INCREMENT DEFAULT '0',
	player_id_1 varchar NOT NULL,
	player_id_2 varchar NOT NULL,
	player_1_elo_start INT NOT NULL DEFAULT '1200',
	player_1_elo_end INT NOT NULL DEFAULT '1200',
	player_2_elo_start INT NOT NULL DEFAULT '1200',
	player_2_elo_end INT NOT NULL DEFAULT '1200',
	game_end_time INT NOT NULL DEFAULT '0',
	winner INT NOT NULL,
	state_start varchar NOT NULL,
	actions varchar NOT NULL,
	PRIMARY KEY (match_id)
);
CREATE INDEX player_id_1_idx ON game (player_id_1);
CREATE INDEX player_id_2_idx ON game (player_id_2);
CREATE INDEX game_end_time_idx ON game (game_end_time);`}

var sqldb = ""

// Initialize the database and create the database if required
func Initialize(database string) {
	sqldb = databasePath + database

	conn, err := sqlite3.Open(sqldb)
	if err != nil {
		log.Printf("Could not connect to the database [%s]", sqldb)
		log.Fatalf("Error message: [%s]", err)
	}

	defer conn.Close()

	err := conn.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction: [%s]", err)
	}

	for _, command := range schemas {
		err = conn.Exec(command)
		if err != nil {
			log.Printf("Could not create the a table using the following command: [%s]", command)
			log.Fatalf("Error message: [%s]", err)
		}
	}

	err := conn.Commit()
	if err != nil {
		log.Fatalf("Error committing transaction: [%s]", err)
	}
}
