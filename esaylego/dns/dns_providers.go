package dns

import (
	"fmt"

	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/acmedns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/alidns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/auroradns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/azure"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/bindman"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/bluecat"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/cloudflare"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/cloudns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/cloudxns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/conoha"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/designate"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/digitalocean"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dnsimple"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dnsmadeeasy"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dnspod"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dode"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dreamhost"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/duckdns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/dyn"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/easydns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/exec"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/exoscale"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/fastdns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/gandi"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/gandiv5"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/gcloud"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/glesys"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/godaddy"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/hostingde"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/httpreq"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/iij"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/inwx"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/joker"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/lightsail"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/linode"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/linodev4"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/mydnsjp"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/namecheap"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/namedotcom"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/namesilo"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/netcup"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/nifcloud"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/ns1"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/oraclecloud"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/otc"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/ovh"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/pdns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/rackspace"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/rfc2136"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/route53"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/sakuracloud"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/selectel"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/stackpath"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/transip"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/vegadns"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/versio"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/vscale"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/vultr"
	"github.com/Green-rainBit/easytls-sdk-go/esaylego/dns/zoneee"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
)

func NewDNSChallengeProviderByName(name string, values map[string]string) (challenge.Provider, error) {
	switch name {
	case "acme-dns":
		return acmedns.NewDNSProviderByValues(values)
	case "alidns":
		return alidns.NewDNSProviderByValues(values)
	case "azure":
		return azure.NewDNSProviderByValues(values)
	case "auroradns":
		return auroradns.NewDNSProviderByValues(values)
	case "bindman":
		return bindman.NewDNSProviderByValues(values)
	case "bluecat":
		return bluecat.NewDNSProviderByValues(values)
	case "cloudflare":
		return cloudflare.NewDNSProviderByValues(values)
	case "cloudns":
		return cloudns.NewDNSProviderByValues(values)
	case "cloudxns":
		return cloudxns.NewDNSProviderByValues(values)
	case "conoha":
		return conoha.NewDNSProviderByValues(values)
	case "designate":
		return designate.NewDNSProviderByValues(values)
	case "digitalocean":
		return digitalocean.NewDNSProviderByValues(values)
	case "dnsimple":
		return dnsimple.NewDNSProviderByValues(values)
	case "dnsmadeeasy":
		return dnsmadeeasy.NewDNSProviderByValues(values)
	case "dnspod":
		return dnspod.NewDNSProviderByValues(values)
	case "dode":
		return dode.NewDNSProviderByValues(values)
	case "dreamhost":
		return dreamhost.NewDNSProviderByValues(values)
	case "duckdns":
		return duckdns.NewDNSProviderByValues(values)
	case "dyn":
		return dyn.NewDNSProviderByValues(values)
	case "fastdns":
		return fastdns.NewDNSProviderByValues(values)
	case "easydns":
		return easydns.NewDNSProviderByValues(values)
	case "exec":
		return exec.NewDNSProviderByValues(values)
	case "exoscale":
		return exoscale.NewDNSProviderByValues(values)
	case "gandi":
		return gandi.NewDNSProviderByValues(values)
	case "gandiv5":
		return gandiv5.NewDNSProviderByValues(values)
	case "glesys":
		return glesys.NewDNSProviderByValues(values)
	case "gcloud":
		return gcloud.NewDNSProviderByValues(values)
	case "godaddy":
		return godaddy.NewDNSProviderByValues(values)
	case "hostingde":
		return hostingde.NewDNSProviderByValues(values)
	case "httpreq":
		return httpreq.NewDNSProviderByValues(values)
	case "iij":
		return iij.NewDNSProviderByValues(values)
	case "inwx":
		return inwx.NewDNSProviderByValues(values)
	case "joker":
		return joker.NewDNSProviderByValues(values)
	case "lightsail":
		return lightsail.NewDNSProviderByValues(values)
	case "linode":
		return linode.NewDNSProviderByValues(values)
	case "linodev4":
		return linodev4.NewDNSProviderByValues(values)
	case "manual":
		return dns01.NewDNSProviderManual()
	case "mydnsjp":
		return mydnsjp.NewDNSProviderByValues(values)
	case "namecheap":
		return namecheap.NewDNSProviderByValues(values)
	case "namedotcom":
		return namedotcom.NewDNSProviderByValues(values)
	case "namesilo":
		return namesilo.NewDNSProviderByValues(values)
	case "netcup":
		return netcup.NewDNSProviderByValues(values)
	case "nifcloud":
		return nifcloud.NewDNSProviderByValues(values)
	case "ns1":
		return ns1.NewDNSProviderByValues(values)
	case "oraclecloud":
		return oraclecloud.NewDNSProviderByValues(values)
	case "otc":
		return otc.NewDNSProviderByValues(values)
	case "ovh":
		return ovh.NewDNSProviderByValues(values)
	case "pdns":
		return pdns.NewDNSProviderByValues(values)
	case "rackspace":
		return rackspace.NewDNSProviderByValues(values)
	case "route53":
		return route53.NewDNSProviderByValues(values)
	case "rfc2136":
		return rfc2136.NewDNSProviderByValues(values)
	case "sakuracloud":
		return sakuracloud.NewDNSProviderByValues(values)
	case "stackpath":
		return stackpath.NewDNSProviderByValues(values)
	case "selectel":
		return selectel.NewDNSProviderByValues(values)
	case "transip":
		return transip.NewDNSProviderByValues(values)
	case "vegadns":
		return vegadns.NewDNSProviderByValues(values)
	case "versio":
		return versio.NewDNSProviderByValues(values)
	case "vultr":
		return vultr.NewDNSProviderByValues(values)
	case "vscale":
		return vscale.NewDNSProviderByValues(values)
	case "zoneee":
		return zoneee.NewDNSProviderByValues(values)
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
