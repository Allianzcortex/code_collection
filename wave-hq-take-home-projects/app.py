from flask import Flask

from utils import db
from views import payrollview
from config import SQLALCHEMY_DATABASE_URI, DEV_DATABASE, TEST_DATABASE


def create_app(test=False):
    app = Flask(__name__)
    if test:
        database = TEST_DATABASE
    else:
        database = DEV_DATABASE

    app.config["SQLALCHEMY_DATABASE_URI"] = SQLALCHEMY_DATABASE_URI + database
    app.debug = True if test else False

    app.register_blueprint(payrollview)
    db.init_app(app)

    return app


if __name__ == "__main__":
    app = create_app()
    app.run(host="0.0.0.0")
