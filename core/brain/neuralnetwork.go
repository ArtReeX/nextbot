package brain

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// NeuralNetwork - структура нейронной сети
type NeuralNetwork struct {
	// количество входящих, скрытых и исходыщих узлов
	NInputs, NHiddens, NOutputs int
	// активации узлов
	InputActivations, HiddenActivations, OutputActivations []float64
	// ElmanRNN-контексты
	Contexts [][]float64
	// веса
	InputWeights, OutputWeights [][]float64
	// последнее изменение весов для импульса
	InputChanges, OutputChanges [][]float64
}

/*
Initialize - функция для инициализации нейронной сети.

Параметры:
«inputs» - это количество входов, которое будет иметь нейронная сеть,
«hiddens» - это количество скрытых узлов,
«outputs» - это количество выходов нейронной сети.
*/
func (nn *NeuralNetwork) Initialize(inputs, hiddens, outputs int) {
	nn.NInputs = inputs + 1   // +1 для смещения
	nn.NHiddens = hiddens + 1 // +1 для смещения
	nn.NOutputs = outputs

	nn.InputActivations = vector(nn.NInputs, 1.0)
	nn.HiddenActivations = vector(nn.NHiddens, 1.0)
	nn.OutputActivations = vector(nn.NOutputs, 1.0)

	nn.InputWeights = matrix(nn.NInputs, nn.NHiddens)
	nn.OutputWeights = matrix(nn.NHiddens, nn.NOutputs)

	for i := 0; i < nn.NInputs; i++ {
		for j := 0; j < nn.NHiddens; j++ {
			nn.InputWeights[i][j] = random(-1, 1)
		}
	}

	for i := 0; i < nn.NHiddens; i++ {
		for j := 0; j < nn.NOutputs; j++ {
			nn.OutputWeights[i][j] = random(-1, 1)
		}
	}

	nn.InputChanges = matrix(nn.NInputs, nn.NHiddens)
	nn.OutputChanges = matrix(nn.NHiddens, nn.NOutputs)
}

// Save - функция позволяет сохранить текущее состояние нейронной сети
func (nn *NeuralNetwork) Save() ([]byte, error) {

	data, err := json.Marshal(nn)
	if err != nil {
		return make([]byte, 0), err
	}

	return data, err
}

// Load - функция позволяет загрузить состояние сети в нейронную сеть
func (nn *NeuralNetwork) Load(data []byte) error {

	err := json.Unmarshal(data, nn)
	if err != nil {
		return err
	}

	return nil
}

/*
SetContexts - функция, которая задаёт количество контекстов для добавления в сеть (для создания рекуррентной нейронной сети).

По умолчанию в нейронной сети нет контекста, поэтому это простая сеть для прямой передачи, при добавлении контекстов сеть ведет себя как SRN Elman (Simple Recurrent Network).

Первый параметр (nContexts) используется для указания количества используемых контекстов, второй параметр (initValues) может использоваться для создания настраиваемых инициализированных контекстов.

Если установлено значение «initValues», первый параметр «nContexts» игнорируется и используются контексты, предусмотренные в «initValues».

При использовании «initValues» обратите внимание, что контексты должны иметь одинаковый размер скрытых узлов + 1 (узел смещения).
*/
func (nn *NeuralNetwork) SetContexts(nContexts int, initValues [][]float64) {
	if initValues == nil {
		initValues = make([][]float64, nContexts)

		for i := 0; i < nContexts; i++ {
			initValues[i] = vector(nn.NHiddens, 0.5)
		}
	}

	nn.Contexts = initValues
}

/*
Update - функция используется для активации нейронной сети.
Учитывая массив входов, он возвращает массив, эквивалентный количеству выходов, со значениями от 0 до 1.
*/
func (nn *NeuralNetwork) Update(inputs []float64) []float64 {
	if len(inputs) != nn.NInputs-1 {
		log.Fatal("Error: wrong number of inputs.")
	}

	for i := 0; i < nn.NInputs-1; i++ {
		nn.InputActivations[i] = inputs[i]
	}

	for i := 0; i < nn.NHiddens-1; i++ {
		var sum float64

		for j := 0; j < nn.NInputs; j++ {
			sum += nn.InputActivations[j] * nn.InputWeights[j][i]
		}

		// вычисление сумм контекстов
		for k := 0; k < len(nn.Contexts); k++ {
			for j := 0; j < nn.NHiddens-1; j++ {
				sum += nn.Contexts[k][j]
			}
		}

		nn.HiddenActivations[i] = sigmoid(sum)
	}

	// обновление контекста
	if len(nn.Contexts) > 0 {
		for i := len(nn.Contexts) - 1; i > 0; i-- {
			nn.Contexts[i] = nn.Contexts[i-1]
		}
		nn.Contexts[0] = nn.HiddenActivations
	}

	for i := 0; i < nn.NOutputs; i++ {
		var sum float64
		for j := 0; j < nn.NHiddens; j++ {
			sum += nn.HiddenActivations[j] * nn.OutputWeights[j][i]
		}

		nn.OutputActivations[i] = sigmoid(sum)
	}

	return nn.OutputActivations
}

// BackPropagate - функция используется при обучении нейронной сети, для обратной передачи ошибок из сетевой активации.
func (nn *NeuralNetwork) BackPropagate(targets []float64, lRate, mFactor float64) float64 {
	if len(targets) != nn.NOutputs {
		log.Fatal("Error: wrong number of target values.")
	}

	outputDeltas := vector(nn.NOutputs, 0.0)
	for i := 0; i < nn.NOutputs; i++ {
		outputDeltas[i] = dsigmoid(nn.OutputActivations[i]) * (targets[i] - nn.OutputActivations[i])
	}

	hiddenDeltas := vector(nn.NHiddens, 0.0)
	for i := 0; i < nn.NHiddens; i++ {
		var e float64

		for j := 0; j < nn.NOutputs; j++ {
			e += outputDeltas[j] * nn.OutputWeights[i][j]
		}

		hiddenDeltas[i] = dsigmoid(nn.HiddenActivations[i]) * e
	}

	for i := 0; i < nn.NHiddens; i++ {
		for j := 0; j < nn.NOutputs; j++ {
			change := outputDeltas[j] * nn.HiddenActivations[i]
			nn.OutputWeights[i][j] = nn.OutputWeights[i][j] + lRate*change + mFactor*nn.OutputChanges[i][j]
			nn.OutputChanges[i][j] = change
		}
	}

	for i := 0; i < nn.NInputs; i++ {
		for j := 0; j < nn.NHiddens; j++ {
			change := hiddenDeltas[j] * nn.InputActivations[i]
			nn.InputWeights[i][j] = nn.InputWeights[i][j] + lRate*change + mFactor*nn.InputChanges[i][j]
			nn.InputChanges[i][j] = change
		}
	}

	var e float64

	for i := 0; i < len(targets); i++ {
		e += 0.5 * math.Pow(targets[i]-nn.OutputActivations[i], 2)
	}

	return e
}

// Train - функция используется для обучения нейронной сети, запуская тренировочную операцию N раз и возвращает вычисленные ошибки при обучении.
func (nn *NeuralNetwork) Train(patterns [][][]float64, iterations int, lRate, mFactor float64, debug bool) []float64 {
	errors := make([]float64, iterations)

	for i := 0; i < iterations; i++ {
		var e float64
		for _, p := range patterns {
			nn.Update(p[0])

			tmp := nn.BackPropagate(p[1], lRate, mFactor)
			e += tmp
		}

		errors[i] = e

		if debug && i%1000 == 0 {
			fmt.Println(i, e)
		}
	}

	return errors
}

// Test - функция тестирования
func (nn *NeuralNetwork) Test(patterns [][][]float64) {
	for _, p := range patterns {
		fmt.Println(p[0], "->", nn.Update(p[0]), " : ", p[1])
	}
}
