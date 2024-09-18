import boto3

secretsmanager = boto3.client('secretsmanager',          
    endpoint_url='http://localstack:4566',                
    region_name='ap-northeast-1'
)

def handler():
    res = secretsmanager.create_secret(
        Name='/app/sayhello',
        Description='',
        SecretString='hello',
    )
    print(res)

    # 状態を保持しているっぽい
    res = secretsmanager.list_secrets()
    print(res)
