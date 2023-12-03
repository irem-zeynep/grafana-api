package secretmanager

import (
	"encoding/json"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

type Secret struct {
	GrafanaClient ClientSecret
	ErrorTopic    TopicSecret
}

type ClientSecret struct {
	Host   string `json:"host"`
	APIKey string `json:"apiKey"`
}

type TopicSecret struct {
	TopicName string `json:"topicName"`
}

const (
	grafanaClientSecretName = "grafana/stage/grafanaclient"
	timeStreamDBSecretName  = "grafana/stage/timestreamdb"
	errorTopicSecretName    = "grafana/stage/errortopic"
)

type secretManager struct {
	cache *secretcache.Cache
}

func Init() *Secret {
	secretCache, err := secretcache.New()
	if err != nil {
		panic(err)
	}

	manager := secretManager{
		cache: secretCache,
	}

	clientSecret := ClientSecret{}
	manager.mapSecret(grafanaClientSecretName, &clientSecret)

	errorTopicSecret := TopicSecret{}
	manager.mapSecret(errorTopicSecretName, &errorTopicSecret)

	return &Secret{
		GrafanaClient: clientSecret,
		ErrorTopic:    errorTopicSecret,
	}
}

func (m *secretManager) mapSecret(secretId string, secret any) {
	secretsStr, err := m.cache.GetSecretString(secretId)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal([]byte(secretsStr), &secret); err != nil {
		panic(err)
	}
}
