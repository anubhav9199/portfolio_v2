
from backend.django_server.django_server.settings import MONGO_CLIENT


class ConnectWithMongo(object):
    def __init__(self, db_name, collection_name):
        self.db = MONGO_CLIENT.get_database(db_name)
        self.collection = self.db.get_collection(collection_name)

    def get_collection(self):
        return self.collection
