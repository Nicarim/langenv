package nodejs

import (
	"github.com/spf13/cobra"
	"fmt"
	"langenv/helpers"
	"langenv/providers"
	"os"
	"path"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Downloads given version, and makes it available for further use for ex. global / local",
	Run: func(cmd *cobra.Command, args []string) {
		versions := providers.GetVersionsAvailable(true)
		if len(args) < 1 {
			fmt.Println("Please provide the version to be downloaded")
			return
		}
		versionByUser := args[0]
		tempDir := helpers.GetAppTempDirectory()
		var foundVersion = &providers.NodeVersion{Version: ""}
		for _, el := range versions {
			if el.Version == versionByUser {
				foundVersion = &el
				break
			}
		}
		if foundVersion.Version == ""{
			fmt.Println("Given version does not exists.")
			return
		}
		outputFile := helpers.DownloadFile(foundVersion.DownloadUrl, tempDir)
		defer os.Remove(outputFile)
		extractDir := path.Join(tempDir, versionByUser)
		os.Mkdir(extractDir, os.ModePerm)

		helpers.ExtractTarFile(outputFile, extractDir)
		versionTargetDir := path.Join(helpers.GetAppVersionsDirectory("nodejs"), versionByUser)
		os.Rename(extractDir, versionTargetDir)
		fmt.Println("New version has been installed into " + versionTargetDir)
	},
}
