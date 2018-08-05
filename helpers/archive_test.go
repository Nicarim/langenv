package helpers

import (
	"testing"
	"io/ioutil"
	"path"
	"os"
)

func TestFlattenFolderOneDeep(t *testing.T) {
	dir, err := ioutil.TempDir("", "tests_flatten")
	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}
	innerDir := path.Join(dir, "inner_dir")
	errInner := os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_dir")
	}
	innerDirFile, errTempFile := ioutil.TempFile(innerDir, "temp_file_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	FlattenFolder(dir, "")

	files, errReadDir := ioutil.ReadDir(dir)
	if errReadDir != nil {
		t.Fatal("Unable to read directory after flattening")
	}

	for _, f := range files {
		if f.IsDir() {
			t.Fatal("Expected files after flattening, got folder " + f.Name())
		}
	}
}

func TestFlattenFolderMultipleOneDeep(t *testing.T) {
	dir, err := ioutil.TempDir("", "tests_flatten")
	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}
	innerDir := path.Join(dir, "inner_dir")
	errInner := os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_dir")
	}
	innerDirFile, errTempFile := ioutil.TempFile(innerDir, "temp_file_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	innerDirFile, errTempFile = ioutil.TempFile(innerDir, "temp_file2_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	FlattenFolder(dir, "")

	files, errReadDir := ioutil.ReadDir(dir)
	if errReadDir != nil {
		t.Fatal("Unable to read directory after flattening")
	}

	for _, f := range files {
		if f.IsDir() {
			t.Fatal("Expected files after flattening, got folder " + f.Name())
		}
	}
}

func TestFlattenFolderMultipleSecondDeep(t *testing.T) {
	dir, err := ioutil.TempDir("", "tests_flatten")
	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}
	innerDir := path.Join(dir, "inner_dir")
	errInner := os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_dir")
	}

	innerDir = path.Join(innerDir, "inner_m _dir")
	errInner = os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_most_dir")
	}

	innerDirFile, errTempFile := ioutil.TempFile(innerDir, "temp_file_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	innerDirFile, errTempFile = ioutil.TempFile(innerDir, "temp_file2_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	FlattenFolder(dir, "")

	files, errReadDir := ioutil.ReadDir(dir)
	if errReadDir != nil {
		t.Fatal("Unable to read directory after flattening")
	}

	for _, f := range files {
		if f.IsDir() {
			t.Fatal("Expected files after flattening, got folder " + f.Name())
		}
	}
}

func TestFlattenFolderMultipleThirdDeep(t *testing.T) {
	dir, err := ioutil.TempDir("", "tests_flatten")
	if err != nil {
		t.Fatal("Unable to create temporary directory")
	}
	innerDir := path.Join(dir, "inner_dir")
	errInner := os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_dir")
	}

	innerDir = path.Join(innerDir, "inner_more_dir")
	errInner = os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_more_dir")
	}


	innerDir = path.Join(innerDir, "inner_most_dir")
	errInner = os.Mkdir(innerDir, os.ModePerm)
	if errInner != nil {
		t.Fatal("Unable to create inner_most_dir")
	}


	innerDirFile, errTempFile := ioutil.TempFile(innerDir, "temp_file_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	innerDirFile, errTempFile = ioutil.TempFile(innerDir, "temp_file2_test")
	if errTempFile != nil {
		t.Fatal("Unable to create temporary file inside inner_dir")
	}
	innerDirFile.Close()

	FlattenFolder(dir, "")

	files, errReadDir := ioutil.ReadDir(dir)
	if errReadDir != nil {
		t.Fatal("Unable to read directory after flattening")
	}

	for _, f := range files {
		if f.IsDir() {
			t.Fatal("Expected files after flattening, got folder " + f.Name())
		}
	}
}