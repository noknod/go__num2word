package num2word


import  ("testing" ; "fmt" ; "reflect" )


const nillArgumentREMsg = 
    "runtime error: invalid memory address or nil pointer dereference"
const nillArgumentMsg = "Error: must panic with \"runtime error: inval" +
    "id memory address or nil pointer dereference\" when given nill for sNu" +
    "mber argument."


func checkIsNillArgumentError(arg *interface {}) bool {
    if arg == nil {
        return false
    } else {
        var r interface {} = *arg
        switch x := r.(type) {
        case error:
            return (fmt.Errorf("%v", x).Error() == nillArgumentREMsg)
        default:
            return false
        }
    }
}


func TestSplitToTriplets(t *testing.T) {
    var (
        sNumber = ""
        maxTripletsCnt = 2
        triplets = SplitToTriplets(&sNumber, &maxTripletsCnt)
    )
    if triplets != nil {
        t.Errorf("Error: must return nil when sNumber is empty string")
    }

    sNumber = "123456"
    maxTripletsCnt = 1
    triplets = SplitToTriplets(&sNumber, &maxTripletsCnt)
    if triplets != nil {
        t.Errorf("Error: must return nil when count of triplets is greater " +
            "then given maxTripletsCnt")
    }
    
    testSet := []struct {
        sNumber string
        maxTripletsCnt int
        triplets []string // предполагамый результат
    }{
        {"1", 2, []string{"001"}},
        {"0001", 2, []string{"001", "000"}},
        {"101", 2, []string{"101"}},
        {"123456", 2, []string{"456", "123"}},
        {"3456", 2, []string{"456", "003"}},
    }

    for _, v := range testSet {
        triplets = SplitToTriplets(&v.sNumber, &v.maxTripletsCnt)
        if !reflect.DeepEqual(triplets, v.triplets) {
            t.Errorf("Split for %s and %d:\nexpected %v, got %v", 
                v.sNumber, v.maxTripletsCnt, v.triplets, triplets)
        }
    }
}

func TestSplitToTripletsNillNumberArgument(t *testing.T) {
    defer func() {
        if r := recover(); !checkIsNillArgumentError(&r) {
            t.Errorf(nillArgumentMsg)
        }
    }()
    var maxTripletsCnt = 2
    _ = SplitToTriplets(nil, &maxTripletsCnt)
}

func TestSplitToTripletsNillMaxCntArgument(t *testing.T) {
    defer func() {
        if r := recover(); !checkIsNillArgumentError(&r) {
            t.Errorf(nillArgumentMsg)
        }
    }()
    var sNumber = "2"
    _ = SplitToTriplets(&sNumber, nil)
}


func TestTripletToWord(t *testing.T) {
    var tripletInfo TripletInfoList = TripletInfoList{
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
    
    testSet := []struct {
        triplet string
        tripletInfo TripletInfoHolder
        word string // предполагамый результат
    }{
        {"000", tripletInfo.Get(0), ""},
        {"000", tripletInfo.Get(2), ""},
        {"001", tripletInfo.Get(0), "один"},
        {"001", tripletInfo.Get(1), "одна тысяча"},
        {"002", tripletInfo.Get(1), "две тысячи"},
        {"020", tripletInfo.Get(0), "двадцать"},
        {"010", tripletInfo.Get(1), "десять тысяч"},    
        {"022", tripletInfo.Get(2), "двадцать два миллиона"},
        {"500", tripletInfo.Get(0), "пятьсот"},
        {"700", tripletInfo.Get(2), "семьсот миллионов"},
        {"201", tripletInfo.Get(0), "двести один"},
        {"301", tripletInfo.Get(1), "триста одна тысяча"},
        {"402", tripletInfo.Get(1), "четыреста две тысячи"},
        {"135", tripletInfo.Get(0), "сто тридцать пять"},
        {"135", tripletInfo.Get(1), "сто тридцать пять тысяч"},
    }

    var (
        word string
        ok bool
    )
    for _, v := range testSet {
        word, ok = TripletToWord(v.triplet, v.tripletInfo)
        if !ok || !(v.word == word) {
            t.Errorf("Word for <%s> and <%s>: expected <%v>, got <%v>", 
                v.triplet, v.tripletInfo.GetCommonWord(), v.word, word)
        }
    }
}


func TestNumToWord(t *testing.T) {
    testSet := []struct {
        number int
        word string // предполагамый результат
    }{
        {0, "ноль"},
        {1, "один"},
        {-1, "минус один"},
        {1000, "одна тысяча"},
        {-2000, "минус две тысячи"},
        {20, "двадцать"},
        {10000, "десять тысяч"},    
        {22000000, "двадцать два миллиона"},
        {542, "пятьсот сорок два"},
        {-201145, "минус двести одна тысяча сто сорок пять"},
        {135001075, "сто тридцать пять миллионов одна тысяча семьдесят пять"},
        {135001, "сто тридцать пять тысяч один"},
    }

    var word string

    for _, v := range testSet {
        word = NumToWord(v.number)
        if !(v.word == word) {
            t.Errorf("Word for <%d>: expected <%v>, got <%v>", 
                v.number, v.word, word)
        }
    }

}