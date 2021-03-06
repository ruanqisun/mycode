package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0032IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0032I
}

type IL8R0032ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0032O
}

type IL8R0032RequestForm struct {
	Form []IL8R0032IDataForm
}

type IL8R0032ResponseForm struct {
	Form []IL8R0032ODataForm
}

type IL8R0032I struct {
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	SeqNo                          int     `json:"SeqNo"`
	Pridnum                        int     `json:"Pridnum"`
	AcctnDt                        string  `json:"AcctnDt"`
	LoanBatBizTypCd                string  `json:"LoanBatBizTypCd"`
	CustNo                         string  `json:"CustNo"`
	BizSn                          string  `json:"BizSn"`
	BatJobStepStusCd               string  `json:"BatJobStepStusCd"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0032O struct {
	Records                      []IL8R0032ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0032ORecords struct {
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	SeqNo                          int     `json:"SeqNo"`
	Pridnum                        int     `json:"Pridnum"`
	AcctnDt                        string  `json:"AcctnDt"`
	LoanBatBizTypCd                string  `json:"LoanBatBizTypCd"`
	CustNo                         string  `json:"CustNo"`
	BizSn                          string  `json:"BizSn"`
	BatJobStepStusCd               string  `json:"BatJobStepStusCd"`
}

// @Desc Build request message
func (o *IL8R0032RequestForm) PackRequest(il8r0032I IL8R0032I) (responseBody []byte, err error) {

	requestForm := IL8R0032RequestForm{
		Form: []IL8R0032IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0032I",
				},

				FormData: il8r0032I,
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
func (o *IL8R0032RequestForm) UnPackRequest(request []byte) (IL8R0032I, error) {
	il8r0032I := IL8R0032I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0032I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0032I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0032ResponseForm) PackResponse(il8r0032O IL8R0032O) (responseBody []byte, err error) {
	responseForm := IL8R0032ResponseForm{
		Form: []IL8R0032ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0032O",
				},
				FormData: il8r0032O,
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
func (o *IL8R0032ResponseForm) UnPackResponse(request []byte) (IL8R0032O, error) {

	il8r0032O := IL8R0032O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0032O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0032O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0032I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
