package filewalk

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func GetFilesInDir(dirsToBeSkipped []string, programmingLanguageSuffix string) []string {
	var files []string

	err := filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Could not read file or dir: %s", err.Error())
			return err
		}

		if d.IsDir() {
			if dirCanBeSkipped(d.Name(), dirsToBeSkipped) {
				return filepath.SkipDir
			}
		} else {
			// it is a file
			if fileSuffixMatchesProgrammingLanguage(path, programmingLanguageSuffix) {
				files = append(files, path)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Could not walk dir: %s", err.Error())
	}

	return files
}

func fileSuffixMatchesProgrammingLanguage(fullPath string, programmingLanguageSuffix string) bool {
	var suffix = path.Ext(fullPath)
	return suffix == programmingLanguageSuffix
}

func dirCanBeSkipped(dirToCheck string, dirsToBeSkipped []string) bool {
	for _, dir := range dirsToBeSkipped {
		if dir == dirToCheck {
			return true
		}
	}
	return false
}
