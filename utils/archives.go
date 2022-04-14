package utils

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func Untar(dest string, reader io.Reader) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()

		switch {
		case err == io.EOF:
			return nil

		case err != nil:
			return err

		case header == nil:
			continue
		}

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			file, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}

			file.Close()
		}
	}
}
