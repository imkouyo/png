package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

type ImageInfo struct {
	header []byte
	width int
	height int
	chunk []ImageChunk
	content []byte

}
type ImageChunk struct {
	chunkLength uint
	chunkType  string
	chunkData []byte
}

func main() {
	var img ImageInfo
	img.getImage("elephant.png")
	img.getHeader()
	fmt.Println(img.content[160:164])
	img.getChunk()
	fmt.Println(len(img.content), "length")
}

func (i *ImageInfo)getImage(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to read image")
	}
	i.content = content
}

func (i *ImageInfo)getHeader() {
	i.header = i.content[:8]
	i.content = i.content[8:]
}

func (i *ImageInfo)getChunk() {
	chunkLength := i.content[:4]
	var length uint = 0
	for i, n := range chunkLength {

		length +=uint(math.Pow(256, float64(3-i)) * float64(n))
	}

	var chunk ImageChunk
	chunk.chunkLength = length
	chunk.chunkData = i.content[:length + 12]
	chunk.chunkType = string(i.content[4:8])
	i.chunk = append(i.chunk, chunk)
	i.content = i.content[length + 12:]
	fmt.Println(length + 12, len(i.content), chunk.chunkType)
	if len(i.content) > 0 {
		i.getChunk()
	}
}