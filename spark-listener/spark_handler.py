from pyspark.sql import *
from pyspark.sql.functions import col

PATH = "./log.json"

spark = SparkSession.builder.appName("demo").getOrCreate()
sc = spark.sparkContext

logsDF = spark.read \
    .option("multiline", "true") \
    .json("./log.json")

logsDF.printSchema()

logsDF.createOrReplaceTempView("people")


logsDF.select(
    col("time"),
    col("level"),
    col("msg"),
    col("user.username").alias("username"),
    col("user.email").alias("email"),
    col("user.bio").alias("bio"),
    col("user.age").alias("возраст")
).show(truncate=False)


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
