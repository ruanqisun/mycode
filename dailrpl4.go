package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL4I
}

type DAILRPL4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL4O
}

type DAILRPL4RequestForm struct {
	Form []DAILRPL4IDataForm
}

type DAILRPL4ResponseForm struct {
	Form []DAILRPL4ODataForm
}

type DAILRPL4I struct {
	LoanAcctNo                  string  `json:"LoanAcctNo"`
	AcctnDt                     string  `json:"AcctnDt"`
	PridnumSeqNo                int     `json:"PridnumSeqNo"`
	AcctiAcctNo                 string  `json:"AcctiAcctNo"`
	OvdueBgnDt                  string  `json:"OvdueBgnDt"`
	CurCd                       string  `json:"CurCd"`
	OvduePrin                   float64  `json:"OvduePrin"`
	OvdueIntr                   float64  `json:"OvdueIntr"`
	CmpdAmt                     float64  `json:"CmpdAmt"`
	OvduePnltint                float64  `json:"OvduePnltint"`
	Cmpndint                    float64  `json:"Cmpndint"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRPL4O struct {
	Records                      []DAILRPL4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL4ORecords struct {
	LoanAcctNo                  string  `json:"LoanAcctNo"`
	AcctnDt                     string  `json:"AcctnDt"`
	PridnumSeqNo                int     `json:"PridnumSeqNo"`
	AcctiAcctNo                 string  `json:"AcctiAcctNo"`
	OvdueBgnDt                  string  `json:"OvdueBgnDt"`
	CurCd                       string  `json:"CurCd"`
	OvduePrin                   float64  `json:"OvduePrin"`
	OvdueIntr                   float64  `json:"OvdueIntr"`
	CmpdAmt                     float64  `json:"CmpdAmt"`
	OvduePnltint                float64  `json:"OvduePnltint"`
	Cmpndint                    float64  `json:"Cmpndint"`
}

// @Desc Build request message
func (o *DAILRPL4RequestForm) PackRequest(dailrpl4I DAILRPL4I) (responseBody []byte, err error) {

	requestForm := DAILRPL4RequestForm{
		Form: []DAILRPL4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL4I",
				},

				FormData: dailrpl4I,
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
func (o *DAILRPL4RequestForm) UnPackRequest(request []byte) (DAILRPL4I, error) {
	dailrpl4I := DAILRPL4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL4ResponseForm) PackResponse(dailrpl4O DAILRPL4O) (responseBody []byte, err error) {
	responseForm := DAILRPL4ResponseForm{
		Form: []DAILRPL4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL4O",
				},
				FormData: dailrpl4O,
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
func (o *DAILRPL4ResponseForm) UnPackResponse(request []byte) (DAILRPL4O, error) {

	dailrpl4O := DAILRPL4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
