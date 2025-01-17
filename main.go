package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/steiler/yang-parser/xpath"
	"github.com/steiler/yang-parser/xpath/grammars/expr"
	"github.com/steiler/yang-parser/xpath/xutils"
)

// call this
// find /home/mava/projects/schema-server/lab/common/yang/ -type f -name "*.yang" | xargs perl -lne '/must \"(.*)\"/ && print $1'

type stats struct {
	total  int
	failed int
}

func (s *stats) String() string {
	pass_ratio := (float32(s.total) - float32(s.failed)) * 100 / float32(s.total)
	return fmt.Sprintf("Pass-Ratio: %.2f%%, Total: %d, Pass %d, Failed: %d", pass_ratio, s.total, s.total-s.failed, s.failed)
}

func main() {

	s := &stats{}

	processor := func(x string) {
		s.total += 1
		xpm, err := expr.NewExprMachine(x, nil)
		if err != nil {
			log.Error(err)
			s.failed += 1
		}
		_ = xpm

		fmt.Println(xpm.PrintMachine())
	}

	ProcessFile(processor, "/home/mava/projects/yang-parser/INPUTDATA")

	fmt.Println(s)
}

func foo(xpm *xpath.Machine, ctxNode xutils.XpathNode) {
	xpath.NewCtxFromMach(xpm, ctxNode)
}

func ProcessFile(f func(string), file string) {
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		f(fileScanner.Text())
	}

	readFile.Close()
}

func ProcessStdIn(f func(string)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		f(scanner.Text())
	}
}
