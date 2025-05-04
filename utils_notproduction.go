//go:build !production

package main

func IsProduction() bool {
	return false
}
