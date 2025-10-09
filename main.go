package main

// os.Args는 string slice임

// os.Args의 첫 번째 원소인 `os.Args[0]`는 명령 자체의 이름이다. 즉, 빌드된 실행 파일의 경로다.
// 그 외의 원소들(1, 2, ...)은 프로그램이 실행될 때 프로그램에 제공된 인수다.

// echo는
// 1. 하나 이상의 문자열 인자를 받습니다
// 2. 각 인자들을 단일 공백(' ')으로 구분하여 합칩니다.
// 3. 합쳐진 문자열의 끝에 **줄바꿈 문자('\n')를 추가합니다
// 4. 최종 결과를 표준 출력(일반적으로 터미널 화면)에 출력합니다.
// 5. -n 옵션: 출력의 마지막에 자동으로 추가되는 줄바꿈 문자를 생략합니다

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nPtr *bool = flag.Bool("n", false, "")
	var ePtr *bool = flag.Bool("e", false, "")
	// var nVal bool = *nPtr // *가 역참조다.
	flag.Parse()

	var res []string = flag.Args() // []string
	var str string
	for i, a := range res {
		if len(os.Args)-1 == 1 {
			// 원소가 하나뿐임
			if *nPtr {
				str += a
			} else {
				str += fmt.Sprintf("%s\n", a)
			}
		} else {
			// 다중 원소
			if i == 0 {
				// 첫 순회
				str += a
			} else {
				// 순회 2번 이상
				if i == len(os.Args)-2 {
					// 마지막 원소
					if *nPtr {
						str += fmt.Sprintf(" %s", a)
					} else {
						str += fmt.Sprintf(" %s\n", a)
					}
				} else {
					str += fmt.Sprintf(" %s", a)
				}
			}
		}
	}

	if *ePtr {
		// TODO: str 에서 내포된 \x를 \\x로 치환해야 한다.
		// TODO: regex 사용하기.
		str = "\"" + str + "\""
		str = strings.ReplaceAll(str, "\n", "\\n")
		str = strings.ReplaceAll(str, "\t", "\\t")

		escapedStr, err := strconv.Unquote(str)
		if err != nil {
			fmt.Printf("Error unquoting string: %v\n", err)
			return
		}
		fmt.Printf("%s", escapedStr)
	} else {
		fmt.Printf("%s", str)
	}
}
