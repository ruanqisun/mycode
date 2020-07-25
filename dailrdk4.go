package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK4I
}

type DAILRDK4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK4O
}

type DAILRDK4RequestForm struct {
	Form []DAILRDK4IDataForm
}

type DAILRDK4ResponseForm struct {
	Form []DAILRDK4ODataForm
}

type DAILRDK4I struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	SeqNo                        int     `json:"SeqNo"`
	TxSn                         string  `json:"TxSn"`
	MakelnStusCd                 string  `json:"MakelnStusCd"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	MakelnAcctNo                 string  `json:"MakelnAcctNo"`
	CurCd                        string  `json:"CurCd"`
	LoanAmt                      float64 `json:"LoanAmt"`
	MakelnOrgNo                  string  `json:"MakelnOrgNo"`
	TxTm                         string  `json:"TxTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK4O struct {
	Records                      []DAILRDK4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK4ORecords struct {
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	SeqNo                        int     `json:"SeqNo"`
	TxSn                         string  `json:"TxSn"`
	MakelnStusCd                 string  `json:"MakelnStusCd"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	MakelnAcctNo                 string  `json:"MakelnAcctNo"`
	CurCd                        string  `json:"CurCd"`
	LoanAmt                      float64 `json:"LoanAmt"`
	MakelnOrgNo                  string  `json:"MakelnOrgNo"`
	TxTm                         string  `json:"TxTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK4RequestForm) PackRequest(dailrdk4I DAILRDK4I) (responseBody []byte, err error) {

	requestForm := DAILRDK4RequestForm{
		Form: []DAILRDK4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK4I",
				},
				FormData: dailrdk4I,
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
func (o *DAILRDK4RequestForm) UnPackRequest(request []byte) (DAILRDK4I, error) {
	dailrdk4I := DAILRDK4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK4ResponseForm) PackResponse(dailrdk4O DAILRDK4O) (responseBody []byte, err error) {
	responseForm := DAILRDK4ResponseForm{
		Form: []DAILRDK4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK4O",
				},
				FormData: dailrdk4O,
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
func (o *DAILRDK4ResponseForm) UnPackResponse(request []byte) (DAILRDK4O, error) {

	dailrdk4O := DAILRDK4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
