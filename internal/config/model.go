package config

type Config struct {
	Server    Server
	Databases Databases
	Mail      Email
}

type Server struct {
	Port string
	Host string
}
type Databases struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Email struct {
	Host     string
	Port     string
	Username string
	Password string
}
