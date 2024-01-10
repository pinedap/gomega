package matchers

import (
	"testing"

	. "github.com/onsi/gomega"
)

///
/// actual: string
///
func TestBeTruthfullyWithFalseAndString(t *testing.T) {
	g := NewWithT(t)

	g.Expect("false").To(BeTruthfully(false))
	g.Expect("False").To(BeTruthfully(false))
	g.Expect("FALSE").To(BeTruthfully(false))
	g.Expect("f").To(BeTruthfully(false))
	g.Expect("F").To(BeTruthfully(false))
	g.Expect("0").To(BeTruthfully(false))

	g.Expect("1").NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match("FaLsE")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match("G")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyFalseWithString(t *testing.T) {
	g := NewWithT(t)

	g.Expect("false").To(BeTruthfullyFalse())
	g.Expect("False").To(BeTruthfullyFalse())
	g.Expect("FALSE").To(BeTruthfullyFalse())
	g.Expect("f").To(BeTruthfullyFalse())
	g.Expect("F").To(BeTruthfullyFalse())
	g.Expect("0").To(BeTruthfullyFalse())

	g.Expect("1").NotTo(BeTruthfullyFalse())

	truth, error := BeTruthfullyFalse().Match("FaLsE")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyFalse().Match("G")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}
func TestBeTruthfullyWithTrueAndString(t *testing.T) {
	g := NewWithT(t)

	g.Expect("true").To(BeTruthfully(true))
	g.Expect("True").To(BeTruthfully(true))
	g.Expect("TRUE").To(BeTruthfully(true))
	g.Expect("t").To(BeTruthfully(true))
	g.Expect("T").To(BeTruthfully(true))
	g.Expect("1").To(BeTruthfully(true))

	g.Expect("0").NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match("TrUe")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match("G")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match("-1")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match("2")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyTrueWithString(t *testing.T) {
	g := NewWithT(t)

	g.Expect("true").To(BeTruthfullyTrue())
	g.Expect("True").To(BeTruthfullyTrue())
	g.Expect("TRUE").To(BeTruthfullyTrue())
	g.Expect("t").To(BeTruthfullyTrue())
	g.Expect("T").To(BeTruthfullyTrue())
	g.Expect("1").To(BeTruthfullyTrue())

	g.Expect("0").NotTo(BeTruthfullyTrue())

	truth, error := BeTruthfullyTrue().Match("TrUe")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyTrue().Match("G")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyTrue().Match("-1")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyTrue().Match("2")
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

///
/// actual: int
///
func TestBeTruthfullyWithTrueAndInt8(t *testing.T) {
	g := NewWithT(t)

	var varInt int8 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndInt16(t *testing.T) {
	g := NewWithT(t)

	var varInt int16 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndInt32(t *testing.T) {
	g := NewWithT(t)

	var varInt int32 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndInt64(t *testing.T) {
	g := NewWithT(t)

	var varInt int64 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndInt(t *testing.T) {
	g := NewWithT(t)

	var varInt int = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyTrueWithInt(t *testing.T) {
	g := NewWithT(t)

	var varInt int = 1
	g.Expect(varInt).To(BeTruthfullyTrue())

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfullyTrue())

	truth, error := BeTruthfullyTrue().Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfullyTrue().Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyTrue().Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfullyTrue().Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndInt8(t *testing.T) {
	g := NewWithT(t)

	var varInt int8 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndInt16(t *testing.T) {
	g := NewWithT(t)

	var varInt int16 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndInt32(t *testing.T) {
	g := NewWithT(t)

	var varInt int32 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndInt64(t *testing.T) {
	g := NewWithT(t)

	var varInt int64 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndInt(t *testing.T) {
	g := NewWithT(t)

	var varInt int = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = -1
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

///
/// actual: bool
///
func TestBeTruthfullyWithFalseAndBool(t *testing.T) {
	g := NewWithT(t)

	g.Expect(false).Should(BeTruthfully(false))
	g.Expect(true).ShouldNot(BeTruthfully(false))
}

func TestBeTruthfullyFalseWithBool(t *testing.T) {
	g := NewWithT(t)

	g.Expect(false).Should(BeTruthfullyFalse())
	g.Expect(true).ShouldNot(BeTruthfullyFalse())
}

func TestBeTruthfullyWithTrueAndBool(t *testing.T) {
	g := NewWithT(t)

	g.Expect(true).Should(BeTruthfully(true))
	g.Expect(false).ShouldNot(BeTruthfully(true))
}

func TestBeTruthfullyTrueWithBool(t *testing.T) {
	g := NewWithT(t)

	g.Expect(true).Should(BeTruthfullyTrue())
	g.Expect(false).ShouldNot(BeTruthfullyTrue())
}

///
/// actual: uint
///
func TestBeTruthfullyWithTrueAndUInt8(t *testing.T) {
	g := NewWithT(t)

	var varInt uint8 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndUInt16(t *testing.T) {
	g := NewWithT(t)

	var varInt uint16 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndUInt32(t *testing.T) {
	g := NewWithT(t)

	var varInt uint32 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndUInt64(t *testing.T) {
	g := NewWithT(t)

	var varInt uint64 = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithTrueAndUInt(t *testing.T) {
	g := NewWithT(t)

	var varInt uint = 1
	g.Expect(varInt).To(BeTruthfully(true))

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfully(true))

	truth, error := BeTruthfully(true).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(true).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(true).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyTrueWithUInt(t *testing.T) {
	g := NewWithT(t)

	var varInt uint = 1
	g.Expect(varInt).To(BeTruthfullyTrue())

	varInt = 0
	g.Expect(varInt).NotTo(BeTruthfullyTrue())

	truth, error := BeTruthfullyTrue().Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfullyTrue().Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfullyTrue().Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndUInt8(t *testing.T) {
	g := NewWithT(t)

	var varInt uint8 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndUInt16(t *testing.T) {
	g := NewWithT(t)

	var varInt uint16 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndUInt32(t *testing.T) {
	g := NewWithT(t)

	var varInt uint32 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndUInt64(t *testing.T) {
	g := NewWithT(t)

	var varInt uint64 = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

func TestBeTruthfullyWithFalseAndUInt(t *testing.T) {
	g := NewWithT(t)

	var varInt uint = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

///
/// actual: byte
///
func TestBeTruthfullyWithFalseAndByte(t *testing.T) {
	g := NewWithT(t)

	var varInt byte = 0
	g.Expect(varInt).To(BeTruthfully(false))

	varInt = 1
	g.Expect(varInt).NotTo(BeTruthfully(false))

	truth, error := BeTruthfully(false).Match(-1)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	truth, error = BeTruthfully(false).Match(2)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())

	varInt = 2
	truth, error = BeTruthfully(false).Match(varInt)
	g.Expect(error).ShouldNot(BeNil())
	g.Expect(truth).Should(BeFalse())
}

