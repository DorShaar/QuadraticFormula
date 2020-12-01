import logging

logging.basicConfig(filename='app.log',
    filemode='a',
    format='%(asctime)s %(process)d - %(levelname)-8s %(message)s',
    datefmt='%d-%m-%Y %H:%M:%S')

logger = logging.getLogger('equation_reader')
logger.setLevel(logging.DEBUG)