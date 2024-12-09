package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(albums)

	album, err := albumById(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(album)

	newAlbum := Album{
		Title:  "Highway to Hell",
		Artist: "AC/DC",
		Price:  259.99,
	}
	id, err := addAlbum(newAlbum)
	if err != nil {
		log.Fatalf("addAlbum error: %v", err)
	}
	log.Printf("Last id: %d\n", id)

}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist =?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumById(id int64) (Album, error) {
	var album Album
	row := db.QueryRow("SELECT * FROM album WHERE id=?", id)
	// order is important while scanning.
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("invalid id: %d", id)
		}
		return album, fmt.Errorf("error while fetching album by id: %d", id)
	}
	return album, nil
}

func addAlbum(album Album) (int64, error) {
	res, err := db.Exec("INSERT INTO album(title, artist, price) VALUES (?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}
	return id, nil
}
