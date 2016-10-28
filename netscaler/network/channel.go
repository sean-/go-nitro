package network

type Channel struct {
  Actflowctl string `json:"actflowctl,omitempty"`
  Actspeed string `json:"actspeed,omitempty"`
  Actthroughput int `json:"actthroughput,omitempty"`
  Autoneg int `json:"autoneg,omitempty"`
  Autonegresult int `json:"autonegresult,omitempty"`
  Backplane string `json:"backplane,omitempty"`
  Bandwidthhigh int `json:"bandwidthhigh,omitempty"`
  Bandwidthnormal int `json:"bandwidthnormal,omitempty"`
  Bdgmuted int `json:"bdgmuted,omitempty"`
  Cleartime int `json:"cleartime,omitempty"`
  Conndistr string `json:"conndistr,omitempty"`
  Description string `json:"description,omitempty"`
  Devicename string `json:"devicename,omitempty"`
  Downtime int `json:"downtime,omitempty"`
  Duplex string `json:"duplex,omitempty"`
  Fctls int `json:"fctls,omitempty"`
  Flags int `json:"flags,omitempty"`
  Flowctl string `json:"flowctl,omitempty"`
  Hamonitor string `json:"hamonitor,omitempty"`
  Hangdetect int `json:"hangdetect,omitempty"`
  Hangreset int `json:"hangreset,omitempty"`
  Hangs int `json:"hangs,omitempty"`
  Id string `json:"id,omitempty"`
  Ifalias string `json:"ifalias,omitempty"`
  Ifnum interface{} `json:"ifnum,omitempty"`
  Indisc int `json:"indisc,omitempty"`
  Intfstate int `json:"intfstate,omitempty"`
  Lacpactoraggregation string `json:"lacpactoraggregation,omitempty"`
  Lacpactorcollecting string `json:"lacpactorcollecting,omitempty"`
  Lacpactordistributing string `json:"lacpactordistributing,omitempty"`
  Lacpactorinsync string `json:"lacpactorinsync,omitempty"`
  Lacpactorportno int `json:"lacpactorportno,omitempty"`
  Lacpactorpriority int `json:"lacpactorpriority,omitempty"`
  Lacpmode string `json:"lacpmode,omitempty"`
  Lacppartneraggregation string `json:"lacppartneraggregation,omitempty"`
  Lacppartnercollecting string `json:"lacppartnercollecting,omitempty"`
  Lacppartnerdefaulted string `json:"lacppartnerdefaulted,omitempty"`
  Lacppartnerdistributing string `json:"lacppartnerdistributing,omitempty"`
  Lacppartnerexpired string `json:"lacppartnerexpired,omitempty"`
  Lacppartnerinsync string `json:"lacppartnerinsync,omitempty"`
  Lacppartnerkey int `json:"lacppartnerkey,omitempty"`
  Lacppartnerportno int `json:"lacppartnerportno,omitempty"`
  Lacppartnerpriority int `json:"lacppartnerpriority,omitempty"`
  Lacppartnerstate string `json:"lacppartnerstate,omitempty"`
  Lacppartnersystemmac string `json:"lacppartnersystemmac,omitempty"`
  Lacppartnersystempriority int `json:"lacppartnersystempriority,omitempty"`
  Lacppartnertimeout string `json:"lacppartnertimeout,omitempty"`
  Lacpportmuxstate string `json:"lacpportmuxstate,omitempty"`
  Lacpportrxstat string `json:"lacpportrxstat,omitempty"`
  Lacpportselectstate string `json:"lacpportselectstate,omitempty"`
  Lacptimeout string `json:"lacptimeout,omitempty"`
  Lamac string `json:"lamac,omitempty"`
  Lamode string `json:"lamode,omitempty"`
  Linkstate int `json:"linkstate,omitempty"`
  Lrminthroughput int `json:"lrminthroughput,omitempty"`
  Mac string `json:"mac,omitempty"`
  Macdistr string `json:"macdistr,omitempty"`
  Media string `json:"media,omitempty"`
  Mode string `json:"mode,omitempty"`
  Mtu int `json:"mtu,omitempty"`
  Outdisc int `json:"outdisc,omitempty"`
  Reqduplex string `json:"reqduplex,omitempty"`
  Reqflowcontrol string `json:"reqflowcontrol,omitempty"`
  Reqmedia string `json:"reqmedia,omitempty"`
  Reqspeed string `json:"reqspeed,omitempty"`
  Reqthroughput int `json:"reqthroughput,omitempty"`
  Rxbytes int `json:"rxbytes,omitempty"`
  Rxdrops int `json:"rxdrops,omitempty"`
  Rxerrors int `json:"rxerrors,omitempty"`
  Rxpackets int `json:"rxpackets,omitempty"`
  Rxstalls int `json:"rxstalls,omitempty"`
  Speed string `json:"speed,omitempty"`
  State string `json:"state,omitempty"`
  Stsstalls int `json:"stsstalls,omitempty"`
  Tagall string `json:"tagall,omitempty"`
  Tagged int `json:"tagged,omitempty"`
  Taggedany int `json:"taggedany,omitempty"`
  Taggedautolearn int `json:"taggedautolearn,omitempty"`
  Throughput int `json:"throughput,omitempty"`
  Trunk string `json:"trunk,omitempty"`
  Txbytes int `json:"txbytes,omitempty"`
  Txdrops int `json:"txdrops,omitempty"`
  Txerrors int `json:"txerrors,omitempty"`
  Txpackets int `json:"txpackets,omitempty"`
  Txstalls int `json:"txstalls,omitempty"`
  Unit int `json:"unit,omitempty"`
  Uptime int `json:"uptime,omitempty"`
  Vlan int `json:"vlan,omitempty"`
  Vmac string `json:"vmac,omitempty"`
  Vmac6 string `json:"vmac6,omitempty"`
}
