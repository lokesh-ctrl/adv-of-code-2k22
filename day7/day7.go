package adv22

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	parent *Dir
	files  []*File
	dirs   []*Dir
	root   *Root
}

type File struct {
	parent *Dir
	name   string
	size   int
}

type Root struct {
	rootDir *Dir
	allDirs []*Dir
}

func (d *Dir) Size() int {
	totalSize := 0
	for _, dir := range d.dirs {
		totalSize += dir.Size()
	}
	for _, file := range d.files {
		totalSize += file.size
	}
	return totalSize
}

func (d *Dir) AddRoot(root Root) {
	d.root = &root
}

func (d *Dir) CheckIfDirExists(name string) (*Dir, bool) {
	for _, dir := range d.dirs {
		if dir.name == name {
			return dir, true
		}
	}
	return nil, false
}

func (d *Dir) CreateSubDir(name string) *Dir {
	dir, exists := d.CheckIfDirExists(name)
	if !exists {
		newDir := &Dir{name: name, root: d.root, parent: d}
		d.dirs = append(d.dirs, newDir)
		d.root.allDirs = append(d.root.allDirs, newDir)
		return newDir
	}
	return dir
}

func (d *Dir) CreateFile(name string, size int) {
	d.files = append(d.files, &File{name: name, size: size, parent: d})
}

func (d *Dir) Print() {
	fmt.Println("dir " + d.name)
	for _, dir := range d.dirs {
		dir.Print()
	}
	for _, file := range d.files {
		file.Print()
	}
}

func (f *File) Print() {
	fmt.Printf("\t file name %s , size - %d \t parent - %s \n", f.name, f.size, f.parent.name)
}

func PrintAllDirsSize(root Root) {
	for _, dir := range root.allDirs {
		fmt.Printf("Dir name - %s \t size- %d \n", dir.name, dir.Size())
	}
}

func GetDirectoriesSizeLessThan(root Root, limitSize int) int {
	totalSize := 0
	for _, dir := range root.allDirs {
		if dir.Size() <= limitSize {
			totalSize += dir.Size()
		}
	}

	return totalSize
}

func Day7P1(inputFile string) {
	root := initializeValues(inputFile)
	PrintAllDirsSize(root)
	fmt.Printf("\n Answer: %d \n", GetDirectoriesSizeLessThan(root, 100000))
}

func initializeValues(inputFile string) Root {
	inputs, _ := parseAndReadLines(inputFile)
	rootDir := &Dir{name: "/"}
	root := &Root{rootDir: rootDir}
	rootDir.AddRoot(*root)

	currentDirectory := rootDir

	for _, input := range inputs {
		splitInputs := strings.Fields(input)
		if splitInputs[0] == "$" {
			if splitInputs[1] == "cd" {
				if splitInputs[2] == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					dirName := splitInputs[2]
					currentDirectory, _ = currentDirectory.CheckIfDirExists(dirName)
				}
			}
		} else {
			if splitInputs[0] == "dir" {
				currentDirectory.CreateSubDir(splitInputs[1])
			} else {
				fileSize, _ := strconv.Atoi(splitInputs[0])
				currentDirectory.CreateFile(splitInputs[1], fileSize)
			}
		}
	}
	return *currentDirectory.root
}

func Day7P2(inputFile string) {
	var directorySizes []int
	root := initializeValues(inputFile)
	for _, dir := range root.allDirs {
		directorySizes = append(directorySizes, dir.Size())
	}

	fmt.Printf("\n Root size: %d \n", root.rootDir.Size())
	answer := findCorrectSizeToDelete(directorySizes, 30000000-(70000000-root.rootDir.Size()))

	fmt.Printf("\n Answer: %d \n", answer)

}

func findCorrectSizeToDelete(sizes []int, limit int) int {
	answer := 0
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] < sizes[j] })
	for _, size := range sizes {
		fmt.Printf("size %d \n", size)
		if size > limit {
			answer = size
		}
	}
	return answer
}

func parseAndReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
