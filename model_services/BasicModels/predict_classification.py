from model_handler import ModelHandler
import numpy as np

from config import Config

class_index_map = [
	"nominal",
	"anomaly"
]

def _get_label(predicted_idx):
	predicted_labels = map(lambda x: class_index_map[x], predicted_idx)
	return list(predicted_labels)

def get_classification_prediction(input):
	input = np.array(input)

	if len(input.shape) == 1:
		input = input.reshape((1, -1))

	model = ModelHandler.get_model(Config.model_type)
	predicted_idx = model.predict(input)
	predicted_labels = _get_label(predicted_idx.tolist())
	return predicted_idx.tolist(), predicted_labels

def train_classifier():
	raise Exception("Not implemented")