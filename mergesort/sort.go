package mergesort

func MergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    arrL := MergeSort(arr[:len(arr)/2])
    arrR := MergeSort(arr[len(arr)/2:])

    return mergeArrays(arrL, arrR)
}

func mergeArrays(arrL, arrR []int) []int {
    mergedArr := make([]int,0, len(arrL)+len(arrR))

    posL, posR := 0, 0
    for i := 0; i < len(arrL)+len(arrR); i++ {
        //left array is traversed, so right can be just concatenated
        if len(arrL) == i-posL {
            mergedArr = append(mergedArr, arrR[i-posR:]...)
            break
        }

        //right array is traversed, so left can be just concatenated
        if len(arrR) == i-posR {
            mergedArr = append(mergedArr, arrL[i-posL:]...)
            break
        }

        //find smaller element
        if arrL[i-posL] < arrR[i-posR] {
            mergedArr = append(mergedArr, arrL[i-posL])
            posR++
        } else {
            mergedArr = append(mergedArr, arrR[i-posR])
            posL++
        }
    }

    return mergedArr
}