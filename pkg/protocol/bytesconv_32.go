//go:build !amd64 && !arm64 && !ppc64 && !ppc64le && !s390x

package protocol

const (
	maxHexIntChars = 7
)
