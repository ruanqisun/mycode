package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDW2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDW2I
}

type DAILRDW2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDW2O
}

type DAILRDW2RequestForm struct {
	Form []DAILRDW2IDataForm
}

type DAILRDW2ResponseForm struct {
	Form []DAILRDW2ODataForm
}

type DAILRDW2I struct {
	TxSn                          string  `json:"TxSn"`
	CustNo                        string  `json:"CustNo"`
	DocMgmtNo                     string  `json:"DocMgmtNo"`
	OperTypCd                     string  `json:"OperTypCd"`
	OperPersonEmpnbr              string  `json:"OperPersonEmpnbr"`
	OperDt                        string  `json:"OperDt"`
	OperTm                        string  `json:"OperTm"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDW2O struct {
	Records                      []DAILRDW2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDW2ORecords struct {
	TxSn                          string  `json:"TxSn"`
	CustNo                        string  `json:"CustNo"`
	DocMgmtNo                     string  `json:"DocMgmtNo"`
	OperTypCd                     string  `json:"OperTypCd"`
	OperPersonEmpnbr              string  `json:"OperPersonEmpnbr"`
	OperDt                        string  `json:"OperDt"`
	OperTm                        string  `json:"OperTm"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDW2RequestForm) PackRequest(dailrdw2I DAILRDW2I) (responseBody []byte, err error) {

	requestForm := DAILRDW2RequestForm{
		Form: []DAILRDW2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDW2I",
				},

				FormData: dailrdw2I,
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
func (o *DAILRDW2RequestForm) UnPackRequest(request []byte) (DAILRDW2I, error) {
	dailrdw2I := DAILRDW2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdw2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdw2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDW2ResponseForm) PackResponse(dailrdw2O DAILRDW2O) (responseBody []byte, err error) {
	responseForm := DAILRDW2ResponseForm{
		Form: []DAILRDW2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDW2O",
				},
				FormData: dailrdw2O,
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
func (o *DAILRDW2ResponseForm) UnPackResponse(request []byte) (DAILRDW2O, error) {

	dailrdw2O := DAILRDW2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdw2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdw2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDW2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
