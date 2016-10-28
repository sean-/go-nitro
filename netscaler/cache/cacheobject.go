package cache

type Cacheobject struct {
  Cachecellappfwmetadataexists string `json:"cachecellappfwmetadataexists,omitempty"`
  Cachecellbasefile string `json:"cachecellbasefile,omitempty"`
  Cachecellcomplex string `json:"cachecellcomplex,omitempty"`
  Cachecellcompressionformat string `json:"cachecellcompressionformat,omitempty"`
  Cachecellcurmisses int `json:"cachecellcurmisses,omitempty"`
  Cachecellcurreaders int `json:"cachecellcurreaders,omitempty"`
  Cachecelldestipverified string `json:"cachecelldestipverified,omitempty"`
  Cachecelldhits int `json:"cachecelldhits,omitempty"`
  Cachecelletaginserted string `json:"cachecelletaginserted,omitempty"`
  Cachecellexpires int `json:"cachecellexpires,omitempty"`
  Cachecellexpiresmillisec int `json:"cachecellexpiresmillisec,omitempty"`
  Cachecellfwpxyobj string `json:"cachecellfwpxyobj,omitempty"`
  Cachecellhits int `json:"cachecellhits,omitempty"`
  Cachecellhttp11 string `json:"cachecellhttp11,omitempty"`
  Cachecellminhit int `json:"cachecellminhit,omitempty"`
  Cachecellminhitflag string `json:"cachecellminhitflag,omitempty"`
  Cachecellmisses int `json:"cachecellmisses,omitempty"`
  Cachecellpolleverytime string `json:"cachecellpolleverytime,omitempty"`
  Cachecellreadywithlastbyte string `json:"cachecellreadywithlastbyte,omitempty"`
  Cachecellreqtime int `json:"cachecellreqtime,omitempty"`
  Cachecellresbadsize string `json:"cachecellresbadsize,omitempty"`
  Cachecellrestime int `json:"cachecellrestime,omitempty"`
  Cachecellweaketag string `json:"cachecellweaketag,omitempty"`
  Cachecontrol string `json:"cachecontrol,omitempty"`
  Cachecurage int `json:"cachecurage,omitempty"`
  Cachedirname string `json:"cachedirname,omitempty"`
  Cacheetag string `json:"cacheetag,omitempty"`
  Cachefilename string `json:"cachefilename,omitempty"`
  Cacheindisk string `json:"cacheindisk,omitempty"`
  Cacheinmemory string `json:"cacheinmemory,omitempty"`
  Cacheresdate string `json:"cacheresdate,omitempty"`
  Cachereshdrsize int `json:"cachereshdrsize,omitempty"`
  Cachereslastmod string `json:"cachereslastmod,omitempty"`
  Cacheressize int `json:"cacheressize,omitempty"`
  Cacheurls string `json:"cacheurls,omitempty"`
  Ceflags int `json:"ceflags,omitempty"`
  Contentgroup string `json:"contentgroup,omitempty"`
  Destipv46 string `json:"destipv46,omitempty"`
  Destport int `json:"destport,omitempty"`
  Flushed string `json:"flushed,omitempty"`
  Force bool `json:"force,omitempty"`
  Group string `json:"group,omitempty"`
  Groupname string `json:"groupname,omitempty"`
  Hitparams interface{} `json:"hitparams,omitempty"`
  Hitvalues interface{} `json:"hitvalues,omitempty"`
  Host string `json:"host,omitempty"`
  Httpcalloutcell string `json:"httpcalloutcell,omitempty"`
  Httpcalloutname string `json:"httpcalloutname,omitempty"`
  Httpcalloutresult string `json:"httpcalloutresult,omitempty"`
  Httpmethod string `json:"httpmethod,omitempty"`
  Httpstatus int `json:"httpstatus,omitempty"`
  Httpstatusoutput int `json:"httpstatusoutput,omitempty"`
  Ignoremarkerobjects string `json:"ignoremarkerobjects,omitempty"`
  Includenotreadyobjects string `json:"includenotreadyobjects,omitempty"`
  Locator int `json:"locator,omitempty"`
  Markerreason string `json:"markerreason,omitempty"`
  Policy int `json:"policy,omitempty"`
  Policyname string `json:"policyname,omitempty"`
  Port int `json:"port,omitempty"`
  Prefetch string `json:"prefetch,omitempty"`
  Prefetchperiod int `json:"prefetchperiod,omitempty"`
  Prefetchperiodmillisec int `json:"prefetchperiodmillisec,omitempty"`
  Returntype string `json:"returntype,omitempty"`
  Rule interface{} `json:"rule,omitempty"`
  Selectorname interface{} `json:"selectorname,omitempty"`
  Selectorvalue interface{} `json:"selectorvalue,omitempty"`
  Totalobjs int `json:"totalobjs,omitempty"`
  Url string `json:"url,omitempty"`
  Warnbucketskip int `json:"warnbucketskip,omitempty"`
}
