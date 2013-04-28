// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.	See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.	If not, see <http://www.gnu.org/licenses/>.

/* Basic ls utility - usage: go run ls.go [args] [dir...] */

package main

import (
	"fmt"
	"os"
	"flag"
	"math"

	"github.com/germ/ansi"
	ce "github.com/germ/checkErr"
)

// Command flags
var (
	humanReadable	bool
	longForm		bool
	colorOutput		bool
	hiddenShown		bool
)

// Flag registration and parsing
func init() {
	flag.BoolVar(&humanReadable,	"h", false, "Show file sizes in more manageable units")
	flag.BoolVar(&longForm,			"l", false, "Display additional file info")
	flag.BoolVar(&colorOutput,		"G", false, "Colorize output")
	flag.BoolVar(&hiddenShown,		"a", false, "Show hidden files")
	flag.Parse()
}
func main() {
	if flag.NArg() == 0 {
		ls(".")
		return
	}

	for i := 0; i < flag.NArg(); i++ {
		ls(flag.Arg(i))
	}
}

func ls(filename string) {
	// Grab a list of files
	file, err := os.Open(filename)
	ce.Exit(err)
	defer file.Close()

	files, err := file.Readdirnames(-1)
	ce.Exit(err)

	// Main control logic
	for _, file := range files {
		if !hiddenShown && (file[0:1] == ".") {
			continue
		}

		// Start escape sequence if needed
		if colorOutput {
			fmt.Print(colorize(file))
		}

		// Long or short form
		if longForm {
			fmt.Println(toLongForm(file))
		} else {
			fmt.Println(file)
		}

		// Close sequence if necessary
		if colorOutput {
			fmt.Print(ansi.Clear)
		}
	}
}

func toLongForm(file string) string {
	var ret string

	// Get file info
	info, err := os.Stat(file)
	if err != nil {
		return file
	}

	// Add permession string
	ret += info.Mode().String() + " "

	// Add file size
	if humanReadable {
		ret += toHuman(info.Size()) + " "
	} else {
		ret += fmt.Sprintf("%10v ", info.Size())
	}

	// Add time and name 
	ret += info.ModTime().Format("02 Jan 15:04 ")
	ret += file

	return ret
}
func toHuman(size int64) string {
	prefixes := []string{"", "K", "M", "G", "T", "Y"}

	// This is just a fancy log
	i, mod := 0, float64(size)
	for ; mod > 1000 && i < len(prefixes); i++ {
		if mod > 1000 {
			mod /= 1000
		}
	}

	// Don't print trailing zeros
	if mod - math.Floor(mod) < 0.1 {
		return fmt.Sprintf("%4v", mod)
	}
	return fmt.Sprintf("%3.1f%v", mod, prefixes[i])
}
func colorize(file string) string {
	info, err := os.Stat(file)
	if err != nil {
		return ""
	}

	var c ansi.AnsiCode
	m := info.Mode()
	if info.IsDir() {
		c = ansi.ColorCyan

	} else if (m & os.ModeSymlink != 0) {
		c = ansi.ColorRed

	} else if (m & os.ModeDevice != 0) {
		c = ansi.ColorGreen

	} else if (m & 0111 != 0) {
		c = ansi.ColorMagenta

	} else {
		return ""
	}

	return ansi.Construct(c)
}
