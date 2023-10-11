package config

// 程序启动时, 会将所有 Config 保存在 Global
type Config struct {
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
