package cli

import (
	"errors"

	"github.com/spf13/cobra"
)


func (s *ScanType) Type() string {
    return "ScanType"
}

func (s *ScanType) Set(v string) error {
    switch v {
    case "web", "dns", "windows", "all":
        *s = ScanType(v)
        return nil
    default:
        return errors.New(`must be one of "all", "web", "dns" or "windows"`)
    }
}

func (s *ScanType) Value() string {
    switch *s {
    case SCAN_WEB:
        return "web"
    case SCAN_DNS:
        return "dns"
    case SCAN_WIN:
        return "windows"
    default:
        return "all"
    }
}

func GetParser(opts *Options) *cobra.Command {
    version := "0.0.1"
    var cmd = &cobra.Command{
        Use: "quickscan",
        Version: version,
        DisableSuggestions : true,
        Short: "quickscan - yet another automated infrastructure scanner",
        Long: `CLI tool built to automate scanning on THM/HTB boxes`,
        Run: func(cmd *cobra.Command, args []string) {
        },

    }

    defaults := GetDefaultCLIOptions()

    cmd.PersistentFlags().StringVarP(&opts.OutDirectory, "out-dir", "o", defaults.OutDirectory, "output directory")
    cmd.PersistentFlags().StringVarP(&opts.TargetURL, defaults.TargetURL, "target", "t", "target URL / IP addr")
    cmd.PersistentFlags().StringVarP(&opts.Wordlist, "wordlist", "w", defaults.Wordlist, "Prefered wordlist")

    cmd.PersistentFlags().BoolVarP(&opts.ShouldRunDorks, "dorks", "d", defaults.ShouldRunDorks, "Run gork module")
    cmd.PersistentFlags().IntVarP(&opts.RateLimit, "rate-limit", "r", defaults.RateLimit, "Prefered rate-limit")

    cmd.Flags().VarP(&opts.Scan, "scan-type", "s", "Prefered scan type (one of `all`, `dns`, `web`, `windows`)")

    return cmd
}
