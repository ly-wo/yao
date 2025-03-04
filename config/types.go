package config

// Config 象传应用引擎配置
type Config struct {
	Mode    string `json:"mode,omitempty" env:"YAO_ENV" envDefault:"production"`    // 象传引擎启动模式 production/development
	Root    string `json:"root,omitempty" env:"YAO_ROOT" envDefault:"."`            // 应用根目录
	Host    string `json:"host,omitempty" env:"YAO_HOST" envDefault:"0.0.0.0"`      // 服务监听地址
	Port    int    `json:"port,omitempty" env:"YAO_PORT" envDefault:"5099"`         // 服务监听端口
	Cert    string `json:"cert,omitempty" env:"YAO_CERT"`                           // HTTPS 证书文件地址
	Key     string `json:"key,omitempty" env:"YAO_KEY"`                             // HTTPS 证书密钥地址
	Log     string `json:"log,omitempty" env:"YAO_LOG"`                             // 服务日志地址
	LogMode string `json:"log_mode,omitempty" env:"YAO_LOG_MODE" envDefault:"TEXT"` // 服务日志模式 JSON|TEXT
	// Session   string        `json:"session,omitempty" env:"YAO_SESSION" envDefault:"memory"`         // 用户会话模式 memory|redis|database
	JWTSecret string        `json:"jwt_secret,omitempty" env:"YAO_JWT_SECRET"`                  // JWT 密钥
	DB        DBConfig      `json:"db,omitempty"`                                               // 数据库配置
	AllowFrom []string      `json:"allowfrom,omitempty" envSeparator:"|" env:"YAO_ALLOW_FROM" ` // Domain list the separator is |
	Session   SessionConfig `json:"session,omitempty"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver    string   `json:"driver,omitempty" env:"YAO_DB_DRIVER" envDefault:"sqlite3"`                        // 数据库驱动 sqlite3| mysql| postgres
	Primary   []string `json:"primary,omitempty" env:"YAO_DB_PRIMARY" envSeparator:"|" envDefault:"./db/yao.db"` // 主库连接DSN
	Secondary []string `json:"secondary,omitempty" env:"YAO_DB_SECONDARY" envSeparator:"|"`                      // 从库连接DSN
	AESKey    string   `json:"aeskey,omitempty" env:"YAO_DB_AESKEY"`                                             // 加密存储KEY
}
