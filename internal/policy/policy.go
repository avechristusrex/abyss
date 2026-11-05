// Package policy maps request descriptors to configured limits.
package policy

import "github.com/khatran/abyss/internal/limiter"

type Descriptor map[string]string
type Policy struct {
	Domain string
	Match  Descriptor
	Limit  limiter.Policy
}
type Resolver interface {
	Resolve(domain string, descriptors []Descriptor) ([]Policy, error)
}
