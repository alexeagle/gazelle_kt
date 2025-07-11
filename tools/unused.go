//go:build tools

package tools

// Workaround for rules_jvm which references go.starlark.net/starlark

import (
	_ "go.starlark.net/starlark"
)