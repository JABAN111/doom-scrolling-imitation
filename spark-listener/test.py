from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType, TimestampType, IntegerType, MapType
from pyspark.sql.functions import col, from_json

# 1. Определяем схему (schema) для вашего JSON
#    Она должна отражать поля, которые есть в массиве объектов вашего log.json
log_schema = StructType([
    StructField("time", TimestampType(), True),
    StructField("level", StringType(), True),
    StructField("msg", StringType(), True),
    # Поле "user" может быть либо вложенным объектом, либо отсутствовать, 
    # поэтому здесь сделаем его MapType, а потом разберём, что внутри
    StructField("user", 
                StructType([
                    StructField("username", StringType(), True),
                    StructField("email", StringType(), True),
                    StructField("bio", StringType(), True),
                    StructField("age", IntegerType(), True),
                ]), 
                True),
    # Некоторые события (например, "follow user", "like post") имеют другие поля на верхнем уровне:
    # добавим опциональные поля, которые встречаются в вашем примере.
    StructField("username", StringType(), True),
    StructField("usernameToFollow", StringType(), True),
    StructField("post", 
                StructType([
                    StructField("id", StringType(), True),
                    StructField("user_id", StringType(), True),
                    StructField("caption", StringType(), True),
                    StructField("created_at", TimestampType(), True),
                ]), 
                True),
    StructField("userID", StringType(), True),
    StructField("postID", StringType(), True),
])

# 2. Создаём SparkSession
spark = SparkSession.builder \
    .appName("StreamingDemo") \
    .getOrCreate()

# 3. Указываем директорию, куда по мере поступления будут падать новые JSON-файлы.
#    Например: ./logs_stream/*.json — все новые файлы в папке logs_stream будут обрабатываться.
streaming_input_path = "log.json"

# 4. Читаем «поток» JSON-объектов из папки (каждый новый файл = новый micro-batch)
logsStreamingDF = (
    spark.readStream
         .schema(log_schema)               # обязательно указываем заранее известную схему
         .option("multiline", "true")      # если ваши входные JSON-файлы — многострочные
         .json(streaming_input_path)       # директория, где появляются файлы
)

# 5. Выполняем необходимые преобразования: выбираем поля, даём алиасы и т.д.
parsedDF = logsStreamingDF.select(
    col("time"),
    col("level"),
    col("msg"),
    col("user.username").alias("user_username"),
    col("user.email").alias("user_email"),
    col("user.bio").alias("user_bio"),
    col("user.age").alias("user_age"),
    col("username"),            # для событий «follow user»
    col("usernameToFollow"),    # для «follow user»
    col("post.id").alias("post_id"),
    col("post.user_id").alias("post_user_id"),
    col("post.caption").alias("post_caption"),
    col("post.created_at").alias("post_created_at"),
    col("userID"),              # для «like post»
    col("postID")               # для «like post»
)

# 6. Подключаем к «выводу» — здесь для простоты будем писать всё в консоль
#    (можно писать в Kafka, Parquet, другое хранилище, базу и т.д.)
query = (
    parsedDF.writeStream
            .outputMode("append")           # «append» — добавляем новые строки
            .format("console")              # выводим прямо в stdout (консоль)
            .option("truncate", "false")    # не урезаем длинные строки в выводе
            .start()                        # запускаем стриминг
)

# 7. После старта микробэтчей — ждём, когда стриминг завершится (обычно Terminate по сигналу)
query.awaitTermination()
