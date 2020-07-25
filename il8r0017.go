package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0017IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0017I
}

type IL8R0017ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0017O
}

type IL8R0017RequestForm struct {
	Form []IL8R0017IDataForm
}

type IL8R0017ResponseForm struct {
	Form []IL8R0017ODataForm
}

type IL8R0017I struct {
	BizEventCd                     string  `json:"BizEventCd"`
	SeqNum                         int  `json:"SeqNum"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	CurCd                          string  `json:"CurCd"`
	AsstClsSubjNo                  string  `json:"AsstClsSubjNo"`
	LiabClsSubjNo                  string  `json:"LiabClsSubjNo"`
	CntptySubjNo                   string  `json:"CntptySubjNo"`
	EfftDt                         string  `json:"EfftDt"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	ValidFlg                       string  `json:"ValidFlg"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	FundFlgCd                      string  `json:"FundFlgCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0017O struct {
	Records                      []IL8R0017ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0017ORecords struct {
	BizEventCd                     string  `json:"BizEventCd"`
	SeqNum                         int  `json:"SeqNum"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	CurCd                          string  `json:"CurCd"`
	AsstClsSubjNo                  string  `json:"AsstClsSubjNo"`
	LiabClsSubjNo                  string  `json:"LiabClsSubjNo"`
	CntptySubjNo                   string  `json:"CntptySubjNo"`
	EfftDt                         string  `json:"EfftDt"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	ValidFlg                       string  `json:"ValidFlg"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	FundFlgCd                      string  `json:"FundFlgCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0017RequestForm) PackRequest(il8r0017I IL8R0017I) (responseBody []byte, err error) {

	requestForm := IL8R0017RequestForm{
		Form: []IL8R0017IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0017I",
				},

				FormData: il8r0017I,
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
func (o *IL8R0017RequestForm) UnPackRequest(request []byte) (IL8R0017I, error) {
	il8r0017I := IL8R0017I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0017I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0017I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0017ResponseForm) PackResponse(il8r0017O IL8R0017O) (responseBody []byte, err error) {
	responseForm := IL8R0017ResponseForm{
		Form: []IL8R0017ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0017O",
				},
				FormData: il8r0017O,
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
func (o *IL8R0017ResponseForm) UnPackResponse(request []byte) (IL8R0017O, error) {

	il8r0017O := IL8R0017O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0017O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0017O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0017I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
