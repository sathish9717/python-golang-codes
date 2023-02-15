func twoSum(nums []int, target int) []int {
    result := []int{}
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                result = append(result, i)
                result = append(result, j)
                break
            }
        }
        if len(result) > 0 {
            break
        }
    }
    return result
func searchRange(nums []int, target int) []int {
    hashmap := make(map[int]int)
    var a []int
    for i, v := range nums {
        hashmap[v] = i
        val, ok := hashmap[target]
        if ok {
            a = append(a, val)
        }

    }
    if len(a) == 0 {
        return []int{-1, -1}
    } else {
        return []int{a[0], a[len(a)-1]}
    }
}

func searchRange2(nums []int, target int) []int {
    L := 0
    var a []int
    for L < len(nums) {
        if nums[L] == target {
            a = append(a, L)
            break
        }
        L++
    }
    R := len(a) - 1
    for R < len(nums) {
        if nums[R] == target {
            a = append(a, R)
            break
        }
        R--
    }
    return a

}

func main() {
    // nums = [5,7,7,8,8,10], target = 8
    a := []int{5, 7, 7, 8, 8, 10}
    t := 8
    searchRange2(a, t)
}
