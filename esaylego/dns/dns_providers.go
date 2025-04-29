package dns

import (
	"easytls-sdk-go/esaylego/dns/acmedns"
	"easytls-sdk-go/esaylego/dns/alidns"
	"easytls-sdk-go/esaylego/dns/auroradns"
	"easytls-sdk-go/esaylego/dns/azure"
	"easytls-sdk-go/esaylego/dns/bindman"
	"easytls-sdk-go/esaylego/dns/bluecat"
	"easytls-sdk-go/esaylego/dns/cloudflare"
	"easytls-sdk-go/esaylego/dns/cloudns"
	"easytls-sdk-go/esaylego/dns/cloudxns"
	"easytls-sdk-go/esaylego/dns/conoha"
	"easytls-sdk-go/esaylego/dns/designate"
	"easytls-sdk-go/esaylego/dns/digitalocean"
	"easytls-sdk-go/esaylego/dns/dnsimple"
	"easytls-sdk-go/esaylego/dns/dnsmadeeasy"
	"easytls-sdk-go/esaylego/dns/dnspod"
	"easytls-sdk-go/esaylego/dns/dode"
	"easytls-sdk-go/esaylego/dns/dreamhost"
	"easytls-sdk-go/esaylego/dns/duckdns"
	"easytls-sdk-go/esaylego/dns/dyn"
	"easytls-sdk-go/esaylego/dns/easydns"
	"easytls-sdk-go/esaylego/dns/exec"
	"easytls-sdk-go/esaylego/dns/exoscale"
	"easytls-sdk-go/esaylego/dns/fastdns"
	"easytls-sdk-go/esaylego/dns/gandi"
	"easytls-sdk-go/esaylego/dns/gandiv5"
	"easytls-sdk-go/esaylego/dns/gcloud"
	"easytls-sdk-go/esaylego/dns/glesys"
	"easytls-sdk-go/esaylego/dns/godaddy"
	"easytls-sdk-go/esaylego/dns/hostingde"
	"easytls-sdk-go/esaylego/dns/httpreq"
	"easytls-sdk-go/esaylego/dns/iij"
	"easytls-sdk-go/esaylego/dns/inwx"
	"easytls-sdk-go/esaylego/dns/joker"
	"easytls-sdk-go/esaylego/dns/lightsail"
	"easytls-sdk-go/esaylego/dns/linode"
	"easytls-sdk-go/esaylego/dns/linodev4"
	"easytls-sdk-go/esaylego/dns/mydnsjp"
	"easytls-sdk-go/esaylego/dns/namecheap"
	"easytls-sdk-go/esaylego/dns/namedotcom"
	"easytls-sdk-go/esaylego/dns/namesilo"
	"easytls-sdk-go/esaylego/dns/netcup"
	"easytls-sdk-go/esaylego/dns/nifcloud"
	"easytls-sdk-go/esaylego/dns/ns1"
	"easytls-sdk-go/esaylego/dns/oraclecloud"
	"easytls-sdk-go/esaylego/dns/otc"
	"easytls-sdk-go/esaylego/dns/ovh"
	"easytls-sdk-go/esaylego/dns/pdns"
	"easytls-sdk-go/esaylego/dns/rackspace"
	"easytls-sdk-go/esaylego/dns/rfc2136"
	"easytls-sdk-go/esaylego/dns/route53"
	"easytls-sdk-go/esaylego/dns/sakuracloud"
	"easytls-sdk-go/esaylego/dns/selectel"
	"easytls-sdk-go/esaylego/dns/stackpath"
	"easytls-sdk-go/esaylego/dns/transip"
	"easytls-sdk-go/esaylego/dns/vegadns"
	"easytls-sdk-go/esaylego/dns/versio"
	"easytls-sdk-go/esaylego/dns/vscale"
	"easytls-sdk-go/esaylego/dns/vultr"
	"easytls-sdk-go/esaylego/dns/zoneee"
	"fmt"

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
