package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX7IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX7I
}

type DAILRSX7ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX7O
}

type DAILRSX7RequestForm struct {
	Form []DAILRSX7IDataForm
}

type DAILRSX7ResponseForm struct {
	Form []DAILRSX7ODataForm
}

type DAILRSX7I struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	SggestQta                    float64 `json:"SggestQta"`
	DcmkTypCd                    string  `json:"DcmkTypCd"`
	DcmkDescr                    string  `json:"DcmkDescr"`
	TxResultStusCd               string  `json:"TxResultStusCd"`
	TxResultDescr                string  `json:"TxResultDescr"`
	CreditRatCd                  string  `json:"CreditRatCd"`
	NegtRsn                      string  `json:"NegtRsn"`
	Bp                           int     `json:"Bp"`
	AdjDrctCd                    string  `json:"AdjDrctCd"`
	HitExcludRuleComnt           string  `json:"HitExcludRuleComnt"`
	BlklistTypCd                 string  `json:"BlklistTypCd"`
	MonIncomeAmt                 float64 `json:"MonIncomeAmt"`
	ScoreCardGrpgInfo            string  `json:"ScoreCardGrpgInfo"`
	EvaltScr                     float64 `json:"EvaltScr"`
	RatCd                        string  `json:"RatCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX7O struct {
	Records                      []DAILRSX7ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX7ORecords struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	SggestQta                    float64 `json:"SggestQta"`
	DcmkTypCd                    string  `json:"DcmkTypCd"`
	DcmkDescr                    string  `json:"DcmkDescr"`
	TxResultStusCd               string  `json:"TxResultStusCd"`
	TxResultDescr                string  `json:"TxResultDescr"`
	CreditRatCd                  string  `json:"CreditRatCd"`
	NegtRsn                      string  `json:"NegtRsn"`
	Bp                           int     `json:"Bp"`
	AdjDrctCd                    string  `json:"AdjDrctCd"`
	HitExcludRuleComnt           string  `json:"HitExcludRuleComnt"`
	BlklistTypCd                 string  `json:"BlklistTypCd"`
	MonIncomeAmt                 float64 `json:"MonIncomeAmt"`
	ScoreCardGrpgInfo            string  `json:"ScoreCardGrpgInfo"`
	EvaltScr                     float64 `json:"EvaltScr"`
	RatCd                        string  `json:"RatCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX7RequestForm) PackRequest(dailrsx7I DAILRSX7I) (responseBody []byte, err error) {

	requestForm := DAILRSX7RequestForm{
		Form: []DAILRSX7IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX7I",
				},

				FormData: dailrsx7I,
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
func (o *DAILRSX7RequestForm) UnPackRequest(request []byte) (DAILRSX7I, error) {
	dailrsx7I := DAILRSX7I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx7I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx7I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX7ResponseForm) PackResponse(dailrsx7O DAILRSX7O) (responseBody []byte, err error) {
	responseForm := DAILRSX7ResponseForm{
		Form: []DAILRSX7ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX7O",
				},
				FormData: dailrsx7O,
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
func (o *DAILRSX7ResponseForm) UnPackResponse(request []byte) (DAILRSX7O, error) {

	dailrsx7O := DAILRSX7O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx7O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx7O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX7I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
