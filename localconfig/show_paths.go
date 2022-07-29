package localconfig

import "locvis/entities"

func GetShowPaths(localConfig entities.LocalConfig) bool {
	var showPathsPointer *bool = localConfig.Paths
	if showPathsPointer == nil {
		// show paths is not present in local config
		// return default true for show paths
		return true
	} else {
		// show paths is locally present
		// return local config value
		var localConfigShowPaths bool = *localConfig.Paths
		return localConfigShowPaths
	}
}
