package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW2I
}

type DAILRYW2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW2O
}

type DAILRYW2RequestForm struct {
	Form []DAILRYW2IDataForm
}

type DAILRYW2ResponseForm struct {
	Form []DAILRYW2ODataForm
}

type DAILRYW2I struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	StartDt                      string  `json:"StartDt"`
	MatrDt                       string  `json:"MatrDt"`
	LoanProdtClsfCd              string  `json:"LoanProdtClsfCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	QuotaType                    string  `json:"QuotaType"`
	QtaCtrlFlg                   string  `json:"QtaCtrlFlg"`
	LoanTypCd                    string  `json:"LoanTypCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	CurCd                        string  `json:"CurCd"`
	AplyLoanDeadl                int  `json:"AplyLoanDeadl"`
	AplyLoanAmtCeilVal           float64  `json:"AplyLoanAmtCeilVal"`
	BegintDayDtrmnManrCd         string  `json:"BegintDayDtrmnManrCd"`
	ExpdayDtrmnManrCd            string  `json:"ExpdayDtrmnManrCd"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	ProdtIntrtCd                 string  `json:"ProdtIntrtCd"`
	TranOvdueOperManrCd          string  `json:"TranOvdueOperManrCd"`
	IntrtAdjManrCd               string  `json:"IntrtAdjManrCd"`
	PmitAdvRepayFlg              string  `json:"PmitAdvRepayFlg"`
	SglacctHighCrdtQta           float64  `json:"SglacctHighCrdtQta"`
	NormlLoanChrgSeqCtrlBnch     string  `json:"NormlLoanChrgSeqCtrlBnch"`
	DvalLoanChrgSeqCtrlBnch      string  `json:"DvalLoanChrgSeqCtrlBnch"`
	RevneCmpdCd                  string  `json:"RevneCmpdCd"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	OvdueIntrtFlotRatio          float64  `json:"OvdueIntrtFlotRatio"`
	OvdueIntrtFlotManrCd         string  `json:"OvdueIntrtFlotManrCd"`
	OvdueIntrtFlotDrctCd         string  `json:"OvdueIntrtFlotDrctCd"`
	OvdueIntrtBpFlotVal          float64  `json:"OvdueIntrtBpFlotVal"`
	EmbIntrtFlotManrCd           string  `json:"EmbIntrtFlotManrCd"`
	EmbIntrtFlotDrctCd           string  `json:"EmbIntrtFlotDrctCd"`
	EmbIntrtBpFlotVal            float64  `json:"EmbIntrtBpFlotVal"`
	EmbLoanPnltintIntrtFlotRatio float64  `json:"EmbLoanPnltintIntrtFlotRatio"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr             string  `json:"FinlModifrEmpnbr"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRYW2O struct {
	Records                      []DAILRYW2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW2ORecords struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	StartDt                      string  `json:"StartDt"`
	MatrDt                       string  `json:"MatrDt"`
	LoanProdtClsfCd              string  `json:"LoanProdtClsfCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	QuotaType                    string  `json:"QuotaType"`
	QtaCtrlFlg                   string  `json:"QtaCtrlFlg"`
	LoanTypCd                    string  `json:"LoanTypCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	CurCd                        string  `json:"CurCd"`
	AplyLoanDeadl                int  `json:"AplyLoanDeadl"`
	AplyLoanAmtCeilVal           float64  `json:"AplyLoanAmtCeilVal"`
	BegintDayDtrmnManrCd         string  `json:"BegintDayDtrmnManrCd"`
	ExpdayDtrmnManrCd            string  `json:"ExpdayDtrmnManrCd"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	ProdtIntrtCd                 string  `json:"ProdtIntrtCd"`
	TranOvdueOperManrCd          string  `json:"TranOvdueOperManrCd"`
	IntrtAdjManrCd               string  `json:"IntrtAdjManrCd"`
	PmitAdvRepayFlg              string  `json:"PmitAdvRepayFlg"`
	SglacctHighCrdtQta           float64  `json:"SglacctHighCrdtQta"`
	NormlLoanChrgSeqCtrlBnch     string  `json:"NormlLoanChrgSeqCtrlBnch"`
	DvalLoanChrgSeqCtrlBnch      string  `json:"DvalLoanChrgSeqCtrlBnch"`
	RevneCmpdCd                  string  `json:"RevneCmpdCd"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	OvdueIntrtFlotRatio          float64  `json:"OvdueIntrtFlotRatio"`
	OvdueIntrtFlotManrCd         string  `json:"OvdueIntrtFlotManrCd"`
	OvdueIntrtFlotDrctCd         string  `json:"OvdueIntrtFlotDrctCd"`
	OvdueIntrtBpFlotVal          float64  `json:"OvdueIntrtBpFlotVal"`
	EmbIntrtFlotManrCd           string  `json:"EmbIntrtFlotManrCd"`
	EmbIntrtFlotDrctCd           string  `json:"EmbIntrtFlotDrctCd"`
	EmbIntrtBpFlotVal            float64  `json:"EmbIntrtBpFlotVal"`
	EmbLoanPnltintIntrtFlotRatio float64  `json:"EmbLoanPnltintIntrtFlotRatio"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr             string  `json:"FinlModifrEmpnbr"`
}

// @Desc Build request message
func (o *DAILRYW2RequestForm) PackRequest(dailryw2I DAILRYW2I) (responseBody []byte, err error) {

	requestForm := DAILRYW2RequestForm{
		Form: []DAILRYW2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW2I",
				},

				FormData: dailryw2I,
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
func (o *DAILRYW2RequestForm) UnPackRequest(request []byte) (DAILRYW2I, error) {
	dailryw2I := DAILRYW2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW2ResponseForm) PackResponse(dailryw2O DAILRYW2O) (responseBody []byte, err error) {
	responseForm := DAILRYW2ResponseForm{
		Form: []DAILRYW2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW2O",
				},
				FormData: dailryw2O,
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
func (o *DAILRYW2ResponseForm) UnPackResponse(request []byte) (DAILRYW2O, error) {

	dailryw2O := DAILRYW2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
