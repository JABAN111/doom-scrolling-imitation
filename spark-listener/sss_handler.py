from minio import Minio
from minio.error import S3Error

# da gvozdami
MINIO_ENDPOINT = "localhost:9000"
MINIO_ACCESS_KEY = "JABA_SUPER_USER_MINIO"
MINIO_SECRET_KEY = "jaba127!368601NO"
BUCKET_NAME = "logs"  
OBJECT_NAME = "log"
LOCAL_FILE = "log.json"

def download_file() -> None:
    client = Minio(
        MINIO_ENDPOINT,
        access_key=MINIO_ACCESS_KEY,
        secret_key=MINIO_SECRET_KEY,
        secure=False
    )
    try:
        if not client.bucket_exists(BUCKET_NAME):
            print(f"Bucket '{BUCKET_NAME}' not exist.")
            return

        client.fget_object(BUCKET_NAME, OBJECT_NAME, LOCAL_FILE)
        print(f"file '{OBJECT_NAME}' saved to '{LOCAL_FILE}'.")

    except S3Error as e:
        print(f"error while working with minio: {e}")

if __name__ == "__main__":
    download_file()
