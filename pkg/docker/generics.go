package docker

// Base interface for docker commands
type BaseCommand interface {
	Handle(...string)
	ExtractOptionsAndParams(opts ...string) ([]string, []string)
}
