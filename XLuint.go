package XLuint

import (
	"strconv"
)

type XLuint struct {
	numberPieces []uint32
}

func (XL *XLuint) SetHex(hexString string) {
	var uintPiece uint64
	XL.numberPieces = XL.numberPieces[:0]
	for i := len(hexString); i > 0; i -= 8 {

		if i >= 8 {
			uintPiece, _ = strconv.ParseUint(hexString[i-8:i], 16, 32)
		} else {
			uintPiece, _ = strconv.ParseUint(hexString[:i], 16, 32)
		}

		XL.numberPieces = append(XL.numberPieces, uint32(uintPiece))
	}
}

func (XL XLuint) GetHex() string {
	var hexString = strconv.FormatUint(uint64(XL.numberPieces[len(XL.numberPieces)-1]), 16)
	for i := len(XL.numberPieces) - 2; i >= 0; i-- {
		addString := strconv.FormatUint(uint64(XL.numberPieces[i]), 16)
		if len(addString) != 8 {
			for j := len(addString); j < 8; j++ {
				addString = "0" + addString
			}
		}
		hexString += addString
	}
	return hexString
}

func INV(number XLuint) XLuint {
	var inversedNumber XLuint
	for i := 0; i < len(number.numberPieces); i++ {
		inversedNumber.numberPieces = append(inversedNumber.numberPieces, ^number.numberPieces[i])
	}
	return inversedNumber
}

func XOR(number1, number2 XLuint) XLuint {
	var xoredNumber XLuint
	var weight uint8
	var iterCount = len(number1.numberPieces)
	var overPieces = number2.numberPieces[iterCount:]

	if len(number1.numberPieces) > len(number2.numberPieces) {
		weight = 1
		iterCount = len(number2.numberPieces)
		overPieces = number1.numberPieces[iterCount:]
	} else if len(number1.numberPieces) < len(number2.numberPieces) {
		weight = 2
	}

	for i := 0; i < iterCount; i++ {
		xoredNumber.numberPieces = append(xoredNumber.numberPieces, number1.numberPieces[i]^number2.numberPieces[i])
	}

	if weight == 0 {
		return xoredNumber
	}
	xoredNumber.numberPieces = append(xoredNumber.numberPieces, overPieces...)
	return xoredNumber
}

func OR(number1, number2 XLuint) XLuint {
	var oredNumber XLuint
	var weight uint8
	var iterCount = len(number1.numberPieces)
	var overPieces = number2.numberPieces[iterCount:]

	if len(number1.numberPieces) > len(number2.numberPieces) {
		weight = 1
		iterCount = len(number2.numberPieces)
		overPieces = number1.numberPieces[iterCount:]
	} else if len(number1.numberPieces) < len(number2.numberPieces) {
		weight = 2
	}

	for i := 0; i < iterCount; i++ {
		oredNumber.numberPieces = append(oredNumber.numberPieces, number1.numberPieces[i]|number2.numberPieces[i])
	}

	if weight == 0 {
		return oredNumber
	}
	oredNumber.numberPieces = append(oredNumber.numberPieces, overPieces...)
	return oredNumber
}

func AND(number1, number2 XLuint) XLuint {
	var andedNumber XLuint
	var iterCount = len(number1.numberPieces)

	if len(number1.numberPieces) > len(number2.numberPieces) {
		iterCount = len(number2.numberPieces)
	}

	for i := 0; i < iterCount; i++ {
		andedNumber.numberPieces = append(andedNumber.numberPieces, number1.numberPieces[i]&number2.numberPieces[i])
	}

	return andedNumber
}

func ShiftR(number XLuint, bits int) XLuint {
	//var cary uint8
	var empty = bits / 32
	var bits32 = bits % 32
	var shiftRnumber XLuint

	if empty >= len(number.numberPieces) {
		shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, 0)
		return shiftRnumber
	}

	number.numberPieces = number.numberPieces[empty:]
	shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, number.numberPieces[0]>>bits32)

	for i := 1; i < len(number.numberPieces); i++ {
		shiftRnumber.numberPieces[i-1] += ((0xffffffff >> (32 - bits32)) & number.numberPieces[i]) << (32 - bits32)
		shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, number.numberPieces[i]>>bits32)
	}

	return shiftRnumber
}

func ShiftL(number XLuint, bits int) XLuint {
	var empty = bits / 32
	var bits32 = bits % 32
	var shiftRnumber XLuint

	for i := 0; i < empty; i++ {
		shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, 0)
	}

	var cary uint32
	for i := 0; i < len(number.numberPieces); i++ {
		shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, (number.numberPieces[i]<<bits32)+cary)
		cary = (^uint32(0xffffffff>>bits32) & number.numberPieces[i]) >> (32 - bits32)
	}

	if number.numberPieces[len(number.numberPieces)-1] & ^uint32(0xffffffff>>bits32) != 0 {
		shiftRnumber.numberPieces = append(shiftRnumber.numberPieces, number.numberPieces[len(number.numberPieces)-1]>>(32-bits32))
	}

	return shiftRnumber
}

func ADD(number1, number2 XLuint) XLuint {
	var cary uint32
	var largestNumber = number1
	var smallerLength = len(number2.numberPieces)
	var resultNumber XLuint

	if len(number1.numberPieces) < len(number2.numberPieces) {
		largestNumber = number2
		smallerLength = len(number1.numberPieces)
	} else if len(number1.numberPieces) == len(number2.numberPieces) && number1.numberPieces[len(number1.numberPieces)-1] < number2.numberPieces[len(number2.numberPieces)-1] {
		largestNumber = number2
		smallerLength = len(number1.numberPieces)
	}

	for i := 0; i < smallerLength; i++ {
		resultNumber.numberPieces = append(resultNumber.numberPieces, number1.numberPieces[i]+number2.numberPieces[i]+cary)
		if number1.numberPieces[i]+number2.numberPieces[i]+cary < number1.numberPieces[i] {
			cary = 1
		} else {
			cary = 0
		}
	}

	for i := 0; i < len(largestNumber.numberPieces)-smallerLength; i++ {
		resultNumber.numberPieces = append(resultNumber.numberPieces, largestNumber.numberPieces[i+smallerLength]+cary)
		if largestNumber.numberPieces[i+smallerLength]+cary < largestNumber.numberPieces[i+smallerLength] {
			cary = 1
		} else {
			cary = 0
		}
	}

	if cary == 1 {
		resultNumber.numberPieces = append(resultNumber.numberPieces, 1)
	}

	return resultNumber
}

func SUB(number1, number2 XLuint) XLuint {
	var cary uint32
	var balance = len(number1.numberPieces) - len(number2.numberPieces)
	var smallerLength = len(number2.numberPieces)
	var resultNumber XLuint

	if len(number1.numberPieces) < len(number2.numberPieces) {
		smallerLength = len(number1.numberPieces)
	} else if len(number1.numberPieces) == len(number2.numberPieces) && number1.numberPieces[len(number1.numberPieces)-1] < number2.numberPieces[len(number2.numberPieces)-1] {
		smallerLength = len(number1.numberPieces)
	}

	for i := 0; i < smallerLength; i++ {
		resultNumber.numberPieces = append(resultNumber.numberPieces, number1.numberPieces[i]-number2.numberPieces[i]-cary)
		if number1.numberPieces[i]-number2.numberPieces[i]-cary > number1.numberPieces[i] {
			cary = 1
		} else {
			cary = 0
		}
	}

	if balance < 0 {
		for i := 0; i < len(number2.numberPieces)-len(number1.numberPieces); i++ {
			resultNumber.numberPieces = append(resultNumber.numberPieces, -number2.numberPieces[i+smallerLength]-cary)
			cary = 1
		}
	} else {
		for i := 0; i < len(number1.numberPieces)-len(number2.numberPieces); i++ {
			resultNumber.numberPieces = append(resultNumber.numberPieces, number1.numberPieces[i+smallerLength]-cary)
			if number1.numberPieces[i+smallerLength]-cary > number1.numberPieces[i+smallerLength] {
				cary = 1
			} else {
				cary = 0
			}
		}
	}

	for i := len(resultNumber.numberPieces) - 1; i >= 0; i-- {
		if resultNumber.numberPieces[i] == 0 && len(resultNumber.numberPieces) != 1 {
			resultNumber.numberPieces = resultNumber.numberPieces[:i]
		} else {
			break
		}
	}

	return resultNumber
}

func MOD(number, modulus XLuint) XLuint {
	var result XLuint
	_, result = DIV(number, modulus)
	return result
}

func MUL(number1, number2 XLuint) XLuint {
	var number1_h, number1_l, number2_h, number2_l XLuint
	if len(number1.numberPieces) < len(number2.numberPieces) {
		for i := 0; i < len(number2.numberPieces)-len(number1.numberPieces); i++ {
			number1.numberPieces = append(number1.numberPieces, 0)
		}
	} else if len(number1.numberPieces) > len(number2.numberPieces) {
		for i := 0; i < len(number1.numberPieces)-len(number2.numberPieces); i++ {
			number2.numberPieces = append(number2.numberPieces, 0)
		}
	}

	if (number1.numberPieces[0]&(0xffff<<16) != 0 || number2.numberPieces[0]&(0xffff<<16) != 0) || len(number1.numberPieces) > 1 {
		number1_h = ShiftR(number1, len(number1.numberPieces)*16)
		number1_l = SUB(number1, ShiftL(number1_h, len(number1.numberPieces)*16))

		number2_h = ShiftR(number2, len(number2.numberPieces)*16)
		number2_l = SUB(number2, ShiftL(number2_h, len(number2.numberPieces)*16))
		number12_h := MUL(number1_h, number2_h)
		number12_l := MUL(number1_l, number2_l)
		number12_hl := SUB(MUL(ADD(number1_h, number1_l), ADD(number2_h, number2_l)), ADD(number12_h, number12_l))
		simpleSum := ADD(ShiftL(number12_h, len(number1.numberPieces)*32), ShiftL(number12_hl, len(number1.numberPieces)*16))
		simpleSum = ADD(simpleSum, number12_l)

		for i := len(simpleSum.numberPieces) - 1; i >= 0; i-- {
			if simpleSum.numberPieces[i] == 0 && len(simpleSum.numberPieces) != 1 {
				simpleSum.numberPieces = simpleSum.numberPieces[:i]
			} else {
				break
			}
		}

		return simpleSum
	}
	simpleRes := XLuint{}
	simpleRes.numberPieces = append(simpleRes.numberPieces, number1.numberPieces[0]*number2.numberPieces[0])
	return simpleRes

}

func DIV(number1, number2 XLuint) (XLuint, XLuint) {
	var quotient, remainder, one XLuint
	quotient.SetHex("0")
	remainder.SetHex("0")
	one.SetHex("1")
	for i := len(number1.numberPieces)*32 - 1; i >= 0; i-- {
		remainder = ADD(ShiftL(remainder, 1), AND(ShiftR(number1, i), one))

		if cMPGE(remainder, number2) {
			remainder = SUB(remainder, number2)
			quotient = ADD(quotient, ShiftL(one, i))
		}

	}
	return quotient, remainder
}

func cMPGE(number1, number2 XLuint) bool {
	if len(number1.numberPieces) > len(number2.numberPieces) {
		return true
	} else if len(number1.numberPieces) < len(number2.numberPieces) {
		return false
	}
	for i := len(number1.numberPieces) - 1; i >= 0; i-- {
		if number1.numberPieces[i] < number2.numberPieces[i] {
			return false
		} else if number1.numberPieces[i] > number2.numberPieces[i] {
			return true
		}
	}
	return true
}
func POW(number, modulus XLuint, power uint) XLuint {
	number = MOD(number, modulus)
	result := number
	for i := uint(1); i < power; i++ {
		result = MOD(MUL(result, number), modulus)
	}
	return result
}
