package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW6I
}

type DAILRYW6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW6O
}

type DAILRYW6RequestForm struct {
	Form []DAILRYW6IDataForm
}

type DAILRYW6ResponseForm struct {
	Form []DAILRYW6ODataForm
}

type DAILRYW6I struct {
	OrgNo                         string  `json:"OrgNo"`
	SeqNo                         int  `json:"SeqNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtVersNo               string  `json:"LoanProdtVersNo"`
	LoanGuarManrCd                string  `json:"LoanGuarManrCd"`
	KeprcdStusCd                  string  `json:"KeprcdStusCd"`
	OrgNm                         string  `json:"OrgNm"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	QtaTypCd                      string  `json:"QtaTypCd"`
	LnrwSmlGroupPeoNo             int  `json:"LnrwSmlGroupPeoNo"`
	LonRvwPeoNo                   int  `json:"LonRvwPeoNo"`
	IdpnApprvrModeFloorAmt        float64  `json:"IdpnApprvrModeFloorAmt"`
	LnrwSmlGroupModeFloorAmt      float64  `json:"LnrwSmlGroupModeFloorAmt"`
	LonRvwModeFloorAmt            float64  `json:"LonRvwModeFloorAmt"`
	IdpnApprvrModeAmtGtValScpCd   string  `json:"IdpnApprvrModeAmtGtValScpCd"`
	LnrwSmlGroupModeAmtGtValScpCd string  `json:"LnrwSmlGroupModeAmtGtValScpCd"`
	LonRvwModeAmtGtValScpCd       string  `json:"LonRvwModeAmtGtValScpCd"`
	IdpnApprvrCeilAmt             float64  `json:"IdpnApprvrCeilAmt"`
	LnrwSmlGroupModeCeilAmt       float64  `json:"LnrwSmlGroupModeCeilAmt"`
	LonRvwModeCeilAmt             float64  `json:"LonRvwModeCeilAmt"`
	EfftDt                        string  `json:"EfftDt"`
	ExpiryDt                      string  `json:"ExpiryDt"`
	Remrk                         string  `json:"Remrk"`
	CrtDt                         string  `json:"CrtDt"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYW6O struct {
	Records                      []DAILRYW6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW6ORecords struct {
	OrgNo                         string  `json:"OrgNo"`
	SeqNo                         int  `json:"SeqNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtVersNo               string  `json:"LoanProdtVersNo"`
	LoanGuarManrCd                string  `json:"LoanGuarManrCd"`
	KeprcdStusCd                  string  `json:"KeprcdStusCd"`
	OrgNm                         string  `json:"OrgNm"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	QtaTypCd                      string  `json:"QtaTypCd"`
	LnrwSmlGroupPeoNo             int  `json:"LnrwSmlGroupPeoNo"`
	LonRvwPeoNo                   int  `json:"LonRvwPeoNo"`
	IdpnApprvrModeFloorAmt        float64  `json:"IdpnApprvrModeFloorAmt"`
	LnrwSmlGroupModeFloorAmt      float64  `json:"LnrwSmlGroupModeFloorAmt"`
	LonRvwModeFloorAmt            float64  `json:"LonRvwModeFloorAmt"`
	IdpnApprvrModeAmtGtValScpCd   string  `json:"IdpnApprvrModeAmtGtValScpCd"`
	LnrwSmlGroupModeAmtGtValScpCd string  `json:"LnrwSmlGroupModeAmtGtValScpCd"`
	LonRvwModeAmtGtValScpCd       string  `json:"LonRvwModeAmtGtValScpCd"`
	IdpnApprvrCeilAmt             float64  `json:"IdpnApprvrCeilAmt"`
	LnrwSmlGroupModeCeilAmt       float64  `json:"LnrwSmlGroupModeCeilAmt"`
	LonRvwModeCeilAmt             float64  `json:"LonRvwModeCeilAmt"`
	EfftDt                        string  `json:"EfftDt"`
	ExpiryDt                      string  `json:"ExpiryDt"`
	Remrk                         string  `json:"Remrk"`
	CrtDt                         string  `json:"CrtDt"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW6RequestForm) PackRequest(dailryw6I DAILRYW6I) (responseBody []byte, err error) {

	requestForm := DAILRYW6RequestForm{
		Form: []DAILRYW6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW6I",
				},

				FormData: dailryw6I,
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
func (o *DAILRYW6RequestForm) UnPackRequest(request []byte) (DAILRYW6I, error) {
	dailryw6I := DAILRYW6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW6ResponseForm) PackResponse(dailryw6O DAILRYW6O) (responseBody []byte, err error) {
	responseForm := DAILRYW6ResponseForm{
		Form: []DAILRYW6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW6O",
				},
				FormData: dailryw6O,
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
func (o *DAILRYW6ResponseForm) UnPackResponse(request []byte) (DAILRYW6O, error) {

	dailryw6O := DAILRYW6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
