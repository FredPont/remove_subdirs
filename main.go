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
 (c) Frederic Pont 2024
*/

package main

import (
	"fmt"

	"remove_subdirs/src/fileio"
)

func main() {

	fileio.Title()

	dirNames := fileio.ListDir("dir_to_clean/") // read tables in table dir

	fmt.Println(dirNames)

	for _, dirName := range dirNames {
		fileio.ProcessDir(dirName)
	}

	//processing.CheckAllTables(tableNames) // verify the path on all the tables in the table dir

	//processing.ProcTable(tableNames) // copy the files

}
