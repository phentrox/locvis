package localconfig

import "locvis/entities"

func GetTop(localConfig entities.LocalConfig) int {
	if localConfig.Top != 0 {
		return localConfig.Top
	}
	return 10
}
