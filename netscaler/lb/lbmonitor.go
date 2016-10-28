package lb

type Lbmonitor struct {
  Acctapplicationid interface{} `json:"acctapplicationid,omitempty"`
  Action string `json:"action,omitempty"`
  Alertretries int `json:"alertretries,omitempty"`
  Application string `json:"application,omitempty"`
  Attribute string `json:"attribute,omitempty"`
  Authapplicationid interface{} `json:"authapplicationid,omitempty"`
  Basedn string `json:"basedn,omitempty"`
  Binddn string `json:"binddn,omitempty"`
  Customheaders string `json:"customheaders,omitempty"`
  Database string `json:"database,omitempty"`
  Destip string `json:"destip,omitempty"`
  Destport int `json:"destport,omitempty"`
  Deviation int `json:"deviation,omitempty"`
  Dispatcherip string `json:"dispatcherip,omitempty"`
  Dispatcherport int `json:"dispatcherport,omitempty"`
  Domain string `json:"domain,omitempty"`
  Downtime int `json:"downtime,omitempty"`
  Dupstate string `json:"dup_state,omitempty"`
  Dupweight int `json:"dup_weight,omitempty"`
  Dynamicinterval int `json:"dynamicinterval,omitempty"`
  Dynamicresponsetimeout int `json:"dynamicresponsetimeout,omitempty"`
  Evalrule string `json:"evalrule,omitempty"`
  Failureretries int `json:"failureretries,omitempty"`
  Filename string `json:"filename,omitempty"`
  Filter string `json:"filter,omitempty"`
  Firmwarerevision int `json:"firmwarerevision,omitempty"`
  Group string `json:"group,omitempty"`
  Hostipaddress string `json:"hostipaddress,omitempty"`
  Hostname string `json:"hostname,omitempty"`
  Httprequest string `json:"httprequest,omitempty"`
  Inbandsecurityid string `json:"inbandsecurityid,omitempty"`
  Interval int `json:"interval,omitempty"`
  Ipaddress interface{} `json:"ipaddress,omitempty"`
  Iptunnel string `json:"iptunnel,omitempty"`
  Kcdaccount string `json:"kcdaccount,omitempty"`
  Lasversion string `json:"lasversion,omitempty"`
  Logonpointname string `json:"logonpointname,omitempty"`
  Lrtm string `json:"lrtm,omitempty"`
  Lrtmconf int `json:"lrtmconf,omitempty"`
  Maxforwards int `json:"maxforwards,omitempty"`
  Metric string `json:"metric,omitempty"`
  Metrictable string `json:"metrictable,omitempty"`
  Metricthreshold int `json:"metricthreshold,omitempty"`
  Metricweight int `json:"metricweight,omitempty"`
  Monitorname string `json:"monitorname,omitempty"`
  Mssqlprotocolversion string `json:"mssqlprotocolversion,omitempty"`
  Multimetrictable interface{} `json:"multimetrictable,omitempty"`
  Netprofile string `json:"netprofile,omitempty"`
  Originhost string `json:"originhost,omitempty"`
  Originrealm string `json:"originrealm,omitempty"`
  Password string `json:"password,omitempty"`
  Productname string `json:"productname,omitempty"`
  Query string `json:"query,omitempty"`
  Querytype string `json:"querytype,omitempty"`
  Radaccountsession string `json:"radaccountsession,omitempty"`
  Radaccounttype int `json:"radaccounttype,omitempty"`
  Radapn string `json:"radapn,omitempty"`
  Radframedip string `json:"radframedip,omitempty"`
  Radkey string `json:"radkey,omitempty"`
  Radmsisdn string `json:"radmsisdn,omitempty"`
  Radnasid string `json:"radnasid,omitempty"`
  Radnasip string `json:"radnasip,omitempty"`
  Recv string `json:"recv,omitempty"`
  Respcode interface{} `json:"respcode,omitempty"`
  Resptimeout int `json:"resptimeout,omitempty"`
  Resptimeoutthresh int `json:"resptimeoutthresh,omitempty"`
  Retries int `json:"retries,omitempty"`
  Reverse string `json:"reverse,omitempty"`
  Rtsprequest string `json:"rtsprequest,omitempty"`
  Scriptargs string `json:"scriptargs,omitempty"`
  Scriptname string `json:"scriptname,omitempty"`
  Secondarypassword string `json:"secondarypassword,omitempty"`
  Secure string `json:"secure,omitempty"`
  Send string `json:"send,omitempty"`
  Servicegroupname string `json:"servicegroupname,omitempty"`
  Servicename string `json:"servicename,omitempty"`
  Sipmethod string `json:"sipmethod,omitempty"`
  Sipreguri string `json:"sipreguri,omitempty"`
  Sipuri string `json:"sipuri,omitempty"`
  Sitepath string `json:"sitepath,omitempty"`
  Snmpcommunity string `json:"snmpcommunity,omitempty"`
  Snmpoid string `json:"Snmpoid,omitempty"`
  Snmpthreshold string `json:"snmpthreshold,omitempty"`
  Snmpversion string `json:"snmpversion,omitempty"`
  Sqlquery string `json:"sqlquery,omitempty"`
  State string `json:"state,omitempty"`
  Storedb string `json:"storedb,omitempty"`
  Storefrontacctservice string `json:"storefrontacctservice,omitempty"`
  Storename string `json:"storename,omitempty"`
  Successretries int `json:"successretries,omitempty"`
  Supportedvendorids interface{} `json:"supportedvendorids,omitempty"`
  Tos string `json:"tos,omitempty"`
  Tosid int `json:"tosid,omitempty"`
  Transparent string `json:"transparent,omitempty"`
  Type string `json:"type,omitempty"`
  Units1 string `json:"units1,omitempty"`
  Units2 string `json:"units2,omitempty"`
  Units3 string `json:"units3,omitempty"`
  Units4 string `json:"units4,omitempty"`
  Username string `json:"username,omitempty"`
  Validatecred string `json:"validatecred,omitempty"`
  Vendorid int `json:"vendorid,omitempty"`
  Vendorspecificacctapplicationids interface{} `json:"vendorspecificacctapplicationids,omitempty"`
  Vendorspecificauthapplicationids interface{} `json:"vendorspecificauthapplicationids,omitempty"`
  Vendorspecificvendorid int `json:"vendorspecificvendorid,omitempty"`
  Weight int `json:"weight,omitempty"`
}
