package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf"
)

func main() {
	b, err := ioutil.ReadFile("saved_model.pb")
	if err != nil {
		panic(err)
	}
	var model protobuf.SavedModel
	err = proto.Unmarshal(b, &model)
	if err != nil {
		panic(err)
	}
	mgs := model.GetMetaGraphs()
	for i, g := range mgs {
		sigDef := g.GetSignatureDef()

		fmt.Printf("Metagraph %d:\n", i)
		val := sigDef["serving_default"]
		fmt.Println("  Inputs:")
		for k, v := range val.GetInputs() {
			fmt.Printf("    %s:\n", k)
			fmt.Println("      Tensor name: \t", v.GetName())
			fmt.Println("      Data type:   \t", v.GetDtype().String())
			fmt.Println("      Tensor shape:\t", v.GetTensorShape().GetDim())
		}
		fmt.Println("  Outputs:")
		for k, v := range val.GetOutputs() {
			fmt.Printf("    %s:\n", k)
			fmt.Println("      Tensor name: \t", v.GetName())
			fmt.Println("      Data type:   \t", v.GetDtype().String())
			fmt.Println("      Tensor shape:\t", v.GetTensorShape().GetDim())
		}
		fmt.Println("_______________")
	}
}
