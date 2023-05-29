# We should set those variables in os.environment[''] or in
# other cloud environement(eg rancher)
# But let's make it simpler here

DEV_DATABASE = "Payroll"
TEST_DATABASE = "PayrollTest"

MYSQL_USER = "root"
MYSQL_PASSWORD = "root"
SQLALCHEMY_DATABASE_URI = f"mysql://{MYSQL_USER}:{MYSQL_PASSWORD}@db/"

CSV_FOLDER = "./static"
