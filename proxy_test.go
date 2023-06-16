package main

import (
	"testing"
)

func TestReadProxyFromFile(t *testing.T) {
	ReadProxyFromFile("proxy")
	if len(proxyList) != 268 {
		t.Fatalf("Expected len for this data: %d\n But got len: %d", 268, len(proxyList))
	}

	t.Log("Function worked properly.")
}

func TestBuildProxyURL(t *testing.T) {
	proxy := BuildProxyURL("1.1.1.1:1111")
	if proxy != "http://1.1.1.1:1111" {
		t.Errorf("Expected result: http://1.1.1.1:1111,\n But got: %v", proxy)
	}

	t.Log(proxy)
}

func TestGetProxy(t *testing.T) {
	ReadProxyFromFile("proxy")
	proxy := GetProxy()
	t.Log(proxy)
}
