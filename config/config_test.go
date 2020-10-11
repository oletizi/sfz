package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	ass := require.New(t)
	path, err := filepath.Abs("../src/test/resources/test.sfz")
	ass.Nil(err)

	log.Printf("path: %v", path)

	config, err := New(path)
	ass.Nil(err)
	ass.NotNil(config)

	// test one-line group
	found := 0
	for _, group := range config.Group() {
		// <group> group=32 lokey=-1 on_locc64=126 on_hicc64=127 off_by=33 volume=-9
		if group.Value("group") == "33" {
			found += 1
			ass.Equal("-1", group.Value("lokey"))
			ass.Equal("1", group.Value("on_hicc64"))
			ass.Equal("0", group.Value("on_locc64"))
			ass.Equal("-9", group.Value("volume"))
		}
	}

	ass.True(found > 0, "Still haven't found what I'm looking for")
}
