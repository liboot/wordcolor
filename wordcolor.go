package wordcolor

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

const (
    MAGIC_NUMBER = 5
    COLOR_LIMIT  = 242
)

type RGB struct {
	red, green, blue int64
}

func (color RGB) Tohex() string {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return r+g+b
}

func (color RGB) Torgb() string {
	r := strconv.Itoa(int(color.red))
	g := strconv.Itoa(int(color.green))
	b := strconv.Itoa(int(color.blue))
	return fmt.Sprintf("rgb(%s)", strings.Join([]string{r, g, b}, ", "))
}

func (color RGB) BgRGB() string {
	if ((int(color.red) * 299 + int(color.green) * 587 + int(color.blue) * 144) > 200000) {
        return "rgb(0, 0, 0)"
    }else {
		return "rgb(255, 255, 255)"
	}
}

func (color RGB) BgHEX() string {
	if ((int(color.red) * 299 + int(color.green) * 587 + int(color.blue) * 144) > 200000) {
        return "000000"
    }else {
		return "ffffff"
	}
}

func Hex2rgb(color string) RGB {
	r, _ := strconv.ParseInt(color[:2], 16, 10)
	g, _ := strconv.ParseInt(color[2:4], 16, 18)
	b, _ := strconv.ParseInt(color[4:], 16, 10)
	return RGB{r,g,b}
}
  
//calculate the word rgb color
func WordColor(word string, spec int) string {
	color := GetColor(word)
	if spec == 0 {
		return color.Tohex()
	} else {
		return color.Torgb()
	}
}

func WordColorRGB(word string) RGB {
	color := GetColor(word)
	return color
}

func t2x(t int64) string {
    result := strconv.FormatInt(t, 16)
    if len(result) == 1{
        result = "0" + result
    }
    return result
}

func GetColor(word string) RGB {
	res := RGB{0,0,0}
	word = strings.Trim(word, " ")
	rgb := [3]int{0, 0, 0}
	bs := []rune(word)
	length := len(bs)

	for i := 0; i < length; i++ {
		level := i / len(rgb)
		rgb[i%3] += int(getHashNum(string(bs[i:i+1])) / getRatio(float64(level)))
	}
	for key, val := range rgb {
		tmp := int64(0)
		if val > 255 {
			tmp = int64(255)
		} else {
			tmp = int64(val)
		}
		if key == 0 {
			res.red = tmp
		} else if key == 1 {
			res.green = tmp
		} else {
			res.blue = tmp
		}
	}
	return res
}

//Just get the original rgb slice
func GetRGB(word string) [3]string {
	word = strings.Trim(word, " ")
	rgb := [3]int{0, 0, 0}
	bs := []rune(word)
	length := len(bs)

	for i := 0; i < length; i++ {
		level := i / len(rgb)
		rgb[i%3] += int(getHashNum(string(bs[i:i+1])) / getRatio(float64(level)))
	}

	res := [3]string{}
	for key, val := range rgb {
		if val > 255 {
			res[key] = "255"
		} else {
			res[key] = strconv.Itoa(val)
		}
	}
	return res
}

//Get the ratio
func getRatio(level float64) float64 {
	return math.Pow(MAGIC_NUMBER, level)
}

//Get hash number.
func getHashNum(c string) float64 {
	ca := getCharCodeAt(c, 0)
	return float64((ca << MAGIC_NUMBER) % COLOR_LIMIT)
}

// get char code at, just as javascript.
// should know the `strings` pacakge
func getCharCodeAt(str string, index int) int {
	//need know the `rune`
	a := []rune(str)
	return int(a[index])
}
