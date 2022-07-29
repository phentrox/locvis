package skip

import "locvis/entities"

func CreateDirsToBeSkipped(localConfigVar entities.LocalConfig) []string {
	var configDirsToBeSkiped []string = localConfigVar.IgnoreDirs
	var dirsToBeSkipped = []string{".git", ".idea", ".mvn"}
	dirsToBeSkipped = append(dirsToBeSkipped, configDirsToBeSkiped...)
	return dirsToBeSkipped
}
