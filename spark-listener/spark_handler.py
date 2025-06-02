from pyspark.sql import *
from pyspark.sql.functions import col
from sss_handler import download_file

PATH = "./log.json"


download_file()
spark = SparkSession.builder.appName("demo").getOrCreate()
sc = spark.sparkContext

logsDF = spark.read \
    .option("multiline", "true") \
    .json("./log.json")

logsDF.printSchema()

logsDF.createOrReplaceTempView("people")



def peopleNames():
    res = spark.sql("SELECT user.username FROM people WHERE user.age BETWEEN 13 AND 30")
    return res

def countPeople():
    return spark.sql("SELECT user.username FROM people WHERE user.age BETWEEN 13 AND 30").count()

peopleNames().show()


def pupu():
    df = spark.createDataFrame(
        [
            ("sue", 32),
            ("li", 3),
            ("bob", 75),
            ("heo", 13),
        ],
        ["first_name", "age"],
    )
    return df.show()

