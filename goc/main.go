package main

import "github.com/atotto/clipboard"
import "github.com/pkg4go/pathx"
import "github.com/spf13/cobra"
import "io/ioutil"
import "fmt"

func main() {
	root := &cobra.Command{
		Use:   "goc",
		Short: "go clipboard",
		Long:  "clipboard CLI by golang",
		Run: func(c *cobra.Command, args []string) {
			if len(args) == 0 {
				text, err := clipboard.ReadAll()
				notError(err)

				fmt.Print(text)
				fmt.Println()
				return
			}

			if len(args) == 1 {
				filename := pathx.Resolve("", args[0])
				data, err := ioutil.ReadFile(filename)
				notError(err)

				err = clipboard.WriteAll(string(data[:]))
				notError(err)
				return
			}
		},
	}

	root.Execute()
}

func notError(e error) {
	if e != nil {
		panic(e)
	}
}
