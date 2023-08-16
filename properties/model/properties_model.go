package properties_model

type YamlConfig struct {
	Database DbProp   `mapstructure:"database"`
	Auth     AuthProp `mapstructure:"auth"`
}

type DbProp struct {
	ConfigMySql MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Address   string `mapstructure:"address"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Port      string `mapstructure:"port"`
	Database  string `mapstructure:"database"`
	Populated bool   `mapstructure:"populated"`
	Migrate   bool   `mapstructure:"migrate"`
}

type AuthProp struct {
	ConfigBasicAuth BasicAuthConfig `mapstructure:"basic"`
}

type BasicAuthConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
