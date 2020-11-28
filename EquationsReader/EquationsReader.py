import csv
import jsonpickle
from setup_logger import logger
from QueueMessage import QueueMessage

class EquationsReader:
    def __init__(self, csv_path):
        self.csv_path = csv_path

    def print_queue_messages(self):
        possible_equations = get_csv_lines(self.csv_path)
        equations = filter_valid_equations(possible_equations)

        for eq in equations:
            logger.info(create_json_queue_message(eq))

# Get list of equations being read from a given csv.
def get_csv_lines(csv_path):
    possible_equations = []

    with open(csv_path, newline='\r\n') as csvfile:
        lines = csv.reader(csvfile)
        for row in lines:
            possible_equations.append(row[0])

    logger.debug("Fetched %d possible equations", len(possible_equations))
    return possible_equations

# Filters only valid equations.
# Validation is simply based on checking if string contains one of the '+' or '-' operators.
def filter_valid_equations(possible_equations):
    equations = []

    for possible_equation in possible_equations:
        if possible_equation.find("+") != -1 or possible_equation.find("-") != -1:
            equations.append(possible_equation)

    logger.info("Found %d seems to be valid equations", len(equations))
    for equation in equations:
        logger.debug(equation)

    return equations

# Create a json string of QueueMessage object from sent equation string. 
def create_json_queue_message(equation):
    message = QueueMessage(equation)
    return jsonpickle.encode(message, indent=True, unpicklable=False)
