import os, io, json
from datetime import date

import pytest
from flask_sqlalchemy import SQLAlchemy

from app import create_app
from models import Record
from utils import db, get_week_range

record1 = {
    "employeeId": "1",
    "payPeriod": {"startDate": "2020-01-01", "endDate": "2020-01-15"},
    "amountPaid": "$300.00",
}

record2 = {
    "employeeId": "1",
    "payPeriod": {"startDate": "2020-01-16", "endDate": "2020-01-31"},
    "amountPaid": "$80.00",
}

record3 = {
    "employeeId": "2",
    "payPeriod": {"startDate": "2020-01-16", "endDate": "2020-01-31"},
    "amountPaid": "$90.00",
}


db = SQLAlchemy(create_app(test=True))


class TestAddFile:
    def setup_method(self, test_method):
        self._clear_table()

    def _clear_table(self):
        # truncate all previous data
        self.conn = db.engine.connect()
        self.conn.execute("TRUNCATE TABLE record")
        self.conn.execute("TRUNCATE TABLE file")

    def test_add_sample_file(self):
        app = create_app(test=True)
        with app.test_client() as client:
            data = {}
            data["file"] = (
                open("fixtures/time-report-1.csv", "rb"),
                "time-report-1.csv",
            )

            res = client.post(
                "/api/upload", data=data, content_type="multipart/form-data"
            )
            assert Record.query.count() == 4

    def teardown_method(self, test_method):
        self.conn.close()


class TestGetSummary:
    def test_get_summary(self):
        app = create_app(test=True)
        with app.test_client() as client:
            res = dict(json.loads(client.get("/api/summary").data))

            assert "payrollReport" in res
            empoloyeeReports = res["payrollReport"]
            assert "employeeReports" in empoloyeeReports
            reports = empoloyeeReports["employeeReports"]

            for record in (record1, record2, record3):
                assert record in reports


class TestReAddFileGetSummary:
    def test_readd_same_file(self):
        app = create_app(test=True)
        with app.test_client() as client:
            data = {}
            data["file"] = (
                open("fixtures/time-report-1.csv", "rb"),
                "time-report-1.csv",
            )
            res = client.post(
                "/api/upload", data=data, content_type="multipart/form-data"
            )

            assert (
                res.data.decode("utf-8")
                == "You have uploaded this file in the past, please upload a new file."
            )

    def test_read_different_file(self):
        app = create_app(test=True)
        with app.test_client() as client:
            data = {}
            data["file"] = (
                open("fixtures/time-report-2.csv", "rb"),
                "time-report-2.csv",
            )
            res = client.post(
                "/api/upload", data=data, content_type="multipart/form-data"
            )

        with app.test_client() as client:
            res = dict(json.loads(client.get("/api/summary").data))
            reports = res["payrollReport"]["employeeReports"]

            # adjust expected response
            record1["amountPaid"] = "$400.00"
            record3["amountPaid"] = "$180.00"
            record4 = {
                "employeeId": "3",
                "payPeriod": {"startDate": "2020-04-01", "endDate": "2020-04-15"},
                "amountPaid": "$100.00",
            }

            for record in (record1, record2, record3, record4):
                assert record in reports


# below is util test
# util is usually function so we use function rather than class to test
testdata = [
    # handle February specially
    ((2020, 2, 20), date(2020, 2, 16), date(2020, 2, 29)),
    ((2014, 2, 20), date(2014, 2, 16), date(2014, 2, 28)),
    ((2021, 1, 10), date(2021, 1, 1), date(2021, 1, 15)),
    ((2014, 9, 30), date(2014, 9, 16), date(2014, 9, 30)),
]


@pytest.mark.parametrize("input_,start_date,end_date", testdata)
def test_get_date_range(input_, start_date, end_date):
    year, month, day = input_
    assert (start_date, end_date) == get_week_range(year, month, day)
