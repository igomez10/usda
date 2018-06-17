package main

import "encoding/xml"

type report struct {
	XMLName xml.Name `xml:"report"`
	Items   []item   `xml:"item"`
}
type item struct {
	XMLName       xml.Name `xml:"item"`
	CommName      string   `xml:"commName"`
	CityName      string   `xml:"cityName"`
	PackageDesc   string   `xml:"packageDesc"`
	VarietyName   string   `xml:"varietyName"`
	SubvarName    string   `xml:"subvarName"`
	GradeDesc     string   `xml:"gradeDesc"`
	ReportDate    string   `xml:"reportDate"`
	LowPriceMin   float64  `xml:"lowPriceMin"`
	HighPriceMax  float64  `xml:"highPriceMax"`
	MostlyLowMin  float64  `xml:"mostlyLowMin"`
	MostlyHighMax float64  `xml:mostlyHighMax`
	OriginName    string   `xml:"originName"`
	DistrictName  string   `xml:"districtName"`
	ItemSizeDesc  string   `xml:"itemSizeDesc"`
	Color         string   `xml:"color"`
	EnvDesc       string   `xml:"envDesc"`
	UnitSale      string   `xml:"unitSale"`
	Quality       string   `xml:"quality"`
	Condition     string   `xml:"condition"`
	Appearance    string   `xml:"appearance"`
	Storage       string   `xml:"storage"`
	Crop          string   `xml:"crop"`
	RePack        string   `xml:"rePack"`
	Transmode     string   `xml:"transMode"`
	Offerings     string   `xml:"offerings"`
	MarketTone    string   `xml:"marketTone"`
	PriceComment  string   `xml:"priceComment"`
	Comment       string   `xml:"comment"`
}
