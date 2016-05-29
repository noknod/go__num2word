package num2word


import  ( "fmt" ; "strconv" )


var (
    hundreds map[byte]string
    tens map[string]string
    ones map[string]string
    tripletInfo TripletInfoList
    maxTripletCnt int
)

func init() {
    hundreds = map[byte]string{
        '0': "", '1': "сто", '2': "двести", '3':  "триста", '4':  "четыреста",
        '5':  "пятьсот", '6':  "шестьсот", '7':  "семьсот", '8':  "восемьсот",
        '9':  "девятьсот",
    }
    tens = map[string]string{
        "0_": "", "10": "десять", "11": "одинадцать", "12": "двенадцать",
        "13": "тринадцать", "14": "четырнадцать", "15": "пятнадцать", 
        "16": "шестнадцать", "17": "семнадцать", "18":"восемнадцать",
        "19": "девятнадцать", "2_": "двадцать", "3_": "тридцать", 
        "4_": "сорок", "5_": "пятьдесят", "6_": "шестьдесят", 
        "7_": "семьдесят", "8_": "восемьдесят", "9_": "девяносто",
    }
    ones = map[string]string{
        "0": "", "1F": "одна", "1M": "один", "2F": "две", "2M": "два", 
        "3": "три", "4": "четыре", "5": "пять", "6": "шесть", "7": "семь", 
        "8": "восемь", "9": "девять",
    }
    tripletInfo = TripletInfoList{
        Info: []TripletInfo{
            TripletInfo{
                NounGender: 'M',
                CommonWord: "",
                WordFor1: "",
                WordFor234: "",
            },
            TripletInfo{
                NounGender: 'F',
                CommonWord: "тысяч",
                WordFor1: "тысяча",
                WordFor234: "тысячи",
            },
            TripletInfo{
                NounGender: 'M',
                CommonWord: "миллионов",
                WordFor1: "миллион",
                WordFor234: "миллиона",
            },
        },
    }
    maxTripletCnt = len(tripletInfo.Info)
}


type TripletInfoHolder interface {
    GetNounGender() byte
    GetCommonWord() string
    GetWordFor1() string
    GetWordFor234() string
}


type TripletInfo struct {
    NounGender byte
    CommonWord string
    WordFor1 string
    WordFor234 string
}
func (i *TripletInfo) GetNounGender() byte {
    return i.NounGender
}
func (i *TripletInfo) GetCommonWord() string {
    return i.CommonWord
}
func (i *TripletInfo) GetWordFor1() string {
    return i.WordFor1
}
func (i *TripletInfo) GetWordFor234() string {
    return i.WordFor234
}

type TripletInfoList struct {
    Info []TripletInfo
}
func (i *TripletInfoList) Get(index int) TripletInfoHolder {
    return &i.Info[index]
}


func SplitToTriplets(sNumber *string, maxTripletCnt *int) (out []string) {
    var (
        length = len(*sNumber)
        triplets_cnt = length / 3
    )
    if length % 3 != 0 {
        triplets_cnt++
    }
    if length == 0 || triplets_cnt > *maxTripletCnt {
        return nil
    }

    var fullNumber = fmt.Sprintf(
        "%0" + strconv.Itoa(triplets_cnt * 3) + "s", *sNumber)
    
    var dummy []string = make([]string, 0)
    for i := 0; i < triplets_cnt; i++ {
        low := i * 3
        high := low + 3
        dummy = append(dummy, fullNumber[low:high])
    }

    for i := len(dummy) - 1; i >= 0 ; i-- {
        out = append(out, dummy[i])
    }
    return
}


func TripletToWord(triplet string, tripletInfo TripletInfoHolder) (out string, ok bool) {
    var (
        words []string = make([]string, 0)
        word string
        suffix string = "common"
    )

    if word, ok = hundreds[triplet[0]]; ok {
        if word != "" {
           words = append(words, word)
       }
    } else {
        return out, false
    }

    if triplet[1] == '1' {
        if word, ok = tens[triplet[1:3]]; ok {
            if word != "" {
                words = append(words, word)
            }
        } else {
            return out, false
        }

    } else {
        if word, ok = tens[string(triplet[1]) + "_"]; ok {
            if word != "" {
                words = append(words, word)
            }
        } else {
            return out, false
        }

        if triplet[2] == '1' {
            if word, ok = ones["1" + string(tripletInfo.GetNounGender())]; ok {
                if word != "" {
                    words = append(words, word)
                }
            } else {
                return out, false
            }
            suffix = "1word"
        } else if triplet[2] == '2' {
            if word, ok = ones["2" + string(tripletInfo.GetNounGender())]; ok {
                if word != "" {
                    words = append(words, word)
                }
            } else {
                return out, false
            }
            suffix = "234word"
        } else {
            if word, ok = ones[string(triplet[2])]; ok {
                if word != "" {
                    words = append(words, word)
                }
            } else {
                return out, false
            }
        }
    }

    for _, word = range words {
        out += word + " "
    }
    if out != "" {
        out = out[:len(out) - 1]
        if suffix == "common" {
            suffix = tripletInfo.GetCommonWord()
        } else if suffix == "1word" {
            suffix = tripletInfo.GetWordFor1()
        } else if suffix == "234word" {
            suffix = tripletInfo.GetWordFor234()
        } else {
            return out, false
        }
        if suffix != "" {
            out += " " + suffix
        }
    }

    ok = true
    return
}


func NumToWord(number int) string {
    var (
        sNumber string
        prefix string
        triplets []string
        word string
        ok bool
        out string
    )

    if number == 0 {
        return "ноль"
    } else if number < 0 {
        prefix = "минус "
        number *= -1
    } else {
        prefix = ""
    }
    sNumber = strconv.Itoa(number)
    triplets = SplitToTriplets(&sNumber, &maxTripletCnt)

    for i, triplet := range triplets {
        word, ok = TripletToWord(triplet, tripletInfo.Get(i))
        if !ok {
            return ""
        }
        if word != "" {
            out = word + " " + out
        }
    }
    out = out[:len(out) - 1]
    return prefix + out
}