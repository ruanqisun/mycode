package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0025IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0025I
}

type IL8R0025ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0025O
}

type IL8R0025RequestForm struct {
	Form []IL8R0025IDataForm
}

type IL8R0025ResponseForm struct {
	Form []IL8R0025ODataForm
}

type IL8R0025I struct {
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	MstkAmt                        float64  `json:"MstkAmt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0025O struct {
	Records                      []IL8R0025ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0025ORecords struct {
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	MstkAmt                        float64  `json:"MstkAmt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0025RequestForm) PackRequest(il8r0025I IL8R0025I) (responseBody []byte, err error) {

	requestForm := IL8R0025RequestForm{
		Form: []IL8R0025IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0025I",
				},

				FormData: il8r0025I,
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
func (o *IL8R0025RequestForm) UnPackRequest(request []byte) (IL8R0025I, error) {
	il8r0025I := IL8R0025I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0025I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0025I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0025ResponseForm) PackResponse(il8r0025O IL8R0025O) (responseBody []byte, err error) {
	responseForm := IL8R0025ResponseForm{
		Form: []IL8R0025ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0025O",
				},
				FormData: il8r0025O,
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
func (o *IL8R0025ResponseForm) UnPackResponse(request []byte) (IL8R0025O, error) {

	il8r0025O := IL8R0025O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0025O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0025O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0025I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
