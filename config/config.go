package config

import "example/internal/common/helper/confighelper"

var defaultConfig = []byte(`
	app: go-example
	env: test
	http_address: 9999
`)

// database_postgres:
// 		host: 0.0.0.0
// 		port: 5434
// 		username: root
// 		password: Abc12345
// 		database: test
// 		schema: golangdb
// 		use_tls: false
// 		tls_mode: prefer
// 		tls_rootca_cert_file:
// 		tls_key_file:
// 		tls_cert_file:
// 		insecure_skip_verify: false
// 		logging_enabled: false
// 		use_logging_db: false
// 		auto_migration: false
// 		max_open_conns: 10

type (
	Config struct {
		App              string            `mapstructure:"app"`
		Env              string            `mapstructure:"env"`
		HttpAddress      uint32            `mapstructure:"http_address"`
		DatabasePostgres SqlDatabaseConfig `mapstructure:"database_postgres"`
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
		AutoMigration       bool     `mapstructure:"auto_migration"`
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
