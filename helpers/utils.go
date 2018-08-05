package helpers

import (
	"regexp"
	"strconv"
	"os"
	"log"
	"path"
	"runtime"
)

const nodeVersionRegexp = `(v\d+\.\d+\.)(\d+)`
const tmpDirName = "tmp"
const versionsDirName = "versions"

func FilterListForMinorVersions(list *[]string) []string {
	newSlice := make([]string, 0)
	r := regexp.MustCompile(nodeVersionRegexp)
	for _, outerElement := range *list {
		versionMatches := r.FindStringSubmatch(outerElement)
		patchVersion, _ := strconv.ParseInt(versionMatches[2], 10, 32)
		hasMatchedExisting := false
		for i, element := range newSlice {
			sliceVersionMatches := r.FindStringSubmatch(element)
			if sliceVersionMatches[1] == versionMatches[1] {
				slicePatchVersion, _ := strconv.ParseInt(sliceVersionMatches[2], 10, 32)
				hasMatchedExisting = true
				if patchVersion > slicePatchVersion {
					newSlice[i] = outerElement
					break
				}
			}
		}
		if !hasMatchedExisting {
			newSlice = append(newSlice, outerElement)
		}
	}
	return newSlice
}

func GetSystemArchitecture() string {
	if runtime.GOARCH == "amd64" {
		return "x64"
	} else if runtime.GOARCH == "386" {
		return "x86"
	}
	log.Fatal("Unsupported system architecture, expected amd64 or 386, got " + runtime.GOARCH)
	return ""
}

func GetAppDirectory() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		log.Fatal("There is no $HOME set for this user, can't create app directory")
		return ""
	}
	appDirName := ".langenv"
	appDirPath := path.Join(homeDir, appDirName)
	if _, err := os.Stat(appDirPath); os.IsNotExist(err) {
		os.Mkdir(appDirPath, os.ModePerm)
	}
	return appDirPath
}

func GetAppTempDirectory() string {
	appDirPath := GetAppDirectory()
	tempAppDirName := tmpDirName
	appTempDirPath := path.Join(appDirPath, tempAppDirName)
	if _, err := os.Stat(appTempDirPath); os.IsNotExist(err) {
		os.Mkdir(appTempDirPath, os.ModePerm)
	}
	return appTempDirPath
}

func GetAppVersionsDirectory(langFolder string) string {
	appDirPath := GetAppDirectory()
	appVersionsDirPath := path.Join(appDirPath, versionsDirName)
	if _, err := os.Stat(appVersionsDirPath); os.IsNotExist(err) {
		os.Mkdir(appVersionsDirPath, os.ModePerm)
	}
	if langFolder == "" {
		return appVersionsDirPath
	}
	appVersionLangDirPath := path.Join(appVersionsDirPath, langFolder)
	if _, err := os.Stat(appVersionLangDirPath); os.IsNotExist(err) {
		os.Mkdir(appVersionLangDirPath, os.ModePerm)
	}
	return appVersionLangDirPath
}
