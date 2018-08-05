package providers

import (
	"github.com/buger/jsonparser"
	"log"
	"langenv/helpers"
	"fmt"
)

const nodeJsRootUrl = "https://nodejs.org/dist"

const IndexUrl = nodeJsRootUrl + "/index.json"

type NodeVersion struct {
	Version     string
	DownloadUrl string
}

func createNodeSrcDownloadUrl(version string) string {
	return fmt.Sprintf("%s/%s/node-%s-linux-%s.tar.gz", nodeJsRootUrl, version, version, helpers.GetSystemArchitecture())
}

func GetVersionsAvailable(includePatchVersions bool) []NodeVersion {
	stringVersionArray := make([]string, 0)
	bodyData, err := helpers.GetBodyFromUrl(IndexUrl)

	if err != nil {
		log.Fatal(err.Error())
		return []NodeVersion{}
	}

	jsonparser.ArrayEach(bodyData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		versionExtracted, errGetString := jsonparser.GetString(value, "version")
		if errGetString != nil {
			log.Fatal("Failed to get version key from index response")
		}

		stringVersionArray = append(stringVersionArray, versionExtracted)
	})
	if !includePatchVersions {
		stringVersionArray = helpers.FilterListForMinorVersions(&stringVersionArray)
	}
	nodeVersionArray := make([]NodeVersion, len(stringVersionArray))
	for i, el := range stringVersionArray {
		nodeVersionArray[i] = NodeVersion{
			Version:     el,
			DownloadUrl: createNodeSrcDownloadUrl(el),
		}
	}
	return nodeVersionArray
}
