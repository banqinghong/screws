package ldap

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/miekg/dns"
)

type DnsRecord struct {
	DataLength string
	Data       string
}

type service struct {
	cli *ldap.Conn
}

func Run() {
	s := &service{}
	l, err := ldap.DialURL(HostOfLdap)
	if err != nil {
		log.Fatal(err)
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: UserOfLdap,
		Password: PasswordOfLdap,
	})
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}

	s.cli = l
	zones, err := s.GetDnsZones()
	if err != nil {
		fmt.Printf("query dns zones failed: %s\n", err)
		return
	}

	//jsonStr, _ := json.Marshal(zones)
	//fmt.Println("json: ", string(jsonStr))

	fmt.Printf("get %d zone\n", len(zones))
	s.QueryDomainRecord("xgimi.com")
	//queryDomainInfo("redis1.qa.xgimi.com")
	//for _, zone := range zones {
	//	fmt.Printf("start %s\n", zone)
	//	s.QueryDomainRecord(zone)
	//}
}

func (s *service) GetDnsZones() ([]string, error) {
	searchRequest := ldap.NewSearchRequest(
		BaseDN, // The base dn to search
		ldap.ScopeSingleLevel, ldap.DerefAlways, 0, 0, false,
		"(objectClass=dnsZone)", // The filter to apply
		[]string{"dc"},          // A list attributes to retrieve
		nil,
	)
	sr, err := s.cli.SearchWithPaging(searchRequest, 500)
	if err != nil {
		return nil, err
	}
	var zones []string
	for _, entry := range sr.Entries {
		zone := entry.GetAttributeValue("dc")
		zones = append(zones, zone)
	}
	return zones, nil
}

func (s *service) QueryDomainRecord(zone string) {
	// filePath := fmt.Sprintf("./result-%s-%d.json", zone, time.Now().Unix())
	baseDN := fmt.Sprintf("DC=%s,%s", zone, BaseDN)
	searchRequest := ldap.NewSearchRequest(
		baseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		"(objectClass=*)", // The filter to apply
		[]string{"dnsRecord", "dNSTombstoned", "name", "instanceType", "objectClass"}, // A list attributes to retrieve
		nil,
	)

	sr, err := s.cli.SearchWithPaging(searchRequest, 500)
	if err != nil {
		fmt.Printf("query zone[%s] domain record failed: %s\n", zone, err)
		return
	}
	fmt.Println("长度： ", len(sr.Entries))
	//jsonStr, _ := json.Marshal(sr)
	//fmt.Println("json: ", string(jsonStr))
	for _, entry := range sr.Entries {
		if len(entry.Attributes) < 1 {
			nameByDN, err := getDomainNameByDN(entry.DN)
			if err != nil {
				fmt.Printf("get name by dn[%s] failed: %s\n", entry.DN, err)
				continue
			}

			hostName, err := isHostName(nameByDN)
			if err != nil {
				fmt.Printf("filter host name [%s] failed: %s\n", nameByDN, err)
				continue
			}
			if hostName {
				fmt.Printf("[%s] is host\n", nameByDN)
				continue
			}

			domainName := nameByDN + "." + zone
			domainInfo, err := queryDomainInfo(domainName)
			if err != nil {
				continue
			}

			content := fmt.Sprintf("%s: %s -- %s\n", domainName, domainInfo.Type, domainInfo.Address)
			fmt.Println("content: ", content)
			// bfile.Append2File(content, filePath)
		}
	}
}

// 过滤 JM IT等开头的办公网机器地址
func isHostName(nameByDN string) (bool, error) {
	return regexp.Match("^(JM|IT|GM|YBIT|jm|YB|it20|WIN|VDI).*", []byte(nameByDN))
}

func getDomainNameByDN(dn string) (string, error) {
	// DC=JM210921004826,DC=xgimi.com,CN=MicrosoftDNS,DC=DomainDnsZones,DC=xgimi,DC=com
	dnSplit := strings.Split(dn, ",")
	if len(dnSplit) < 6 {
		return "", errors.New("parse dn error")
	}
	nameSplit := strings.Split(dnSplit[0], "=")
	if len(nameSplit) != 2 {
		return "", errors.New("parse dn name error")
	}
	return nameSplit[1], nil
}

type domainInfo struct {
	Address string
	Type    string
	Name    string
}

//func queryDomainAddr(domain string) (*domainInfo, error) {
//	domainInfo := &domainInfo{
//		Name: domain,
//	}
//	ipByteList, err := net.LookupIP(domain)
//	if err != nil {
//		return nil, err
//	}
//	if len(ipByteList) == 1 {
//		domainInfo.Type = "A"
//		domainInfo.Address = fmt.Sprintf("%s", ipByteList[0])
//		return domainInfo, nil
//	}
//	cname, err  := net.LookupCNAME(domain)
//	if err != nil {
//		return nil, err
//	}
//	domainInfo.Type = "CNAME"
//	domainInfo.Address = cname
//	return domainInfo, nil
//}

func queryDomainInfo(domain string) (*domainInfo, error) {
	domainInfo := &domainInfo{
		Name: domain,
	}
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	m := &dns.Msg{}
	m.SetQuestion(domain+".", dns.TypeA)
	r, _, err := c.Exchange(m, DomainServer)
	if err != nil {
		fmt.Printf("query dns[%s] failed: %s\n", domain, err)
		return nil, err
	}

	//jsonStr, _ := json.Marshal(r)
	//fmt.Println(string(jsonStr))
	if len(r.Answer) < 1 {
		return nil, errors.New("no record found")
	}

	switch record := r.Answer[0].(type) {
	case *dns.A:
		domainInfo.Type = "A"
		domainInfo.Address = record.A.String()
	case *dns.CNAME:
		domainInfo.Type = "CNAME"
		domainInfo.Address = record.Target
	default:
		fmt.Println("answer", r.Answer[0].String())
		return nil, errors.New("unknown type")
	}
	return domainInfo, nil
}
