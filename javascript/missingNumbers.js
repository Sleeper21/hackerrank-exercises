// Given two arrays of integers, find which elements in the second array are missing from the first array.

// Example
// arr = [7, 2, 5, 3, 5]
// brr = [7, 2, 3, 5, 5, 8]

// The array is the orginal list. The numbers missing are

// .

// Notes

//     If a number occurs multiple times in the lists, you must ensure that the frequency of that number in both lists is the same. If that is not the case, then it is also a missing number.
//     Return the missing numbers sorted ascending.
//     Only include a missing number once, even if it is missing multiple times.
//     The difference between the maximum and minimum numbers in the original list is less than or equal to 


let arr = [7, 2, 5, 3, 5]
let brr = [7, 2, 3, 5, 5, 8]
let checkedNums = []

function missingNumbers(arr, brr) {
  let output = []

  for (let i = 0; i < brr.length; i++) {
    for (let j = 0; j < arr.length; j++) {
      if ( brr[i] == arr[i] ) {
        checkedNums.push(checkedNums, brr[i])
        continue
      }  

    }
    
  }

  return output
}

console.log(missingNumbers(arr, brr));