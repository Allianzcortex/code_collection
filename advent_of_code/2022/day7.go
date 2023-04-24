package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func day7Part1(input []string) int {
	var curDir []string

	dirLevel := make(map[string][]string)
	dirList := []string{}
	fileSizeMatch := make(map[string]int64)

	for _, line := range input {
		switch {
		case strings.HasPrefix(line, "$ cd"):
			target := line[5:]
			// if target == "/" {
			// 	curDir = append(curDir, "/")
			// } else
			if target == ".." {
				// TODO : maybe make curDir a []string is easier
				curDir = curDir[:len(curDir)-1]
			} else {
				curDir = append(curDir, target)
			}
		case strings.HasPrefix(line, "$ ls"):
			continue
		default:
			dir := curDir[0] + strings.Join(curDir[1:], "/")
			fileName := strings.Fields(line)[1]
			if strings.HasPrefix(line, "dir") {
				temp := fileName
				if dir != "/" {
					temp = "/" + temp
				}
				dirLevel[dir] = append(dirLevel[dir], dir+temp)
				dirList = append(dirList, dir+temp)
			} else {
				if dir != "/" {
					fileName = "/" + fileName
				}
				temp := dir + fileName
				dirLevel[dir] = append(dirLevel[dir], temp)
				fileSize := strings.Fields(line)[0]
				size, _ := strconv.ParseInt(fileSize, 10, 64)
				fileSizeMatch[temp] = size
			}
		}
	}

	accumulate(dirLevel, &fileSizeMatch, "/")
	fmt.Println(dirLevel)
	fmt.Println(fileSizeMatch)
	fmt.Println(dirList)

	val := 0
	for _, dir := range dirList {
		dirSize := fileSizeMatch[dir]
		if dirSize <= 100000 {
			val += int(dirSize)
		}
	}
	return val
}

// maybe I should define a category with `dirSizeMatch`, but let's use pointer to make life easier...
// fileName can be an exact filename or a directory name that starts with "/"
func accumulate(dirLevel map[string][]string, fileSizeMatch *map[string]int64, fileOrDirName string) int64 {
	if val, exist := (*fileSizeMatch)[fileOrDirName]; exist {
		return val
	}
	// if we can reach this line of code, it means it must be a directory
	sum := 0
	for _, subDirOrFile := range dirLevel[fileOrDirName] {
		sum += int(accumulate(dirLevel, fileSizeMatch, subDirOrFile))
	}
	(*fileSizeMatch)[fileOrDirName] = int64(sum)
	return int64(sum)

}
