package main

import (
	"github.com/codegangsta/cli"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-resize-image"
	app.Version = "1.0.0"
	app.Usage = "Convert designated image size"
	app.Author = "nikuyoshi"
	app.Email = "nikuyoshi@gmail.com"
	app.Action = func(c *cli.Context) {
		imageFileDir, _ := path.Split(c.Args()[0])
		resizeMyImage(c.Args()[0], imageFileDir, c.Args()[1], c.Args()[2])
	}
	app.Run(os.Args)
}

func resizeMyImage(imageFileFullPath string, compressionImageDir string, x string, y string) {

	openedFile, err := os.Open(imageFileFullPath)
	if err != nil {
		panic(err)
	}
	defer openedFile.Close()

	image, err := jpeg.Decode(openedFile)
	if err != nil {
		panic(err)
	}

	uintX, err := strconv.ParseUint(x, 10, 0)
	if err != nil {
		panic(err)
	}
	uintY, err := strconv.ParseUint(y, 10, 0)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	m := resize.Resize(uint(uintX), uint(uintY), image, resize.Lanczos3)

	resizedImage, err := os.Create(compressionImageDir + "/after.jpg")
	if err != nil {
		panic(err)
	}
	defer resizedImage.Close()

	jpeg.Encode(resizedImage, m, nil)
	log.Println("Wrote out after.jpg")

}
