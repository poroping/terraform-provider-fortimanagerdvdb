package utils

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ValidateConvIPMask2CIDR(oNewIP string, oOldIPList []string) []string {
	oOldNetmask := strings.Split(oOldIPList[1], ".")
	if len(oOldNetmask) > 1 {
		prefixSize, _ := net.IPMask(net.ParseIP(oOldIPList[1]).To4()).Size()
		oOldIPList[1] = strconv.Itoa(prefixSize)
	}
	oOldIP := strings.Join(oOldIPList, "/")
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return []string{ip, strconv.Itoa(prefixSize)}
		}
	}
	return strings.Split(oOldIP, " ")
}

func Ipv4Split(s string) []string {
	ip, m, err := net.ParseCIDR(s)
	if err != nil {
		log.Printf("[WARN] Unable to parse CIDR, %v", err)
	}
	mask := fmt.Sprintf("%d.%d.%d.%d", m.Mask[0], m.Mask[1], m.Mask[2], m.Mask[3])
	return []string{ip.String(), mask}
	// mask := strings.Split(l[1], ".")
	// if len(mask) != 0 {
	// 	prefixSize, _ := net.IPMask(net.ParseIP(l[1]).To4()).Size()
	// 	l2 := []string{l[0], strconv.Itoa(prefixSize)}
	// 	return l2
	// }
	// return l
}

func Ipv4Read(d *schema.ResourceData, get string, ip []string) *[]string {
	if current, ok := d.GetOk(get); ok {
		if s, ok := current.([]string); ok {
			tmp2 := strings.Join(s, "/")
			tmp3 := ValidateConvIPMask2CIDR(tmp2, ip)
			return &tmp3
		}
	}
	return nil
}

func Ipv4NetmaskListToCidr(ipList []string) string {
	ip := ipList[0]
	mask := ipList[1]
	prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
	return fmt.Sprintf("%s/%s", ip, strconv.Itoa(prefixSize))

}

func escapeFilter(filter string) string {
	filter = strings.ReplaceAll(filter, "_", "-")
	filter = strings.ReplaceAll(filter, "fosid", "id")
	filter = strings.ReplaceAll(filter, "&", "&filter=")
	filter = strings.ReplaceAll(filter, ".", "\\.")
	filter = strings.ReplaceAll(filter, "\\", "\\\\")
	filter = "filter=" + filter
	return filter
}

func SortSubtable(result []map[string]interface{}, fieldname string) {
	sort.Slice(result, func(i, j int) bool {
		v1 := result[i][fieldname]
		v2 := result[j][fieldname]
		if reflect.TypeOf(v1).Name() == "string" {
			return v1.(string) < v2.(string)
		} else if reflect.TypeOf(v1).Name() == "int" {
			return int(v1.(int)) < int(v2.(int))
		} else if reflect.TypeOf(v1).Name() == "int64" {
			return int(v1.(int64)) < int(v2.(int64))
		} else if reflect.TypeOf(v1).Name() == "float64" {
			return int(v1.(float64)) < int(v2.(float64))
		} else {
			println("error") // TODO: make this error less crap
			return true
		}
	})
}

func ParseDownloadedPemCertificate(v string) (*map[string]interface{}, error) {
	der, rest := pem.Decode([]byte(v))

	if der == nil {
		return nil, fmt.Errorf("error, valid certificate not found: %v", rest)
	}

	cert, err := x509.ParseCertificate(der.Bytes)

	if err != nil {
		return nil, fmt.Errorf("error parsing certificate: %v", err)
	}

	o := map[string]interface{}{
		"signature_algorithm":  cert.SignatureAlgorithm.String(),
		"public_key_algorithm": cert.PublicKeyAlgorithm.String(),
		"serial_number":        cert.SerialNumber.String(),
		"is_ca":                cert.IsCA,
		"version":              cert.Version,
		"issuer":               cert.Issuer.String(),
		"subject":              cert.Subject.String(),
		"not_before":           cert.NotBefore.Format(time.RFC3339),
		"not_after":            cert.NotAfter.Format(time.RFC3339),
		"sha1_fingerprint":     fmt.Sprintf("%x", sha1.Sum(cert.Raw)),
		"is_valid":             time.Now().Before(cert.NotAfter) && time.Now().After(cert.NotBefore),
	}

	return &o, nil
}

func ParseMkey(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	if i, ok := v.(float64); ok {
		return strconv.Itoa(int(i))
	}
	if i, ok := v.(int64); ok {
		return strconv.Itoa(int(i))
	}
	if i, ok := v.(int); ok {
		return strconv.Itoa(i)
	}
	return ""
}

func CheckVer(curVer, addedVer, removedVer string) bool {
	if curVer == "" || (addedVer == "" && removedVer == "") {
		return true
	}

	cur, err := semver.NewVersion(curVer)
	if err != nil {
		log.Fatal(err)
	}

	add := "*"
	rm := "*"

	if addedVer != "" {
		add = fmt.Sprintf(">= %s", addedVer)
	}

	if removedVer != "" {
		rm = fmt.Sprintf("< %s", removedVer)
	}

	c, err := semver.NewConstraint(strings.Join([]string{add, rm}, ","))
	if err != nil {
		log.Fatal(err)
	}
	return c.Check(cur)
}

func AttributeVersionWarning(attribute, version string) diag.Diagnostic {
	e := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  fmt.Sprintf("Attribute %q not supported in FortiOS version %q. Attempting to use anyway.", attribute, version),
		Detail:   fmt.Sprintf("Attribute %q not suppported in FortiOS version %q. Attempting to use anyway. This could cause errors downstream. Remove offending attribute or overwrite version in provider configuration.", attribute, version),
	}
	e.Validate()
	log.Printf("[WARN] Attribute %q not supported in version %q. Attempting to use anyway.", attribute, version)
	return e
}
