/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/mengseeker/swagger2markdown/swagger"

	"github.com/spf13/cobra"
)

var (
	customTemplate string
	inputFile      string
	inputFormat    string
	outputFile     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swagger2markdown",
	Short: "transform swagger into markdown",
	Long: `transform swagger into markdown.
only support swagger 2.0`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			inputData []byte
			output    io.Writer
			err       error
		)

		if inputFile == "" {
			inputData, err = io.ReadAll(os.Stdin)
		} else {
			inputData, err = os.ReadFile(inputFile)
		}
		if err != nil {
			return err
		}

		if len(inputData) == 0 {
			return errors.New("empty swagger")
		}

		if inputFormat == "" {
			inputData = bytes.TrimLeft(inputData, " \n\r\t ")
			if inputFile[0] == '{' {
				inputFormat = "json"
			}
		}

		if outputFile == "" {
			output = os.Stdout
		} else {
			output, err = os.Create(outputFile)
			if err != nil {
				return err
			}
			defer output.(*os.File).Close()
		}

		return swagger.Execute(inputData, inputFormat, output, customTemplate)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&customTemplate, "template", "m", "", "custom template file")
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file, can be json or yaml format, default read from stdin")
	rootCmd.Flags().StringVarP(&inputFormat, "inputFormat", "f", "", "input file format, json or yaml, default auto detect")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file, default print to stdout")
}
