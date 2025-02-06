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
	"time"

	"remove_subdirs/src/fileio"

	"github.com/schollz/progressbar/v3"
)

func main() {

	fileio.Title()
	t0 := time.Now()

	dirNames := fileio.ListDir("dir_to_clean/") // read tables in table dir

	fmt.Println(dirNames)

	bar := progressbar.Default(int64(len(dirNames)))
	for _, dirName := range dirNames {
		bar.Add(1)
		fileio.ProcessDir(dirName)
	}

	fmt.Println("\ndone !")
	fmt.Println("Elapsed time : ", time.Since(t0))
}
