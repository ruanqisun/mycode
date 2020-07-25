package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0034IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0034I
}

type IL8R0034ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0034O
}

type IL8R0034RequestForm struct {
	Form []IL8R0034IDataForm
}

type IL8R0034ResponseForm struct {
	Form []IL8R0034ODataForm
}

type IL8R0034I struct {
	CtrtNo                         string  `json:"CtrtNo"`
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	CustNo                         string  `json:"CustNo"`
	CustName                       string  `json:"CustName"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	ExecIntrt                      float64  `json:"ExecIntrt"`
	IntrtFlotManrCd                string  `json:"IntrtFlotManrCd"`
	BpFlotVal                      float64  `json:"BpFlotVal"`
	IntrtFlotRatio                 float64  `json:"IntrtFlotRatio"`
	CtrtTypCd                      string  `json:"CtrtTypCd"`
	LoanProdtNm                    string  `json:"LoanProdtNm"`
	LoanGuarManrCd                 string  `json:"LoanGuarManrCd"`
	CurCd                          string  `json:"CurCd"`
	AprvAmt                        float64  `json:"AprvAmt"`
	CtrtAmt                        float64  `json:"CtrtAmt"`
	AprvDeadlMons                  int  `json:"AprvDeadlMons"`
	CtrtDeadlMons                  int  `json:"CtrtDeadlMons"`
	LoanUsageCd                    string  `json:"LoanUsageCd"`
	BrwmnyUsageComnt               string  `json:"BrwmnyUsageComnt"`
	SpmeCntLoanRiskClsfCd          string  `json:"SpmeCntLoanRiskClsfCd"`
	SpmeCntLoanRiskClsfComnt       string  `json:"SpmeCntLoanRiskClsfComnt"`
	SpdmnyManrCd                   string  `json:"SpdmnyManrCd"`
	AutonPaymtLmt                  float64  `json:"AutonPaymtLmt"`
	PayWayCd                       string  `json:"PayWayCd"`
	AutonPaymtAcctNo               string  `json:"AutonPaymtAcctNo"`
	AutonPaymtAcctNm               string  `json:"AutonPaymtAcctNm"`
	PaymtRptsCycCd                 string  `json:"PaymtRptsCycCd"`
	PaymtRptsCycQty                int  `json:"PaymtRptsCycQty"`
	RepayManrCd                    string  `json:"RepayManrCd"`
	IntStlManrCd                   string  `json:"IntStlManrCd"`
	IntacrManrCd                   string  `json:"IntacrManrCd"`
	TranOvdueGraceDays             int  `json:"TranOvdueGraceDays"`
	ChrgManrCd                     string  `json:"ChrgManrCd"`
	EntrstChrgAcctNo               string  `json:"EntrstChrgAcctNo"`
	AcctNm                         string  `json:"AcctNm"`
	RepaySorcComnt                 string  `json:"RepaySorcComnt"`
	IntSubsidyRatio                float64  `json:"IntSubsidyRatio"`
	IntSubsidyExpiryDt             string  `json:"IntSubsidyExpiryDt"`
	BilsAddr                       string  `json:"BilsAddr"`
	PostCd                         string  `json:"PostCd"`
	EMail                          string  `json:"EMail"`
	CotactTelNo                    string  `json:"CotactTelNo"`
	MobileNo                       string  `json:"MobileNo"`
	DsptSltnManrCd                 string  `json:"DsptSltnManrCd"`
	ArbitCommNm                    string  `json:"ArbitCommNm"`
	ArbitCommAddr                  string  `json:"ArbitCommAddr"`
	MakelnPmseCondComnt            string  `json:"MakelnPmseCondComnt"`
	OthApntMtsComnt                string  `json:"OthApntMtsComnt"`
	KepacctOrgNo                   string  `json:"KepacctOrgNo"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0034O struct {
	Records                      []IL8R0034ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0034ORecords struct {
	CtrtNo                         string  `json:"CtrtNo"`
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	CustNo                         string  `json:"CustNo"`
	CustName                       string  `json:"CustName"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	ExecIntrt                      float64  `json:"ExecIntrt"`
	IntrtFlotManrCd                string  `json:"IntrtFlotManrCd"`
	BpFlotVal                      float64  `json:"BpFlotVal"`
	IntrtFlotRatio                 float64  `json:"IntrtFlotRatio"`
	CtrtTypCd                      string  `json:"CtrtTypCd"`
	LoanProdtNm                    string  `json:"LoanProdtNm"`
	LoanGuarManrCd                 string  `json:"LoanGuarManrCd"`
	CurCd                          string  `json:"CurCd"`
	AprvAmt                        float64  `json:"AprvAmt"`
	CtrtAmt                        float64  `json:"CtrtAmt"`
	AprvDeadlMons                  int  `json:"AprvDeadlMons"`
	CtrtDeadlMons                  int  `json:"CtrtDeadlMons"`
	LoanUsageCd                    string  `json:"LoanUsageCd"`
	BrwmnyUsageComnt               string  `json:"BrwmnyUsageComnt"`
	SpmeCntLoanRiskClsfCd          string  `json:"SpmeCntLoanRiskClsfCd"`
	SpmeCntLoanRiskClsfComnt       string  `json:"SpmeCntLoanRiskClsfComnt"`
	SpdmnyManrCd                   string  `json:"SpdmnyManrCd"`
	AutonPaymtLmt                  float64  `json:"AutonPaymtLmt"`
	PayWayCd                       string  `json:"PayWayCd"`
	AutonPaymtAcctNo               string  `json:"AutonPaymtAcctNo"`
	AutonPaymtAcctNm               string  `json:"AutonPaymtAcctNm"`
	PaymtRptsCycCd                 string  `json:"PaymtRptsCycCd"`
	PaymtRptsCycQty                int  `json:"PaymtRptsCycQty"`
	RepayManrCd                    string  `json:"RepayManrCd"`
	IntStlManrCd                   string  `json:"IntStlManrCd"`
	IntacrManrCd                   string  `json:"IntacrManrCd"`
	TranOvdueGraceDays             int  `json:"TranOvdueGraceDays"`
	ChrgManrCd                     string  `json:"ChrgManrCd"`
	EntrstChrgAcctNo               string  `json:"EntrstChrgAcctNo"`
	AcctNm                         string  `json:"AcctNm"`
	RepaySorcComnt                 string  `json:"RepaySorcComnt"`
	IntSubsidyRatio                float64  `json:"IntSubsidyRatio"`
	IntSubsidyExpiryDt             string  `json:"IntSubsidyExpiryDt"`
	BilsAddr                       string  `json:"BilsAddr"`
	PostCd                         string  `json:"PostCd"`
	EMail                          string  `json:"EMail"`
	CotactTelNo                    string  `json:"CotactTelNo"`
	MobileNo                       string  `json:"MobileNo"`
	DsptSltnManrCd                 string  `json:"DsptSltnManrCd"`
	ArbitCommNm                    string  `json:"ArbitCommNm"`
	ArbitCommAddr                  string  `json:"ArbitCommAddr"`
	MakelnPmseCondComnt            string  `json:"MakelnPmseCondComnt"`
	OthApntMtsComnt                string  `json:"OthApntMtsComnt"`
	KepacctOrgNo                   string  `json:"KepacctOrgNo"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0034RequestForm) PackRequest(il8r0034I IL8R0034I) (responseBody []byte, err error) {

	requestForm := IL8R0034RequestForm{
		Form: []IL8R0034IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0034I",
				},

				FormData: il8r0034I,
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
func (o *IL8R0034RequestForm) UnPackRequest(request []byte) (IL8R0034I, error) {
	il8r0034I := IL8R0034I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0034I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0034I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0034ResponseForm) PackResponse(il8r0034O IL8R0034O) (responseBody []byte, err error) {
	responseForm := IL8R0034ResponseForm{
		Form: []IL8R0034ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0034O",
				},
				FormData: il8r0034O,
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
func (o *IL8R0034ResponseForm) UnPackResponse(request []byte) (IL8R0034O, error) {

	il8r0034O := IL8R0034O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0034O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0034O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0034I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
