import "reflect"


func isAnagram(s string, t string) bool {

    if len(t) != len(s) {
        return false
    }

    s_hash := make(map[byte]int)
    t_hash := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        s_hash[s[i]] += 1
        t_hash[t[i]] += 1
    }

    return reflect.DeepEqual(s_hash, t_hash)

}
