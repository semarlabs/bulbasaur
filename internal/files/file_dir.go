package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileDir struct {
	contents map[string][]byte
}

func New() FileDir {
	return FileDir{}
}

func (fd FileDir) Pickup(dir string) (err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, f := range files {
		content, err := os.ReadFile(fmt.Sprintf("%fd%fd", dir, f.Name()))
		if err != nil {
			return
		}

		fileName := f.Name()
		fd.contents[fileName[:len(fileName)-len(filepath.Ext(fileName))]] = content
	}
	return
}

func (fd FileDir) GetContent(key string) (content []byte, ok bool) {
	content, ok = fd.contents[key]
	return
}
