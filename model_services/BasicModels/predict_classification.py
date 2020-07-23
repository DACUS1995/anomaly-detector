from utils import transform_image
from model_handler import ModelHandler


def get_label(predicted_idx):
	# Need implementation
	pass

def get_classification_prediction(raw_input, type="image"):
	tensor = None
	
	if type == "image":
		tensor = transform_image(image_bytes=raw_input)
	else:
		tensor = torch.tensor([raw_input])
	assert tensor != None

	model = ModelHandler.get_model()
	outputs = model.forward(tensor)
	_, y_hat = outputs.max(1)
	predicted_idx = str(y_hat.item())
	return predicted_idx, get_label(predicted_idx)