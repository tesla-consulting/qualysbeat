package beater

import (
	"fmt"
	"time"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/fs015/qualysbeat/config"
	"net/http"
	"io/ioutil"
	xj "github.com/basgys/goxml2json"
	"github.com/tidwall/gjson"
	"strings"

)

// qualysbeat configuration.
type qualysbeat struct { 
	done   chan struct{}
	config config.Config
	client beat.Client
}

func RetList(user string,password string,cliente string)string{
	req, err := http.NewRequest("GET", "https://qualysapi.qg2.apps.qualys.eu/api/2.0/fo/scan/?action=list", nil)
	if err != nil {
		// handle err
		logp.Info("RetList: Error setting the list from Qualys")
		panic("Error to setting list")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("X-Requested-With", "Curl Sample")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		logp.Info("RetList: Error getting the list from Qualys")
		panic("Error to getting list")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xml := strings.NewReader(string(body))
	json, _ := xj.Convert(xml)
	scan := gjson.Get(json.String(), "SCAN_LIST_OUTPUT.RESPONSE.SCAN_LIST.SCAN.#(STATUS.STATE==\"Finished\")#").Array()[0]
	ref := gjson.Get(scan.String(),"REF").String()
	//fmt.Println(ref)
	return  ref
}

func RetScan(ref string,user string,passw string)string{
	req, err := http.NewRequest("GET", "https://qualysapi.qg2.apps.qualys.eu/api/2.0/fo/scan/?action=fetch&output_format=json_extended&scan_ref="+ref+"&mode=extended", nil)
	if err != nil {
		// handle err
		logp.Info("RetScan: Error setting the scan from Qualys")
		panic("Error to setting scan")
	}
	req.SetBasicAuth(user, passw)
	req.Header.Set("X-Requested-With", "Curl Sample")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		logp.Info("RetScan: Error getting the scan from Qualys")
		panic("Error to getting scan")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}


// New creates an instance of qualysbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &qualysbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts qualysbeat.
func (bt *qualysbeat) Run(b *beat.Beat) error {
	logp.Info("qualysbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	path := "/etc/qualysbeat/qualys.conf"
	dat, err_re := ioutil.ReadFile(path)
	if err_re != nil{
		fmt.Println("========================================================================")
		fmt.Println("						ERRORE NELL'APRIRE IL FILE CONF					 ")
		fmt.Println("								PATH:",path,"							 ")
		fmt.Println("								ERR:",err_re,"							 ")
		fmt.Println("========================================================================")
		logp.Info("RetList: Error to open path "+path)
		panic("Error to open path")
	}

	user := gjson.Get(string(dat),"qualys.user").String()
	password := gjson.Get(string(dat),"qualys.password").String()
	cliente := gjson.Get(string(dat),"qualys.cliente").String()		

	fmt.Println("========================================================================")
	fmt.Println("						SETTAGGIO										 ")
	fmt.Println("								PATH:",path,"							 ")
	fmt.Println("								CLIENTE:",cliente,"						 ")
	fmt.Println("========================================================================")
	if len(user) <1 || len(cliente) < 1 || len(password) < 1 {
		logp.Info("Insufficient credentials to run the query")
		panic("Insufficient credentials to run the query")
	}
	
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
		scanlist := RetList(user,password,cliente)

		out := RetScan(scanlist,user,password)

		ltemp:= "{\"lista\":"+out+"}"
		list := gjson.Get(ltemp,"lista")

		n_item := len(list.Array())
		fmt.Println(n_item)
		for i,v := range list.Array(){

			if i>2 && i<n_item-1 {
	
				result := make(common.MapStr)

				vulnerability := make(map[string]interface{})
				severity := gjson.Get(v.String(),"severity")
				typed := gjson.Get(v.String(),"type")
				istance := gjson.Get(v.String(),"istance")
				pci_vuln := gjson.Get(v.String(),"pci_vuln")
				bugtraq_id := gjson.Get(v.String(),"bugtraq_id")
				associated_malware := gjson.Get(v.String(),"associated_malware")
				exploitability := gjson.Get(v.String(),"exploitability")
				solution := gjson.Get(v.String(),"solution")
				impact := gjson.Get(v.String(),"impact")
				title := gjson.Get(v.String(),"title")
				results := gjson.Get(v.String(),"results")
				netbios := gjson.Get(v.String(),"netbios")
				vulnerability["severity"] = severity.String()
				vulnerability["type"] = typed.String()
				vulnerability["istance"] = istance.String()
				vulnerability["pci_vuln"] = pci_vuln.String()
				vulnerability["bugtraq_id"] = bugtraq_id.String()
				vulnerability["associated_malware"] = associated_malware.String()
				vulnerability["exploitability"] = exploitability.String()
				vulnerability["solution"] = solution.String()
				vulnerability["impact"] = impact.String()
				vulnerability["title"] = title.String()
				vulnerability["results"] = results.String()
				vulnerability["netbios"] = netbios.String()
					
				host := make(map[string]interface{})
				host["os"] = gjson.Get(v.String(),"os").String()
				host["ssl"] = gjson.Get(v.String(),"ssl").String()
				host["fqdn"] = gjson.Get(v.String(),"fqdn").String()
				host["dns"] = gjson.Get(v.String(),"dns").String()
				host["protocol"] = gjson.Get(v.String(),"protocol").String()
				host["port"] = gjson.Get(v.String(),"port").String()
				host["ip"] = gjson.Get(v.String(),"ip").String()
				vulnerability["host"] = host
					
				score := make(map[string]interface{})
				
				score["version"] = "2.0"
				score["temporal"] = gjson.Get(v.String(),"cvss_temporal").String()
				score["base"] = gjson.Get(v.String(),"cvss_base").String()
				vulnerability["score"] = score
					
				scanner := make(map[string]interface{})
					
				scanner["vendor"] = "Qualys"
				scanner["id"] = gjson.Get(v.String(),"qid").String()
				vulnerability["scanner"] = scanner
					
				vulnerability["id"] = gjson.Get(v.String(),"cve_id").String()
				vulnerability["enumeration"] = "CVE"
				vulnerability["description"] =  gjson.Get(v.String(),"threat").String()
				vulnerability["category"] =  gjson.Get(v.String(),"category").String()

				result.Put("vulnerability", vulnerability)

				event := beat.Event{
					Timestamp: time.Now(),
					Fields: result,
				}
				bt.client.Publish(event)
				logp.Info("Event sent")
				counter++
			}
		}
	}
}

// Stop stops qualysbeat.
func (bt *qualysbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
