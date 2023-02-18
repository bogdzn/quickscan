package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func loadDirectories(target string) string {
    fullpath := "./scans/" + target
    os.MkdirAll(fullpath, 0755)

    /* TODO(bogdan): build a directory for each scan category */
    return fullpath
}


func GetParser(opts *ScanOptions) *cobra.Command {
    version := "0.0.1"

    var createConfig = &cobra.Command{
        Use: "setup",
        Version: version,
        DisableSuggestions: true,
        Short: "create default config file",
        Long: "CLI tool to automate your recon conveniently",
        Run: func (cmd *cobra.Command, args []string) {

            settings := GetDefaultScans()
            asJson, err := json.Marshal(settings); if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
            }

            err = ioutil.WriteFile(opts.Config, asJson, 0644); if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
            }

        },
    }

    var scan = &cobra.Command{
        Use: "quickscan",
        Version: version,
        DisableSuggestions : true,
        Short: "quickscan - yet another automated infrastructure scanner",
        Long: `CLI tool built to automate scanning on THM/HTB boxes`,
        Run: func(cmd *cobra.Command, args []string) {
        },

    }

    defaults := GetDefaultScanOptions()

    createConfig.PersistentFlags().StringVarP(&opts.OutDirectory, "out-dir", "o", defaults.OutDirectory, "output directory")
    createConfig.PersistentFlags().StringVarP(&opts.Config, "config", "c", defaults.Config, ".json filepath")

    return scan
}
