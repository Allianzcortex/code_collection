from flask_sqlalchemy import SQLAlchemy

from utils import db


class Record(db.Model):
    """
    we store `startDate/endDate` in the table so it's easier
    for us to group by id/weekrange
    """
    __tablename__ = "record"

    id = db.Column(db.Integer, primary_key=True)
    date = db.Column(db.Date, nullable=False)
    hours = db.Column(db.Float, nullable=False)
    employeeId = db.Column(db.Integer, nullable=False)
    jobGroup = db.Column(db.String(1), nullable=False)
    sourceId = db.Column(db.String(12), nullable=False)
    startDate = db.Column(db.Date)
    endDate = db.Column(db.Date)

    # TODO: add __repr__ later


class Ratio(db.Model):
    """
    ratio can be changed at anytime so we create a table for it
    """

    __table__name = "ratio"
    id = db.Column(db.Integer, primary_key=True)
    category = db.Column(db.String(10))
    price = db.Column(db.Integer)


class File(db.Model):
    """
    In theory, filename is always unique so we can also use it as the primary key
    but we still try to avoid it

    In the future, if we use other storage stystem(eg. s3) we may want to add a new
    column about stored URL

    """

    __table__name = "file"
    id = db.Column(db.Integer, primary_key=True)
    filename = db.Column(db.String(100), nullable=False)
