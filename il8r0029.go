package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0029IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0029I
}

type IL8R0029ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0029O
}

type IL8R0029RequestForm struct {
	Form []IL8R0029IDataForm
}

type IL8R0029ResponseForm struct {
	Form []IL8R0029ODataForm
}

type IL8R0029I struct {
	AplyBizNo                      string  `json:"AplyBizNo"`
	CustNo                         string  `json:"CustNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	LoanProdtVersNo                string  `json:"LoanProdtVersNo"`
	ExecIntrt                      float64  `json:"ExecIntrt"`
	SpcyBnchmkIntrt                float64  `json:"SpcyBnchmkIntrt"`
	BaseIntrtTypCd                 string  `json:"BaseIntrtTypCd"`
	IntrtTypCd                     string  `json:"IntrtTypCd"`
	IntrtNo                        string  `json:"IntrtNo"`
	IntrtFlotDrctCd                string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                      float64  `json:"BpFlotVal"`
	IntrtFlotRatio                 float64  `json:"IntrtFlotRatio"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	FilesTrmCd                     string  `json:"FilesTrmCd"`
	FixdIntrtPricingManrCd         string  `json:"FixdIntrtPricingManrCd"`
	IntrtFlotManrCd                string  `json:"IntrtFlotManrCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0029O struct {
	Records                      []IL8R0029ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0029ORecords struct {
	AplyBizNo                      string  `json:"AplyBizNo"`
	CustNo                         string  `json:"CustNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	LoanProdtVersNo                string  `json:"LoanProdtVersNo"`
	ExecIntrt                      float64  `json:"ExecIntrt"`
	SpcyBnchmkIntrt                float64  `json:"SpcyBnchmkIntrt"`
	BaseIntrtTypCd                 string  `json:"BaseIntrtTypCd"`
	IntrtTypCd                     string  `json:"IntrtTypCd"`
	IntrtNo                        string  `json:"IntrtNo"`
	IntrtFlotDrctCd                string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                      float64  `json:"BpFlotVal"`
	IntrtFlotRatio                 float64  `json:"IntrtFlotRatio"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	FilesTrmCd                     string  `json:"FilesTrmCd"`
	FixdIntrtPricingManrCd         string  `json:"FixdIntrtPricingManrCd"`
	IntrtFlotManrCd                string  `json:"IntrtFlotManrCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0029RequestForm) PackRequest(il8r0029I IL8R0029I) (responseBody []byte, err error) {

	requestForm := IL8R0029RequestForm{
		Form: []IL8R0029IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0029I",
				},

				FormData: il8r0029I,
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
func (o *IL8R0029RequestForm) UnPackRequest(request []byte) (IL8R0029I, error) {
	il8r0029I := IL8R0029I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0029I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0029I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0029ResponseForm) PackResponse(il8r0029O IL8R0029O) (responseBody []byte, err error) {
	responseForm := IL8R0029ResponseForm{
		Form: []IL8R0029ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0029O",
				},
				FormData: il8r0029O,
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
func (o *IL8R0029ResponseForm) UnPackResponse(request []byte) (IL8R0029O, error) {

	il8r0029O := IL8R0029O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0029O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0029O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0029I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
