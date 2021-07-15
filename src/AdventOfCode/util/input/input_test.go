package input

import (
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

/*************************************/
/* Local function                    */
/*************************************/

// Compare two variables
func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Get working directory
func wd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return wd
}

// Make a text file which have only single line text data
// It returns file path
func makeSingleLineFile(in string) string {
	filePath := wd() + "/testdata_singlelinetext"

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	_, err = f.WriteString(in)
	if err != nil {
		log.Fatal(err)
		f.Close()
		return ""
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return f.Name()
}

/*************************************/
/* Test                              */
/*************************************/

func TestReadLines(t *testing.T) {
	in := wd() + "/testdata1"
	want := []string{
		"test 1",
		"test 2",
		"test 3",
	}

	got := ReadLines(in)
	if !equal(got, want) {
		t.Errorf("\ngot  = %v;\nwant = %v", got, want)
	}
}

func TestReadLinesAsync(t *testing.T) {
	in := wd() + "/testdata1"
	want := []string{
		"test 1",
		"test 2",
		"test 3",
	}

	var got []string
	ch := make(chan string)
	err := make(chan error)
	go ReadLinesAsync(in, ch, err)
	end := false
	for {
		if end {
			break
		}

		select {
		case e := <-err:
			if e != nil {
				t.Errorf("\ne = %v", e)
				end = true
				break
			}
		case v, ok := <-ch:
			if !ok {
				end = true
				break
			}
			got = append(got, v)
		}
	}

	if !equal(got, want) {
		t.Errorf("\ngot  = %v;\nwant = %v", got, want)
	}
}

func TestReadLine(t *testing.T) {
	str := strings.Repeat("x", 65536)
	in := makeSingleLineFile(str)
	want := str

	got := ReadLine(in)
	if got != want {
		t.Errorf("\ngot  = %v;\nwant = %v", got, want)
	}
}
