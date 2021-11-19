package main

import (
  "fmt"
)

func CalculateCurrentAmounts(averageAmount uint64, nodeAmount uint64) []uint64{
  result := make([]uint64, 0)
  for {
    if (averageAmount > nodeAmount){
      result[len(result) - 1] += nodeAmount
      nodeAmount = 0
      break
    }
    result = append(result, averageAmount)
    nodeAmount -= averageAmount
    averageAmount *= 2
  }
  return result
}
