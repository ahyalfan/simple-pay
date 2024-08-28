package config

type Config struct {
	Server    Server
	Databases Databases
	Mail      Email
	Redis     Redis
	Midtrans  Midtrans
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

type Redis struct {
	Addr     string
	Password string
	DB       string
}

type Midtrans struct {
	Key    string
	IsProd bool
}
