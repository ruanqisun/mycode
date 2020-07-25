package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0024IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0024I
}

type IL8R0024ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0024O
}

type IL8R0024RequestForm struct {
	Form []IL8R0024IDataForm
}

type IL8R0024ResponseForm struct {
	Form []IL8R0024ODataForm
}

type IL8R0024I struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	SoftProdtNm                    string  `json:"SoftProdtNm"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0024O struct {
	Records                      []IL8R0024ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0024ORecords struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	SoftProdtNm                    string  `json:"SoftProdtNm"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0024RequestForm) PackRequest(il8r0024I IL8R0024I) (responseBody []byte, err error) {

	requestForm := IL8R0024RequestForm{
		Form: []IL8R0024IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0024I",
				},

				FormData: il8r0024I,
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
func (o *IL8R0024RequestForm) UnPackRequest(request []byte) (IL8R0024I, error) {
	il8r0024I := IL8R0024I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0024I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0024I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0024ResponseForm) PackResponse(il8r0024O IL8R0024O) (responseBody []byte, err error) {
	responseForm := IL8R0024ResponseForm{
		Form: []IL8R0024ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0024O",
				},
				FormData: il8r0024O,
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
func (o *IL8R0024ResponseForm) UnPackResponse(request []byte) (IL8R0024O, error) {

	il8r0024O := IL8R0024O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0024O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0024O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0024I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
