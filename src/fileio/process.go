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
)

func ProcessDir(dir string) {
	sourceDir := "dir_to_clean/" + dir
	outDir := "results/" + dir
	// copy the directory to the results directory
	CopyDir(sourceDir, outDir)

	allfiles, err := ListFilesRecursive(sourceDir)
	if err != nil {
		fmt.Println(err)
	}
	newFileNames := AppendSuffixToDuplicates(allfiles) // list of initial files with a suffix 001, 002... for duplicates

	for i, file := range allfiles {
		newFilename := newFileNames[i]
		CopyFile(file, outDir+"/"+newFilename)
	}

}
