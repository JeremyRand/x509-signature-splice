package certinject

import "gopkg.in/hlandau/easyconfig.v1/cflag"
import "github.com/hlandau/xlog"


// This package is used to add and remove certificates to the system trust 
// store.
// Currently only supports the system NSS store.

var log, Log = xlog.New("ncdns.certinject")

var (
	flagGroup          = cflag.NewGroup(nil, "certstore")
	nssSharedFlag      = cflag.Bool(flagGroup, "nss-shared", false, "Synchronize TLS certs to the NSS shared trust store?  This enables HTTPS to work with Chromium/Chrome.  Only use if you've set up null HPKP in Chromium/Chrome as per documentation.  If you haven't set up null HPKP, or if you access ncdns from browsers not based on Chromium, this is unsafe and should not be used.")
	certExpirePeriod   = cflag.Int(flagGroup, "expire", 60 * 30, "Duration (in seconds) after which TLS certs will be removed from the trust store.  Making this smaller than the DNS TTL (default 600) may cause TLS errors.")
)

// Injects the given cert into all configured trust stores.
func InjectCert(derBytes []byte) {

	if nssSharedFlag.Value() {
		injectCertNssShared(derBytes)
	}
}

// Cleans expired certs from all configured trust stores.
func CleanCerts() {

	if nssSharedFlag.Value() {
		cleanCertsNssShared()
	}

}
