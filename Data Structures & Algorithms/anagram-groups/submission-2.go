import (
    "slices"
)

func sortString(w string) string {
    runes := []rune(w)
    slices.Sort(runes)
    return string(runes)
}

func groupAnagrams(strs []string) [][]string {
    groups := map[string][]string{}
    
    for _, w := range strs {
        key := sortString(w)
        if val, ok := groups[key]; ok {
            groups[key] = append(val, w)
        } else {
            groups[key] = []string{w}
        }
    }

    var out [][]string
    for _, vals := range groups {
        out = append(out, vals)
    }

    return out
}
