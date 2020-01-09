/*
Copyright © 2020 Paul Nelson Baker

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "git-export-zip",
	Short: "Export the current git HEAD as a zip file",
	Long:  `Above the `,
	Run: func(cmd *cobra.Command, args []string) {
		gitRootDirCmd := exec.Command("git", "rev-parse", "--show-toplevel")
		rootDirBytes, err := gitRootDirCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Could not determine current git hash: %v", err)
		}
		gitBaseFilename := filepath.Base(strings.TrimSuffix(string(rootDirBytes), "\n"))
		projectParentDirectory := path.Dir(string(rootDirBytes))
		gitHeadHashCmd := exec.Command("git", "log", "--pretty=format:'%h'", "-n", "1")
		hashBytes, err := gitHeadHashCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Could not determine HEAD commit hash: %v", err)
		}
		hashValue := string(hashBytes[1 : len(hashBytes)-1])
		formattedTimestamp := time.Now().UTC().Format("01-02-2006")
		projectName := strings.ToLower(strings.ReplaceAll(gitBaseFilename, " ", "_"))
		outputArchiveFilename := fmt.Sprintf("%s/%s-%s-%s.zip",
			projectParentDirectory, projectName, formattedTimestamp, hashValue)

		gitArchiveCmd := exec.Command("git", "archive", "-o", outputArchiveFilename, "HEAD")
		gitArchiveCmd.Stdout = os.Stdout
		gitArchiveCmd.Stdin = os.Stdin
		gitArchiveCmd.Stderr = os.Stderr
		if err := gitArchiveCmd.Run(); err != nil {
			log.Fatalf("Could not export archive: %v", err)
		} else {
			log.Printf("Exported to: %s", outputArchiveFilename)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-export-zip.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".git-export-zip" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".git-export-zip")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
