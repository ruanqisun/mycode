package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH7IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH7I
}

type DAILRDH7ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH7O
}

type DAILRDH7RequestForm struct {
	Form []DAILRDH7IDataForm
}

type DAILRDH7ResponseForm struct {
	Form []DAILRDH7ODataForm
}

type DAILRDH7I struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	SeqNo                        int     `json:"SeqNo"`
	WrtoffBizAplyNo              string  `json:"WrtoffBizAplyNo"`
	PrinWrtoffTypCd              string  `json:"PrinWrtoffTypCd"`
	IntrWrtoffTypCd              string  `json:"IntrWrtoffTypCd"`
	ApprvStusCd                  string  `json:"ApprvStusCd"`
	AplyDt                       string  `json:"AplyDt"`
	ApyfOrgNo                    string  `json:"ApyfOrgNo"`
	ApyfCustMgrEmpnbr            string  `json:"ApyfCustMgrEmpnbr"`
	DubilNo                      string  `json:"DubilNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustNm                       string  `json:"CustNm"`
	OprOrgNo                     string  `json:"OprOrgNo"`
	OprCustMgrEmpnbr             string  `json:"OprCustMgrEmpnbr"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	CurCd                        string  `json:"CurCd"`
	LoanAmt                      float64 `json:"LoanAmt"`
	LoanBal                      float64  `json:"LoanBal"`
	LoanIntrt                    float64  `json:"LoanIntrt"`
	LoanDeadl                    int     `json:"LoanDeadl"`
	GrantDt                      string  `json:"GrantDt"`
	MatrDt                       string  `json:"MatrDt"`
	DuePrin                      float64  `json:"DuePrin"`
	DuePrinBgnDt                 string  `json:"DuePrinBgnDt"`
	OvdueIntr                    float64  `json:"OvdueIntr"`
	IntrOvdueBgnDt               string  `json:"IntrOvdueBgnDt"`
	OvduePnltint                 float64  `json:"OvduePnltint"`
	PnltintBgnDt                 string  `json:"PnltintBgnDt"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	WrtoffPrin                   float64  `json:"WrtoffPrin"`
	WrtoffReasCd                 string  `json:"WrtoffReasCd"`
	WrtoffReasComnt              string  `json:"WrtoffReasComnt"`
	BizDescr                     string  `json:"BizDescr"`
	ImgFileId                    string  `json:"ImgFileId"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH7O struct {
	Records                      []DAILRDH7ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH7ORecords struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	SeqNo                        int     `json:"SeqNo"`
	WrtoffBizAplyNo              string  `json:"WrtoffBizAplyNo"`
	PrinWrtoffTypCd              string  `json:"PrinWrtoffTypCd"`
	IntrWrtoffTypCd              string  `json:"IntrWrtoffTypCd"`
	ApprvStusCd                  string  `json:"ApprvStusCd"`
	AplyDt                       string  `json:"AplyDt"`
	ApyfOrgNo                    string  `json:"ApyfOrgNo"`
	ApyfCustMgrEmpnbr            string  `json:"ApyfCustMgrEmpnbr"`
	DubilNo                      string  `json:"DubilNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustNm                       string  `json:"CustNm"`
	OprOrgNo                     string  `json:"OprOrgNo"`
	OprCustMgrEmpnbr             string  `json:"OprCustMgrEmpnbr"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	CurCd                        string  `json:"CurCd"`
	LoanAmt                      float64 `json:"LoanAmt"`
	LoanBal                      float64  `json:"LoanBal"`
	LoanIntrt                    float64  `json:"LoanIntrt"`
	LoanDeadl                    int     `json:"LoanDeadl"`
	GrantDt                      string  `json:"GrantDt"`
	MatrDt                       string  `json:"MatrDt"`
	DuePrin                      float64  `json:"DuePrin"`
	DuePrinBgnDt                 string  `json:"DuePrinBgnDt"`
	OvdueIntr                    float64  `json:"OvdueIntr"`
	IntrOvdueBgnDt               string  `json:"IntrOvdueBgnDt"`
	OvduePnltint                 float64  `json:"OvduePnltint"`
	PnltintBgnDt                 string  `json:"PnltintBgnDt"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	WrtoffPrin                   float64  `json:"WrtoffPrin"`
	WrtoffReasCd                 string  `json:"WrtoffReasCd"`
	WrtoffReasComnt              string  `json:"WrtoffReasComnt"`
	BizDescr                     string  `json:"BizDescr"`
	ImgFileId                    string  `json:"ImgFileId"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH7RequestForm) PackRequest(dailrdh7I DAILRDH7I) (responseBody []byte, err error) {

	requestForm := DAILRDH7RequestForm{
		Form: []DAILRDH7IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH7I",
				},

				FormData: dailrdh7I,
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
func (o *DAILRDH7RequestForm) UnPackRequest(request []byte) (DAILRDH7I, error) {
	dailrdh7I := DAILRDH7I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh7I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh7I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH7ResponseForm) PackResponse(dailrdh7O DAILRDH7O) (responseBody []byte, err error) {
	responseForm := DAILRDH7ResponseForm{
		Form: []DAILRDH7ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH7O",
				},
				FormData: dailrdh7O,
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
func (o *DAILRDH7ResponseForm) UnPackResponse(request []byte) (DAILRDH7O, error) {

	dailrdh7O := DAILRDH7O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh7O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh7O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH7I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
