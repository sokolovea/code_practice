package main

import (
	"bufio"
	"fmt"
	mathrand "math/rand"
	"os"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

type CensusFile struct {
	fileDescriptor *os.File
	writer         *bufio.Writer
}

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	populationCount int
	cenusWritersMap map[byte]*CensusFile
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.populationCount
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	writer := c.cenusWritersMap[name[0]].writer
	writer.WriteString(name)
	writer.WriteByte('\n')
	c.populationCount++
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	for _, value := range c.cenusWritersMap {
		value.writer.Flush()
		err := value.fileDescriptor.Close()
		if err != nil {
			panic(err)
		}
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	census := Census{populationCount: 0, cenusWritersMap: make(map[byte]*CensusFile, len(alphabet))}
	err := os.Mkdir("census", 0755)
	// if err != nil {
	// 	panic(err)
	// }
	err = os.Chdir("census")
	if err != nil {
		panic(err)
	}
	for _, c := range alphabet {
		census.cenusWritersMap[byte(c)] = new(CensusFile)
		fileStructPtr := census.cenusWritersMap[byte(c)]
		fileStructPtr.fileDescriptor, err = os.Create(string(c) + ".txt")
		if err != nil {
			panic(err)
		}
		fileStructPtr.writer = bufio.NewWriter(fileStructPtr.fileDescriptor)
	}
	return &census
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for range 1024 {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
