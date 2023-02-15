type Twoin struct {
	arr    []int
	target int
}
type testcases2 struct {
	cases string
	input Twoin
	want  []int
}

func TestTwoSum(t *testing.T) {

	a := []testcases2{{
		cases: "testcase1",
		input: Twoin{
			arr:    []int{1, 2, 3, 4},
			target: 10,
		},
	}}

	for _, v := range a {
		result := twoSum(v.input.arr, v.input.target)
		if !reflect.DeepEqual(result, v.want) {
			t.Error()
		}
	}
}

type searchRg struct {
	testcase string
	input    struct {
		arr    []int
		target int
	}
	want []int
}

func TestSearchRange(t *testing.T) {
	a := []searchRg{
		{
			testcase: "case 1",
			input: struct {
				arr    []int
				target int
			}{[]int{}, 0},
			want: []int{-1, -1},
		},
		{
			testcase: "case 2",
			input: struct {
				arr    []int
				target int
			}{[]int{5, 7, 7, 8, 8, 10}, 8},
			want: []int{3, 4},
		},
		{
			testcase: "case 3",
			input: struct {
				arr    []int
				target int
			}{[]int{5, 7, 7, 8, 8, 10}, 6},
			want: []int{-1, -1},
		},
	}
	for _, v := range a {
		result := searchRange(v.input.arr, v.input.target)
		if !reflect.DeepEqual(result, v.want) {
			t.Errorf("%v", v.testcase)
		}
	}
}
