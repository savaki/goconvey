// This script updates version 1.* GoConvey test suites to accurately utilize
// the GoConvey 2.* DSL. Please, for the love of Mike, before running this
// script on your code do the following:
//
// 1. Make sure you are writing your tests according to conventions specified
//    in the GoConvey documentation and examples. If you don't do this you'll
//    have to tweak this script to match your naming/formatting style
// 2. Run `go fmt` on all your code.
// 3. Commit all of your code to source control
//
// After running the script please compile and run your tests and ensure that
// you get the same results (number of passes/failures/skips, etc...)

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	flag.StringVar(&root, "root", ".", rootDescription)
	flag.Parse()
	log.Println("Start at:", root)

	_, this, _, _ := runtime.Caller(0)
	thisFolder := filepath.Dir(this)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		currentFolder, err := filepath.Abs(filepath.Dir(path))
		if currentFolder == thisFolder {
			log.Println("Skipping:", currentFolder)
			return filepath.SkipDir
		}

		if strings.HasSuffix(info.Name(), "_test.go") {
			log.Println("Rewriting file:", info.Name())
			lines, err := readLines(path)
			if err != nil {
				return err
			}
			updated := rewrite(lines)
			err = writeLines(updated, path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil && err != filepath.SkipDir {
		panic(err)
	}
}

func rewrite(inLines []string) []string {
	outLines := make([]string, len(inLines))

	for number, line := range inLines {
		if strings.HasPrefix(dedent(line), "Convey(\"") && strings.HasSuffix(line, ", t, func() {") {
			// top-level `Convey`
			outLines[number] = strings.Replace(line, ", t, func() {", ", t, func(c *Context, so Assert) {", 1)
		} else if strings.HasPrefix(dedent(line), "Convey(\"") {
			// 'Convey' call
			outLines[number] = strings.Replace(line, ", func() {", ", c, func() {", 1)
		} else if strings.HasPrefix(dedent(line), "SkipConvey(\"") {
			// 'SkipConvey' call
			outLines[number] = strings.Replace(line, ", func() {", ", c, func() {", 1)
		} else if strings.HasPrefix(dedent(line), "So(") && strings.Contains(line, "Should") {
			// 'So'
			outLines[number] = strings.Replace(line, "So(", "so(", 1)
		} else if strings.HasPrefix(dedent(line), "SkipSo(") && strings.Contains(line, "Should") {
			// 'SkipSo'
			outLines[number] = strings.Replace(line, "SkipSo(", "c.Skipso(", 1) // NOT HAPPY WITH THIS...
		} else if strings.HasPrefix(dedent(line), "Reset(") {
			// 'Reset'
			outLines[number] = strings.Replace(line, "Reset(", "Reset(c, ", 1)
		} else {
			// no-op, just copy the line as-is
			outLines[number] = line
		}
	}

	return outLines
}

func dedent(line string) string {
	return strings.Replace(line, "\t", "", -1)
}

func readLines(path string) ([]string, error) {
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

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

var root string

const rootDescription = "The top-level directory to begin scanning for *_test.go files which will be updated to the latest GoConvey DSL style."
