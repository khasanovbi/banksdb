package paymentsystem

// Following constants represent known payment systems.
const (
	AmericanExpress         PaymentSystem = "American Express"
	Dankort                 PaymentSystem = "Dankort"
	DinersClubInternational PaymentSystem = "Diners Club International"
	Discover                PaymentSystem = "Discover"
	InstaPayment            PaymentSystem = "InstaPayment"
	InterPayment            PaymentSystem = "InterPayment"
	JCB                     PaymentSystem = "JCB"
	LankaPay                PaymentSystem = "LankaPay"
	Maestro                 PaymentSystem = "Maestro"
	MasterCard              PaymentSystem = "MasterCard"
	Mir                     PaymentSystem = "Mir"
	NPSPridnestrovie        PaymentSystem = "NPS Pridnestrovie"
	RuPay                   PaymentSystem = "RuPay"
	Troy                    PaymentSystem = "Troy"
	TUnion                  PaymentSystem = "T-Union"
	UATP                    PaymentSystem = "UATP"
	UnionPay                PaymentSystem = "UnionPay"
	Verve                   PaymentSystem = "Verve"
	Visa                    PaymentSystem = "Visa"
)

// https://en.wikipedia.org/wiki/Payment_card_number#Issuer_identification_number_(IIN)
// https://www.discoverglobalnetwork.com/downloads/IPP_VAR_Compliance.pdf
// https://www.barclaycard.co.uk/business/files/BIN-Rules-UK.pdf
//nolint: gomnd
var rawPaymentSystems = map[PaymentSystem][]paymentSystemInfo{
	AmericanExpress: {{prefixes: []int{34, 37}, lengthChecker: &exactLengthChecker{Exact: 15}}},
	Dankort:         {{prefixes: []int{5019, 4571}, lengthChecker: &exactLengthChecker{Exact: 16}}},
	DinersClubInternational: {
		{prefixes: []int{36}, lengthChecker: &rangeLengthChecker{From: 14, To: 19}},
		{
			prefixes:      []int{3095},
			prefixRanges:  []prefixRange{{from: 300, to: 305}, {from: 38, to: 39}},
			lengthChecker: &rangeLengthChecker{From: 16, To: 19},
		},
	},
	Discover: {
		{
			prefixes: []int{601174},
			prefixRanges: []prefixRange{
				{from: 601100, to: 601103},
				{from: 601105, to: 601109},
				{from: 60112, to: 60114},
				{from: 601177, to: 601179},
				{from: 601186, to: 601199},
				{from: 6440, to: 6505},
				{from: 650601, to: 650609},
				{from: 650611, to: 659999},
			},
			lengthChecker: &rangeLengthChecker{From: 16, To: 19},
		},
	},
	InterPayment: {{prefixes: []int{636}, lengthChecker: &rangeLengthChecker{From: 16, To: 19}}},
	InstaPayment: {{prefixRanges: []prefixRange{{from: 637, to: 639}}, lengthChecker: &exactLengthChecker{Exact: 16}}},
	JCB: {
		{
			prefixRanges: []prefixRange{
				{from: 3088, to: 3094},
				{from: 3096, to: 3102},
				{from: 3112, to: 3120},
				{from: 3158, to: 3159},
				{from: 3337, to: 3349},
				{from: 3528, to: 3589},
			},
			lengthChecker: &rangeLengthChecker{From: 16, To: 19},
		},
	},
	LankaPay: {{prefixes: []int{357111}, lengthChecker: &exactLengthChecker{Exact: 16}}},
	Maestro: {
		{
			prefixes:      []int{50},
			prefixRanges:  []prefixRange{{from: 56, to: 69}},
			lengthChecker: &rangeLengthChecker{From: 12, To: 19},
		},
	},
	MasterCard: {
		{
			prefixRanges:  []prefixRange{{from: 2221, to: 2720}, {from: 51, to: 55}},
			lengthChecker: &exactLengthChecker{Exact: 16},
		},
	},
	Mir: {{prefixRanges: []prefixRange{{from: 2200, to: 2204}}, lengthChecker: &exactLengthChecker{Exact: 16}}},
	NPSPridnestrovie: {
		{prefixRanges: []prefixRange{{from: 6054740, to: 6054744}}, lengthChecker: &exactLengthChecker{Exact: 16}},
	},
	RuPay: {
		{
			prefixes:      []int{6521, 6522}, // NOTE: Remove 60, to prefer maestro cards
			lengthChecker: &exactLengthChecker{Exact: 16},
		},
	},
	Troy:   {{prefixRanges: []prefixRange{{from: 979200, to: 979289}}, lengthChecker: &exactLengthChecker{Exact: 16}}},
	TUnion: {{prefixes: []int{31}, lengthChecker: &exactLengthChecker{Exact: 19}}},
	UATP:   {{prefixes: []int{1}, lengthChecker: &exactLengthChecker{Exact: 15}}},
	UnionPay: {
		{
			prefixes: []int{810},
			prefixRanges: []prefixRange{
				{from: 622126, to: 622925},
				{from: 624, to: 626},
				{from: 6282, to: 6288},
				{from: 8110, to: 8171},
			},
			lengthChecker: &rangeLengthChecker{From: 16, To: 19},
		},
	},
	Verve: {
		{
			prefixRanges:  []prefixRange{{from: 506099, to: 506198}, {from: 650002, to: 650027}},
			lengthChecker: &oneOfLengthChecker{16, 19},
		},
	},
	Visa: {{prefixes: []int{4}, lengthChecker: &exactLengthChecker{Exact: 16}}},
}
