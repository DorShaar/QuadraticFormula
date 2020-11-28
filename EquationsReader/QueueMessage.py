import uuid
from json import JSONEncoder

class QueueMessage:
	def __init__(self, equation):
		self.correlation_id = uuid.uuid4()
		self.equation = equation