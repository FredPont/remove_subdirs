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

package main

import (
	"fmt"
	"reflect"
	"remove_subdirs/src/fileio"
	"testing"
)

func TestAppendSuffixToDuplicates(t *testing.T) {
	tests := []struct {
		files []string
		want  []string
	}{
		{[]string{"file1.txt", "file2.txt", "file3.txt"}, []string{"file1.txt", "file2.txt", "file3.txt"}},
		{[]string{"file1.txt", "file1.txt", "file3.txt"}, []string{"file1.txt", "file1_001.txt", "file3.txt"}},
		{[]string{"file1", "file2.txt", "file3.txt"}, []string{"file1", "file2.txt", "file3.txt"}},
		{[]string{"file1.txt", "file1.txt", "file3.txt", "file2", "file2"}, []string{"file1.txt", "file1_001.txt", "file3.txt", "file2", "file2_001"}},
		{[]string{"dir1/dir2/dir3/file1.txt", "dir1/dir2/dir3/dir4/file1.txt", "file1.txt"}, []string{"file1.txt", "file1_001.txt", "file1_002.txt"}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("IsRemote=%d", i), func(t *testing.T) {
			got := fileio.AppendSuffixToDuplicates(tc.files)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}

		})
	}

}
