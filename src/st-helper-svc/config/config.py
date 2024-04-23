import configparser

config = configparser.ConfigParser()
config.read('configs.ini')

OBB_PAT = config['DEFAULT']['OPENBB_PAT']
GROQ_API_KEY = config['DEFAULT']['GROQ_API_KEY']