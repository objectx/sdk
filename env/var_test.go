package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {

	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)
	defer os.RemoveAll(sdkmanDir)

	err = os.Mkdir(sdkmanDir+"/var", 0755)
	check(err)

	expected := "1.0.0"
	err = ioutil.WriteFile(sdkmanDir+varFile, []byte(expected), os.ModePerm)
	check(err)

	actual := GetVersion(sdkmanDir)
	assert.Equal(t, expected, actual)
}
