package cli

/* Get default CLI options for interfaces, etc... */
func GetDefaultScanOptions() ScanOptions {
    opts := ScanOptions{
        OutDirectory: "scan",
        TargetDomain: "scanme.nmap.org",
        TargetIP: []string{},
        Config: "config.json",
        RateLimit: 256,
        Proxy: "",
    }

    return opts
}


func getDefaultCommonScans() []Scan {
    return []Scan{
        {
            Name: "nmap",
            Payload: "nmap  -p0- -v -A -T4 --max-rate {{ ratelimit }} {{ ip|domain }}",
        },
    }
}


func getDefaultOsintScans() []Scan {
    return []Scan{
        {
            Name: "amass",
            Payload: "amass enum -passive -d {{ domain }}",
        },
        {
            Name: "google dorks",
            Payload: "gork -t {{ domain }}",
        },
    }
}


func getDefaultWebScans() []Scan {
    return []Scan{
        {
            Name: "nikto",
            Payload: "nikto -h {{ domain }}",
        },
        {
            Name: "whatweb",
            Payload: "whatweb -v -a 3 {{ domain }}",
        },
        {
            Name: "dalfox",
            Payload: "dalfox url http://{{ domain }}",
        },
        {
            Name: "dirsearch",
            Payload: "dirsearch -r -u http://{{ domain }}",
        },
    }
}


func getDefaultDnsScans() []Scan {
    return []Scan{
        {
            Name: "dig-reverse-lookup",
            Payload: "dig axfr @{{ ip }}",
        },
    }
}


func getDefaultWindowsScans() []Scan {
    return []Scan{
        {
            Name: "enum4linux",
            Payload: "enum4linux {{ ip }}",
        },
        {
            Name: "smbmap",
            Payload: "smbmap -H {{ ip }}",
        },
    }
}


/* Load default scans offered by the CLI */
func GetDefaultScans() ScanSettings {
    settings := ScanSettings{
        Scans: []ScanSuite{
            {
                Category: "commons",
                Scans: getDefaultCommonScans(),
            },
            {
                Category: "osint",
                Scans: getDefaultOsintScans(),
            },
            {
                Category: "web",
                Scans: getDefaultWebScans(),
            },
            {
                Category: "windows",
                Scans: getDefaultWindowsScans(),
            },
            {
                Category: "dns",
                Scans: getDefaultDnsScans(),
            },
        },
    }

    return settings
}
