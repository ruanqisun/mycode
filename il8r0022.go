package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0022IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0022I
}

type IL8R0022ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0022O
}

type IL8R0022RequestForm struct {
	Form []IL8R0022IDataForm
}

type IL8R0022ResponseForm struct {
	Form []IL8R0022ODataForm
}

type IL8R0022I struct {
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	BizEventCd                     string  `json:"BizEventCd"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	UnuslRsn                       string  `json:"UnuslRsn"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0022O struct {
	Records                      []IL8R0022ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0022ORecords struct {
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	BizEventCd                     string  `json:"BizEventCd"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	UnuslRsn                       string  `json:"UnuslRsn"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0022RequestForm) PackRequest(il8r0022I IL8R0022I) (responseBody []byte, err error) {

	requestForm := IL8R0022RequestForm{
		Form: []IL8R0022IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0022I",
				},

				FormData: il8r0022I,
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
func (o *IL8R0022RequestForm) UnPackRequest(request []byte) (IL8R0022I, error) {
	il8r0022I := IL8R0022I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0022I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0022I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0022ResponseForm) PackResponse(il8r0022O IL8R0022O) (responseBody []byte, err error) {
	responseForm := IL8R0022ResponseForm{
		Form: []IL8R0022ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0022O",
				},
				FormData: il8r0022O,
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
func (o *IL8R0022ResponseForm) UnPackResponse(request []byte) (IL8R0022O, error) {

	il8r0022O := IL8R0022O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0022O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0022O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0022I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
