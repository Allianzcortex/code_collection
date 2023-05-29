import os, csv
from datetime import date

from flask import Blueprint, request, jsonify
from schemas import SummarySchema
from models import Record, File
from utils import db, get_week_range
from config import CSV_FOLDER

payrollview = Blueprint("payroll", __name__)
summary_schema = SummarySchema(many=True)


@payrollview.route("/api/upload", methods=["POST"])
def handle_upload_csv():
    if "file" not in request.files:
        return "Please attach the file you want to upload", 400
    file_ = request.files["file"]
    filename = file_.filename
    if db.session.query(File.query.filter(File.filename == filename).exists()).scalar():
        return "You have uploaded this file in the past, please upload a new file.", 400
    else:
        db.session.add(File(filename=filename))
        db.session.commit()
    path = os.path.join(CSV_FOLDER, filename)
    file_.save(path)
    save_csv(path)

    return jsonify("The CSV file has been uploaded successfully.")


def save_csv(file_):
    with open(file_) as input_:
        reader = csv.reader(
            input_,
        )
        # a bit tricky here
        sourceId = file_.split("-")[-1][:-4]
        # skip header
        next(reader)
        for row in reader:
            date_, hours, id_, group = row
            day, month, year = [int(item) for item in date_.split("/")]
            start_date, end_date = get_week_range(year, month, day)

            # TODO : another way is to add all records in a list and use bulk_save_objects
            # but let's do it later
            db.session.add(
                Record(
                    date=date(year, month, day),
                    hours=float(hours),
                    employeeId=int(id_),
                    jobGroup=group,
                    sourceId=sourceId,
                    startDate=start_date,
                    endDate=end_date,
                )
            )
        db.session.commit()


@payrollview.route("/api/summary", methods=["GET"])
def get_summary():
    """
    It's open for discussion what's the best practice here, there're 2 options:

    1. Use ORM to retrieve data from tables and join them via code
    2. Execute raw SQL query for data

    I choose 2 for the following 2 reasons:

    1. The SQL is not so complex, just join 2 tables together so it's intuitive
    2. Read is a high-frequency action so the endpoint will be hit more. If there is
       bottleneck in the future then the issue can be addressed easier(eg choose a
       distributed database / add Joint Index)

    """

    query = """
    select A.employeeId,A.startDate,A.endDate, A.x*B.price as amount
    from
    (select employeeId,startDate,endDate,jobGroup,sum(hours) as x,count(*) 
    from record group by employeeId,startDate,endDate,jobGroup) as A
    join 
    (select category,price from ratio) as B
    on A.jobGroup = B.category
    order by A.employeeId,A.startDate;
    """

    records = db.session.execute(query)
    results = [dict(record) for record in records.mappings().all()]

    return {"payrollReport": summary_schema.dump(results)}
