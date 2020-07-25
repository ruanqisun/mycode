package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRKH3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH3I
}

type DAILRKH3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH3O
}

type DAILRKH3RequestForm struct {
	Form []DAILRKH3IDataForm
}

type DAILRKH3ResponseForm struct {
	Form []DAILRKH3ODataForm
}

type DAILRKH3I struct {
	NmSnglTypCd                  string  `json:"NmSnglTypCd"`
	OrgNo                        string  `json:"OrgNo"`
	NmSnglNo                     string  `json:"NmSnglNo"`
	NmSnglSeqNo                  int     `json:"NmSnglSeqNo"`
	SorcDescr                    string  `json:"SorcDescr"`
	RecmndEmpnbr                 string  `json:"RecmndEmpnbr"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	MobileNo                     string  `json:"MobileNo"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperProcesCd                 string  `json:"OperProcesCd"`
	OperDt                       string  `json:"OperDt"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	RplshInfo                    string  `json:"RplshInfo"`
	TxSn                         string  `json:"TxSn"`
	OprOrgNo                     string  `json:"OprOrgNo"`
	OprorEmpnbr                  string  `json:"OprorEmpnbr"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRKH3O struct {
	Records                      []DAILRKH3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRKH3ORecords struct {
	NmSnglTypCd                  string  `json:"NmSnglTypCd"`
	OrgNo                        string  `json:"OrgNo"`
	NmSnglNo                     string  `json:"NmSnglNo"`
	NmSnglSeqNo                  int     `json:"NmSnglSeqNo"`
	SorcDescr                    string  `json:"SorcDescr"`
	RecmndEmpnbr                 string  `json:"RecmndEmpnbr"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	MobileNo                     string  `json:"MobileNo"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperProcesCd                 string  `json:"OperProcesCd"`
	OperDt                       string  `json:"OperDt"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	RplshInfo                    string  `json:"RplshInfo"`
	TxSn                         string  `json:"TxSn"`
	OprOrgNo                     string  `json:"OprOrgNo"`
	OprorEmpnbr                  string  `json:"OprorEmpnbr"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRKH3RequestForm) PackRequest(dailrkh3I DAILRKH3I) (responseBody []byte, err error) {

	requestForm := DAILRKH3RequestForm{
		Form: []DAILRKH3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH3I",
				},

				FormData: dailrkh3I,
			},
		},
	}

	responseBody, err = json.Marshal(requestForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.REQPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing request message
func (o *DAILRKH3RequestForm) UnPackRequest(request []byte) (DAILRKH3I, error) {
	dailrkh3I := DAILRKH3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRKH3ResponseForm) PackResponse(dailrkh3O DAILRKH3O) (responseBody []byte, err error) {
	responseForm := DAILRKH3ResponseForm{
		Form: []DAILRKH3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH3O",
				},
				FormData: dailrkh3O,
			},
		},
	}

	responseBody, err = json.Marshal(responseForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.RSPPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing response message
func (o *DAILRKH3ResponseForm) UnPackResponse(request []byte) (DAILRKH3O, error) {

	dailrkh3O := DAILRKH3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRKH3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
