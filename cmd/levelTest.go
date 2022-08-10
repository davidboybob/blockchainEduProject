package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

var level1QuestionsMap = map[string]bool{
	"블록체인은 레코드 유지 및 계약 적용 기술이며 이전 기록을 변경하기가 극도로 어렵게 만드는 암호화를 사용 한다.":               true,
	"블록체인 참가자는 공유 내용에서 변경 내용을 추적하여 작업 흐름을 알 수 있다.":                                 true,
	"블록체인은 중앙 집중식 데이터베이스다.":                                                        false, //분산 데이터베이스이다.
	"분산 데이터베이스의 장점은 액세스와 일관성을 제어하기가 쉽다는 것이다.":                                      false, // 중앙 집중식 데이터베이스의 장점이다.
	"분산 데이터베이스의 장점은 각 참가자에게 데이터베이스의 복사본이 있다는 것이다.":                                 true,
	"블록체인 참가자 간에 분산된 블록체인 네트워크를 컨소시엄 네트워크라고 한다.":                                   true,
	"블록체인은 차(minus)의 규칙을 사용하여 데이터를 전체 노드에서 일관되도록 유지한다.":                            false, //합의 규칙을 사용한다.
	"블록체인은 암호화를 사용하여 참가자가 데이터를 신뢰할 수 있도록 한다.":                                      true,
	"블록체인은 개별 참가자 또는 소수의 참가자가 기록을 수정하지 못하도록 방지한다.":                                 true,
	"블록체인은 탈중앙화되어 있으므로 탈중앙식 데이터베이스를 사용할 수 있는 솔루션이 가장 효율적으로 작동한다.":                  true,
	"블록체인을 사용하면 중앙 권리자를 둬야 한다.":                                                    false, // 필요 없다.
	"블록체인의 데이터는 상태를 나타낸다. 이는 암호 화폐와 같은 디지털 토큰이 블록체인에 적합한 이유이다.":                    true,
	"블록체인은 트랜잭션을 사용하여 데이터의 상태를 한 값에서 다른 값으로 변경한다.":                                 true,
	"작업 증명, 지분 증명 및 권한 증명을 비롯한 여러 블록체인 합의 알고리즘이 있다. 각 알고리즘은 서로 다은 방식으로 일관성을 해결한다.": true,
	"블록은 트랜잭션 정보를 저장하는 브록체인 내의 데이터 클러스터이다.":                                        true,
	"블록체인은 암호화 해시를 사용하여 블록 간에 링크를 만든다.":                                            true,
	"암호화 해시는 임의 크기의 데이터를 고정된 크기의 비트 표현에 매핑하는 알고리즘이다.":                              true,
	"비트코인은 SHA-256 해시 알고리즘을 사용한다":                                                  true,
}

type QuestionContent struct {
	errMsg string
	label  string
}

func levelTestGetInput(qc QuestionContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(qc.errMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold}}",
	}

	prompt := promptui.Prompt{
		Label:     qc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v \n", err)
		os.Exit(1)
	}

	return result
}

type Queue []interface{}

//IsEmpty - 큐가 비어있는지 확인하는 함수.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//Enqueue - 큐에 값을 추가하는 함수.
func (q *Queue) EnQueue(data interface{}) {
	*q = append(*q, data)
}

//Dequeue - 큐에 첫번째 요소를 반환하고 제거하는 함수.
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	data := (*q)[0] // 큐에 첫번째 값을 가져옴.
	*q = (*q)[1:]   // 큐에 첫번째 데이터를 제거함.
	fmt.Printf("Dequeue: %v\n", data)
	return data
}

//입력 받은 count 만큼 문제를 만든다
func CreateQuestionBank(count int) *Queue {
	//여러가지 자료구조를 사용해보기위해 작성한 함수, 추후 최적화 필요

	//1. 무작위로 문제를 내기 위해 순서를 섞는다.
	questionTotalCnt := len(level1QuestionsMap)
	fmt.Printf("전체 문제 수 %d \n", questionTotalCnt)
	questionSlice := make([]int, questionTotalCnt)
	for i := 0; i < questionTotalCnt; i++ {
		questionSlice[i] = i
	}

	fmt.Println("섞기 전 : ", questionSlice)

	rand.Seed(time.Now().UnixNano()) //random 값 생성을 위한 seed 설정 (default가 시간이긴 하다.)

	rand.Shuffle(len(questionSlice), func(i, j int) { questionSlice[i], questionSlice[j] = questionSlice[j], questionSlice[i] })

	fmt.Println("섞은 후 : ", questionSlice)

	q := Queue{}

	for question, answer := range level1QuestionsMap {
		fmt.Println(question, answer)
	}

	var idx int = 0
	for index := range level1QuestionsMap {
		for index2 := range questionSlice[0:count] {
			if idx == index2 {
				levelTestPropmptContent := QuestionContent{
					"Please, Do it Again!",
					index,
				}

				levelTestAnswer := levelTestGetInput(levelTestPropmptContent)

				fmt.Println(levelTestAnswer)
				// fmt.Println(index, "=>", element)
				q.EnQueue(index)
			}
		}
		idx++
	}

	//사용할때는 Dequeue 해서 사용
	//q.Dequeue()

	return &q
}
