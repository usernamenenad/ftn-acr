import logging


class Logger(object):
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            logging.basicConfig(
                level=logging.INFO,
                format="%(asctime)s %(levelname)s %(message)s",
                datefmt="%Y-%m-%d %H:%M:%S",
            )
            cls._instance = super(Logger, cls).__new__(cls, *args, **kwargs)

        return cls._instance

    def log(self, message: str) -> None:
        logging.info(message)
