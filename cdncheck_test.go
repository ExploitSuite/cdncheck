package cdncheck

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCDNCheckValid(t *testing.T) {
	client := New()

	found, provider, itemType, err := client.Check(net.ParseIP("173.245.48.12"))
	fmt.Println(found, provider, itemType, err)
	require.Equal(t, "cloudflare", provider, "could not get correct provider")
	require.Equal(t, "waf", itemType, "could not get correct item type")
	require.Nil(t, err, "Could not check ip in ranger")
	require.True(t, found, "Could not check cloudlfare ip blacklist")

	found, _, _, err = client.Check(net.ParseIP("127.0.0.1"))
	require.Nil(t, err, "Could not check ip in ranger")
	require.False(t, found, "Localhost IP found in blacklist")
}

func TestCheckDomain(t *testing.T) {
	client := New()

	valid, provider, itemType, err := client.CheckDomainWithFallback("www.gap.com")
	fmt.Println(valid, provider, itemType, err)
	require.Nil(t, err, "could not check")
	require.True(t, valid, "could not check domain")
	require.Equal(t, "akamai", provider, "could not get correct provider")
	require.Equal(t, "waf", itemType, "could not get correct itemType")
}
