package options

var Options struct {
	Port        int
	TLSCert     string
	TLSKey      string
	Provision   int
	Deprovision int
	Bind        int
	Unbind      int
}
