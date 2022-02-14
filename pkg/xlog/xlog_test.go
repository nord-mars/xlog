package xlog_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/nord-mars/xlog"
)

// Read one line from file
func readOneLine(filename *string) (string, error) {
	rez := ""
	file, err := os.OpenFile(*filename, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return rez, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rez = scanner.Text()
	}

	if err := file.Close(); err != nil {
		return rez, err
	}
	return rez, err
}

// clean files after test
func cleanFiles(filemask *string) {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob(path + "/" + *filemask)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		//fmt.Println(f)
		if err := os.Remove(f); err != nil {
			log.Fatal(err)
		}
	}
}

//
func isFileExist(t *testing.T, filename *string) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// is file created ?
	if _, err := os.Stat(path + "/" + *filename); os.IsNotExist(err) {
		t.Fatalf("File not exist: [%v]", err)
	}
}

// Test: table driven - xlog flags
func Test_FlagLdate(t *testing.T) {
	scenario := []struct {
		input  int
		expect string
	}{
		{input: log.Ldate, expect: `\d{4}\/\d{2}\/\d{2}`},
		{input: log.Ltime, expect: `\d{2}:\d{2}:\d{2}`},
		{input: log.Lmicroseconds, expect: `\d{2}:\d{2}:\d{2}\.\d{6}`},
		{input: log.Lmsgprefix | log.Ldate | xlog.LINE_PID, expect: `\[\d{1,15}\]`},
		{input: log.Ldate | xlog.LINE_PID, expect: `^\[\d{1,15}\]`},
		{input: xlog.LINE_CALL, expect: `xlog_test`},

		{input: xlog.FILE_DATE, expect: `INFO`},
		{input: xlog.FILE_TIME, expect: `INFO`},
		{input: xlog.FILE_PID, expect: `INFO`},
	}

	filename := "xlog_test.log"
	filemask := "xlog*.log"

	for _, s := range scenario {

		// Creatw log and write one line
		Log := xlog.New(filename, 10, s.input)
		Log.Write(0, xlog.INFO, fmt.Sprintf("flag: %014b", s.input))
		defer func() {
			cleanFiles(&filemask)
		}()

		// is file created ?
		isFileExist(t, &filename)

		// check file content
		line, err := readOneLine(&filename)
		if err != nil {
			t.Fatal(err)
		}
		//fmt.Printf("%014b: %v\n", s.input, line)

		//
		re := regexp.MustCompile(s.expect)
		if !re.Match([]byte(line)) {
			t.Errorf("Did not get expected result for input: [%014b]. Expected: [%v], got: [%v] ",
				s.input, s.expect, line)
		}
	}

	time.Sleep(1 * time.Second)
}

//
func Benchmark_Logger(b *testing.B) {
	filename := "xlog_test.log"
	filemask := "xlog*.log"

	Log := xlog.New(filename, 10, log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix|xlog.LINE_CALL|xlog.LINE_PID|xlog.FILE_PID|xlog.FILE_DATE|xlog.FILE_TIME)
	defer func() {
		cleanFiles(&filemask)
	}()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Log.Write(0, xlog.INFO, "Benchmark test - write 1 line to log")
	}
}
