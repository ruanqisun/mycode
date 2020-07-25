package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0016IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0016I
}

type IL8R0016ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0016O
}

type IL8R0016RequestForm struct {
	Form []IL8R0016IDataForm
}

type IL8R0016ResponseForm struct {
	Form []IL8R0016ODataForm
}

type IL8R0016I struct {
	BizEventCd                     string  `json:"BizEventCd"`
	BizEventNm                     string  `json:"BizEventNm"`
	BizEventTypCd                  string  `json:"BizEventTypCd"`
	BizEventTypNm                  string  `json:"BizEventTypNm"`
	EfftDt                         string  `json:"EfftDt"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0016O struct {
	Records                      []IL8R0016ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0016ORecords struct {
	BizEventCd                     string  `json:"BizEventCd"`
	BizEventNm                     string  `json:"BizEventNm"`
	BizEventTypCd                  string  `json:"BizEventTypCd"`
	BizEventTypNm                  string  `json:"BizEventTypNm"`
	EfftDt                         string  `json:"EfftDt"`
	ValidFlg                       string  `json:"ValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0016RequestForm) PackRequest(il8r0016I IL8R0016I) (responseBody []byte, err error) {

	requestForm := IL8R0016RequestForm{
		Form: []IL8R0016IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0016I",
				},

				FormData: il8r0016I,
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
func (o *IL8R0016RequestForm) UnPackRequest(request []byte) (IL8R0016I, error) {
	il8r0016I := IL8R0016I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0016I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0016I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0016ResponseForm) PackResponse(il8r0016O IL8R0016O) (responseBody []byte, err error) {
	responseForm := IL8R0016ResponseForm{
		Form: []IL8R0016ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0016O",
				},
				FormData: il8r0016O,
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
func (o *IL8R0016ResponseForm) UnPackResponse(request []byte) (IL8R0016O, error) {

	il8r0016O := IL8R0016O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0016O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0016O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0016I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
