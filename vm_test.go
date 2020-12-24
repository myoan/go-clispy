package main

import (
	"testing"
)

func TestVirtualMachine_Push(t *testing.T) {
	vm := NewVirtualMachine()
	vm.Push(1)
	if vm.sp != 1 {
		t.Errorf("sp not incremented")
	}
	if vm.stack[0] != 1 {
		t.Errorf("pushed value not match")
	}
}

func TestVirtualMachine_Pop(t *testing.T) {
	vm := NewVirtualMachine()
	_, err := vm.Pop()
	if err == nil || err != ErrEmptyStack {
		t.Errorf("message not match: expect is 'Stack is empty', but actual is '%v'\n", err)
	}
	vm.Push(1)
	actual, err := vm.Pop()
	if err != nil {
		t.Errorf("'%v'\n", err)
	}
	if actual != 1 {
		t.Errorf("result not matched: expect is 1, but actual is %d\n", actual)
	}
}

func TestVM_Exec(t *testing.T) {
	testcase := []struct {
		in     *InstructionSet
		msg    string
		expect int
	}{
		{
			in: &InstructionSet{
				insts: []*Instruction{
					{
						iType:  InsPush,
						value1: 1,
					},
					{
						iType:  InsPush,
						value1: 2,
					},
					{
						iType: InsAdd,
					},
				},
				ft: nil,
			},
			msg:    "add",
			expect: 3,
		},
		{
			in: &InstructionSet{
				insts: []*Instruction{
					{
						iType:  InsPush,
						value1: 1,
					},
					{
						iType:  InsPush,
						value1: 2,
					},
					{
						iType: InsSub,
					},
					{
						iType:  InsPush,
						value1: 3,
					},
					{
						iType:  InsPush,
						value1: 4,
					},
					{
						iType: InsMul,
					},
					{
						iType: InsAdd,
					},
				},
				ft: nil,
			},
			msg:    "nested",
			expect: 11,
		},
		{
			in: &InstructionSet{
				insts: []*Instruction{
					{
						iType:  InsPush,
						value1: 1,
					},
					{
						iType:  InsPush,
						value1: 2,
					},
					{
						iType: InsLt,
					},
					{
						iType:  InsIf,
						value1: 5,
					},
					{
						iType:  InsPush,
						value1: 3,
					},
					{
						iType:  InsJump,
						value1: 8,
					},
					{
						iType:  InsPush,
						value1: 4,
					},
				},
				ft: nil,
			},
			msg:    "jump",
			expect: 3,
		},
		{
			in: &InstructionSet{
				insts: []*Instruction{
					{
						iType:  InsPush,
						value1: 1,
					},
					{
						iType:  InsCall,
						value1: 0,
					},
				},
				ft: &NFunctionTable{
					funcs: []*NFunction{
						{
							insts: []*Instruction{
								{
									iType:  InsPush,
									value1: 1,
								},
								{
									iType: InsAdd,
								},
							},
						},
					},
				},
			},
			msg:    "call function",
			expect: 2,
		},
	}
	for _, tt := range testcase {
		vm := NewVirtualMachine()
		vm.Exec(tt.in)
		if tt.expect != vm.Result() {
			t.Errorf("[Error: %s] expect: %d, but actual: %d\n", tt.msg, tt.expect, vm.Result())
		}
	}
}
