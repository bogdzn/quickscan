package cli



/// CLI Options to be sent to the parser
type ScanOptions struct {

    /// Output directory
    OutDirectory string

    /// Target Domain to be scanned
    TargetDomain string

    /// Target IPs to be scanned
    TargetIP []string

    /// JSON config filepath
    Config string

    /// Proxy to use
    Proxy string

    /// Prefered rate-limit
    RateLimit int
}

/// Struct holding the various scan categories, this is basically what is
/// unmarshalled from the config JSON file
type ScanSettings struct {

    /// Scan suite defined by the user
    Scans []ScanSuite `json:scans`

}


/// Maps a category (web, dns, windows, whatever) to an array of scans
type ScanSuite struct {

    /// Category name
    Category string `json:category`

    /// scans to be run
    Scans []Scan `json:scans`
}


/// A scan-type
type Scan struct {
    /// Name of the tool used for the scan
    Name string `json:name`

    /// Function / scan-type to be ran
    Payload string  `json:payload`
}
