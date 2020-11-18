package config

// Configuration struct
type Configuration struct {
	DBMS     string `default:"mysql"`
	USER     string `default:"root"`
	PASS     string `default:"hogehoge"`
	PROTOCOL string `default:"tcp(mysql:3306)"`
	DBNAME   string `default:"music_app"`
}
