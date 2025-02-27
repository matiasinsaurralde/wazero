package wazeroir

import "fmt"

type SignLessInt byte

const (
	SignLessInt32 SignLessInt = iota
	SignLessInt64
)

func (s SignLessInt) String() (ret string) {
	switch s {
	case SignLessInt32:
		ret = "i32"
	case SignLessInt64:
		ret = "i64"
	}
	return
}

type SignFulInt byte

const (
	SignFulInt32 SignFulInt = iota
	SignFulInt64
	SignFulUint32
	SignFulUint64
)

func (s SignFulInt) String() (ret string) {
	switch s {
	case SignFulUint32:
		ret = "i32"
	case SignFulUint64:
		ret = "i64"
	case SignFulInt32:
		ret = "u32"
	case SignFulInt64:
		ret = "u64"
	}
	return
}

type Float byte

const (
	Float32 Float = iota
	Float64
)

func (s Float) String() (ret string) {
	switch s {
	case Float32:
		ret = "f32"
	case Float64:
		ret = "f64"
	}
	return
}

type SignLessType byte

const (
	SignLessTypeI32 SignLessType = iota
	SignLessTypeI64
	SignLessTypeF32
	SignLessTypeF64
	SignLessTypeUnknown
)

func (s SignLessType) String() (ret string) {
	switch s {
	case SignLessTypeI32:
		ret = "i32"
	case SignLessTypeI64:
		ret = "i64"
	case SignLessTypeF32:
		ret = "f32"
	case SignLessTypeF64:
		ret = "f64"
	case SignLessTypeUnknown:
		ret = "unknown"
	}
	return
}

type SignFulType byte

const (
	SignFulTypeInt32 SignFulType = iota
	SignFulTypeUint32
	SignFulTypeInt64
	SignFulTypeUint64
	SignFulTypeFloat32
	SignFulTypeFloat64
)

func (s SignFulType) String() (ret string) {
	switch s {
	case SignFulTypeInt32:
		ret = "i32"
	case SignFulTypeUint32:
		ret = "u32"
	case SignFulTypeInt64:
		ret = "i64"
	case SignFulTypeUint64:
		ret = "u64"
	case SignFulTypeFloat32:
		ret = "f32"
	case SignFulTypeFloat64:
		ret = "f64"
	}
	return
}

type Operation interface {
	Kind() OperationKind
}

type OperationKind byte

func (o OperationKind) String() (ret string) {
	switch o {
	case OperationKindUnreachable:
		ret = "Unreachable"
	case OperationKindLabel:
		ret = "Label"
	case OperationKindBr:
		ret = "Br"
	case OperationKindBrIf:
		ret = "BrIf"
	case OperationKindBrTable:
		ret = "BrTable"
	case OperationKindCall:
		ret = "Call"
	case OperationKindCallIndirect:
		ret = "CallIndirect"
	case OperationKindDrop:
		ret = "Drop"
	case OperationKindSelect:
		ret = "Select"
	case OperationKindPick:
		ret = "Pick"
	case OperationKindSwap:
		ret = "Swap"
	case OperationKindGlobalGet:
		ret = "GlobalGet"
	case OperationKindGlobalSet:
		ret = "GlobalSet"
	case OperationKindLoad:
		ret = "Load"
	case OperationKindLoad8:
		ret = "Load8"
	case OperationKindLoad16:
		ret = "Load16"
	case OperationKindLoad32:
		ret = "Load32"
	case OperationKindStore:
		ret = "Store"
	case OperationKindStore8:
		ret = "Store8"
	case OperationKindStore16:
		ret = "Store16"
	case OperationKindStore32:
		ret = "Store32"
	case OperationKindMemorySize:
		ret = "MemorySize"
	case OperationKindMemoryGrow:
		ret = "MemoryGrow"
	case OperationKindConstI32:
		ret = "ConstI32"
	case OperationKindConstI64:
		ret = "ConstI64"
	case OperationKindConstF32:
		ret = "ConstF32"
	case OperationKindConstF64:
		ret = "ConstF64"
	case OperationKindEq:
		ret = "Eq"
	case OperationKindNe:
		ret = "Ne"
	case OperationKindEqz:
		ret = "Eqz"
	case OperationKindLt:
		ret = "Lt"
	case OperationKindGt:
		ret = "Gt"
	case OperationKindLe:
		ret = "Le"
	case OperationKindGe:
		ret = "Ge"
	case OperationKindAdd:
		ret = "Add"
	case OperationKindSub:
		ret = "Sub"
	case OperationKindMul:
		ret = "Mul"
	case OperationKindClz:
		ret = "Clz"
	case OperationKindCtz:
		ret = "Ctz"
	case OperationKindPopcnt:
		ret = "Popcnt"

	case OperationKindDiv:
		ret = "Div"
	case OperationKindRem:
		ret = "Rem"
	case OperationKindAnd:
		ret = "And"
	case OperationKindOr:
		ret = "Or"
	case OperationKindXor:
		ret = "Xor"
	case OperationKindShl:
		ret = "Shl"
	case OperationKindShr:
		ret = "Shr"
	case OperationKindRotl:
		ret = "Rotl"
	case OperationKindRotr:
		ret = "Rotr"
	case OperationKindAbs:
		ret = "Abs"
	case OperationKindNeg:
		ret = "Neg"
	case OperationKindCeil:
		ret = "Ceil"
	case OperationKindFloor:
		ret = "Floor"
	case OperationKindTrunc:
		ret = "Trunc"
	case OperationKindNearest:
		ret = "Nearest"
	case OperationKindSqrt:
		ret = "Sqrt"
	case OperationKindMin:
		ret = "Min"
	case OperationKindMax:
		ret = "Max"
	case OperationKindCopysign:
		ret = "Copysign"
	case OperationKindI32WrapFromI64:
		ret = "I32WrapFromI64"
	case OperationKindITruncFromF:
		ret = "ITruncFromF"
	case OperationKindFConvertFromI:
		ret = "FConvertFromI"
	case OperationKindF32DemoteFromF64:
		ret = "F32DemoteFromF64"
	case OperationKindF64PromoteFromF32:
		ret = "F64PromoteFromF32"
	case OperationKindI32ReinterpretFromF32:
		ret = "I32ReinterpretFromF32"
	case OperationKindI64ReinterpretFromF64:
		ret = "I64ReinterpretFromF64"
	case OperationKindF32ReinterpretFromI32:
		ret = "F32ReinterpretFromI32"
	case OperationKindF64ReinterpretFromI64:
		ret = "F64ReinterpretFromI64"
	case OperationKindExtend:
		ret = "Extend"
	}
	return
}

const (
	OperationKindUnreachable OperationKind = iota
	OperationKindLabel
	OperationKindBr
	OperationKindBrIf
	OperationKindBrTable
	OperationKindCall
	OperationKindCallIndirect
	OperationKindDrop
	OperationKindSelect
	OperationKindPick
	OperationKindSwap
	OperationKindGlobalGet
	OperationKindGlobalSet
	OperationKindLoad
	OperationKindLoad8
	OperationKindLoad16
	OperationKindLoad32
	OperationKindStore
	OperationKindStore8
	OperationKindStore16
	OperationKindStore32
	OperationKindMemorySize
	OperationKindMemoryGrow
	OperationKindConstI32
	OperationKindConstI64
	OperationKindConstF32
	OperationKindConstF64
	OperationKindEq
	OperationKindNe
	OperationKindEqz
	OperationKindLt
	OperationKindGt
	OperationKindLe
	OperationKindGe
	OperationKindAdd
	OperationKindSub
	OperationKindMul
	OperationKindClz
	OperationKindCtz
	OperationKindPopcnt
	OperationKindDiv
	OperationKindRem
	OperationKindAnd
	OperationKindOr
	OperationKindXor
	OperationKindShl
	OperationKindShr
	OperationKindRotl
	OperationKindRotr
	OperationKindAbs
	OperationKindNeg
	OperationKindCeil
	OperationKindFloor
	OperationKindTrunc
	OperationKindNearest
	OperationKindSqrt
	OperationKindMin
	OperationKindMax
	OperationKindCopysign
	OperationKindI32WrapFromI64
	OperationKindITruncFromF
	OperationKindFConvertFromI
	OperationKindF32DemoteFromF64
	OperationKindF64PromoteFromF32
	OperationKindI32ReinterpretFromF32
	OperationKindI64ReinterpretFromF64
	OperationKindF32ReinterpretFromI32
	OperationKindF64ReinterpretFromI64
	OperationKindExtend
)

type Label struct {
	FrameID uint32
	Kind    LabelKind
}

func (l *Label) String() (ret string) {
	switch l.Kind {
	case LabelKindHeader:
		ret = fmt.Sprintf(".L%d", l.FrameID)
	case LabelKindElse:
		ret = fmt.Sprintf(".L%d_else", l.FrameID)
	case LabelKindContinuation:
		ret = fmt.Sprintf(".L%d_cont", l.FrameID)
	}
	return
}

type LabelKind = byte

const (
	LabelKindHeader LabelKind = iota
	LabelKindElse
	LabelKindContinuation
)

func (l *Label) asBranchTarget() *BranchTarget {
	return &BranchTarget{Label: l}
}

func (l *Label) asBranchTargetDrop() *BranchTargetDrop {
	return &BranchTargetDrop{Target: l.asBranchTarget()}
}

type BranchTarget struct {
	Label *Label
}

func (b *BranchTarget) IsReturnTarget() bool {
	return b.Label == nil
}

func (b *BranchTarget) String() (ret string) {
	if b.IsReturnTarget() {
		ret = ".return"
	} else {
		ret = b.Label.String()
	}
	return
}

type BranchTargetDrop struct {
	Target *BranchTarget
	ToDrop *InclusiveRange
}

func (b *BranchTargetDrop) String() (ret string) {
	if b.ToDrop != nil {
		ret = fmt.Sprintf("%s(drop %d..%d)", b.Target, b.ToDrop.Start, b.ToDrop.End)
	} else {
		ret = b.Target.String()
	}
	return
}

type OperationUnreachable struct{}

func (o *OperationUnreachable) Kind() OperationKind {
	return OperationKindUnreachable
}

type OperationLabel struct {
	Label *Label
}

func (o *OperationLabel) Kind() OperationKind {
	return OperationKindLabel
}

type OperationBr struct {
	Target *BranchTarget
}

func (o *OperationBr) Kind() OperationKind {
	return OperationKindBr
}

type OperationBrIf struct {
	Then, Else *BranchTargetDrop
}

func (o *OperationBrIf) Kind() OperationKind {
	return OperationKindBrIf
}

type InclusiveRange struct {
	Start, End int
}

type OperationBrTable struct {
	Targets []*BranchTargetDrop
	Default *BranchTargetDrop
}

func (o *OperationBrTable) Kind() OperationKind {
	return OperationKindBrTable
}

type OperationCall struct {
	FunctionIndex uint32
}

func (o *OperationCall) Kind() OperationKind {
	return OperationKindCall
}

type OperationCallIndirect struct {
	TypeIndex, TableIndex uint32
}

func (o *OperationCallIndirect) Kind() OperationKind {
	return OperationKindCallIndirect
}

type OperationDrop struct{ Range *InclusiveRange }

func (o *OperationDrop) Kind() OperationKind {
	return OperationKindDrop
}

type OperationSelect struct{}

func (o *OperationSelect) Kind() OperationKind {
	return OperationKindSelect
}

type OperationPick struct{ Depth int }

func (o *OperationPick) Kind() OperationKind {
	return OperationKindPick
}

type OperationSwap struct{ Depth int }

func (o *OperationSwap) Kind() OperationKind {
	return OperationKindSwap
}

type OperationGlobalGet struct{ Index uint32 }

func (o *OperationGlobalGet) Kind() OperationKind {
	return OperationKindGlobalGet
}

type OperationGlobalSet struct{ Index uint32 }

func (o *OperationGlobalSet) Kind() OperationKind {
	return OperationKindGlobalSet
}

type MemoryImmediate struct {
	Alignment, Offest uint32
}

type OperationLoad struct {
	Type SignLessType
	Arg  *MemoryImmediate
}

func (o *OperationLoad) Kind() OperationKind {
	return OperationKindLoad
}

type OperationLoad8 struct {
	Type SignFulInt
	Arg  *MemoryImmediate
}

func (o *OperationLoad8) Kind() OperationKind {
	return OperationKindLoad8
}

type OperationLoad16 struct {
	Type SignFulInt
	Arg  *MemoryImmediate
}

func (o *OperationLoad16) Kind() OperationKind {
	return OperationKindLoad16
}

type OperationLoad32 struct {
	Signed bool
	Arg    *MemoryImmediate
}

func (o *OperationLoad32) Kind() OperationKind {
	return OperationKindLoad32
}

type OperationStore struct {
	Type SignLessType
	Arg  *MemoryImmediate
}

func (o *OperationStore) Kind() OperationKind {
	return OperationKindStore
}

type OperationStore8 struct {
	Type SignLessInt
	Arg  *MemoryImmediate
}

func (o *OperationStore8) Kind() OperationKind {
	return OperationKindStore8
}

type OperationStore16 struct {
	Type SignLessInt
	Arg  *MemoryImmediate
}

func (o *OperationStore16) Kind() OperationKind {
	return OperationKindStore16
}

type OperationStore32 struct {
	Arg *MemoryImmediate
}

func (o *OperationStore32) Kind() OperationKind {
	return OperationKindStore32
}

type OperationMemorySize struct{}

func (o *OperationMemorySize) Kind() OperationKind {
	return OperationKindMemorySize
}

type OperationMemoryGrow struct{ Aligment uint64 }

func (o *OperationMemoryGrow) Kind() OperationKind {
	return OperationKindMemoryGrow
}

type OperationConstI32 struct{ Value uint32 }

func (o *OperationConstI32) Kind() OperationKind {
	return OperationKindConstI32
}

type OperationConstI64 struct{ Value uint64 }

func (o *OperationConstI64) Kind() OperationKind {
	return OperationKindConstI64
}

type OperationConstF32 struct{ Value float32 }

func (o *OperationConstF32) Kind() OperationKind {
	return OperationKindConstF32
}

type OperationConstF64 struct{ Value float64 }

func (o *OperationConstF64) Kind() OperationKind {
	return OperationKindConstF64
}

type OperationEq struct{ Type SignLessType }

func (o *OperationEq) Kind() OperationKind {
	return OperationKindEq
}

type OperationNe struct{ Type SignLessType }

func (o *OperationNe) Kind() OperationKind {
	return OperationKindNe
}

type OperationEqz struct{ Type SignLessInt }

func (o *OperationEqz) Kind() OperationKind {
	return OperationKindEqz
}

type OperationLt struct{ Type SignFulType }

func (o *OperationLt) Kind() OperationKind {
	return OperationKindLt
}

type OperationGt struct{ Type SignFulType }

func (o *OperationGt) Kind() OperationKind {
	return OperationKindGt
}

type OperationLe struct{ Type SignFulType }

func (o *OperationLe) Kind() OperationKind {
	return OperationKindLe
}

type OperationGe struct{ Type SignFulType }

func (o *OperationGe) Kind() OperationKind {
	return OperationKindGe
}

type OperationAdd struct{ Type SignLessType }

func (o *OperationAdd) Kind() OperationKind {
	return OperationKindAdd
}

type OperationSub struct{ Type SignLessType }

func (o *OperationSub) Kind() OperationKind {
	return OperationKindSub
}

type OperationMul struct{ Type SignLessType }

func (o *OperationMul) Kind() OperationKind {
	return OperationKindMul
}

type OperationClz struct{ Type SignLessInt }

func (o *OperationClz) Kind() OperationKind {
	return OperationKindClz
}

type OperationCtz struct{ Type SignLessInt }

func (o *OperationCtz) Kind() OperationKind {
	return OperationKindCtz
}

type OperationPopcnt struct{ Type SignLessInt }

func (o *OperationPopcnt) Kind() OperationKind {
	return OperationKindPopcnt
}

type OperationDiv struct{ Type SignFulType }

func (o *OperationDiv) Kind() OperationKind {
	return OperationKindDiv
}

type OperationRem struct{ Type SignFulInt }

func (o *OperationRem) Kind() OperationKind {
	return OperationKindRem
}

type OperationAnd struct{ Type SignLessInt }

func (o *OperationAnd) Kind() OperationKind {
	return OperationKindAnd
}

type OperationOr struct{ Type SignLessInt }

func (o *OperationOr) Kind() OperationKind {
	return OperationKindOr
}

type OperationXor struct{ Type SignLessInt }

func (o *OperationXor) Kind() OperationKind {
	return OperationKindXor
}

type OperationShl struct{ Type SignLessInt }

func (o *OperationShl) Kind() OperationKind {
	return OperationKindShl
}

type OperationShr struct{ Type SignFulInt }

func (o *OperationShr) Kind() OperationKind {
	return OperationKindShr
}

type OperationRotl struct{ Type SignLessInt }

func (o *OperationRotl) Kind() OperationKind {
	return OperationKindRotl
}

type OperationRotr struct{ Type SignLessInt }

func (o *OperationRotr) Kind() OperationKind {
	return OperationKindRotr
}

type OperationAbs struct{ Type Float }

func (o *OperationAbs) Kind() OperationKind {
	return OperationKindAbs
}

type OperationNeg struct{ Type Float }

func (o *OperationNeg) Kind() OperationKind {
	return OperationKindNeg
}

type OperationCeil struct{ Type Float }

func (o *OperationCeil) Kind() OperationKind {
	return OperationKindCeil
}

type OperationFloor struct{ Type Float }

func (o *OperationFloor) Kind() OperationKind {
	return OperationKindFloor
}

type OperationTrunc struct{ Type Float }

func (o *OperationTrunc) Kind() OperationKind {
	return OperationKindTrunc
}

type OperationNearest struct{ Type Float }

func (o *OperationNearest) Kind() OperationKind {
	return OperationKindNearest
}

type OperationSqrt struct{ Type Float }

func (o *OperationSqrt) Kind() OperationKind {
	return OperationKindSqrt
}

type OperationMin struct{ Type Float }

func (o *OperationMin) Kind() OperationKind {
	return OperationKindMin
}

type OperationMax struct{ Type Float }

func (o *OperationMax) Kind() OperationKind {
	return OperationKindMax
}

type OperationCopysign struct{ Type Float }

func (o *OperationCopysign) Kind() OperationKind {
	return OperationKindCopysign
}

type OperationI32WrapFromI64 struct{}

func (o *OperationI32WrapFromI64) Kind() OperationKind {
	return OperationKindI32WrapFromI64
}

type OperationITruncFromF struct {
	InputType  Float
	OutputType SignFulInt
}

func (o *OperationITruncFromF) Kind() OperationKind {
	return OperationKindITruncFromF
}

type OperationFConvertFromI struct {
	InputType  SignFulInt
	OutputType Float
}

func (o *OperationFConvertFromI) Kind() OperationKind {
	return OperationKindFConvertFromI
}

type OperationF32DemoteFromF64 struct{}

func (o *OperationF32DemoteFromF64) Kind() OperationKind {
	return OperationKindF32DemoteFromF64
}

type OperationF64PromoteFromF32 struct{}

func (o *OperationF64PromoteFromF32) Kind() OperationKind {
	return OperationKindF64PromoteFromF32
}

type OperationI32ReinterpretFromF32 struct{}

func (o *OperationI32ReinterpretFromF32) Kind() OperationKind {
	return OperationKindI32ReinterpretFromF32
}

type OperationI64ReinterpretFromF64 struct{}

func (o *OperationI64ReinterpretFromF64) Kind() OperationKind {
	return OperationKindI64ReinterpretFromF64
}

type OperationF32ReinterpretFromI32 struct{}

func (o *OperationF32ReinterpretFromI32) Kind() OperationKind {
	return OperationKindF32ReinterpretFromI32
}

type OperationF64ReinterpretFromI64 struct{}

func (o *OperationF64ReinterpretFromI64) Kind() OperationKind {
	return OperationKindF64ReinterpretFromI64
}

type OperationExtend struct{ Signed bool }

func (o *OperationExtend) Kind() OperationKind {
	return OperationKindExtend
}
