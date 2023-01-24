package cli


type ScanType string

/// define all available scan types
const (
    /// Run every available scans
    SCAN_ALL ScanType = "all"

    /// only run web scans
    SCAN_WEB ScanType = "web"

    /// only run dns scans
    SCAN_DNS ScanType = "dns"

    /// only run windows scans
    SCAN_WIN ScanType =  "windows"
)


/// CLI Options to be sent to the parser
type Options struct {

    /// Output directory
    OutDirectory string

    /// Target URL or IP to be scanned
    TargetURL string

    /// Scan type to be ran
    Scan ScanType

    /// Specify wether or not to run google dorks (default: false)
    ShouldRunDorks bool

    /// Prefered Wordlist
    Wordlist string

    /// Proxy to use
    Proxy string

    /// Prefered rate-limit
    RateLimit int
}
