package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK5I
}

type DAILRDK5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK5O
}

type DAILRDK5RequestForm struct {
	Form []DAILRDK5IDataForm
}

type DAILRDK5ResponseForm struct {
	Form []DAILRDK5ODataForm
}

type DAILRDK5I struct {
	RelaAcctId                   string  `json:"RelaAcctId"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	RelaAcctBlngtoRwCategCd      string  `json:"RelaAcctBlngtoRwCategCd"`
	CustNo                       string  `json:"CustNo"`
	RelaAcctNoTypCd              string  `json:"RelaAcctNoTypCd"`
	DfltRepayAcctFlg             string  `json:"DfltRepayAcctFlg"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	RelaAcctNo                   string  `json:"RelaAcctNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	RelaDt                       string  `json:"RelaDt"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	RelaAcctOpnAcctBnkBnkNo      string  `json:"RelaAcctOpnAcctBnkBnkNo"`
	RelaAcctOpnAcctBnkBnkName    string  `json:"RelaAcctOpnAcctBnkBnkName"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK5O struct {
	Records                      []DAILRDK5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK5ORecords struct {
	RelaAcctId                   string  `json:"RelaAcctId"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	RelaAcctBlngtoRwCategCd      string  `json:"RelaAcctBlngtoRwCategCd"`
	CustNo                       string  `json:"CustNo"`
	RelaAcctNoTypCd              string  `json:"RelaAcctNoTypCd"`
	DfltRepayAcctFlg             string  `json:"DfltRepayAcctFlg"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	RelaAcctNo                   string  `json:"RelaAcctNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	RelaDt                       string  `json:"RelaDt"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	RelaAcctOpnAcctBnkBnkNo      string  `json:"RelaAcctOpnAcctBnkBnkNo"`
	RelaAcctOpnAcctBnkBnkName    string  `json:"RelaAcctOpnAcctBnkBnkName"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK5RequestForm) PackRequest(dailrdk5I DAILRDK5I) (responseBody []byte, err error) {

	requestForm := DAILRDK5RequestForm{
		Form: []DAILRDK5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK5I",
				},
				FormData: dailrdk5I,
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
func (o *DAILRDK5RequestForm) UnPackRequest(request []byte) (DAILRDK5I, error) {
	dailrdk5I := DAILRDK5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK5ResponseForm) PackResponse(dailrdk5O DAILRDK5O) (responseBody []byte, err error) {
	responseForm := DAILRDK5ResponseForm{
		Form: []DAILRDK5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK5O",
				},
				FormData: dailrdk5O,
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
func (o *DAILRDK5ResponseForm) UnPackResponse(request []byte) (DAILRDK5O, error) {

	dailrdk5O := DAILRDK5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
