package main

import (
	"fmt"
	"testing"
)

func IsSameInstruction(actual, expect *Instruction) bool {
	if actual.iType != expect.iType {
		return false
	}
	if actual.value1 != expect.value1 {
		return false
	}
	if actual.value2 != expect.value2 {
		return false
	}
	return true
}

func IsSameInstructionSet(actual, expect *InstructionSet) bool {
	if actual == nil {
		return false
	}
	if actual.Size() != expect.Size() {
		return false
	}
	for i, ains := range actual.insts {
		eins := expect.insts[i]
		if !IsSameInstruction(ains, eins) {
			return false
		}
	}
	return true
}

func ShowInstructionSet(is *InstructionSet) {
	for _, i := range is.insts {
		fmt.Printf("[%d] (%d, %d)\n", i.iType, i.value1, i.value2)
	}
}

func IsSameStringArray(actual, expect []string) bool {
	return true
}

func IsSameFunctionTable(actual, expect *FunctionTable) bool {
	if len(actual.tbl) != len(actual.funcs) {
		return false
	}
	if len(actual.tbl) != len(expect.tbl) {
		return false
	}
	if len(actual.funcs) != len(expect.funcs) {
		return false
	}

	for i, fname := range actual.tbl {
		if fname != expect.tbl[i] {
			return false
		}
	}
	return true
}

func TestFunctionTable_Marge(t *testing.T) {
	testcase := []struct {
		msg    string
		src    *FunctionTable
		dst    *FunctionTable
		expect *FunctionTable
	}{
		{
			msg:    "empty",
			src:    &FunctionTable{},
			dst:    &FunctionTable{},
			expect: &FunctionTable{},
		},
		{
			msg: "no change",
			src: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			dst: &FunctionTable{
				tbl:   []string{},
				funcs: []*Function{},
			},
			expect: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
		},
		{
			msg: "simple add",
			src: &FunctionTable{
				funcs: []*Function{},
				tbl:   []string{},
			},
			dst: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			expect: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
		},
		{
			msg: "add table",
			src: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			dst: &FunctionTable{
				tbl:   []string{"fuga"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			expect: &FunctionTable{
				tbl:   []string{"hoge", "fuga"},
				funcs: []*Function{{insts: []*Instruction{}}, {insts: []*Instruction{}}},
			},
		},
		{
			msg: "store only uniq",
			src: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			dst: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
			expect: &FunctionTable{
				tbl:   []string{"hoge"},
				funcs: []*Function{{insts: []*Instruction{}}},
			},
		},
	}

	for _, tt := range testcase {
		err := tt.src.Merge(tt.dst)
		if err != nil {
			t.Errorf("[Error] actual is nil\n")
		} else if !IsSameFunctionTable(tt.src, tt.expect) {
			t.Errorf("[Error] %s\n", tt.msg)
		}
	}
}

func TestCompile(t *testing.T) {
	testcase := []struct {
		in     *Node
		msg    string
		expect *InstructionSet
	}{
		{
			in: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "param",
						children: []*Node{
							{
								nodeType: Add,
								vari:     "+",
								children: []*Node{
									{
										nodeType: Num,
										value:    1,
										children: []*Node{},
									},
									{
										nodeType: Num,
										value:    2,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
			msg: "add",
			expect: &InstructionSet{
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
		},
		{
			in: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "param",
						children: []*Node{
							{
								nodeType: Add,
								vari:     "+",
								children: []*Node{
									{
										nodeType: Non,
										vari:     "param",
										children: []*Node{
											{
												nodeType: Sub,
												vari:     "-",
												children: []*Node{
													{
														nodeType: Num,
														value:    1,
														children: []*Node{},
													},
													{
														nodeType: Num,
														value:    2,
														children: []*Node{},
													},
												},
											},
										},
									},
									{
										nodeType: Non,
										vari:     "param",
										children: []*Node{
											{
												nodeType: Mul,
												vari:     "*",
												children: []*Node{
													{
														nodeType: Num,
														value:    3,
														children: []*Node{},
													},
													{
														nodeType: Num,
														value:    4,
														children: []*Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			msg: "nested",
			expect: &InstructionSet{
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
		},
		{
			in: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: If,
								vari:     "if",
								children: []*Node{
									{
										nodeType: Non,
										vari:     "paren",
										children: []*Node{
											{
												nodeType: Lt,
												vari:     "<",
												children: []*Node{
													{
														nodeType: Num,
														value:    1,
														children: []*Node{},
													},
													{
														nodeType: Num,
														value:    2,
														children: []*Node{},
													},
												},
											},
										},
									},
									{
										nodeType: Num,
										value:    3,
										children: []*Node{},
									},
									{
										nodeType: Num,
										value:    4,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
			msg: "jump",
			expect: &InstructionSet{
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
		},
		/*
			{
				in: &Node{
					nodeType: Non,
					children: []*Node{
						{
							nodeType: Non,
							vari:     "paren",
							children: []*Node{
								{
									nodeType: Defun,
									vari:     "DEFUN",
									children: []*Node{
										{
											nodeType: Var,
											vari:     "incr",
											children: []*Node{},
										},
										{
											nodeType: Args,
											vari:     "n",
											children: []*Node{},
										},
										{
											nodeType: Non,
											vari:     "paren",
											children: []*Node{
												{
													nodeType: Add,
													vari:     "+",
													children: []*Node{
														{
															nodeType: Var,
															vari:     "n",
															children: []*Node{},
														},
														{
															nodeType: Num,
															vari:     "",
															value:    1,
															children: []*Node{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						{
							nodeType: Non,
							vari:     "paren",
							children: []*Node{
								{
									nodeType: Var,
									vari:     "incr",
									children: []*Node{
										{
											nodeType: Num,
											vari:     "",
											value:    1,
											children: []*Node{},
										},
									},
								},
							},
						},
					},
				},
				msg: "defun",
				expect: &InstructionSet{
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
					ft: nil,
				},
			},
		*/
	}
	for _, tt := range testcase {
		actual := Compile(tt.in)
		if actual == nil {
			t.Errorf("[Error] actual is nil\n")
		} else if !IsSameInstructionSet(actual, tt.expect) {
			t.Errorf("[Error] %s\n", tt.msg)
			fmt.Println("-- actual")
			ShowInstructionSet(actual)
			fmt.Println("-- expect")
			ShowInstructionSet(tt.expect)
		}
	}
}
