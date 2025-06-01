from pyspark.sql import *

spark = SparkSession.builder.appName("demo").getOrCreate()

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




# df.where(col("life_stage").isin(["teenager", "adult"])).show()