package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0021IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0021I
}

type IL8R0021ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0021O
}

type IL8R0021RequestForm struct {
	Form []IL8R0021IDataForm
}

type IL8R0021ResponseForm struct {
	Form []IL8R0021ODataForm
}

type IL8R0021I struct {
	SumId                          string  `json:"SumId"`
	TxAcctnDt                      string  `json:"TxAcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	ReqeRcrdacctAcctnDt            string  `json:"ReqeRcrdacctAcctnDt"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	CurCd                          string  `json:"CurCd"`
	DebitCrdtFlgCd                 string  `json:"DebitCrdtFlgCd"`
	TxAmt                          float64  `json:"TxAmt"`
	SubjNo                         string  `json:"SubjNo"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0021O struct {
	Records                      []IL8R0021ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0021ORecords struct {
	SumId                          string  `json:"SumId"`
	TxAcctnDt                      string  `json:"TxAcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	ReqeRcrdacctAcctnDt            string  `json:"ReqeRcrdacctAcctnDt"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	CurCd                          string  `json:"CurCd"`
	DebitCrdtFlgCd                 string  `json:"DebitCrdtFlgCd"`
	TxAmt                          float64  `json:"TxAmt"`
	SubjNo                         string  `json:"SubjNo"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0021RequestForm) PackRequest(il8r0021I IL8R0021I) (responseBody []byte, err error) {

	requestForm := IL8R0021RequestForm{
		Form: []IL8R0021IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0021I",
				},

				FormData: il8r0021I,
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
func (o *IL8R0021RequestForm) UnPackRequest(request []byte) (IL8R0021I, error) {
	il8r0021I := IL8R0021I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0021I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0021I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0021ResponseForm) PackResponse(il8r0021O IL8R0021O) (responseBody []byte, err error) {
	responseForm := IL8R0021ResponseForm{
		Form: []IL8R0021ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0021O",
				},
				FormData: il8r0021O,
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
func (o *IL8R0021ResponseForm) UnPackResponse(request []byte) (IL8R0021O, error) {

	il8r0021O := IL8R0021O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0021O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0021O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0021I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
