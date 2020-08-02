from sklearn.tree import DecisionTreeClassifier
from enum import Enum
from joblib import dump, load


class ModelTypes(Enum):
	KNN = 1


class ModelHandler:
	__model = None
	@staticmethod 
	def get_model(model_type):
		if ModelHandler.__model == None:
			ModelHandler(model_type)
		return ModelHandler.__model

	def __init__(self, model_type):
		""" Virtually private constructor. """
		if ModelHandler.__model != None:
			raise Exception("This class is a singleton!")
		else:
			model = None

			if model_type == ModelTypes.KNN:
				model = load("./saved_models/KNN.joblib")
			else:
				raise Exception("Unknown model type requested")
			
			assert model is not None
			ModelHandler.__model = model
			