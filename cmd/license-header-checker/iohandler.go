/* MIT License

Copyright (c) 2022 Lluis Sanchez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ioHandler implements the ioHandle interface defined in the process package
type ioHandler struct{}

func (s *ioHandler) ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

func (s *ioHandler) WalkDir(path string, walkDirFn fs.WalkDirFunc) error {
	return filepath.WalkDir(path, walkDirFn)
}

func (s *ioHandler) ReplaceFileContent(name string, content string) error {
	err := os.Remove(name)
	if err != nil {
		return fmt.Errorf("failed deleting the file: %w", err)
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}

	if err = writer.Flush(); err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}

	return nil
}
