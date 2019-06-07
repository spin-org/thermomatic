// Package common implements utilities & functionality commonly consumed by the
// rest of the packages.
package common

import "errors"

// ErrNotImplemented is raised throughout the codebase of the challenge to
// denote implementations to be done by the candidate.
var ErrNotImplemented = errors.New("not implemented")
