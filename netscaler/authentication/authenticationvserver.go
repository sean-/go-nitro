package authentication

type Authenticationvserver struct {
  Appflowlog string `json:"appflowlog,omitempty"`
  Authentication string `json:"authentication,omitempty"`
  Authenticationdomain string `json:"authenticationdomain,omitempty"`
  Backupvserver string `json:"backupvserver,omitempty"`
  Cachetype string `json:"cachetype,omitempty"`
  Cachevserver string `json:"cachevserver,omitempty"`
  Clttimeout int `json:"clttimeout,omitempty"`
  Comment string `json:"comment,omitempty"`
  Curaaausers int `json:"curaaausers,omitempty"`
  Curstate string `json:"curstate,omitempty"`
  Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
  Downstateflush string `json:"downstateflush,omitempty"`
  Failedlogintimeout int `json:"failedlogintimeout,omitempty"`
  Groupextraction bool `json:"groupextraction,omitempty"`
  Httpprofilename string `json:"httpprofilename,omitempty"`
  Ip string `json:"ip,omitempty"`
  Ipv46 string `json:"ipv46,omitempty"`
  Listenpolicy string `json:"listenpolicy,omitempty"`
  Listenpriority int `json:"listenpriority,omitempty"`
  Maxloginattempts int `json:"maxloginattempts,omitempty"`
  Name string `json:"name,omitempty"`
  Newname string `json:"newname,omitempty"`
  Ngname string `json:"ngname,omitempty"`
  Policy string `json:"policy,omitempty"`
  Port int `json:"port,omitempty"`
  Precedence string `json:"precedence,omitempty"`
  Priority int `json:"priority,omitempty"`
  Range int `json:"range,omitempty"`
  Redirect string `json:"redirect,omitempty"`
  Redirecturl string `json:"redirecturl,omitempty"`
  Rule string `json:"rule,omitempty"`
  Secondary bool `json:"secondary,omitempty"`
  Servicename string `json:"servicename,omitempty"`
  Servicetype string `json:"servicetype,omitempty"`
  Somethod string `json:"somethod,omitempty"`
  Sopersistence string `json:"sopersistence,omitempty"`
  Sopersistencetimeout int `json:"sopersistencetimeout,omitempty"`
  Sothreshold int `json:"sothreshold,omitempty"`
  State string `json:"state,omitempty"`
  Status int `json:"status,omitempty"`
  Tcpprofilename string `json:"tcpprofilename,omitempty"`
  Td int `json:"td,omitempty"`
  Type string `json:"type,omitempty"`
  Value string `json:"value,omitempty"`
  Vstype int `json:"vstype,omitempty"`
  Weight int `json:"weight,omitempty"`
}
