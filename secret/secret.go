package secret

import (
	"encoding/json"
	"log"
	"onemore/awsgo"
	"onemore/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		log.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)

	return dataSecret, nil
}