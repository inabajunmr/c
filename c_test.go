package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	assert(t, "41", 41)
	assert(t, "5+2", 7)
	assert(t, "5-2", 3)
	assert(t, "5-2+3", 6)
	assert(t, "6/2", 3)
	assert(t, "6*2", 12)
	assert(t, "(3+3)*8/(12*2)", 2)
	assert(t, "6*2", 12)
	assert(t, "-6*-2", 12)
	assert(t, "(-3+3+6)*+8/(-12*-2)", 2)
	assert(t, "3==3", 1)
	assert(t, "3>=(1+2)", 1)
	assert(t, "4<=(1+2)", 0)
	assert(t, "(4*2)!=(1+2)", 1)
	assert(t, "(1==1)!=(1!=1)", 1)
}

func assert(t *testing.T, source string, expectedCode int) {

	f, err := os.Create("./build/ctest.s")
	if err != nil {
		t.Fatal()
	}
	f.Write([]byte(compile(source)))
	exec.Command("ls").Run()

	exec.Command("cc", "-o", "./build/ctest", "./build/ctest.s").Run()
	cmd := exec.Command("./build/ctest")
	cmd.Run()
	code := cmd.ProcessState.ExitCode()
	if code != expectedCode {
		t.Errorf("source is %v, code is %v", source, code)
	}
}
