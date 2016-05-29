package num2word


import  ( "fmt" ; "strconv" )


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
