// Package param provides a parser for DNN hyperparameter and image process parameter.
package param

// Hyperparam defines network hyperparameter.
type Hyperparam struct {
	// epoch defines how many iterations should be exeucte
	// for training all data.
	epoch int

	// batch defines how many images will be organised
	// as a batch for forward-propagation.
	batch int

	// subdivisions defines the ratio of dividing a batch into mini-batch.
	//
	// For example: if batch=128, subdivisions=4, that each mini-batch size will be
	// 128/4, which equals to 32.
	subdivisions int

	// height defines the network input's image height.
	//
	// 0 or negative for unlimited (for SPPNet or other flexible input network)
	height int

	// width defines the network input's image height.
	//
	// 0 or negative for unlimited (for SPPNet or other flexible input network)
	width int

	// channels defines the input image's channel number. Many traditional
	// object detection/recognition neural network only accept the image
	// which has 1 channel or 3 channels(RGB).
	//
	// So the channels var will have only two posibilities, 1 or 3.
	channels int

	// momentum defines the momentum of changing the learning_rate_delta while back-
	// propagation process.
	//
	// For more detail, please read this paper:
	// https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.57.5612&rep=rep1&type=pdf
	momentum float64

	// decay defines the learning rate decay rate, decay on learning rate will be applied
	// after each 100 iterations.
	//
	// 0 or negative means decay strategy will not be applied on learning rate.
	decay float64

	// learningRate defines the ratio of changing the network paramater on back-propagation.
	learningRate float64

	// coreNum defines how many CPU cores will be used for transfer image from disk to memory.
	//
	// Avaliable on PyTorch-based docker image (Equals to num_workers parameter).
	//
	// Cannot exceed the maximum number of CPU cores of target device.
	coreNum int
}

// HyperparamParser is a hyperparameter parser interface
type HyperparamParser interface {
	ParseHyperParam(hyperparamFile string) (*Hyperparam, error)
}

// GetEpoch returns the epoch member.
func (hp *Hyperparam) GetEpoch() int {
	return hp.epoch
}

// GetBatch returns the batch member.
func (hp *Hyperparam) GetBatch() int {
	return hp.batch
}

// GetSubDivision returns the subdivision member.
func (hp *Hyperparam) GetSubDivision() int {
	return hp.subdivisions
}

// GetHeight returns the height member.
func (hp *Hyperparam) GetHeight() int {
	return hp.height
}

// GetWidth returns the width member.
func (hp *Hyperparam) GetWidth() int {
	return hp.width
}

// GetChannel returns the channels member.
func (hp *Hyperparam) GetChannel() int {
	return hp.channels
}

// GetMomentum returns the momentum member.
func (hp *Hyperparam) GetMomentum() float64 {
	return hp.momentum
}

// GetDecay returns the decay member.
func (hp *Hyperparam) GetDecay() float64 {
	return hp.decay
}

// GetLR returns the learningRate member.
func (hp *Hyperparam) GetLR() float64 {
	return hp.learningRate
}

// GetCoreNum returns the coreNum member.
func (hp *Hyperparam) GetCoreNum() int {
	return hp.coreNum
}
