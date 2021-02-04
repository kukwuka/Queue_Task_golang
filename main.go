package main

//
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//----------------------------------------------------------------------------------------------------------------------

type Line struct {
	mainChan     chan InLine
	writersCount int
}

func (l *Line) DoTask(Ws int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < Ws; i++ {
			select {
			case value, ok := <-l.mainChan:
				if !ok {
					l.mainChan = nil
				}
				ArrLen := len(value.arrOfRandom)
				fmt.Printf("RoutineId :%v time: %v min: %v avarage :%v max:%v  \n",
					value.goroutineId,
					value.queueTime,
					value.arrOfRandom[0],
					value.arrOfRandom[ArrLen/2],
					value.arrOfRandom[ArrLen-1],
				)
			}
			if l.mainChan == nil {
				break
			}
		}
	}()
}

func (l *Line) AddTask(AS, IC int) {
	for i := 0; i < l.writersCount; i++ {
		//Add task to chan
		var NewTask AddTaskIntoChain = &Task{i, AS, IC}
		wg.Add(1)
		go NewTask.CreateTask(l.mainChan)
	}
}
func (l *Line) CloseChanel() {
	wg.Wait()
	close(l.mainChan)

}

type Queue interface {
	DoTask(Ws int)
	AddTask(AS, IC int)
	CloseChanel()
}

//---------------------------------------------------------------------------------------------------------------------
type InLine struct {
	goroutineId int
	arrOfRandom []int
	queueTime   time.Time
}

//---------------------------------------------------------------------------------------------------------------------

type Task struct {
	goroutineId int
	arrSize     int
	iterCount   int
}

func (t *Task) createRandArr() []int {
	//Create Array of rand int
	arrToWork := make([]int, t.arrSize, t.arrSize)
	for i := 0; i < t.arrSize; i++ {

		arrToWork[i] = rand.Intn(t.arrSize)
	}
	return arrToWork
}

func (t *Task) CreateTask(ch chan<- InLine) {
	defer wg.Done()
	for i := 0; i < t.iterCount; i++ {
		ch <- InLine{t.goroutineId, t.createRandArr(), time.Now()}

	}
}

type AddTaskIntoChain interface {
	CreateTask(ch chan<- InLine)
}

//---------------------------------------------------------------------------------------------------------------------
var wg sync.WaitGroup

func main() {
	//Check input arguments
	var WritersCount, ArrSize, IterCount int
	WritersCount, ArrSize, IterCount = CheckInputArguments()
	var WaitingOutputStrings = WritersCount * IterCount
	//Create main Queue
	var MaInQueue Queue = &Line{make(chan InLine, 2), WritersCount}
	MaInQueue.AddTask(ArrSize, IterCount)
	MaInQueue.DoTask(WaitingOutputStrings)

	MaInQueue.CloseChanel()
}

func CheckInputArguments() (int, int, int) {
	argsWithoutProg := os.Args[1:]
	WC, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		fmt.Print("Write only int")
		log.Fatal(err)
	}
	AS, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Print("Write only int")
		log.Fatal(err)
	}
	IC, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		fmt.Print("Write only int")
		log.Fatal(err)
	}
	return WC, AS, IC
}
