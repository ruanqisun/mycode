package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0033IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0033I
}

type IL8R0033ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0033O
}

type IL8R0033RequestForm struct {
	Form []IL8R0033IDataForm
}

type IL8R0033ResponseForm struct {
	Form []IL8R0033ODataForm
}

type IL8R0033I struct {
	CustNo                         string  `json:"CustNo"`
	CustName                       string  `json:"CustName"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	FsttmBizDt                     string  `json:"FsttmBizDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0033O struct {
	Records                      []IL8R0033ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0033ORecords struct {
	CustNo                         string  `json:"CustNo"`
	CustName                       string  `json:"CustName"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	FsttmBizDt                     string  `json:"FsttmBizDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0033RequestForm) PackRequest(il8r0033I IL8R0033I) (responseBody []byte, err error) {

	requestForm := IL8R0033RequestForm{
		Form: []IL8R0033IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0033I",
				},

				FormData: il8r0033I,
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
func (o *IL8R0033RequestForm) UnPackRequest(request []byte) (IL8R0033I, error) {
	il8r0033I := IL8R0033I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0033I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0033I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0033ResponseForm) PackResponse(il8r0033O IL8R0033O) (responseBody []byte, err error) {
	responseForm := IL8R0033ResponseForm{
		Form: []IL8R0033ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0033O",
				},
				FormData: il8r0033O,
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
func (o *IL8R0033ResponseForm) UnPackResponse(request []byte) (IL8R0033O, error) {

	il8r0033O := IL8R0033O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0033O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0033O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0033I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
