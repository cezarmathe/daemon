package daemon

const (
	pkgPrefix = "github.com/cezarmathe/daemon: "
)

type ErrNoSocketFound struct {
	Address string
}

func (ns *ErrNoSocketFound) Error() string {
	return pkgPrefix + "no socket at " + ns.Address
}
