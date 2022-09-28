package ins

import (
	"errors"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func GetMQTTClient(config *MQTTConfigurations) (MQTT.Client, error) {
	if config == nil {
		return nil, errors.New("MQTTConfigurations is nil")
	}
	opts := MQTT.NewClientOptions()
	opts.AddBroker(config.Broker)
	opts.SetClientID(config.ClientId)
	if 0 < len(config.User) {
		opts.SetUsername(config.User)
		if 0 < len(config.Password) {
			opts.SetPassword(config.Password)
		}
	}

	tlsconfig := NewTLSConfig(&config.Cacertfile, &config.Certfile, &config.Keyfile)
	opts.SetTLSConfig(tlsconfig)

	opts.SetCleanSession(config.Cleansess)

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		err := token.Error()

		return nil, err
	}

	return client, nil
}
