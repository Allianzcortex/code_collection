import calendar
from datetime import date

from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()


def get_week_range(year,month,day):
    """
    Given year/month/day format),
    return the the range biweekly in datetime.date format.
    
    e.g. 2020.01.01 -> week range: 2020.01.01-2020.01.15
         2020.01.31 -> week range: 2020.01.16-2020.01.31
    """

    if 1 <= day <= 15:
        start_day, end_day = 1, 15
    else:
        start_day = 16
        end_day = calendar.monthrange(year, month)[-1]

    return date(year, month, start_day), date(year, month, end_day)
