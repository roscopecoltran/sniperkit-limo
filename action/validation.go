// Copyright © 2017 Makoto Ito
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action
/*
import (
	"errors"																					// go-core
	"fmt"																						// go-core
	"github.com/sirupsen/logrus" 																// logs-logrus
	"github.com/ynqa/word-embedding/utils/fileio" 												// ai-word-embed
	"github.com/ynqa/word-embedding/utils/set" 													// ai-word-embed
	//"github.com/davecgh/go-spew/spew" 														// debug-print
	//"github.com/k0kubun/pp" 																	// debug-print
)

func inputFileIsExist() bool {
	return fileio.FileExists(inputFile)
}

func outputFileIsExist() bool {
	return fileio.FileExists(outputFile)
}

func inputVectorFileIsExist() bool {
	return fileio.FileExists(inputVectorFile)
}

func validateCommonParams() error {
	if dimension <= 0 {
		err := errors.New("Set dimension > 0")
		log.WithError(err).WithFields(
			logrus.Fields{
				"src.file": 			"action/validation.go", 
				"method.prev": 			"if dimension <= 0",
				"method.name": 			"validateCommonParams()", 
				"var.learningRate": 	learningRate, 
				"var.window": 			window, 
				"var.dimension": 		dimension, 
				}).Error("could not validate common parameters, set dimension > 0...")
		return err
	} else if window <= 0 {
		err := errors.New("Set window > 0")
		log.WithError(err).WithFields(
			logrus.Fields{
				"src.file": 			"action/validation.go", 
				"method.prev": 			"if window <= 0",
				"method.name": 			"validateCommonParams()", 
				"var.learningRate": 	learningRate, 
				"var.dimension": 		dimension, 
				"var.window": 			window, 
				}).Error("could not validate common parameters, set window > 0")
		return err 
	} else if learningRate <= 0 {
		err := errors.New("Set learning rate > 0")
		log.WithError(err).WithFields(
			logrus.Fields{
				"src.file": 			"action/validation.go", 
				"method.prev": 			"if learningRate <= 0",
				"method.name": 			"validateCommonParams()", 
				"var.window": 			window, 
				"var.dimension": 		dimension, 
				"var.learningRate": 	learningRate, 
				}).Error("could not validate common parameters, set learning rate > 0")
		return err
	}
	return nil
}

func validateWord2vecParams() error {
	if maxDepth < 0 {
		return errors.New("Set maxDepth >= 0")
	} else if sampleSize <= 0 {
		return errors.New("Set sampleSize > 0")
	}

	if validSubModel := set.New("skip-gram", "cbow"); !validSubModel.Contain(subModel) {
		return fmt.Errorf("Set model from: skip-gram|cbow, instead of %s", subModel)
	}

	if validOptimizer := set.New("ns", "hs"); !validOptimizer.Contain(optimizer) {
		return fmt.Errorf("Set approx from: hs|ns, instead of %s", optimizer)
	}
	return nil
}

*/