package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg4go/convert"
	"github.com/pkg4go/pathx"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "jpgc",
		Short: "jpg compress",
		Long:  "compress jpg image",
		Run: func(c *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("invalid params - jpgc filename quality")
				os.Exit(1)
				return
			}

			fp := pathx.Resolve("", args[0])
			quality, err := convert.Int(args[1])
			notError(err)

			file, err := os.Open(fp)
			notError(err)

			defer file.Close()

			img, err := jpeg.Decode(file)
			notError(err)

			_, filename := filepath.Split(fp)
			ext := filepath.Ext(filename)
			filename = strings.Replace(filename, filepath.Ext(filename), "", -1) + "-" + args[1] + ext

			newImg, err := os.Create("./" + filename)
			notError(err)

			defer newImg.Close()

			err = jpeg.Encode(newImg, img, &jpeg.Options{
				Quality: quality,
			})
			notError(err)
		},
	}

	root.Execute()
}

func notError(e error) {
	if e != nil {
		panic(e)
	}
}
