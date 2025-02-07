/*
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

 Written by Frederic PONT.
 (c) Frederic Pont 2025
*/

package fileio

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ListDir lists all directories in a directory
func ListDir(dir string) []string {
	var dirsList []string
	// Open the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	// List the directories
	for _, e := range entries {
		//fmt.Println(e.Name())
		if e.IsDir() {
			dirsList = append(dirsList, e.Name())
		}
	}

	return dirsList
}

func ListDirWithSymlinks(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirNames []string
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			dirNames = append(dirNames, entry.Name()) // Add the symlink to the list
		} else if entry.IsDir() {
			dirNames = append(dirNames, entry.Name()) // Add the directory to the list
		}
	}
	return dirNames, nil
}

// ListFiles lists all files in a directory
func ListFiles(dir string) []string {
	var filesList []string
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name())
		filesList = append(filesList, f.Name())
	}
	return filesList
}

func ListFilesRecursive(dir string) ([]string, error) {
	var files []string
	if isSymlink(dir) {
		dir = openSymLink(dir)
	}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func openSymLink(symlinkPath string) string {
	// Read the target of the symlink
	target, err := os.Readlink(symlinkPath)
	if err != nil {
		fmt.Println("Error reading symlink:", err)
		return ""
	}

	fmt.Println("Symlink points to:", target)
	return target
}

func isSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		println("Error reading symlink:", err)
		return false // Return false
	}
	return info.Mode()&os.ModeSymlink != 0 // Check if the mode indicates a symlink
}
