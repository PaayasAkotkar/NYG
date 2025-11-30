
/**Paayas Akotkar formulas
 * @copyright @2024
 */




/**Paayas Akotkar formulas
 * @copyright @2024
 */

/***Api's for working with array */
export namespace RangeArray {
    /* required enum*/
    export enum Acoordon {
        Aindexes, /**get all the indexes of the duplicate number*/
        Findex, /**get first index of the total number of duplicate number */
        Foccurence, /**get first occurence of the duplicate number*/
        Aoccurence, /**get all the duplicate number*/
    }
    // end of enum

    /**
     * removes the duplicate value from an array
     * note: it is recommend to use a variable to get the values
     * becuase its not cpp void which manipulates the passed param
     * @param DupValArray: a range that contains duplicate value
     * @returns removes the total number of dupicate array from the range 
    */
    export function RemoveDuplicateVal(DupValArray: number[] | any[]): void {
        DupValArray.sort((a, b) => a - b) // important

        // removing duplicate val 

        DupValArray = DupValArray.filter((e, i) => (e != DupValArray[i - 1]))

        // another way of removing duplicate val:
        // DupValArray = DupValArray.filter((e, i) => {
        //     while (e != DupValArray[i - 1]) {
        //         return DupValArray
        //     }
        // })
    }
    // end of RemoveDuplicateVal method

    /**  
     *  note: it is recommend to use a variable to see the change
     * @param ar: an integeral range
     * @param valueToRemove : value to remove from the range
     * @returns removes the element from the range
     * 
    */
    export function erase(ar: number[] | any[], valueToRemove: number | any) {
        ar.sort((a, b) => a - b) // important
        // removing duplicate val 
        ar = ar.filter((e, i) => (valueToRemove != ar[i]))
        return [...ar]
    }
    // end of RemoveValue method

    /**  HHArray= Helping Hand Array
     * a class that provides a shorthand for getting duplicate value and their value index either all or first or last occurence
     * note current it is only preferable to work with created methods here 
     * 
    */
    export class HHArray {
        /**note: the fields are filled with the created methods in this module*/
        firstDupNumber: number | any = 0
        LastDupNumber: number | any = 0
        DupNumbers: number[] | any = [] /**  list of duplicate numbers // note not like 2,2,3,.. */
        ADupNumbers: number[] | any = [] // list of all the duplicate numbers // like 2,2,3,3... */
        ADupNumbersindex: number[] | any = [] // list of all the duplicate numbers indexes // like [2, 2,3,3] then output 0, 1, 2, 3 */
        RefArray: number[] = []  /**  passed array through constructor */
        /**note that this will require you to pass an array to compare with and an enum Acoordon */
        FoundIndexes: number[] = [] /**  either matched indexes or matched indexes values  on desired enum Acoordon in the RefArray*/
        // end of fields 

        /**  actually it is lazy work 
         * it is recommend to use  the methods if got confuse
        */
        constructor(ArrayParam: number[], /**this is the parent array that will fills the given field */
            CompareParam: number[], /**to compare array with the array param*/
            a: Acoordon
        ) {
            this.RefArray = ArrayParam
            this.firstDupNumber = FDuplicate(this.RefArray, false, true)
            this.LastDupNumber = FDuplicate(this.RefArray, false, false, true)
            this.DupNumbers = FDuplicate(this.RefArray, false, false, false, true)
            DTIN(ArrayParam, this.ADupNumbers, true)
            DTIN(ArrayParam, this.ADupNumbersindex)
            FDNIA(ArrayParam, CompareParam, a, this.FoundIndexes)
        }
        // end of constructor 


    } // end of class

    /** returns the total number of count of duplicate numbers
      * if 2,2 then the count is 1 meaning that 2 is repeated once*/
    export function count(ar: number[],
        CountOf: number, /**count of duplicate number to search in ar  */
        bDupValIndex: boolean = false, /**get first occurence index of the CountOf */
        /**  note that the output may be change if you wish to perform a sort
         * not recommended
        */
        bSortGivenArray: boolean = false
    ) {
        let countIs = 0
        let IDupVal = 0 // for the  duplicate value index

        // kind of unorthodox becuase the client may get confuse the passed array and outcome
        // for inputs like: [1, 2, 2, 0, 0, 2, 4, 4] // now without sort the human eye will say 2 and after sort the computer will output 0
        if (bSortGivenArray) {
            ar.sort((a, b) => a - b) // sorting the array so that the duplicate numbers are next to each other
        }
        for (let i = 0; i < ar.length; i++) {
            let m = ar.findIndex(e => e == CountOf)
            if (ar[i] == ar[m]) {
                countIs = (i - m)
                //  duplicate value index
                IDupVal = ar.findIndex(e => e == ar[i])
            }

        }

        // if there is no duplicate number
        if (countIs == 0) {
            return -1
        }
        // to return the index of the duplicate value
        if (bDupValIndex) {
            return IDupVal
        }

        // to return the count of duplicate value
        return countIs
    }
    // end of count method 

    /**  returns all the number of indexes of the duplicate numbers in an array or if prefered returns all the duplicate numbers in the array 
       *  DTIN= Duplicate Total Index Numbers
       * @param 
       * */
    export function DTIN(ar: number[], /**  an array to find duplicate value or index */
        fresh_ar: number[], /** to store the total number of duplicate index*/
        bListDuplicateValues: boolean = false, /**returns all the duplicate vlaues in an ar // note: not the index of those values*/
        bSort: boolean = false /**note: may change the output and not recommend*/
    ) {

        if (bSort) {
            ar.sort((a, b) => a - b)
        }

        // an array to compare with all the duplicate values of the parent array
        let z: number[] = []
        FDuplicate(ar, false, false, false, true, z) // get all the duplicate numbers

        // console.log("z: ", z)

        // run through the two array
        // comparing as if comparing the 2d array
        // this is important and is the only way to compare index value with each index value at 
        for (let i = 0; i < ar.length; ++i) {
            for (let j = 0; j < z.length; ++j) {
                if (z[j] == ar[i]) {
                    // returns all the duplicate values
                    if (bListDuplicateValues && bSort || bListDuplicateValues) {
                        fresh_ar.push(ar[i])
                    }
                    // returns all the duplicate values indexes 
                    else if (bSort || !bListDuplicateValues) { fresh_ar.push(i) }
                }
            }
        }
        // console.log("fresh_z", fresh_ar)
    }
    // end of DTIN method 

    /** returns the first duplicate value in the array or if desired to have last duplicate number that can be get too or 
      * wish to get all the duplciate number
      * though you can view the total number of duplicate values in the array too
      * F stands for Find
     */
    export function FDuplicate(ar: number[], /**array parameter*/
        viewDuplicateNumbers: boolean = false, /**to view total number of duplicate number in the array parameter*/
        FDuplicateNumber: boolean = false, /**to view first duplicate number in the array parameter*/
        LDuplicateNumber: boolean = false, /**to get last duplicate number in the array parameter*/
        bBothTheValues: boolean = false,  /**to get both the duplicate numbers in the array parameter*/
        DuplicateNumbers: number[] = [], /**you can even pass an another array to get both the duplicate number*/
        bSortAr: boolean = false, /**note the output may change and you may get confuse and it is not recommend*/
    ) {
        if (bSortAr) { ar.sort((a, b) => a - b) }
        let StoreBothValues: number[] = []
        let cDuplicateNumberIndex = 0
        // looping for first duplicate value
        for (let i = 0; i < ar.length; i++) {
            // get the index of the duplicate number
            let m = count(ar, i, true)

            // unless m has duplicate numbers
            // get the count of those duplicate index
            if (m != -1) {
                // count of duplicate number
                cDuplicateNumberIndex = count(ar, ar[m])
                //  valid for input like: col1 = [1, 2, 2, 0, 0, 2, 4, 4] / becuase count will sort the array
                StoreBothValues.push(ar[m]) // storing those values
                if (viewDuplicateNumbers) {
                    // console.log("the count of duplicate value: ", cDuplicateNumber)
                    // console.log("found duplicate numbers: ", ar[m])
                    // console.log("index: ", cDuplicateNumberIndex)
                }
            }
        }
        // looping for last duplicate value
        let LastDupNum = 0
        for (let i = 0; i < ar.length; i++) {
            let h = ar.findIndex((e) => e == ar[i])
            let co = count(ar, ar[i], true)

            if (co != -1) {
                LastDupNum = h
                // console.log(LastDupNum)
            }
        }



        // note: returning StoreBothValues at index 0 and len-1 may not return the duplicate value
        switch (true) {
            case FDuplicateNumber: {
                return ar[cDuplicateNumberIndex]
            } break
            case LDuplicateNumber: {
                return ar[LastDupNum]
            } break
            case bBothTheValues && DuplicateNumbers != ar: {
                for (let i = 0; i < StoreBothValues.length; ++i) {

                    DuplicateNumbers[i] = StoreBothValues[i]  // just storing the values in the array
                }
                return StoreBothValues
            } break
            case bBothTheValues && DuplicateNumbers == ar: {
                // console.log("please provide a fresh array")
                return -1
            } break
            default: {
                // console.log("please provide true to any of the one boolean value")
                return -1
            }
        }
    }
    // end of FDuplicate method

    /**  returns the index value and index at the parent array that has been compared with child array 
     * FDNIA= Find Duplicate Number Index Array */
    export function FDNIA(
        /** no need worry about the parent and child definition you can pass as you wish to pass*/
        parent: number[], /**an array of highest len*/
        child: number[], /**an array of lowest len than parent*/
        ToGet: Acoordon, /**an enum to proceed to desired outcome*/
        fresh_ar: number[] /**to store desired outcome*/
    ) {
        let x: number[] = [] // to store dup num of parent
        let y: number[] = []  // to store dup num of child
        DTIN(parent, x)
        DTIN(child, y)
        // console.log("x: ", x)
        // console.log("x: ", y)

        // operator plus is important as it allows any cases to work with
        switch (+ToGet) {
            case Acoordon.Aindexes:
                {
                    // lopping sequnce: from child to parent
                    // returns all the indexes found in the comparision of the highest len
                    if (parent.length > child.length) {
                        for (let i = 0; i < y.length; ++i) {
                            for (let j = 0; j < x.length; ++j) {
                                if (parent[j] == child[i]) {
                                    // console.log(j)
                                    fresh_ar.push(j)
                                }
                            }
                        }
                    }
                    else if (child.length >= parent.length) {
                        for (let i = 0; i < x.length; ++i) {
                            for (let j = 0; j < y.length; ++j) {
                                if (child[j] == parent[i]) {
                                    // console.log(j)
                                    fresh_ar.push(j)
                                }
                            }
                        }
                    }
                } break
            case Acoordon.Aoccurence:
                // returns the value of all the first occurence of the duplicate number in the array that is iterated on the highest len
                {
                    if (parent.length > child.length) {
                        for (let i = 0; i < y.length; ++i) {
                            for (let j = 0; j < x.length; ++j) {
                                if (parent[j] == child[i]) {
                                    //  // console.log("dup num: ", parent[j])
                                    fresh_ar.push(parent[j])
                                }
                            }
                        }
                    }
                    else if (child.length >= parent.length) {
                        for (let i = 0; i < x.length; ++i) {
                            for (let j = 0; j < y.length; ++j) {
                                if (child[j] == parent[i]) {
                                    // console.log(child[j])
                                    fresh_ar.push(child[j])
                                }
                            }
                        }
                    }
                } break
            case Acoordon.Findex:
                // returns  the index of all the first occurence of the duplicate number in the array that is iterated on the highest len
                {
                    if (x.length > y.length) {
                        for (let i = 0; i < x.length; ++i) {
                            if (parent[i] == child[i]) {
                                // console.log("pf: ", i)
                                fresh_ar.push(i)
                            }
                        }
                    } else if (y.length >= x.length) {
                        for (let i = 0; i < y.length; ++i)
                            if (child[i] == parent[i]) {
                                // console.log("cf: ", i)
                                fresh_ar.push(i)
                            }
                    }
                } break
            case Acoordon.Foccurence:
                // returns the value of all the first occurence of the duplicate number in the array that is iterated on the highest len
                {
                    if (x.length > y.length) {
                        for (let i = 0; i < x.length; ++i) {
                            if (parent[i] == child[i]) {
                                // console.log("pf: ", parent[i])
                                fresh_ar.push(parent[i])
                            }
                        }
                    } else if (y.length >= x.length) {
                        for (let i = 0; i < y.length; ++i)
                            if (child[i] == parent[i]) {
                                // console.log("cf: ", child[i])
                                fresh_ar.push(child[i])
                            }
                    }
                } break
            default: {
                // console.log("please give true to any one of the boolean value")
            }
        }
        return fresh_ar
    }
    // end of FDNIA method 
    /** swaps two array with each 
    * and can also keep second array without swapping
    */
    export function swap(
        ar: number[], numsToswapWith: number[],
        bKeepnumsToswapAsItis: boolean = false,/**doesnt swap values of ar in the numsToswapWith  */
    ) {
        let temp = [...ar] // take all the number of ar2
        // do normal swapping

        ar.splice(0, ar.length, ...numsToswapWith)
        if (bKeepnumsToswapAsItis) { }
        else {
            numsToswapWith.splice(0, temp.length, ...temp)
        }
    }
    // end of method swap 

    /** swaps the array value at the inx 
     * Vswap: variant swap
     * does normal swapping but on as per index
     * if provided only value it doesnt only swap it with the value index
     * note not to pass the index that is out of bounds
     * recommendation: pass index,
     * one of working with this method is like pass and pass
     * example: for [5,4,3,2,1]
     *       let c = Vswap(ar, ar[i], true, false, i + 1)
     *      let d = Vswap(ar, c, true, false, i + 1) 
     */
    export function Vswap(ar: number[],
        i: number, /**value to swap 
    note that if not provided the index this will be your index and value too  */
        bIndex: boolean = false, /** above value to swap at index */
        bReverse: boolean = false, /**whether are you passing the negative index value */
        inx: number = -1 /**index to swap i and -1 is the value for the reverse iteration */
    ) {

        if (inx == -1 && bIndex == true && !bReverse) {
            // console.log("please pass the positive index ")
            return -1
        }
        if (inx >= ar.length) {
            // console.log("invalid index")
            return -1
        }
        // console.log(inx)
        if (!bIndex) {
            let temp = ar[i]
            ar[i] = i
            i = temp
            // console.log("i: ", i)
        } else {
            let temp = ar[inx]
            ar[inx] = i
            i = temp
        }
        return i
    }
    // end of method Vswap

    /**
     *  swaps the value inside of an array
     *  you just have to pass the index 
     *  V2swap: V2 variant two swap
     * @param ar: an array param for swapping the index value
     * @param domain: an index swap domain with range
     * @param range : an index swaps with domain 
     * example: [1,2,3]=>[dom, range]=[0,2] => [3,2,1]
     */
    export function V2swap(ar: number[], domain: number, range: number) {
        let g: number, h: number
        // if (domain == range) {
        //     // console.log("cant swap with same index")
        // }
        if (domain > ar.length || range > ar.length) {
            // console.log("index out of bounds")
        }

        g = Vswap(ar, ar[domain], true, false, range)
        h = Vswap(ar, g, true, false, domain)

    }

    // end of V2swap method

    /**
     * 
     * @param first: begin of an collection
     * @param last : end of an collection
     * @param init : initialization 
     * @returns: the sum of the given value init and the elements in the range [first,last]
     */
    export function accumulate(first: number, last: number, init: number) {
        for (; first != last; ++first) {
            init = init + first
        }
        return init
    }

    /**
     * 
     * @param ar : a parent array to find a median
     * @returns : an average 
     */
    export function median(ar: number[]): number {
        let med: number = 0
        let sum_of_element: number = 0
        for (let i = 0; i < ar.length; ++i) {
            sum_of_element += i
        }
        med = Math
            .round(sum_of_element / ar.length)
        return med
    }

    /**
     * 
     * @param begin: begining of an collection
     * @param end : end of an collection
     * @returns distance of the element; for example, [3,6] the distance =3
     */
    export function distance(begin: number, end: number) {
        let result = 0
        while (begin != end) {
            ++begin
            ++result
        }
        return result
    }

    export function RShuffle(ar: number[]) {
        // console.log("passed ar: ", ar)

        let storeOldar: number[] = [...ar] // storing the values of ar before passing
        // if current number greater than the last number
        // do sorting in descending
        // if (ar[0] > ar[ar.length]) {
        //     ar.sort((a, b) => b - a)
        // } else {
        //     ar.sort((a, b) => a - b)
        // }
        ar.sort((a, b) => b - a)
        // console.log("sorted ar: ", ar)
        storeOldar.sort((a, b) => b - a)
        // console.log("sorted ar: ", storeOldar)
        // lets find median number of array
        let med = median(ar)
        return -1
    }
}