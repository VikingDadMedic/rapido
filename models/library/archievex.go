/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package library

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Archiever object
type Archivex interface {
	Create(name string) error
	Add(name string, file []byte) error
	AddFile(name string) error
	AddAll(dir string, includeCurrentFolder bool) error
	Close() error
}

// ArchiveWriteFunc is the closure used by an archive's AddAll method to actually put a file into an archive
// Note that for directory entries, this func will be called with a nil 'file' param
type ArchiveWriteFunc func(info os.FileInfo, file io.Reader, entryName string) (err error)

// Implements *zip.Writer
type ZipFile struct {
	Writer *zip.Writer
	Name   string
	file   *os.File
}

// CreateWriter is used to ?
func (z *ZipFile) createWriter(name string) (io.Writer, error) {
	header := &zip.FileHeader{
		Name:   name,
		Flags:  1 << 11, // use utf8 encoding the file Name
		Method: zip.Deflate,
	}
	return z.Writer.CreateHeader(header)
}

// Creates a new .zip file
func (z *ZipFile) Create(name string) error {
	// check extension .zip
	if strings.HasSuffix(name, ".zip") != true {
		if strings.HasSuffix(name, ".tar.gz") == true {
			name = strings.Replace(name, ".tar.gz", ".zip", -1)
		} else {
			name = name + ".zip"
		}
	}
	z.Name = name
	file, err := os.Create(z.Name)
	if err != nil {
		return err
	}
	z.Writer = zip.NewWriter(file)
	z.file = file
	return nil
}

// Adds byte in archive zip
func (z *ZipFile) Add(name string, file []byte) error {
	iow, err := z.createWriter(name)
	if err != nil {
		return err
	}
	_, err = iow.Write(file)
	return err
}

// Adds file from dir in archive
func (z *ZipFile) AddFile(name string) error {
	zippedFile, err1 := z.createWriter(name)
	if err1 != nil {
		return err1
	}

	file, _ := os.Open(filepath.Join(name))
	fileReader := bufio.NewReader(file)
	blockSize := 512 * 1024 // 512kb
	bytes := make([]byte, blockSize)
	for {
		readedBytes, err2 := fileReader.Read(bytes)
		if err2 != nil {
			if err2.Error() == "EOF" {
				break
			}
			if err2.Error() != "EOF" {
				return err2
			}
		}
		if readedBytes >= blockSize {
			zippedFile.Write(bytes)
			continue
		}
		zippedFile.Write(bytes[:readedBytes])
	}
	return nil
}

// Adds a file to zip with a name
func (z *ZipFile) AddFileWithName(name string, filepath string) error {
	zippedFile, err1 := z.createWriter(name)
	if err1 != nil {
		return err1
	}

	file, err2 := os.Open(filepath)
	defer file.Close()
	if err2 != nil {
		return err2
	}
	fileReader := bufio.NewReader(file)
	blockSize := 512 * 1024 // 512kb
	bytes := make([]byte, blockSize)
	for {
		readedBytes, err3 := fileReader.Read(bytes)
		if err3 != nil {
			if err3.Error() == "EOF" {
				break
			}
			if err3.Error() != "EOF" {
				return err3
			}
		}
		if readedBytes >= blockSize {
			zippedFile.Write(bytes)
			continue
		}
		zippedFile.Write(bytes[:readedBytes])
	}
	return nil
}

// Adds all files from dir in archive, recursively.
// Directories receive a zero-size entry in the archive, with a trailing slash in the header name, and no compression
func (z *ZipFile) AddAll(dir string, includeCurrentFolder bool) error {
	dir = path.Clean(dir)
	return addAll(dir, dir, includeCurrentFolder, func(info os.FileInfo, file io.Reader, entryName string) (err error) {

		// Creates a header based off of the fileinfo
		header, err1 := zip.FileInfoHeader(info)
		if err1 != nil {
			return err1
		}

		// If it's a file, sets the compression method to deflate (leave directories uncompressed)
		if !info.IsDir() {
			header.Method = zip.Deflate
		}

		// Sets the header's name to what we want--it may not include the top folder
		header.Name = entryName

		// Adds a trailing slash if the entry is a directory
		if info.IsDir() {
			header.Name += "/"
		}

		// Get a writer in the archive based on our header
		writer, err2 := z.Writer.CreateHeader(header)
		if err2 != nil {
			return err2
		}

		// If we have a file to write (i.e., not a directory) then pipes the file into the archive writer
		if file != nil {
			_, err3 := io.Copy(writer, file)
			if err3 != nil {
				return err3
			}
		}
		return nil
	})
}

// Closes zip file
func (z *ZipFile) Close() error {
	err := z.Writer.Close()
	z.file.Close()
	return err
}

func getSubDir(dir string, rootDir string, includeCurrentFolder bool) (subDir string) {
	subDir = strings.Replace(dir, rootDir, "", 1)
	// Removes leading slashes, since this is intentionally a subdirectory.
	if len(subDir) > 0 && subDir[0] == os.PathSeparator {
		subDir = subDir[1:]
	}
	subDir = path.Join(strings.Split(subDir, string(os.PathSeparator))...)
	if includeCurrentFolder {
		parts := strings.Split(rootDir, string(os.PathSeparator))
		subDir = path.Join(parts[len(parts)-1], subDir)
	}
	return
}

// Recursively goes down through directories and add each file and directory to an archive, based on an ArchiveWriteFunc given to it
func addAll(dir string, rootDir string, includeCurrentFolder bool, writerFunc ArchiveWriteFunc) error {
	// Gets a list of all entries in the directory, as []os.FileInfo
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Loops through all entries
	for _, info := range fileInfos {
		full := filepath.Join(dir, info.Name())
		// If the entry is a file, gets an io.Reader for it
		var file *os.File
		var reader io.Reader
		if !info.IsDir() {
			file, err = os.Open(full)
			if err != nil {
				return err
			}
			reader = file
		}

		// Writes the entry into the archive
		subDir := getSubDir(dir, rootDir, includeCurrentFolder)
		entryName := path.Join(subDir, info.Name())
		if err := writerFunc(info, reader, entryName); err != nil {
			if file != nil {
				file.Close()
			}
			return err
		}

		if file != nil {
			if err := file.Close(); err != nil {
				return err
			}
		}

		// If the entry is a directory, recurses into it
		if info.IsDir() {
			addAll(full, rootDir, includeCurrentFolder, writerFunc)
		}
	}
	return nil
}
