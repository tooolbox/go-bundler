package urlLoader

import (
	"io"
	"os"
	"path/filepath"

	"github.com/lvl5hm/go-bundler/jsLoader"
)

func LoadFile(fileName, bundleDir string) ([]byte, []string, error) {
	ext := filepath.Ext(fileName)
	objectName := jsLoader.CreateVarNameFromPath(fileName)

	dstFileName := bundleDir + "/" + objectName + ext
	err := copyFile(dstFileName, fileName)
	if err != nil {
		return nil, nil, err
	}

	res := "moduleFns." + objectName + "=" +
		"function(){return {exports:'" + objectName + ext + "'}};"

	return []byte(res), nil, nil
}

func copyFile(dst, src string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()
	_, err = io.Copy(to, from)
	return err
}
