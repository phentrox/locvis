package entities

// LocalConfig paths is a bool pointer because the default is true
// If it is not present in the local config it would be parsed as false
// With the pointer we can check for nil and set it to true
type LocalConfig struct {
	Language    string   `yaml:"lang"`
	Top         int      `yaml:"top"`
	Paths       *bool    `yaml:"paths"`
	IgnoredDirs []string `yaml:"ignored-dirs"`
}
