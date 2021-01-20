package terraform

import (
	"context"

	"github.com/hashicorp/terraform-exec/tfinstall"

	"clairvoyance/log"
)

var TerraformBinaryContext = context.Background()

//Identify if lastest (or specified) version binary is installed in a certain directory.
func DetectBinary(installDir string, version string) string {
	if version == "" {
		tfbinary, err := tfinstall.Find(TerraformBinaryContext, tfinstall.LatestVersion(installDir, false))
		if err != nil {
			log.Errorf("[DetectBinary] Could not install %s to %s.", version, installDir)
		}
		return tfbinary
	} else {
		tfbinary, err := tfinstall.Find(TerraformBinaryContext, tfinstall.ExactVersion(version, installDir))
		if err != nil {
			log.Errorf("[DectectBinary] Could not install %s to %s.", version, installDir)
		}
		return tfbinary
	}
	// validate if installed (check /usr/bin/terraform) - need to expand for Windows
	return "/usr/bin/terraform"
}

func InstallTerraform(installDir string, version string) {
}
