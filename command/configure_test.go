package command

import (
	"bytes"
	"github.com/redwebcreation/nest/common"
	"github.com/redwebcreation/nest/global"
	"github.com/redwebcreation/nest/util"
	"os"
	"testing"
)

type Set struct {
	Strategy   string
	Provider   string
	Repository string
	Dir        string
	Error      error
}

var dataset = []Set{
	{"remote", "github", "felixdorn/config-test", "", nil},
	{"remote", "gitlab", "felixdorn/config-test", "", nil},
	{"remote", "bitbucket", "felixdorn/config-test", "", nil},
	{"invalidStrategy", "github", "felixdorn/config-test", "", common.ErrInvalidStrategy},
	{"remote", "invalidProvider", "felixdorn/config-test", "", common.ErrInvalidProvider},
	{"remote", "github", "invalidRepository", "", common.ErrInvalidRepository},
}

func TestConfigureCommandUsingFlags(t *testing.T) {
	cmd := NewConfigureCommand()

	for _, data := range dataset {
		cmd.Flags().Set("strategy", data.Strategy)
		cmd.Flags().Set("provider", data.Provider)
		cmd.Flags().Set("repository", data.Repository)

		global.ConfigLocatorConfigFile = util.TmpFile().Name()

		err := cmd.Execute()
		if err != data.Error {
			if data.Error == nil {
				t.Errorf("Expected no error, got %s", err)
			} else {
				t.Errorf("Expected %s, got %s", data.Error, err)
			}
		}

		_ = os.Remove(global.ConfigLocatorConfigFile)
	}

}

func TestConfigureCommandInteractively(t *testing.T) {
	cmd := NewConfigureCommand()

	originalStdin := util.Stdin

	for _, data := range dataset {
		if data.Error != nil {
			continue
		}

		util.Stdin = bytes.NewBufferString(data.Strategy + "\n" + data.Provider + "\n" + data.Repository + "\n")

		global.ConfigLocatorConfigFile = util.TmpFile().Name()

		err := cmd.Execute()
		if err != data.Error {
			if data.Error == nil {
				t.Errorf("Expected no error, got %s", err)
			} else {
				t.Errorf("Expected %s, got %s", data.Error, err)
			}
		}

		_ = os.Remove(global.ConfigLocatorConfigFile)
	}

	util.Stdin = originalStdin
}
