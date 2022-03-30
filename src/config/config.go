package config

import (
	"strconv"
	"time"

	c "github.com/joaopandolfi/blackwhale/configurations"
	"github.com/unrolled/secure"
	"golang.org/x/xerrors"
)

// Configs Struct
type Config struct {
	File           map[string]string
	AESKey         string
	SystemID       int
	Propertyes     c.Configurations
	Server         server     `json:"server"`
	Monitoring     monitoring `json:"monitoring"`
	SnakeByDefault bool
}

type server struct {
	Port         string
	Host         string
	TimeoutWrite time.Duration
	TimeoutRead  time.Duration
	Debug        bool
	Security     security
}

type monitoring struct {
}

type security struct {
	TLSCert    string
	TLSKey     string
	Opsec      secure.Options
	Resethash  string
	BcryptCost int //10,11,12,13,14
	JWTSecret  string
	AESKey     string
}

// Config global
var cfg *Config

// Get Config
func Get() Config {
	if cfg == nil {
		panic(xerrors.Errorf("config not loaded"))
	}

	return *cfg
}

// Load config
func Load(args []string) {
	cfile := "./config.json"
	cfg = &Config{
		File: c.LoadJsonFile(cfile),
	}
	c.LoadConfig(c.LoadFromFile(cfile))
	cfg.Propertyes = c.Configuration
	systemID, _ := strconv.Atoi(cfg.File["SYSTEM_ID"])
	cfg.SystemID = systemID
	cfg.AESKey = cfg.File["AES_KEY"]
	cfg.Server.Security.AESKey = cfg.AESKey
	cfg.SnakeByDefault, _ = strconv.ParseBool(cfg.File["SNAKE_DEFAULT"])
}

func Inject(c *Config) {
	cfg = c
}
