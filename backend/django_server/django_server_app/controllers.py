
from backend.django_server.common_lib.mongodb import portfolio


def get_resume_data_controller():
    _data_map = portfolio.ResumeDataCollection.get_resume_data()

    data = {"job": [], "education": [], "other": []}
    for _data in _data_map:
        if _data["type"] == "job":
            data["job"].append(_data)
        elif _data["type"] == "education":
            data["education"].append(_data)
        else:
            data["other"].append(_data)

    return data
