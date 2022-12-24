package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	//log.SetOutput(io.Discard)

	path, err := filepath.Abs("../../test/data/test.sfz")
	assert.Nil(t, err)

	log.Println("path: ", path)

	config, err := New(path)
	assert.Nil(t, err)
	assert.NotNil(t, config)

	groups := config.Group()
	log.Printf("Groups count: %d", len(groups))
	assert.Equal(t, 60, len(groups))

	samples := config.Sample()
	log.Printf("Samples count: %d", len(samples))

	regions := config.Region()
	log.Printf("Regions: %d", len(regions))

	controls := config.Control()
	log.Printf("Controls: %d", len(controls))

	curves := config.Curve()
	log.Printf("Curves: %d", len(curves))

	effects := config.Effect()
	log.Printf("Effects: %d", len(effects))

	global := config.Global()
	log.Printf("Global: %v", global)

	masters := config.Master()
	log.Printf("Masters: %d", len(masters))

	midis := config.Midi()
	log.Printf("Midis: %d", len(midis))

	// test one-line group
	found := 0
	for _, group := range config.Group() {
		// <group> group=32 lokey=-1 on_locc64=126 on_hicc64=127 off_by=33 volume=-9
		if group.Value("group") == "33" {
			found += 1
			assert.Equal(t, "-1", group.Value("lokey"))
			assert.Equal(t, "1", group.Value("on_hicc64"))
			assert.Equal(t, "0", group.Value("on_locc64"))
			assert.Equal(t, "-9", group.Value("volume"))
		}
	}

	assert.True(t, found > 0, "Still haven't found what I'm looking for")
}
