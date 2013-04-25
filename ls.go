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

/* Basic ls utility - usage: go run ls.go [args] [dir...] */

package main

import (
  "fmt"
  "os"
  "flag"
  "strings"
)

/* command line flags */
var hidden_files_toggle = flag.Bool("h", false, "show hidden files and directories")

func ls(filename string) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Printf("ls: cannot access: %s: No such file or directory\n", filename)
    os.Exit(1)
  }
  files, err := file.Readdirnames(-1)
  if err != nil {
    fmt.Printf("ls: cannot access: %s: No such file or directory\n", filename)
    os.Exit(1)
  }
  for j:=0; j < len(files); j++ {
    var hidden_files = strings.HasPrefix(files[j], ".")
    if *hidden_files_toggle == true {
      fmt.Printf("%s\n", files[j])
    } else {
      if hidden_files == false {
        fmt.Printf("%s\n", files[j])
      }
    }
  }
  defer file.Close()
}

func main() {
  flag.Parse()
  if flag.NArg() == 0 {
    ls(".")
  } else {
    for i := 0; i < flag.NArg(); i++ {
      ls(flag.Arg(i))
    }
  }
}

