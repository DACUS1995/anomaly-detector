from model_handler import ModelHandler
import numpy as np

class_index_map = [
	"nominal",
	"anomaly"
]

def get_label(predicted_idx):
	predicted_labels = np.apply_along_axis(lambda x: class_index_map[x], 0, predicted_idx)
	return predicted_labels

def get_classification_prediction(input):
	input = np.array(input)

	if len(input.shape) == 1:
		input = input.reshape((1, -1))

	model = ModelHandler.get_model()
	predicted_idx = model.predict(input)
	predicted_labels = get_label(outputs)
	return predicted_idx.tolist(), get_label(outputs).tolist()