package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFetchXMLRegion(t *testing.T) {
	const nRegions = 8
	const nStations = 111

	region, err := fetchXMLRegion()
	if err != nil {
		t.Error("failed to fetch the full region list")
	}
	if len(region.Region) != nRegions {
		t.Errorf("failed to fetch all the regions (%v instead of %v)", len(region.Region), nRegions)
	}
	stationCount := 0
	for _, stations := range region.Region {
		for _, station := range stations.Stations {
			t.Logf("%v (%v) %v\n", stations.RegionID, station.AreaID, station.Name)
			stationCount += 1
		}
	}
	if stationCount < nStations {
		t.Errorf("failed to fetch all the stations (%v instead of %v)", stationCount, nStations)
	}
}

func TestNewAsset(t *testing.T) {
	const nAreas = 47
	const nRegions = 7
	const nStations = 111

	asset, err := NewAsset()
	if err != nil {
		t.Errorf("failed to parse the asset %s", err)
	}

	// Area
	if len(asset.Regions) != nRegions {
		t.Errorf("wrong number of regions (%v instead of %v)", len(asset.Regions), nRegions)
	}
	areaCount := 0
	for region := range asset.Regions {
		for range asset.Regions[region] {
			areaCount += 1
		}
	}
	if areaCount != nAreas {
		t.Errorf("wrong number of areas (%v instead of %v)", areaCount, nAreas)
	}

	// Coordinate
	if len(asset.Coordinates) != nAreas {
		t.Errorf("wrong number of coordinates (%v instead of %v)", len(asset.Coordinates), nAreas)
	}

	// Station
	if len(asset.Stations) != nStations {
		t.Errorf("wrong number of stations (%v instead of %v)", len(asset.Stations), nStations)
	}

	// Versions
	if len(asset.Versions.Apps) != 26 {
		t.Errorf("wrong number of apps (%v instead of %v)", len(asset.Versions.Apps), 26)
	}
	if len(asset.Versions.Models) != 241 {
		t.Errorf("wrong number of models (%v instead of %v)", len(asset.Versions.Models), 241)
	}
	if len(asset.Versions.SDKs) != 10 {
		t.Errorf("wrong number of sdks (%v instead of %v)", len(asset.Versions.SDKs), 10)
	}
}

func TestGenerateGPSForAreaID(t *testing.T) {
	asset, _ := NewAsset()
	var gpstests = []struct {
		in  string
		out bool
	}{
		{
			"JP13",
			true,
		},
		{
			"NONEXISTENT",
			false,
		},
	}
	for _, tt := range gpstests {
		got := asset.GenerateGPSForAreaID(tt.in)
		if !tt.out && len(got) != 0 { // todo check gps
			t.Errorf("%v => want %v", got, tt.out)
		} else if tt.out {
			gps := strings.Split(got, ",")
			lat, _ := strconv.ParseFloat(gps[0], 64)
			lng, _ := strconv.ParseFloat(gps[1], 64)
			c := asset.Coordinates[tt.in]
			deltaLimit := 1.0 / 40.0
			if math.Abs(c.Lat-lat) > deltaLimit {
				t.Errorf("wrong lat: %v => want %v", lat, c.Lat)
			} else if math.Abs(c.Lng-lng) > deltaLimit {
				t.Errorf("wrong lng: %v => want %v", lng, c.Lng)
			}
		}
	}
}

func TestGetAreaIDByStationID(t *testing.T) {
	asset, _ := NewAsset()
	var areatests = []struct {
		in  string
		out string
	}{
		{
			"TBS",
			"JP13",
		},
		{
			"MBS",
			"JP27",
		},
		{
			"NONEXISTENT",
			"",
		},
	}
	for _, tt := range areatests {
		got := asset.GetAreaIDByStationID(tt.in)
		if got != tt.out {
			t.Errorf("%v => want %v", got, tt.out)
		}
	}
}

func TestGetStationIDsByAreaID(t *testing.T) {
	asset, _ := NewAsset()
	var stationtests = []struct {
		in  string
		out []string
	}{
		{
			"JP13",
			[]string{
				"FMJ",
				"FMT",
				"INT",
				"JOAB",
				"JOAK",
				"JOAK-FM",
				"JORF",
				"LFR",
				"QRR",
				"RN1",
				"RN2",
				"TBS",
			},
		},
		{
			"NONEXISTENT",
			[]string{},
		},
	}
	for _, tt := range stationtests {
		got := asset.GetStationIDsByAreaID(tt.in)
		less := func(a, b string) bool { return a < b }
		if !cmp.Equal(got, tt.out, cmpopts.SortSlices(less)) {
			t.Errorf("%v => want %v", got, tt.out)
		}
	}
}

func TestNewDevice(t *testing.T) {
	a, _ := NewAsset()
	device := a.NewDevice()

	if device.AppName != "aSmartPhone7a" {
		t.Errorf("%v => want %v", device.AppName, "aSmartPhone7a")
	}
	if device.Connection != "wifi" {
		t.Errorf("%v => want %v", device.Connection, "wifi")
	}
	var got string
	got = device.UserID
	if len(got) != 32 {
		t.Errorf("%v => want %v", len(got), 32)
	}
	got = device.AppVersion
	if m, _ := regexp.Match(`^7\.[2-5]\.[0-9]{1,2}$`, []byte(got)); !m {
		t.Errorf("invalid AppVersion: %v", got)
	}
	got = device.Name
	if m, _ := regexp.Match(`^[0-9]{2}\..*$`, []byte(got)); !m {
		t.Errorf("invalid Name: %v", got)
	}
	got = device.UserAgent
	if !strings.HasPrefix(got, "Dalvik/2.1.0") {
		t.Errorf("invalid UserAgent: %v", got)
	}
}