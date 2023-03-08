package config

type (
	Db struct {
		Mysql *Mysql
		Redis *Redis
	}
	Redis struct {
		DB       int    `mapstructure:"db"`
		PoolSize int    `mapstructure:"pool-size"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
	}
	Mysql struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"dbName"`
	}
	Jwt struct {
		ExpiresTime int    `mapstructure:"expires-time"`
		SecretKey   string `mapstructure:"secret-key"`
	}
	Config struct {
		Database *Db     `mapstructure:"database"`
		Jwt      *Jwt    `mapstructure:"jwt"`
		Oauth2   *Oauth2 `mapstructure:"oauth2"`
	}
	Oauth2 struct {
		Client []*Oauth2Client `mapstructure:"client"`
	}
	Oauth2Client struct {
		ClientID     string   `mapstructure:"id"`
		ClientSecret string   `mapstructure:"secret"`
		Name         string   `mapstructure:"name"`
		Domain       string   `mapstructure:"domain"`
		Scope        []*Scope `mapstructure:"scope"`
	}
	Scope struct {
		id    int8   `mapstructure:"id"`
		title string `mapstructure:"title"`
	}
)
