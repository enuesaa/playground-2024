package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)
  
func TestHello(t *testing.T) {
	snaps.MatchSnapshot(t, hello())
}

func TestOther(t *testing.T) {
	snaps.MatchSnapshot(t, 5, 10, 20, 25)
}
