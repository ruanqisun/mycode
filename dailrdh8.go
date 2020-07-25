package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH8IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH8I
}

type DAILRDH8ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH8O
}

type DAILRDH8RequestForm struct {
	Form []DAILRDH8IDataForm
}

type DAILRDH8ResponseForm struct {
	Form []DAILRDH8ODataForm
}

type DAILRDH8I struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	TxSn                         string  `json:"TxSn"`
	BizAplyNo                    string  `json:"BizAplyNo"`
	BizTypCd                     string  `json:"BizTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	SugstnDescr                  string  `json:"SugstnDescr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH8O struct {
	Records                      []DAILRDH8ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH8ORecords struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	TxSn                         string  `json:"TxSn"`
	BizAplyNo                    string  `json:"BizAplyNo"`
	BizTypCd                     string  `json:"BizTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	ApprvSugstnCd                string  `json:"ApprvSugstnCd"`
	SugstnDescr                  string  `json:"SugstnDescr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH8RequestForm) PackRequest(dailrdh8I DAILRDH8I) (responseBody []byte, err error) {

	requestForm := DAILRDH8RequestForm{
		Form: []DAILRDH8IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH8I",
				},

				FormData: dailrdh8I,
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
func (o *DAILRDH8RequestForm) UnPackRequest(request []byte) (DAILRDH8I, error) {
	dailrdh8I := DAILRDH8I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh8I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh8I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH8ResponseForm) PackResponse(dailrdh8O DAILRDH8O) (responseBody []byte, err error) {
	responseForm := DAILRDH8ResponseForm{
		Form: []DAILRDH8ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH8O",
				},
				FormData: dailrdh8O,
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
func (o *DAILRDH8ResponseForm) UnPackResponse(request []byte) (DAILRDH8O, error) {

	dailrdh8O := DAILRDH8O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh8O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh8O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH8I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
