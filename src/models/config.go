package models

type ServerSettings struct {
	ListenAddress string `yaml:"listen_address"`
}

type DatabaseSettings struct {
}

type AzureSettings struct {
}

type FileSettings struct {
	MaxSize int64 `yaml:"max_size"`
}

type Config struct {
	Server   ServerSettings   `yaml:"server_settings"`
	Database DatabaseSettings `yaml:"database_settings"`
	Azure    AzureSettings    `yaml:"azure_settings"`
	File     FileSettings     `yaml:"file_settings"`
}
