package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRCP1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRCP1I
}

type DAILRCP1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRCP1O
}

type DAILRCP1RequestForm struct {
	Form []DAILRCP1IDataForm
}

type DAILRCP1ResponseForm struct {
	Form []DAILRCP1ODataForm
}

type DAILRCP1I struct {
	OrgNo                          string  `json:"OrgNo"`
	SeqNo                          int     `json:"SeqNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	CurCd                          string  `json:"CurCd"`
	MgctManrCd                     string  `json:"MgctManrCd"`
	BalCtrlTypCd                   string  `json:"BalCtrlTypCd"`
	BalCtrlAmt                     float64 `json:"BalCtrlTypCd"`
	BalCtrlBgnDt                   string  `json:"BalCtrlBgnDt"`
	BalCtrlEndDt                   string  `json:"BalCtrlEndDt"`
	NoneDeadlBalCtrlEfftDt         string  `json:"NoneDeadlBalCtrlEfftDt"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	CrtEmpnbr                      string  `json:"CrtEmpnbr"`
	CrtDt                          string  `json:"CrtDt"`
	ApprvEmpnbr                    string  `json:"ApprvEmpnbr"`
	ApprvDt                        string  `json:"ApprvDt"`
	ApprvSugstnCd                  string  `json:"ApprvSugstnCd"`
	ApprvSugstnComnt               string  `json:"ApprvSugstnComnt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type DAILRCP1O struct {
	Records                      []DAILRCP1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRCP1ORecords struct {
	OrgNo                          string  `json:"OrgNo"`
	SeqNo                          int     `json:"SeqNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	CurCd                          string  `json:"CurCd"`
	MgctManrCd                     string  `json:"MgctManrCd"`
	BalCtrlTypCd                   string  `json:"BalCtrlTypCd"`
	BalCtrlAmt                     float64 `json:"BalCtrlTypCd"`
	BalCtrlBgnDt                   string  `json:"BalCtrlBgnDt"`
	BalCtrlEndDt                   string  `json:"BalCtrlEndDt"`
	NoneDeadlBalCtrlEfftDt         string  `json:"NoneDeadlBalCtrlEfftDt"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	CrtEmpnbr                      string  `json:"CrtEmpnbr"`
	CrtDt                          string  `json:"CrtDt"`
	ApprvEmpnbr                    string  `json:"ApprvEmpnbr"`
	ApprvDt                        string  `json:"ApprvDt"`
	ApprvSugstnCd                  string  `json:"ApprvSugstnCd"`
	ApprvSugstnComnt               string  `json:"ApprvSugstnComnt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRCP1RequestForm) PackRequest(dailrcp1I DAILRCP1I) (responseBody []byte, err error) {

	requestForm := DAILRCP1RequestForm{
		Form: []DAILRCP1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRCP1I",
				},

				FormData: dailrcp1I,
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
func (o *DAILRCP1RequestForm) UnPackRequest(request []byte) (DAILRCP1I, error) {
	dailrcp1I := DAILRCP1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrcp1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrcp1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRCP1ResponseForm) PackResponse(dailrcp1O DAILRCP1O) (responseBody []byte, err error) {
	responseForm := DAILRCP1ResponseForm{
		Form: []DAILRCP1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRCP1O",
				},
				FormData: dailrcp1O,
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
func (o *DAILRCP1ResponseForm) UnPackResponse(request []byte) (DAILRCP1O, error) {

	dailrcp1O := DAILRCP1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrcp1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrcp1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRCP1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
