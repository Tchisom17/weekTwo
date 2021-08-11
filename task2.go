package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(nums []float64) float64{
	var sum float64
	for _, v := range nums{
		sum+= v
	}
	return sum
}
func subtract(nums []float64) float64{
	var sub float64 = nums[0]*2
	for _, v := range nums{
		sub-= v
	}
	return sub
}
func multiply(nums []float64) float64{
	var product float64 = 1
	for _, v := range nums{
		product*= v
	}
	return product
}
func divide(nums []float64) float64{
	var div float64 = nums[0]*nums[0]
	for _, v := range nums{
		div/= v
	}
	return div
}
func calc(s ...string) []float64{
	var result []float64
	var convert []float64
	for _,val:=range s{
		//DIVIDE
		if strings.ContainsAny(val, "/"){
			dt:=fmt.Sprint(strings.ReplaceAll(val, "/", " "))
			arr:= strings.Split(dt, " ")
			for _,j:=range arr{
				if num, err:= strconv.ParseFloat(j, 64); err == nil{
					convert=append(convert, num)
				}
			}
			result=append(result,divide(convert))
		}
		//ADDITION
		if strings.ContainsAny(val, "+"){
			var  b []float64
			dt:=fmt.Sprint(strings.ReplaceAll(val, "+", " "))
			arr:= strings.Split(dt, " ")
			for _,j:= range arr{
				if num, err:= strconv.ParseFloat(j, 64); err == nil{
					b = append(b, num)
				}
			}
			result = append(result, add(b))
		}
		//SUBTRACTION
		if strings.ContainsAny(val, "-"){
			var  minus []float64
			m:=fmt.Sprint(strings.ReplaceAll(val, "-", " "))
			arr:= strings.Split(m, " ")
			for _,j:= range arr{
				if num, err:= strconv.ParseFloat(j, 64); err == nil{
					minus = append(minus, num)
				}
			}
			result = append(result, subtract(minus))
		}
		//MULTIPLICATION
		if strings.ContainsAny(val, "*"){
			var times []float64
			n:=fmt.Sprint(strings.ReplaceAll(val, "*", " "))
			arr:= strings.Split(n, " ")
			for _,j:=range arr{
				if num, err:= strconv.ParseFloat(j, 64); err == nil{
					times = append(times, num)
				}
			}

			result=append(result, multiply(times))
		}
	}
	return result
}
func main(){
	//var p =[]float64{10, 5}
	//fmt.Println(add(p))
	//fmt.Println(subtract(p))
	//fmt.Println(multiply(p))
	//fmt.Println(divide(p))
	fmt.Println(calc("10/5", "3+3", "4-2", "4*2"))
}
