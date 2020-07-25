package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL3I
}

type DAILRPL3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL3O
}

type DAILRPL3RequestForm struct {
	Form []DAILRPL3IDataForm
}

type DAILRPL3ResponseForm struct {
	Form []DAILRPL3ODataForm
}

type DAILRPL3I struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	PridnumSeqNo                 int     `json:"PridnumSeqNo"`
	IntStlDt                     string  `json:"IntStlDt"`
	CurCd                        string  `json:"CurCd"`
	IntStlTotlAmt                float64 `json:"IntStlTotlAmt"`
	IntStlPrin                   float64 `json:"IntStlPrin"`
	IntStlIntr                   float64 `json:"IntStlIntr"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRPL3O struct {
	Records                      []DAILRPL3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL3ORecords struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	PridnumSeqNo                 int     `json:"PridnumSeqNo"`
	IntStlDt                     string  `json:"IntStlDt"`
	CurCd                        string  `json:"CurCd"`
	IntStlTotlAmt                float64 `json:"IntStlTotlAmt"`
	IntStlPrin                   float64 `json:"IntStlPrin"`
	IntStlIntr                   float64 `json:"IntStlIntr"`
}

// @Desc Build request message
func (o *DAILRPL3RequestForm) PackRequest(dailrpl3I DAILRPL3I) (responseBody []byte, err error) {

	requestForm := DAILRPL3RequestForm{
		Form: []DAILRPL3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL3I",
				},

				FormData: dailrpl3I,
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
func (o *DAILRPL3RequestForm) UnPackRequest(request []byte) (DAILRPL3I, error) {
	dailrpl3I := DAILRPL3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL3ResponseForm) PackResponse(dailrpl3O DAILRPL3O) (responseBody []byte, err error) {
	responseForm := DAILRPL3ResponseForm{
		Form: []DAILRPL3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL3O",
				},
				FormData: dailrpl3O,
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
func (o *DAILRPL3ResponseForm) UnPackResponse(request []byte) (DAILRPL3O, error) {

	dailrpl3O := DAILRPL3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
