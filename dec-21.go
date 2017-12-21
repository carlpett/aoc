package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type image struct {
	size   int
	pixels [][]string
}

func newImage(size int) *image {
	img := &image{}
	img.size = size
	img.pixels = make([][]string, img.size)
	for r := 0; r < img.size; r++ {
		img.pixels[r] = make([]string, img.size)
	}
	return img
}
func imageFromString(s string) *image {
	size := strings.Count(s, "/") + 1
	i := image{pixels: make([][]string, size), size: size}
	for idx, row := range strings.Split(s, "/") {
		i.pixels[idx] = strings.Split(row, "")
	}
	return &i
}

func (i *image) ruleString() string {
	s := make([]string, i.size)
	for idx, r := range i.pixels {
		s[idx] = strings.Join(r, "")
	}
	return strings.Join(s, "/")
}

func (i *image) String() string {
	s := make([]string, i.size)
	for idx, r := range i.pixels {
		s[idx] = strings.Join(r, "")
	}
	return strings.Join(s, "\n")
}

func (i *image) copy() image {
	c := image{}
	c.size = i.size
	c.pixels = make([][]string, c.size)
	for r := 0; r < i.size; r++ {
		c.pixels[r] = make([]string, c.size)
		copy(c.pixels[r], i.pixels[r])
	}
	return c
}

func main() {
	tS := time.Now()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	rules := make(map[string]*image, len(lines))
	for _, l := range lines {
		ruleStr := strings.Split(l, " => ")
		rules[ruleStr[0]] = imageFromString(ruleStr[1])
	}
	fmt.Printf("Setup in %v\n", time.Since(tS))

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(imageFromString(".#./..#/###"), rules, 5), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(imageFromString(".#./..#/###"), rules, 18), time.Since(tB))
}

func solve(img *image, rules map[string]*image, iter int) int {
	for i := 0; i < iter; i++ {
		subSize := getSubSize(img)

		next := newImage(img.size + img.size/subSize)
		subimages := getSubImages(img, subSize)
		for x, subRow := range subimages {
			for y, sub := range subRow {
				enhanced := enhance(sub, rules)
				setSubImage(next, enhanced, x*enhanced.size, y*enhanced.size)
			}
		}
		img = next
	}
	return strings.Count(img.ruleString(), "#")
}

func getSubSize(img *image) int {
	if img.size%2 == 0 {
		return 2
	}
	return 3
}

func enhance(img *image, rules map[string]*image) *image {
	e, found := rules[img.ruleString()]
	if found {
		return e
	}
	e, found = rules[flip(img, horizontal).ruleString()]
	if found {
		return e
	}
	e, found = rules[flip(img, vertical).ruleString()]
	if found {
		return e
	}
	r1 := rotate(img)
	e, found = rules[r1.ruleString()]
	if found {
		return e
	}
	r2 := flip(r1, vertical)
	e, found = rules[r2.ruleString()]
	if found {
		return e
	}
	r3 := rotate(r2)
	e, found = rules[r3.ruleString()]
	if found {
		return e
	}
	r4 := flip(r2, horizontal)
	e, found = rules[r4.ruleString()]
	if found {
		return e
	}

	panic("Did not find enhancement!")
}

func getSubImages(img *image, subSize int) [][]*image {
	s := img.size / subSize
	subs := make([][]*image, s)
	for r := 0; r < s; r++ {
		subs[r] = make([]*image, s)
		for c := 0; c < s; c++ {
			subs[r][c] = newImage(subSize)
			for p := 0; p < subSize; p++ {
				copy(subs[r][c].pixels[p], img.pixels[r*subSize+p][c*subSize:(c+1)*subSize])
			}
		}
	}

	return subs
}

func setSubImage(img *image, sub *image, x0, y0 int) {
	for r := x0; r < x0+sub.size; r++ {
		copy(img.pixels[r][y0:y0+sub.size], sub.pixels[r-x0])
	}
}

type dim int

const (
	vertical dim = iota
	horizontal
)

func flip(img *image, dim dim) *image {
	flipped := img.copy()
	switch dim {
	case vertical:
		flipped.pixels[0], flipped.pixels[flipped.size-1] = flipped.pixels[flipped.size-1], flipped.pixels[0]
	case horizontal:
		for r := 0; r < flipped.size; r++ {
			flipped.pixels[r][0], flipped.pixels[r][flipped.size-1] = flipped.pixels[r][flipped.size-1], flipped.pixels[r][0]
		}
	}
	return &flipped
}
func rotate(img *image) *image {
	rotated := newImage(img.size)
	for r := 0; r < img.size; r++ {
		for c := 0; c < img.size; c++ {
			rotated.pixels[r][c] = img.pixels[img.size-c-1][img.size-r-1]
		}
	}
	return rotated
}
