from marshmallow import Schema, fields, pre_dump, post_dump


class DateRangeSchema(Schema):
    startDate = fields.Date()
    endDate = fields.Date()


class SummarySchema(Schema):

    """
    The expected format is 3-layers : payrollReport->employeeReport->single report

    We can add fields.Nested() to render 3 layers as below

    We can also use `post_dump` to add new key for json

    I use the 2nd way here
    """

    employeeId = fields.Str()
    payPeriod = fields.Nested(DateRangeSchema())
    amountPaid = fields.Str()

    @pre_dump
    def handleDateRange(self, data, **kwargs):
        # handle scheme change
        # Firstly, combine startDate and endDate into payPeriod
        keys = ["startDate", "endDate"]
        data["payPeriod"] = {key: data.pop(key) for key in keys}

        # Secondly, convert amount to string and add $ prefix
        data["amountPaid"] = f"${format(data['amount'],'.2f')}"
        return data

    @post_dump(pass_many=True)
    def addNewField(self, data, many, **kwargs):
        return {"employeeReports": data}
