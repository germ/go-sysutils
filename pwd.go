// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

/* Basic pwd utility - usage: go run pwd.go */

package main

import (
  "fmt"
  "os"
)

func main() {
  wd, err := os.Getwd()
  if err != nil {
    fmt.Printf("pwd: cannot find directory\n")
    os.Exit(1)
  }
  fmt.Printf("current working directory: %s\n", wd)
}

