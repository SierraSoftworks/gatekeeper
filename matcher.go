package gatekeeper

import "regexp"

var permissionReplacementRegex = regexp.MustCompile("\\:\\w+")

// Matcher provides tools to check whether a set of permissions,
// in a specific context, fulfil a permissions request.
type Matcher struct {
	Permissions []string
	Context     map[string]string
}

// NewMatcher creates a new permissions matcher which can be configured
// using the WithPermissions and WithContext methods.
func NewMatcher() *Matcher {
	return &Matcher{
		Permissions: []string{},
		Context:     map[string]string{},
	}
}

// WithPermissions creates a new matcher which includes the provided
// permissions.
func (m *Matcher) WithPermissions(permissions []string) *Matcher {
	return &Matcher{
		Permissions: permissions,
		Context:     m.Context,
	}
}

// WithContext creates a new matcher which includes the provided
// context data for named replacements.
func (m *Matcher) WithContext(context map[string]string) *Matcher {
	return &Matcher{
		Permissions: m.Permissions,
		Context:     context,
	}
}

// Can determines whether the context and permissions in the matcher
// are capable of fulfilling a permission.
func (m *Matcher) Can(permission string) bool {
	for _, p := range m.Permissions {
		if permission == p {
			return true
		}
	}

	filledPerm := permissionReplacementRegex.ReplaceAllStringFunc(permission, func(match string) string {
		name := match[1:]
		replacement, hasReplacement := m.Context[name]
		if !hasReplacement {
			return match
		}

		return replacement
	})

	if filledPerm == permission {
		return false
	}

	for _, p := range m.Permissions {
		if filledPerm == p {
			return true
		}
	}

	return false
}

// CanAll allows you to determine whether a context fulfils a number of
// permissions.
func (m *Matcher) CanAll(permissions ...string) bool {
	for _, p := range permissions {
		if !m.Can(p) {
			return false
		}
	}

	return true
}
