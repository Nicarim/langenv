package helpers

import (
	"os/exec"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path"
)

func ExtractTarFile(path string, output string) *exec.Cmd {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("tar -xvf %s -C %s", path, output))
	fmt.Println("Extracting file " + path)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	FlattenFolder(output, "")
	return cmd
}

// Removes unnecessary folders inside a path - mostly happens when extracting files, inside archive it often happens
// that there is another folder in it - and we want to get rid of it.
func FlattenFolder(searchPath string, rootPath string) {
	allFilesInDir, err := ioutil.ReadDir(searchPath)
	if err != nil {
		log.Fatal("Couldn't read the directory")
		return
	}
	if len(allFilesInDir) < 1 {
		log.Fatal("Nothing was found in directory")
	} else if len(allFilesInDir) == 1 {
		if allFilesInDir[0].IsDir() {
			if rootPath == "" {
				FlattenFolder(path.Join(searchPath, allFilesInDir[0].Name()), searchPath)
			} else {
				FlattenFolder(path.Join(searchPath, allFilesInDir[0].Name()), rootPath)
				errRemove := os.Remove(searchPath)
				if errRemove != nil {
					log.Fatal("Couldn't remove old folder after flattening.")
				}
			}
		} else {
			if rootPath != "" {
				oldFileName := path.Join(searchPath, allFilesInDir[0].Name())
				newFileName := path.Join(rootPath, allFilesInDir[0].Name())
				errRename := os.Rename(oldFileName, newFileName)
				if errRename != nil {
					log.Fatal("Couldn't rename the file inside the directory")
				}
				errRemove := os.Remove(searchPath)
				if errRemove != nil {
					log.Fatal("Couldn't remove old folder after flattening.")
				}
			}
		}
	} else {
		if rootPath != "" {
			for _, f := range allFilesInDir {
				errRename := os.Rename(path.Join(searchPath, f.Name()), path.Join(rootPath, f.Name()))
				if errRename != nil {
					log.Fatal("Couldn't rename the file inside the directory")
				}
			}
			errRemove := os.Remove(searchPath)
			if errRemove != nil {
				log.Fatal("Couldn't remove old folder after flattening.")
			}
		}
	}
}
