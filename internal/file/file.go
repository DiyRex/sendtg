package file

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"os/user"
)

func PrepareFile(path string) (string, func(), error) {
	path = expandPath(path)

	info, err := os.Stat(path)
	if err != nil {
		return "", nil, err
	}

	if info.IsDir() {
		dirAbs, err := filepath.Abs(path)
		if err != nil {
			return "", nil, err
		}

		// Extract basename of directory (e.g., Reference)
		zipName := filepath.Base(dirAbs) + ".zip"
		zipPath := filepath.Join(os.TempDir(), zipName)

		err = zipFolder(dirAbs, zipPath)
		if err != nil {
			return "", nil, err
		}

		cleanup := func() {
			_ = os.Remove(zipPath)
		}
		return zipPath, cleanup, nil
	}

	return path, func() {}, nil
}




func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return path // fallback if user cannot be resolved
		}
		return filepath.Join(usr.HomeDir, path[1:])
	}
	return path
}

func zipFolder(folder, output string) error {
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	return filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		relPath := strings.TrimPrefix(path, folder+"/")
		w, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	})
}
