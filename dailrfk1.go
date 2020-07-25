package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRFK1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRFK1I
}

type DAILRFK1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRFK1O
}

type DAILRFK1RequestForm struct {
	Form []DAILRFK1IDataForm
}

type DAILRFK1ResponseForm struct {
	Form []DAILRFK1ODataForm
}

type DAILRFK1I struct {
	MakelnAplySn    		     string  `json:"MakelnAplySn"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OthConsmTypCd                string  `json:"OthConsmTypCd"`
	UsageRplshComnt              string  `json:"UsageRplshComnt"`
	DpstAcctNo                   string  `json:"DpstAcctNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanAmt                      float64 `json:"LoanAmt"`
	LoanCurCd                    string  `json:"LoanCurCd"`
	Fee                          float64 `json:"Fee"`
	RprincPayIntCycCd            string  `json:"RprincPayIntCycCd"`
	LoanDeadl                    int     `json:"LoanDeadl"`
	LoanDeadlUnitCd              string  `json:"LoanDeadlUnitCd"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	MakelnAplyStusCd             string  `json:"MakelnAplyStusCd"`
	AplyMakelnDt                 string  `json:"AplyMakelnDt"`
	ActlMakelnTm                 string  `json:"ActlMakelnTm"`
	ExecIntrt                    float64 `json:"ExecIntrt"`
	RgtsintTypCd                 string  `json:"RgtsintTypCd"`
	RgtsintNo                    string  `json:"RgtsintNo"`
	LoanUsageCmtd                string  `json:"LoanUsageCmtd"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	BnchmkIntrt                  float64 `json:"BnchmkIntrt"`
	IntrtFlotDrctCd              string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                    float64 `json:"BpFlotVal"`
	IntrtFlotRatio               float64 `json:"IntrtFlotRatio"`
	LoanIntrtAdjCycQty           int     `json:"LoanIntrtAdjCycQty"`
	LoanIntrtAdjCycCd            string  `json:"LoanIntrtAdjCycCd"`
	FixdIntrtPricingManrCd       string  `json:"FixdIntrtPricingManrCd"`
	IntrtFlotManrCd              string  `json:"IntrtFlotManrCd"`
	RelaAcctOpnAcctBnkBnkNo      string  `json:"RelaAcctOpnAcctBnkBnkNo"`
	RelaAcctOpnAcctBnkBnkName    string  `json:"RelaAcctOpnAcctBnkBnkName"`
	RelaAcctBlngtoRwCategCd      string  `json:"RelaAcctBlngtoRwCategCd"`
	RelaAcctNoTypCd              string  `json:"RelaAcctNoTypCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	RsdnceProvScAdcmCd		     string  `json:"RsdnceProvScAdcmCd"`
	RsdnceCityScAdcmCd           string  `json:"RsdnceCityScAdcmCd"`
	RsdnceCntyScAdcmCd           string  `json:"RsdnceCntyScAdcmCd"`
	RsdnceTwnScAdcmCd            string  `json:"RsdnceTwnScAdcmCd"`
	RsdnceVillgrpScAdcmCd        string  `json:"RsdnceVillgrpScAdcmCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	QtaChgSn                     string  `json:"QtaChgSn"`
	LoanIntrtAdjManrCd           string  `json:"LoanIntrtAdjManrCd"`
	CareerTypCd       		     string  `json:"CareerTypCd"`
	EMail        			     string  `json:"EMail"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRFK1O struct {
	Records                      []DAILRFK1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRFK1ORecords struct {
	MakelnAplySn    		     string  `json:"MakelnAplySn"`
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OthConsmTypCd                string  `json:"OthConsmTypCd"`
	UsageRplshComnt              string  `json:"UsageRplshComnt"`
	DpstAcctNo                   string  `json:"DpstAcctNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanAmt                      float64 `json:"LoanAmt"`
	LoanCurCd                    string  `json:"LoanCurCd"`
	Fee                          float64 `json:"Fee"`
	RprincPayIntCycCd            string  `json:"RprincPayIntCycCd"`
	LoanDeadl                    int     `json:"LoanDeadl"`
	LoanDeadlUnitCd              string  `json:"LoanDeadlUnitCd"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	MakelnAplyStusCd             string  `json:"MakelnAplyStusCd"`
	AplyMakelnDt                 string  `json:"AplyMakelnDt"`
	ActlMakelnTm                 string  `json:"ActlMakelnTm"`
	ExecIntrt                    float64 `json:"ExecIntrt"`
	RgtsintTypCd                 string  `json:"RgtsintTypCd"`
	RgtsintNo                    string  `json:"RgtsintNo"`
	LoanUsageCmtd                string  `json:"LoanUsageCmtd"`
	BaseIntrtTypCd               string  `json:"BaseIntrtTypCd"`
	IntrtNo                      string  `json:"IntrtNo"`
	BnchmkIntrt                  float64 `json:"BnchmkIntrt"`
	IntrtFlotDrctCd              string  `json:"IntrtFlotDrctCd"`
	BpFlotVal                    float64 `json:"BpFlotVal"`
	IntrtFlotRatio               float64 `json:"IntrtFlotRatio"`
	LoanIntrtAdjCycQty           int     `json:"LoanIntrtAdjCycQty"`
	LoanIntrtAdjCycCd            string  `json:"LoanIntrtAdjCycCd"`
	FixdIntrtPricingManrCd       string  `json:"FixdIntrtPricingManrCd"`
	IntrtFlotManrCd              string  `json:"IntrtFlotManrCd"`
	RelaAcctOpnAcctBnkBnkNo      string  `json:"RelaAcctOpnAcctBnkBnkNo"`
	RelaAcctOpnAcctBnkBnkName    string  `json:"RelaAcctOpnAcctBnkBnkName"`
	RelaAcctBlngtoRwCategCd      string  `json:"RelaAcctBlngtoRwCategCd"`
	RelaAcctNoTypCd              string  `json:"RelaAcctNoTypCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	RsdnceProvScAdcmCd		     string  `json:"RsdnceProvScAdcmCd"`
	RsdnceCityScAdcmCd           string  `json:"RsdnceCityScAdcmCd"`
	RsdnceCntyScAdcmCd           string  `json:"RsdnceCntyScAdcmCd"`
	RsdnceTwnScAdcmCd            string  `json:"RsdnceTwnScAdcmCd"`
	RsdnceVillgrpScAdcmCd        string  `json:"RsdnceVillgrpScAdcmCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	QtaChgSn                     string  `json:"QtaChgSn"`
	LoanIntrtAdjManrCd           string  `json:"LoanIntrtAdjManrCd"`
	CareerTypCd       		     string  `json:"CareerTypCd"`
	EMail        			     string  `json:"EMail"`
}

// @Desc Build request message
func (o *DAILRFK1RequestForm) PackRequest(dailrfk1I DAILRFK1I) (responseBody []byte, err error) {

	requestForm := DAILRFK1RequestForm{
		Form: []DAILRFK1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRFK1I",
				},

				FormData: dailrfk1I,
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
func (o *DAILRFK1RequestForm) UnPackRequest(request []byte) (DAILRFK1I, error) {
	dailrfk1I := DAILRFK1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrfk1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrfk1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRFK1ResponseForm) PackResponse(dailrfk1O DAILRFK1O) (responseBody []byte, err error) {
	responseForm := DAILRFK1ResponseForm{
		Form: []DAILRFK1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRFK1O",
				},
				FormData: dailrfk1O,
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
func (o *DAILRFK1ResponseForm) UnPackResponse(request []byte) (DAILRFK1O, error) {

	dailrfk1O := DAILRFK1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrfk1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrfk1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRFK1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
