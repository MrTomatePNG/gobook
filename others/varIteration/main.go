package variteration

import "os"

var rmdirs []func()

func main() {
	for _, d := range temDirs() {
		dir := d
		os.Mkdir(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func temDirs() (dirs []string) {
	panic("unimplemented")
}
