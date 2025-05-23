package main

import (
	"github.com/TriM-Organization/bedrock-chunk-diff/define"
	"github.com/TriM-Organization/bedrock-chunk-diff/timeline"
	"github.com/TriM-Organization/bedrock-world-operator/block"
	"github.com/TriM-Organization/bedrock-world-operator/chunk"
	operator_define "github.com/TriM-Organization/bedrock-world-operator/define"
)

func main() {
	r := operator_define.Dimension(operator_define.DimensionIDOverworld).Range()

	c := chunk.NewChunk(block.AirRuntimeID, r)
	for _, value := range c.Sub() {
		value.Layer(0)
	}

	t, err := timeline.Open("ss", false, false)
	if err != nil {
		panic(err)
	}
	defer t.CloseTimelineDB()

	ctl, err := t.NewChunkTimeline(define.DimChunk{}, false)
	if err != nil {
		panic(err)
	}

	err = ctl.Append(c, nil, true)
	if err != nil {
		panic(err)
	}

	err = ctl.Save()
	if err != nil {
		panic(err)
	}
}
