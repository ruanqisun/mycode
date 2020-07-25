package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRED2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRED2I
}

type DAILRED2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRED2O
}

type DAILRED2RequestForm struct {
	Form []DAILRED2IDataForm
}

type DAILRED2ResponseForm struct {
	Form []DAILRED2ODataForm
}

type DAILRED2I struct {
	ChgSn                        string  `json:"ChgSn"`
	CustNo                       string  `json:"CustNo"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	OperSorcCd                   string  `json:"OperSorcCd"`
	OperBefQtaStusCd             string  `json:"OperBefQtaStusCd"`
	OperAftQtaStusCd             string  `json:"OperAftQtaStusCd"`
	QtaOperTypCd                 string  `json:"QtaOperTypCd"`
	CurCd                        string  `json:"CurCd"`
	ChgBefAvalQta                float64 `json:"ChgBefAvalQta"`
	ChgAftQta                    float64 `json:"ChgAftQta"`
	ChgAftAvalQta                float64 `json:"ChgAftAvalQta"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED2O struct {
	Records                      []DAILRED2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED2ORecords struct {
	ChgSn                        string  `json:"ChgSn"`
	CustNo                       string  `json:"CustNo"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	OperSorcCd                   string  `json:"OperSorcCd"`
	OperBefQtaStusCd             string  `json:"OperBefQtaStusCd"`
	OperAftQtaStusCd             string  `json:"OperAftQtaStusCd"`
	QtaOperTypCd                 string  `json:"QtaOperTypCd"`
	CurCd                        string  `json:"CurCd"`
	ChgBefAvalQta                float64 `json:"ChgBefAvalQta"`
	ChgAftQta                    float64 `json:"ChgAftQta"`
	ChgAftAvalQta                float64 `json:"ChgAftAvalQta"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRED2RequestForm) PackRequest(dailred2I DAILRED2I) (responseBody []byte, err error) {

	requestForm := DAILRED2RequestForm{
		Form: []DAILRED2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED2I",
				},

				FormData: dailred2I,
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
func (o *DAILRED2RequestForm) UnPackRequest(request []byte) (DAILRED2I, error) {
	dailred2I := DAILRED2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailred2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRED2ResponseForm) PackResponse(dailred2O DAILRED2O) (responseBody []byte, err error) {
	responseForm := DAILRED2ResponseForm{
		Form: []DAILRED2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED2O",
				},
				FormData: dailred2O,
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
func (o *DAILRED2ResponseForm) UnPackResponse(request []byte) (DAILRED2O, error) {

	dailred2O := DAILRED2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailred2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRED2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
