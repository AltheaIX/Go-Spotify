package main

import "testing"

func TestGenerateNumber(t *testing.T) {
	randomInt := GenerateNumber(10, 20)
	if randomInt < 10 || randomInt > 20 {
		t.Fatalf("Expected to below or equals with 20 and higher or equals with 10.\n But got: %d", randomInt)
	}

	t.Log(randomInt)
}

func TestGetEmailDomain(t *testing.T) {
	domainList, err := GetEmailDomain()
	if err != nil {
		t.Error(err)
	}

	t.Log(domainList)
}
