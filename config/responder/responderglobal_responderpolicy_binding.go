package responder

type Responderglobalresponderpolicybinding struct {
  Flowtype int `json:"flowtype,omitempty"`
  Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
  Invoke bool `json:"invoke,omitempty"`
  Labelname string `json:"labelname,omitempty"`
  Labeltype string `json:"labeltype,omitempty"`
  Numpol int `json:"numpol,omitempty"`
  Policyname string `json:"policyname,omitempty"`
  Priority int `json:"priority,omitempty"`
  Type string `json:"type,omitempty"`
}