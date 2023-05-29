
## What's this project designed for ?

It's to answer the question of [wave-payroll-challenges](https://github.com/wvchallenges/se-challenge-payroll)


## How to run the program:

`Prerequiste`:

``` bash
a. *nix environment(Windows users can check dev.sh for commands)

b. docker
```

`Commands`:

``` bash
# build and run the program
> ./dev.sh start

# stop the program
> ./dev.sh stop

# run all tests
> ./dev.sh test

# get summary of rcords
> ./dev.sh summary

# upload a file
> ./dev.sh send $filename
# note that currently filename doesn't support path, so the following command doesn't work
> ./dev.sh send docs/sample.csv
# instead you can use 
> cp docs/sample.csv ./ && ./dev.sh send sample.csv

```

Here is a list of operations that can be used to experience all endpoints:

```bash

# start the server first
> ./dev.sh start

# open another terminal
# 1. run all tests
> ./dev.sh test

# 2. get current summary
> ./dev.sh summary

# 3. try to upload our first file
# There are 2 files inside the folder
> ./dev.sh send time-report-1.csv

# 4. get current summary
> ./dev.sh summary

# 5. try to upload a file with same name
> ./dev.sh send time-report-1.csv

# 6. try to upload the 2nd file
> ./dev.sh send time-report-2.csv

# 7. get current summary
> ./dev.sh summary

```

---


## Q && A:

```markdown
Q1: How did you test that your implementation was correct?
```

> A:
This app uses different configurations in `dev/test` environment(we can add `prod` later). In test environment, we use a separate database with same table structure of dev. <br><br>
> THe unit tests will mock the client and send different requests to backend to test whether we can receive expected response.

```markdown
Q2: If this application was destined for a production environment, what would you add or change?
```
>A:
If we want to put the app into prod, firstly we want to know the amount of our potential customers and `pv/uv` so we can use a proper server configuration.<br><br>
> As for the the architech, there are two useful strategies I have witnessed:<br><br>
>a.  For `WRITE`:<br><br>
We will not parse the file synchronously. The file wil be parsed by acynchronous service like `Celery` so a huge CSV file won't block the user and server resources.<br><br>
>b.  For `READ`:<br><br>
We will use `redis` to store data that will be read high-frequency in advance. In such a case , it means we want to find
the company with most employees that have access with `summary` endpoint and store them in the cache. The app will use cache strategy like `write through` and try to read cache firstly. The app should handle the update of cache when new records are inserted.<br><br>

```bash
Q3: What compromises did you have to make as a result of the time constraints of this challenge?
```

> A: There are many like:<br>
> 1. As doc mentioned, the input is always valid so there is not much handle exception here. If there is more time I should handle edge cases in a more elegant way.<br><br>
> 2. There is some devops issues that shows warnings when building the app. If there is more tiem they should also be fixed

---

## About the app

1. I'm familiar with tech frameworks like `Django/Flask/Gin/Spring`. I choose flask for a quick prototype demo, I choose `MySQL` as main database, `sqlalchemy` for model layer and `marshmallow` for serialization. They can help  build a prototype quickly and maintain the possibility of scaling with the increasing of data.

2. There are 3 tables `record`/`ratio`/`file` in database. Those tables are designed based on the consideration:

> a. This is a `low-write,high-read` scenario.
We try to store all the data used for reading into database/other storage directly
and allow latency when uploading a file and parsing it.

> b. The table should also be extensible and handle new features in the future. For example, we may need to clear all records of a CSV file if users find they upload the wrong file. That's why I add a `sourceId` column in `record`. We may also use other file host services like S3 to store upload files so we create a `file` table.

You can check `init.sql` and `models.py`for more details.

`record` table:

```
 id | date       | hours | employeeId | jobGroup | sourceId | startDate  | endDate    |
+----+------------+-------+------------+----------+----------+------------+------------+
| 32 | 2023-11-14 |   7.5 |          1 | A        | 42       | 2023-11-01 | 2023-11-15 |
```

`ratio` table:

```
| id | category | price |
+----+----------+-------+
|  1 | A        |    20 |
|  2 | B        |    30 |
+----+----------+-------+
```

`file` table:

```
| id | filename          |
+----+-------------------+
|  1 | time-report-1.csv |
|  2 | time-report-2.csv |
+----+-------------------+
```