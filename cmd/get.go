/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Pickle Rick",
	Long:  `This get command will call GitHub repo in order to return the desired Pickle Rick`,
	Run: func(cmd *cobra.Command, args []string) {
		var pickleName = "arakaki-pickle.png"

		if len(args) >= 1 && args[0] != "" {
			pickleName = args[0]
		}

		URL := "https://raw.githubusercontent.com/cameracode/ricksofpickle/Develop/" + pickleName + ".png"

		fmt.Println("Try to get: `" + pickleName + "` Pickle Rick...")

		// Get the data
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		// defer closes connection
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(pickleName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			// defer closes connection
			defer out.Close()

			// Write response body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in " + out.Name() + "!")
		} else {
			fmt.Println("Error:  " + pickleName + " not exists! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
