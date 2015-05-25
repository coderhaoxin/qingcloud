package main

import "path/filepath"
import "io/ioutil"
import "strconv"
import "reflect"
import "fmt"
import "os"

func toInt(i interface{}) int {
	if i == nil {
		return 0
	}

	s := i.(string)
	num, e := strconv.Atoi(s)

	if e != nil {
		return 0
	}

	return num
}

func toBool(i interface{}) bool {
	if i == nil {
		return false
	}

	return i.(bool)
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}

	return i.(string)
}

func readfile(path string) []byte {
	if !filepath.IsAbs(path) {
		cwd, _ := os.Getwd()
		path = filepath.Join(cwd, path)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return bytes
}

// utils for test
func equal(args ...interface{}) {
	if len(args)%2 != 0 {
		panic(fmt.Errorf("not matched args"))
	}

	step := len(args) / 2
	for i := 0; i < step; i++ {
		if !reflect.DeepEqual(args[i], args[i+step]) {
			panic(fmt.Errorf("expect %v to equal %v", args[i], args[i+step]))
		}
	}

}
