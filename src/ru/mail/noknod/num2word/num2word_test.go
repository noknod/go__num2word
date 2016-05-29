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
        fmt.Println(v.sNumber, v.maxTripletsCnt, v.triplets[0])
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
        }/* else {
            var s string = fmt.Errorf("%v", r).Error()
            fmt.Println("\t__", reflect.TypeOf(r), "\n\t", r, "\n\t==" + s + "\n")
        }*/
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
