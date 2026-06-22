func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    m, n := len(num1), len(num2)
    res := make([]int, m+n)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            digA := int(num1[i]-'0')
            digB := int(num2[j]-'0')
            res[i+j+1] += digA * digB

            if res[i+j+1] >= 10 {
                carry := res[i+j+1]/10
                res[i+j+1] = res[i+j+1]%10
                res[i+j] += carry
            }
        }
    }

    var out []rune
    for i, num := range res {
        if i == 0 && num == 0 {
            continue
        }

        out = append(out, rune(num+'0'))
    }

    return string(out)
}
