import "strconv"

func getHashedStr(str string) string {

    min_val := int(rune('a'))

    var counter [26]int
    for _, char := range str {
        char_value := int(char) - min_val
        counter[char_value] ++
    }

    hashedStr := ""
    for _, val := range counter {
        hashedStr += "." + strconv.Itoa(val)
    }

    return hashedStr
}


func groupAnagrams(strs []string) [][]string {
    groupedStrings := [][]string{}
    mapHashedStrings := make(map[string][]string)

    if len(strs) == 1 {
        return [][]string{strs}
    } 

    for _, str := range strs {
        hashedStr := getHashedStr(str)
        if _, exists := mapHashedStrings[hashedStr]; exists {
            mapHashedStrings[hashedStr] = append(mapHashedStrings[hashedStr], str)
        } else {
            mapHashedStrings[hashedStr] = []string{str}
        } 
             
    }

    for _, groupedAnagrams := range mapHashedStrings {
        groupedStrings = append(groupedStrings, groupedAnagrams)
    }

    return groupedStrings
    
}
