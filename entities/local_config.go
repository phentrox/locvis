package entities

type LocalConfig struct {
	Language    string   `yaml:"lang"`
	Top         string   `yaml:"top"`
	IgnoredDirs []string `yaml:"ignored-dirs"`
}
