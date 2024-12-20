package internal

import "testing"

func TestExample(t *testing.T) {
  if 1+1 != 2 {
    t.Fatalf("math broken")
  }
}
