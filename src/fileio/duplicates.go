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
	"path/filepath"
	"strings"
)

// AppendSuffixToDuplicates append number to duplicate file names : file_001.xxx, file002.xxx etc...
func AppendSuffixToDuplicates(filePaths []string) []string {
	countMap := make(map[string]int)
	result := make([]string, len(filePaths)) // list of new file names with eventually a suffix for duplicates
	var duplicates []string

	for i, filePath := range filePaths {
		filename := filepath.Base(filePath) // extract the filename from path
		if count, exists := countMap[filename]; exists {
			ext := filepath.Ext(filename)
			name := strings.TrimSuffix(filename, ext)
			count++
			suffix := fmt.Sprintf("_%03d", count-1) // -1 to start with 001 not 002
			result[i] = name + suffix + ext         //create a new file name
			countMap[filename] = count
			duplicates = append(duplicates, filePath)
		} else {
			result[i] = filename
			countMap[filename] = 1
		}
	}
	println()
	fmt.Println(len(duplicates), "duplicates found !")
	fmt.Println(duplicates)
	return result
}
