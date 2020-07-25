package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYX2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX2I
}

type DAILRYX2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYX2O
}

type DAILRYX2RequestForm struct {
	Form []DAILRYX2IDataForm
}

type DAILRYX2ResponseForm struct {
	Form []DAILRYX2ODataForm
}

type DAILRYX2I struct {
	CustNo                        string  `json:"CustNo"`
	SeqNo                         int     `json:"SeqNo"`
	CustName                      string  `json:"CustName"`
	CtrtNo                        string  `json:"CtrtNo"`
	LoanDubilNo                   string  `json:"LoanDubilNo"`
	RgtsintClsfNo                 string  `json:"RgtsintClsfNo"`
	RgtsintNo                     string  `json:"RgtsintNo"`
	RgtsintUseStusCd              string  `json:"RgtsintUseStusCd"`
	UseDt                         string  `json:"UseDt"`
	CnclDt                        string  `json:"CnclDt"`
	RgtsintMatrDt                 string  `json:"RgtsintMatrDt"`
	ExecFinhDt                    string  `json:"ExecFinhDt"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYX2O struct {
	Records                      []DAILRYX2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYX2ORecords struct {
	CustNo                        string  `json:"CustNo"`
	SeqNo                         int     `json:"SeqNo"`
	CustName                      string  `json:"CustName"`
	CtrtNo                        string  `json:"CtrtNo"`
	LoanDubilNo                   string  `json:"LoanDubilNo"`
	RgtsintClsfNo                 string  `json:"RgtsintClsfNo"`
	RgtsintNo                     string  `json:"RgtsintNo"`
	RgtsintUseStusCd              string  `json:"RgtsintUseStusCd"`
	UseDt                         string  `json:"UseDt"`
	CnclDt                        string  `json:"CnclDt"`
	RgtsintMatrDt                 string  `json:"RgtsintMatrDt"`
	ExecFinhDt                    string  `json:"ExecFinhDt"`
}

// @Desc Build request message
func (o *DAILRYX2RequestForm) PackRequest(dailryx2I DAILRYX2I) (responseBody []byte, err error) {

	requestForm := DAILRYX2RequestForm{
		Form: []DAILRYX2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX2I",
				},

				FormData: dailryx2I,
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
func (o *DAILRYX2RequestForm) UnPackRequest(request []byte) (DAILRYX2I, error) {
	dailryx2I := DAILRYX2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryx2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYX2ResponseForm) PackResponse(dailryx2O DAILRYX2O) (responseBody []byte, err error) {
	responseForm := DAILRYX2ResponseForm{
		Form: []DAILRYX2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYX2O",
				},
				FormData: dailryx2O,
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
func (o *DAILRYX2ResponseForm) UnPackResponse(request []byte) (DAILRYX2O, error) {

	dailryx2O := DAILRYX2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryx2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryx2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYX2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
