package dns

type Dnsprofile struct {
  Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
  Cacherecords string `json:"cacherecords,omitempty"`
  Dnsanswerseclogging string `json:"dnsanswerseclogging,omitempty"`
  Dnserrorlogging string `json:"dnserrorlogging,omitempty"`
  Dnsextendedlogging string `json:"dnsextendedlogging,omitempty"`
  Dnsprofilename string `json:"dnsprofilename,omitempty"`
  Dnsquerylogging string `json:"dnsquerylogging,omitempty"`
  Referencecount int `json:"referencecount,omitempty"`
}
