package faketls

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_readServerHello(t *testing.T) {
	p := filepath.Join("testdata", "server_hello")
	entries, err := os.ReadDir(p)
	if err != nil {
		t.Fatal(err)
	}

	clientRandom := map[string][32]uint8{
		"alexbers.hex": {
			0xa1, 0x32, 0xe3, 0x91, 0x60, 0x83, 0xb3, 0x14, 0xc1, 0xb9, 0x74, 0xd0, 0x57, 0x85, 0xe8, 0xee,
			0x70, 0x45, 0x6e, 0x5f, 0x86, 0x6d, 0x96, 0x57, 0xd5, 0x0a, 0x5c, 0x08, 0xea, 0x38, 0x31, 0x8e,
		},
		"mtgv2.hex": {
			0xca, 0x36, 0xbb, 0xa8, 0x33, 0x80, 0x9a, 0x33, 0xaa, 0x62, 0x7e, 0xbb, 0x5a, 0x32, 0xa1, 0x01,
			0x02, 0xd1, 0xa6, 0x1e, 0x1e, 0x6c, 0x58, 0xa4, 0x61, 0xd9, 0x34, 0x57, 0x4d, 0x2e, 0x2e, 0xa3,
		},
	}
	secret := []uint8{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	for _, entry := range entries {
		if entry.IsDir() {
			t.Error("Unexpected directory in testdata")
			continue
		}
		fileName := entry.Name()

		t.Run(fileName, func(t *testing.T) {
			a := require.New(t)

			data, err := os.ReadFile(filepath.Join(p, fileName))
			a.NoError(err)

			decode, err := hex.DecodeString(strings.TrimSpace(string(data)))
			a.NoError(err)

			r := bytes.NewReader(decode)
			a.NoError(readServerHello(r, clientRandom[entry.Name()], secret))

			a.Zero(r.Len(), "should read all data")
		})
	}
}
