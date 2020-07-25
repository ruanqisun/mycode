package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDW1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDW1I
}

type DAILRDW1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDW1O
}

type DAILRDW1RequestForm struct {
	Form []DAILRDW1IDataForm
}

type DAILRDW1ResponseForm struct {
	Form []DAILRDW1ODataForm
}

type DAILRDW1I struct {
	DocMgmtNo                     string  `json:"DocMgmtNo"`
	MobileNo                      string  `json:"MobileNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	CustNo                        string  `json:"CustNo"`
	ClsfLablNm                    string  `json:"ClsfLablNm"`
	BizSn                         string  `json:"BizSn"`
	CustName                      string  `json:"CustName"`
	AtchTypCd                     string  `json:"AtchTypCd"`
	IndvCrtfTypCd                 string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                    string  `json:"IndvCrtfNo"`
	CtrtNo                        string  `json:"CtrtNo"`
	LoanDubilNo                   string  `json:"LoanDubilNo"`
	AchvsValidFlg                 string  `json:"AchvsValidFlg"`
	DocPrdusStageCd               string  `json:"DocPrdusStageCd"`
	DocTypNm                      string  `json:"DocTypNm"`
	DocInfoTypCd                  string  `json:"DocInfoTypCd"`
	DocNm                         string  `json:"DocNm"`
	DocNo                         string  `json:"DocNo"`
	DocPgNo                       string  `json:"DocPgNo"`
	DocPath                       string  `json:"DocPath"`
	DocCrtDt                      string  `json:"DocCrtDt"`
	DocCrtTm                      string  `json:"DocCrtTm"`
	DocDescr                      string  `json:"DocDescr"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDW1O struct {
	Records                      []DAILRDW1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDW1ORecords struct {
	DocMgmtNo                     string  `json:"DocMgmtNo"`
	MobileNo                      string  `json:"MobileNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	CustNo                        string  `json:"CustNo"`
	ClsfLablNm                    string  `json:"ClsfLablNm"`
	BizSn                         string  `json:"BizSn"`
	CustName                      string  `json:"CustName"`
	AtchTypCd                     string  `json:"AtchTypCd"`
	IndvCrtfTypCd                 string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                    string  `json:"IndvCrtfNo"`
	CtrtNo                        string  `json:"CtrtNo"`
	LoanDubilNo                   string  `json:"LoanDubilNo"`
	AchvsValidFlg                 string  `json:"AchvsValidFlg"`
	DocPrdusStageCd               string  `json:"DocPrdusStageCd"`
	DocTypNm                      string  `json:"DocTypNm"`
	DocInfoTypCd                  string  `json:"DocInfoTypCd"`
	DocNm                         string  `json:"DocNm"`
	DocNo                         string  `json:"DocNo"`
	DocPgNo                       string  `json:"DocPgNo"`
	DocPath                       string  `json:"DocPath"`
	DocCrtDt                      string  `json:"DocCrtDt"`
	DocCrtTm                      string  `json:"DocCrtTm"`
	DocDescr                      string  `json:"DocDescr"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDW1RequestForm) PackRequest(dailrdw1I DAILRDW1I) (responseBody []byte, err error) {

	requestForm := DAILRDW1RequestForm{
		Form: []DAILRDW1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDW1I",
				},

				FormData: dailrdw1I,
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
func (o *DAILRDW1RequestForm) UnPackRequest(request []byte) (DAILRDW1I, error) {
	dailrdw1I := DAILRDW1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdw1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdw1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDW1ResponseForm) PackResponse(dailrdw1O DAILRDW1O) (responseBody []byte, err error) {
	responseForm := DAILRDW1ResponseForm{
		Form: []DAILRDW1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDW1O",
				},
				FormData: dailrdw1O,
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
func (o *DAILRDW1ResponseForm) UnPackResponse(request []byte) (DAILRDW1O, error) {

	dailrdw1O := DAILRDW1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdw1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdw1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDW1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
