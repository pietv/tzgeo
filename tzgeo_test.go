package tzgeo

import (
	"math"
	"testing"
)

func TestLookupKnownZones(t *testing.T) {
	cases := []struct {
		name         string
		wantLat, wLo float64
	}{
		// From zone1970.tab: +5230+01322 Europe/Berlin (short form).
		{"Europe/Berlin", 52 + 30.0/60, 13 + 22.0/60},
		// From zone1970.tab: +744144-0944945 America/Resolute (long form).
		{"America/Resolute", 74 + 41.0/60 + 44.0/3600, -(94 + 49.0/60 + 45.0/3600)},
		// From zone1970.tab: +2232+08822 Asia/Kolkata (short form).
		{"Asia/Kolkata", 22 + 32.0/60, 88 + 22.0/60},
	}
	for _, c := range cases {
		lat, lon, ok := Lookup(c.name)
		if !ok {
			t.Errorf("%s: not found", c.name)
			continue
		}
		if math.Abs(lat-c.wantLat) > 1e-9 {
			t.Errorf("%s: lat = %v, want %v", c.name, lat, c.wantLat)
		}
		if math.Abs(lon-c.wLo) > 1e-9 {
			t.Errorf("%s: lon = %v, want %v", c.name, lon, c.wLo)
		}
	}
}

func TestLookupSouthernWestern(t *testing.T) {
	// zone1970.tab: AR -3436-05827 America/Argentina/Buenos_Aires.
	lat, lon, ok := Lookup("America/Argentina/Buenos_Aires")
	if !ok {
		t.Fatal("America/Argentina/Buenos_Aires not found")
	}
	if lat >= 0 {
		t.Errorf("lat should be negative (southern), got %v", lat)
	}
	if lon >= 0 {
		t.Errorf("lon should be negative (western), got %v", lon)
	}
}

func TestLookupMissing(t *testing.T) {
	if _, _, ok := Lookup("Mars/Olympus"); ok {
		t.Error("bogus zone should not resolve")
	}
	// US/Pacific is a legacy alias kept in tzdata's "backward" file,
	// not in zone1970.tab — confirm we return ok=false rather than
	// guessing.
	if _, _, ok := Lookup("US/Pacific"); ok {
		t.Error("legacy alias should not be in zone1970.tab")
	}
}
