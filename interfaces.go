package openinsider

type Client interface {
	LatestClusterBuys() ([]Disclosure, error)
	Screen(*Filter) ([]Disclosure, error)
}
