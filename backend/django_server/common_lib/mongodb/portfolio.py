from backend.django_server.common_lib import mongo_params
from backend.django_server.common_lib.mongo_utils import clean_data
from backend.django_server.common_lib.mongodb import ConnectWithMongo


class ResumeDataCollection:
    def __init__(self) -> None:
        self._db_name = mongo_params.PORTFOLIO_DB
        self._collection_name = mongo_params.RESUME_DATA_COLLECTION
        self._collection = ConnectWithMongo(self._db_name, self._collection_name).get_collection()

    @classmethod
    @clean_data
    def get_resume_data(cls):
        _db = ResumeDataCollection()
        collection = _db._collection
        cursor = collection.find({})
        return cursor
