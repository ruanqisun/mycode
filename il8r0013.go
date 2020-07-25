package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0013IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0013I
}

type IL8R0013ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0013O
}

type IL8R0013RequestForm struct {
	Form []IL8R0013IDataForm
}

type IL8R0013ResponseForm struct {
	Form []IL8R0013ODataForm
}

type IL8R0013I struct {
	AcctnDt                        string  `json:"AcctnDt"`
	SubjNo                         string  `json:"SubjNo"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SubjNm                         string  `json:"SubjNm"`
	SubjStusCd                     string  `json:"SubjStusCd"`
	SubjBal                        float64  `json:"SubjBal"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0013O struct {
	Records                      []IL8R0013ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0013ORecords struct {
	AcctnDt                        string  `json:"AcctnDt"`
	SubjNo                         string  `json:"SubjNo"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SubjNm                         string  `json:"SubjNm"`
	SubjStusCd                     string  `json:"SubjStusCd"`
	SubjBal                        float64  `json:"SubjBal"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0013RequestForm) PackRequest(il8r0013I IL8R0013I) (responseBody []byte, err error) {

	requestForm := IL8R0013RequestForm{
		Form: []IL8R0013IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0013I",
				},

				FormData: il8r0013I,
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
func (o *IL8R0013RequestForm) UnPackRequest(request []byte) (IL8R0013I, error) {
	il8r0013I := IL8R0013I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0013I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0013I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0013ResponseForm) PackResponse(il8r0013O IL8R0013O) (responseBody []byte, err error) {
	responseForm := IL8R0013ResponseForm{
		Form: []IL8R0013ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0013O",
				},
				FormData: il8r0013O,
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
func (o *IL8R0013ResponseForm) UnPackResponse(request []byte) (IL8R0013O, error) {

	il8r0013O := IL8R0013O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0013O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0013O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0013I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
