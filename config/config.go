package config

import "example/internal/common/helper/confighelper"

var defaultConfig = []byte(`
app: go-example
env: dev
mode: debug
http_address: 9999
database_postgres:
  host: 0.0.0.0
  port: 5434
  username: root
  password: Abc12345
  database: test
  schema: golangdb
  use_tls: false
  tls_mode: prefer
  tls_rootca_cert_file:
  tls_key_file:
  tls_cert_file:
  insecure_skip_verify: false
  logging_enabled: true
  use_logging_db: false
  use_logging_file: false
  auto_migration: true
  max_open_conns: 10
database_log:
  host: 0.0.0.0
  port: 5434
  username: root
  password: Abc12345
  database: test
  schema: logs
  use_tls: false
  tls_mode: prefer
  tls_rootca_cert_file:
  tls_key_file:
  tls_cert_file:
  insecure_skip_verify: false
  logging_enabled: true
  use_logging_db: false
  use_logging_file: false
  auto_migration: false
  max_open_conns: 10
redis_client:
  host: 0.0.0.0
  port: 6379
  username: 
  password: Abc12345
  db:
jwt:
  key_secret: ~!Messi!@#$Ronaldo%^&Marco832574Reus*()
  token_life_time: 60
  refresh_token_life_time: 240
`)

type (
	Config struct {
		App              string            `mapstructure:"app"`
		Env              string            `mapstructure:"env"`
		Mode             string            `mapstructure:"mode"`
		HttpAddress      uint32            `mapstructure:"http_address"`
		DatabasePostgres SqlDatabaseConfig `mapstructure:"database_postgres"`
		DatabaseLog      SqlDatabaseConfig `mapstructure:"database_log"`
		RedisClient      RedisConfig       `mapstructure:"redis_client"`
		Jwt              JwtConfig         `mapstructure:"jwt"`
	}

	SqlDatabaseConfig struct {
		Addrs               []string `mapstructure:"addrs"`
		Host                string   `mapstructure:"host"`
		Port                uint32   `mapstructure:"port"`
		Username            string   `mapstructure:"username"`
		Password            string   `mapstructure:"password"`
		Database            string   `mapstructure:"database"`
		Schema              string   `mapstructure:"schema"`
		UseTls              bool     `mapstructure:"use_tls"`
		TlsMode             string   `mapstructure:"tls_mode"`
		TlsRootCACertFile   string   `mapstructure:"tls_rootca_cert_file"`
		TlsKeyFile          string   `mapstructure:"tls_key_file"`
		TlsCertFile         string   `mapstructure:"tls_cert_file"`
		TlsRootCACertBase64 string   `mapstructure:"tls_rootca_cert_base64"`
		TlsKeyBase64        string   `mapstructure:"tls_key_base64"`
		TlsCertBase64       string   `mapstructure:"tls_cert_base64"`
		InsecureSkipVerify  bool     `mapstructure:"insecure_skip_verify"`
		MaxOpenConns        int      `mapstructure:"max_open_conns"`
		LoggingEnabled      bool     `mapstructure:"logging_enabled"`
		UseLoggingDb        bool     `mapstructure:"use_logging_db"`
		UseLoggingFile      bool     `mapstructure:"use_logging_file"`
		AutoMigration       bool     `mapstructure:"auto_migration"`
	}

	RedisConfig struct {
		Address  string `mapstructure:"addrs"`
		Host     string `mapstructure:"host"`
		Port     uint32 `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	}

	JwtConfig struct {
		JwtSecret            string `mapstructure:"key_secret"`
		TokenLifeTime        int64  `mapstructure:"token_life_time"`
		RefreshTokenLifeTime int64  `mapstructure:"refresh_token_life_time"`
	}
)

func LoadEnvironment() (*Config, error) {
	var cfg = &Config{}
	err := confighelper.Load(cfg, defaultConfig)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
