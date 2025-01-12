package ouidb

import (
	"fmt"
	"math/rand"
	"testing"
)

var db *OuiDB

func lookup(t *testing.T, mac, expectedOrg string, expectedError error) {
	if db == nil {
		t.Fatal("database not initialized")
	}
	v, err := db.Lookup(mac)
	if err != expectedError {
		t.Fatalf("parse: %s: %s", mac, err.Error())
	}
	if v != expectedOrg {
		t.Fatalf("lookup: input %s, expect %q, got %q", mac, expectedOrg, v)
	}
	//t.Logf("%s => %s\n", mac, v)
}

func string48(b [6]byte) string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
		b[0], b[1], b[2], b[3], b[4], b[5])
}

func string24(b [3]byte) string {
	return fmt.Sprintf("%02x:%02x:%02x:aa:bb:cc", b[0], b[1], b[2])
}

func invalid(t *testing.T, mac string) {
	if db == nil {
		t.Fatal("database not initialized")
	}
	v, err := db.Lookup(mac)
	if err == nil {
		t.Fatalf("didn't fail on invalid %s, got %s", mac, v)
	}
}

func TestInitialization(t *testing.T) {
	var err error
	db, err = New()
	if err != nil {
		t.Fatal(err)
	}

	if db == nil {
		t.Fatal("can't load database file oui.txt")
	}
}

func TestLookup24(t *testing.T) {
	lookup(t, "60:03:08:a0:ec:a6", "Apple", nil)
}

func TestLookup36(t *testing.T) {
	lookup(t, "00:1B:C5:00:E1:55", "VigorEle", nil)
}

func TestLookup40(t *testing.T) {
	lookup(t, "20-52-45-43-56-aa", "Receive", nil)
}

func TestLookupUnknown(t *testing.T) {
	lookup(t, "ff:ff:00:a0:ec:a6", "", NotFoundErr)
}

func TestFormatSingleZero(t *testing.T) {
	lookup(t, "0:25:9c:42:0:62", "Cisco-Li", nil)
}

func TestFormatUppercase(t *testing.T) {
	lookup(t, "0:25:9C:42:C2:62", "Cisco-Li", nil)
}

func TestInvalidMAC1(t *testing.T) {
	invalid(t, "00:25-:9C:42:C2:62")
}

func TestLookupAll48(t *testing.T) {
	for _, b := range db.blocks48 {
		lookup(t, string48(b.oui), b.Organization(), nil)
	}
}

func TestLookupAll24(t *testing.T) {
	for _, b := range db.blocks24 {
		lookup(t, string24(b.oui), b.Organization(), nil)
	}
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := db.blocks24[rand.Intn(len(db.blocks24))]
		db.Lookup(string24(b.oui))
	}
}
