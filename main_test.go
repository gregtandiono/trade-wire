package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("########### TEARING DOWN & SEEDING ###########")
	var t *testing.T
	seedDataBase(t)
	os.Exit(m.Run())
}
