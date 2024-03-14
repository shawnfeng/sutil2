// Copyright 2014 The sutil Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sfloat

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SFloat struct {
	// 精度位数
	precNum  int
	scaleNum int

	// 浮点数精读定义，参考binance返回的字符串就是8位精度
	precision   float64
	scaleFactor float64
}

func NewSFloat() *SFloat {
	return NewSFloatBySet(8, 8)
}

func NewSFloatBySet(precNum int, scaleNum int) *SFloat {
	return &SFloat{
		precNum:     precNum,
		scaleNum:    scaleNum,
		precision:   1 / math.Pow(10, float64(precNum)),
		scaleFactor: math.Pow(10, float64(scaleNum)),
	}
}

func (m *SFloat) String() string {
	return fmt.Sprintf("precNum:%d scaleNum:%d precision:%.20f scaleFactor:%.20f",
		m.precNum, m.scaleNum, m.precision, m.scaleFactor,
	)
}

// Getter for PrecNum
func (m *SFloat) GetPrecNum() int {
	return m.precNum
}

// Getter for ScaleNum
func (m *SFloat) GetScaleNum() int {
	return m.scaleNum
}

// Getter for Precision
func (m *SFloat) GetPrecision() float64 {
	return m.precision
}

// Getter for ScaleFactor
func (m *SFloat) GetScaleFactor() float64 {
	return m.scaleFactor
}

func (m *SFloat) IsEqual(a, b float64) bool {
	return math.Abs(a-b) < m.precision
}

func (m *SFloat) IsMoreEqual(a, b float64) bool {
	return a-b > 0 || m.IsEqual(a, b)
}

func (m *SFloat) IsLessEqual(a, b float64) bool {
	return a-b < 0 || m.IsEqual(a, b)
}

func (m *SFloat) IsZero(a float64) bool {
	return math.Abs(a) < m.precision
}

func GetFloatPrecision2(f float64) int {
	decimals := FmtMaxDecimalFloat(f)
	return GetFloatPrecision(decimals)
}

func (m *SFloat) FmtFloatInPrec(v float64) string {
	return strconv.FormatFloat(v, 'f', m.precNum, 64)
}

func (m *SFloat) MsgFloat(f float64) string {
	f = RoundToDecimalPlace(f, m.precNum)
	fStr := m.FmtFloatInPrec(f)

	return TrimFloatStr(fStr)
}

func (m *SFloat) MsgFloat4(f float64) string {
	f = RoundToDecimalPlace(f, 4)
	fStr := m.FmtFloatInPrec(f)

	return TrimFloatStr(fStr)
}

func TrimFloatStr(s string) string {
	s0 := strings.TrimRight(s, "0")
	return strings.TrimRight(s0, ".")
}

func GetMinPrice(p0, p1 string) string {
	fp0, err := ParseFloat64(p0)
	if err != nil {
		return ""
	}

	fp1, err := ParseFloat64(p1)
	if err != nil {
		return ""
	}

	p := math.Min(fp0, fp1)
	return strconv.FormatFloat(p, 'f', 8, 64)
}

func (m *SFloat) IntScaleDefault(f float64) int {
	return int(math.Round(f * m.scaleFactor))
}

func (m *SFloat) FloatUnScaleDefault(i int) float64 {
	return float64(i) / m.scaleFactor
}

func (m *SFloat) Int64ScaleDefault(f float64) int64 {
	return int64(math.Round(f * m.scaleFactor))
}

func (m *SFloat) FloatUnScaleDefault64(i int64) float64 {
	return float64(i) / m.scaleFactor
}

func GetIntList(intBegin, intEnd int, intStep int) []int {

	if intBegin >= intEnd {
		return nil
	}

	if intStep <= 0 {
		return nil
	}

	intPriceList := make([]int, 0)
	for i := intBegin; i <= intEnd; i += intStep {
		intPriceList = append(intPriceList, i)
	}

	if intPriceList[len(intPriceList)-1] < intEnd {
		intPriceList = append(intPriceList, intEnd)
	}

	return intPriceList
}

func FmtMaxDecimalFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (m *SFloat) RoundToDefault(value float64) float64 {
	return RoundToDecimalPlace(value, m.precNum)
}

func RoundToDecimalPlace(value float64, decimalPlaces int) float64 {
	pow := math.Pow10(decimalPlaces)
	temp := value * pow
	truncated := math.Round(temp) / pow
	return truncated
}

func FloorFloatNDecimal(value float64, decimalPlaces int) float64 {
	pow := math.Pow10(decimalPlaces)
	temp := value * pow
	truncated := math.Floor(temp) / pow
	return truncated
}

func GetMaxPrice(p0, p1 string) string {
	fp0, err := ParseFloat64(p0)
	if err != nil {
		return ""
	}

	fp1, err := ParseFloat64(p1)
	if err != nil {
		return ""
	}

	p := math.Max(fp0, fp1)
	return strconv.FormatFloat(p, 'f', 8, 64)
}

func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseFloat64Must(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func GetFloatPrecision(f string) int {
	f = strings.TrimSpace(f)
	f = strings.Trim(f, "0")
	parts := strings.Split(f, ".")
	if len(parts) < 2 {
		return 0
	}
	return len(parts[1])
}
