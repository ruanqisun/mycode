package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0023IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0023I
}

type IL8R0023ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0023O
}

type IL8R0023RequestForm struct {
	Form []IL8R0023IDataForm
}

type IL8R0023ResponseForm struct {
	Form []IL8R0023ODataForm
}

type IL8R0023I struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	ProdtNm                        string  `json:"ProdtNm"`
	EfftDt                         string  `json:"EfftDt"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0023O struct {
	Records                      []IL8R0023ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0023ORecords struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	ProdtNm                        string  `json:"ProdtNm"`
	EfftDt                         string  `json:"EfftDt"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0023RequestForm) PackRequest(il8r0023I IL8R0023I) (responseBody []byte, err error) {

	requestForm := IL8R0023RequestForm{
		Form: []IL8R0023IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0023I",
				},

				FormData: il8r0023I,
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
func (o *IL8R0023RequestForm) UnPackRequest(request []byte) (IL8R0023I, error) {
	il8r0023I := IL8R0023I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0023I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0023I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0023ResponseForm) PackResponse(il8r0023O IL8R0023O) (responseBody []byte, err error) {
	responseForm := IL8R0023ResponseForm{
		Form: []IL8R0023ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0023O",
				},
				FormData: il8r0023O,
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
func (o *IL8R0023ResponseForm) UnPackResponse(request []byte) (IL8R0023O, error) {

	il8r0023O := IL8R0023O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0023O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0023O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0023I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
