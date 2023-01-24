package cli

/* Get default CLI options for interfaces, etc... */
func GetDefaultCLIOptions() Options {
    opts := Options{
        OutDirectory: "quickscan",
        TargetURL: "https://scanme.nmap.org",
        Scan: SCAN_ALL,
        ShouldRunDorks: false,
        Wordlist: "/usr/share/seclists/Discovery/Web-Content/quickhits.txt",
        RateLimit: 256,
        Proxy: "",
    }

    return opts
}
