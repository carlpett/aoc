package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"flag"
	"log"
	"runtime"
	"runtime/pprof"
)

type image struct {
	size   int
	pixels [][]string
	rs     string
}

func newImage(size int) *image {
	img := &image{
		size:   size,
		pixels: make([][]string, size),
	}
	for r := 0; r < img.size; r++ {
		img.pixels[r] = make([]string, img.size)
	}
	return img
}
func imageFromString(s string) *image {
	size := strings.Count(s, "/") + 1
	i := image{pixels: make([][]string, size), size: size, rs: s}
	for idx, row := range strings.Split(s, "/") {
		i.pixels[idx] = strings.Split(row, "")
	}
	return &i
}

func (i *image) ruleString() string {
	if i.rs != "" {
		return i.rs
	}
	b := new(bytes.Buffer) //bytes.NewBuffer(make([]byte, 0, i.size*i.size))
	for idx, r := range i.pixels {
		for _, s := range r {
			b.WriteString(s)
		}
		if idx < i.size-1 {
			b.WriteString("/")
		}
	}
	i.rs = b.String()
	return i.rs
}

func (i *image) String() string {
	s := make([]string, i.size)
	for idx, r := range i.pixels {
		s[idx] = strings.Join(r, "")
	}
	return strings.Join(s, "\n")
}

func (i *image) copy() *image {
	c := image{}
	c.size = i.size
	c.pixels = make([][]string, c.size)
	c.rs = i.rs
	for r := 0; r < i.size; r++ {
		c.pixels[r] = make([]string, c.size)
		copy(c.pixels[r], i.pixels[r])
	}
	return &c
}

var profile = flag.Bool("profile", false, "")

func main() {
	flag.Parse()
	if *profile {
		fmt.Println("Profiling")
		f, err := os.Create("dec-21.cpu")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

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

	if *profile {
		pprof.StopCPUProfile()

		f, err := os.Create("dec-21.mem")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
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

var enhancedMemory = make(map[string]*image)

func enhance(img *image, rules map[string]*image) *image {
	if mem, ok := enhancedMemory[img.ruleString()]; ok {
		return mem //.copy()
	}
	needle := img.copy()
	funcs := []func(*image){
		func(img *image) {}, // Noop
		func(img *image) { flip(img, horizontal) },
		func(img *image) { flip(img, vertical) },
		func(img *image) { rotate(img) },
		func(img *image) { flip(img, vertical) },
		func(img *image) { rotate(img) },
		func(img *image) { flip(img, horizontal) },
	}
	for _, fn := range funcs {
		fn(needle)
		if e, found := rules[needle.ruleString()]; found {
			enhancedMemory[img.ruleString()] = e
			return e
		}
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
	diagonal
)

func flip(img *image, dim dim) {
	switch dim {
	case vertical:
		img.pixels[0], img.pixels[img.size-1] = img.pixels[img.size-1], img.pixels[0]
	case horizontal:
		for r := 0; r < img.size; r++ {
			img.pixels[r][0], img.pixels[r][img.size-1] = img.pixels[r][img.size-1], img.pixels[r][0]
		}
	case diagonal:
		for r := 0; r < (img.size+1)/2; r++ {
			for c := 0; c < (img.size+1)/2; c++ {
				img.pixels[r][c], img.pixels[img.size-c-1][img.size-r-1] = img.pixels[img.size-c-1][img.size-r-1], img.pixels[r][c]
			}
		}
	}
	img.rs = ""
}
func rotate(img *image) {
	flip(img, diagonal)
	flip(img, horizontal)
}
