package inputformats

import (
	"bufio"
	"io"
	"strings"
)

type VhostFormat struct{}

func NewVhostFormat() *VhostFormat {
	return &VhostFormat{}
}

var _ Format = &VhostFormat{}

func (v *VhostFormat) Name() string {
	return "vhost"
}

func (v *VhostFormat) Parse(input io.Reader, callback func(url string) bool) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !callback(line) {
			break
		}
	}
	return scanner.Err()
}
