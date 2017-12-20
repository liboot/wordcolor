package wordcolor

import (
	"testing"
)

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh%2F11.3.md
func TestConvert(t *testing.T) {
	txt := "Hello world"
	_fixture := "rgb(146, 124, 79)"
	rgb := WordColor(txt, 1)
	if _fixture != rgb {
		t.Error(rgb)
	}
}

func TestGetRGB(t *testing.T) {
	txt := "Hello world"
	_fixture := [3]string{"146", "124", "79"}
	rgb := GetRGB(txt)
	if _fixture != rgb {
		t.Error(rgb)
	}
}

func TestUnicodeConvert(t *testing.T) {
	txt := "李牧"
	_fixture := "rgb(240, 160, 0)"
	if _fixture != WordColor(txt, 0) {
		t.Error("Convert unicode char error")
	}
}

func TestWithJSSpec(t *testing.T) {
	_fixture := "rgb(188, 174, 18)"
	rgb := WordColor("words", 1)
	if rgb != _fixture {
		t.Errorf("Except %v, but got %v", _fixture, rgb)
	}
}

func TestLongWord(t *testing.T) {
	_fixture := "rgb(104, 184, 167)"
	rgb := WordColor("longlonglongwords", 1)
	if rgb != _fixture {
		t.Logf("Except: %v, but get %v", _fixture, rgb)
		t.Error("Convert long word error")
	}
}

func TestGetCharCodeAt(t *testing.T) {
	code := getCharCodeAt("longlonglongwords", 0)
	if code != 108 {
		t.Errorf("Should be %v, but got %v", 108, code)
	}
}
