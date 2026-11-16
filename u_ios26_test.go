package tls

import (
	"reflect"
	"testing"
)

func TestIOS26IsSafari263(t *testing.T) {
	ios, err := UTLSIdToSpec(HelloIOS_26)
	if err != nil {
		t.Fatal(err)
	}
	safari, err := UTLSIdToSpec(HelloSafari_26_3)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ios.CipherSuites, safari.CipherSuites) || len(ios.Extensions) != len(safari.Extensions) {
		t.Fatal("iOS 26 sends Safari 26.3's ClientHello")
	}
	if HelloIOS_Auto != HelloIOS_26 {
		t.Fatal("HelloIOS_Auto is iOS 26")
	}
}
