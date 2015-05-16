package freetype

import (
	"log"
	"os"
	"strings"
	"testing"
)

var (
	fonts = []string{
		"/usr/share/fonts/truetype/ttf-droid/DroidSans.ttf",
		"/usr/share/fonts/truetype/droid/DroidSans.ttf",
		"/usr/share/fonts/truetype/DroidSans.ttf",
	}
)

func TestMain(m *testing.M) {
	SetupTest()
	os.Exit(m.Run())
	TeardownTest()
}

var (
	lib      *Library
	face     *Face
	fileName string
	err      error
)

func SetupTest() {
	lib, err = InitFreeType()
	if err != nil {
		log.Fatal(err)
	}

	for i := range fonts {
		_, err := os.Open(fonts[i])
		if err == nil {
			fileName = fonts[i]
			break
		}
	}

	if len(fileName) == 0 {
		log.Fatalf("no font file found from the list: %s", strings.Join(fonts, ", "))
	}

	face, err = NewFace(lib, fileName, 0)
	if err != nil {
		log.Fatal(err)
	}
	if face == nil {
		log.Fatal("face should not be nil")
	}
}

func TeardownTest() {
	if face != nil {
		err = face.Done()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = lib.Done()
	if err != nil {
		log.Fatal(err)
	}
}
