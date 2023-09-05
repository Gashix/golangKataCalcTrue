package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func safeDivide(a,b int)(){
	defer func ()  {
		if r:=recover(); r!=nil{
			fmt.Println("деление на ноль недопустимо!")
			duration := 3 * time.Second
						time.Sleep(duration)
						os.Exit(1)
		}
	}()
	result := a/b
	fmt.Println(result)
}

func rimArab(rim string) (arab int) {
	rimNums := [20]string{"I","II","III","IV","V","VI","VII","VIII","IX","X","XI","XII","XIII","XIV","XV","XVI","XVII","XVIII","XIX","XX"}
	for index, num := range rimNums {
		if num == rim{
			arab = index+1
		}
	}
	return
}

func arabRim(arab int) (rim string) {
	carta := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10:"X", 
		11:"XI", 12:"XII", 13:"XIII", 14:"XIV", 15:"XV", 16:"XVI", 17:"XVII", 18:"XVIII", 19:"XIX", 20:"XX",
		21:"XXI", 22:"XXII", 23:"XXIII", 24:"XXIV", 25:"XXV", 26:"XXVI", 27:"XXVII", 28:"XXVIII", 29:"XXIX", 30:"XXX",
		31:"XXXI", 32:"XXXII", 33:"XXXIII", 34:"XXXIV", 35:"XXXV", 36:"XXXVI", 37:"XXXVII", 38:"XXXVIII", 39:"XXXIX", 40:"XL",
		41:"XLI", 42:"XLII", 43:"XLIII", 44:"XLIV", 45:"XLV", 46:"XLVI", 47:"XLVII", 48:"XLVIII", 49:"XLIX", 50:"L",
		51:"LI", 52:"LII", 53:"LIII", 54:"LIV", 55:"LV", 56:"LVI", 57:"LVII", 58:"LVIII", 59:"LIX", 60:"LX",
		61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
		71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
		81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}

	for key, num := range carta {
		if key == arab{
			rim = num
		}
	}
	return
}

func main(){
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		line := scan.Text()
		if line == "exit"{os.Exit(0)}
	
		refirst := regexp.MustCompile(`(\d\s\d)|([I,II,III,IV,V,VI,VII,VIII,IX,X])(\s)([I,II,III,IV,V,VI,VII,VIII,IX,X])`)
		foundfirst := refirst.FindAllString(line,-1)
		if foundfirst != nil{
			fmt.Println("цифры пишутся слитно!")
			duration := 3 * time.Second
			time.Sleep(duration)
			os.Exit(1)			
		} else{
			line1 := strings.ReplaceAll(line, " ", "")			
		
		//Римские вычисления
			reRim := regexp.MustCompile(`(^IX|^IV|^III|^II|^I|^VIII|^VII|^VI|^V|^X)(\+|-|\*|\/)(IX$|IV$|III$|II$|I$|VIII$|VII$|VI$|V$|X$)`)
			foundRim := reRim.FindAllString(line1,-1)
			resultArab:= 0
			if foundRim != nil{
				for _, rimElement := range line1 {
					switch rimElement {
					case '+':
						rimSplit := strings.Split(line1, "+")
						ls  := rimSplit[0]
						arab1 := rimArab(ls)
						rs := rimSplit[1]
						arab2 := rimArab(rs)
						resultArab = arab1+arab2
					case '-':
						rimSplit := strings.Split(line1, "-")
						ls  := rimSplit[0]
						arab1 := rimArab(ls)
						rs := rimSplit[1]
						arab2 := rimArab(rs)
						resultArab = arab1-arab2
					case '*':
						rimSplit := strings.Split(line1, "*")
						ls  := rimSplit[0]
						arab1 := rimArab(ls)
						rs := rimSplit[1]
						arab2 := rimArab(rs)
						resultArab = arab1*arab2
					case '/':
						rimSplit := strings.Split(line1, "/")
						ls  := rimSplit[0]
						arab1 := rimArab(ls)
						rs := rimSplit[1]
						arab2 := rimArab(rs)
						resultArab = arab1/arab2
					}
				}
				if resultArab < 1 {
					fmt.Println("Ответ меньше одного, что недопустимо для римских цифр!!!")
					duration := 3 * time.Second
					time.Sleep(duration)
					os.Exit(1)					
				} else{
					resultRim := arabRim(resultArab)
					fmt.Println(resultRim)				
				}

			} else {
				//Арабские вычисления
					re:=regexp.MustCompile(`(^10|^[1-9]{1}|^-10|^-[1-9]{1})(\+|-|\*|\/)(10$|[1-9]{1}$|-10$|-[1-9]{1}$)`)
					found := re.FindAllString(line1, -1)

					re_zero:=regexp.MustCompile(`(^|[^0-9])0`)
					foundzero:=re_zero.FindAllString(line1, -1)

					if foundzero != nil{
						fmt.Println("число 0 недопустимо, входные данные могут быть только от 1 до 10")
						duration := 3 * time.Second
						time.Sleep(duration)
						os.Exit(1)		
					}

					if found == nil {
						fmt.Printf("Выражение написано неверно!\n")
						duration := 3 * time.Second
						time.Sleep(duration)
						os.Exit(1)						
					} else{
						reminus :=regexp.MustCompile(`\d-\d`)
						found_minus := reminus.FindAllString(line1,-1)
						
						for _, element := range line1 {
							switch element {
							case '+':							
								lineSplit := strings.Split(line1, "+")
								ls, _ := strconv.Atoi(lineSplit[0])							
								rs, _ := strconv.Atoi(lineSplit[1])								
								fmt.Println(ls + rs)
							case '*':
								lineSplit := strings.Split(line1, "*")
								ls, _ := strconv.Atoi(lineSplit[0])
								rs, _ := strconv.Atoi(lineSplit[1])
								fmt.Println(ls * rs)
							case '/':
								lineSplit := strings.Split(line1, "/")
								ls, _ := strconv.Atoi(lineSplit[0])
								rs, _ := strconv.Atoi(lineSplit[1])
								safeDivide(ls,rs)
							case '-':
								if found_minus != nil{
									lineSplit := strings.Split(line1, "-")
									ls, _ := strconv.Atoi(lineSplit[0])
									rs, _ := strconv.Atoi(lineSplit[1])
									fmt.Println(ls - rs)
								}
							}
						}

					}				
	
			}
		}		
	}	
}





