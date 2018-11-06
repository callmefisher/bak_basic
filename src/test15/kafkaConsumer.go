package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/qiniu/log.v1"
	"io/ioutil"
	"strings"
	"sync"
	"util"
)

var (
	brokers       = flag.String("brokers", "127.0.0.1:9092,127.0.0.1:9093", "The Kafka brokers to connect to brokers")
	certFile      = flag.String("certificate", "", "The optional certificate file for client authentication")
	keyFile       = flag.String("key", "", "The optional key file for client authentication")
	caFile        = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
	verifySsl     = flag.Bool("verify", false, "Optional verify ssl certificates chain")
	topic         = flag.String("topic", "", "consumer which topic")
	consumerGroup = flag.String("cg", "", "consumer group")
)

func init() {
	flag.Parse()
	if *topic == "" || *consumerGroup == "" {
		log.Fatal("please input topic and consumer group")
	}
	log.Info("consumer init finished")
}

func _joinGroup(clientId string) error {
	config := sarama.NewConfig()
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}

	config.Version = sarama.V0_10_1_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Return.Errors = true
	config.ClientID = clientId

	//log.Info("brokers1:", strings.Split(*brokers, ","))

	client, err := sarama.NewClient(strings.Split(*brokers, ","), config)
	if err != nil {
		log.Error("alloc kafka client fail:", err)
		return err
	}
	//log.Info("brokers2:", strings.Split(*brokers, ","))

	req := &sarama.JoinGroupRequest{
		Version:        1,
		GroupId:        *consumerGroup,
		ProtocolType:   "consumer",
		SessionTimeout: 10000,
	}

	if err := req.AddGroupProtocolMetadata("range", &sarama.ConsumerGroupMemberMetadata{
		Topics: strings.Split(*topic, ","),
	}); err != nil {
		log.Error("err1:", err)
		return err
	}

	broker, err := client.Coordinator(*consumerGroup)
	if err != nil {
		log.Error("err2:", err)
		return err
	}
	resp, err := broker.JoinGroup(req)
	if err != nil {
		log.Error("err3:", err)
		return err
	} else if resp.Err != sarama.ErrNoError {
		log.Error("err4:", resp.Err)
		return resp.Err
	}

	log.Info("Consumer:", clientId, " joined group:", *consumerGroup, " memberId:", resp.MemberId, " genId:",
		resp.GenerationId, " len:", len(resp.Members))

	return nil
}

func newConsumer(clientId string) error {

	config := sarama.NewConfig()
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}

	config.Version = sarama.V2_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Return.Errors = true
	config.ClientID = clientId

	// Start with a client
	client, err := sarama.NewClient(strings.Split(*brokers, ","), config)
	if err != nil {
		log.Error("error1:", err)
		return err
	}
	defer func() { client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(*consumerGroup, client)
	if err != nil {
		log.Error("error2:", err)
		return err
	}

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := strings.Split(*topic, ",")
		handler := exampleConsumerGroupHandler{}
		log.Info("consumer:", clientId, " topics", topics)
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var s2 util.AccessLogEntry
		json.Unmarshal(msg.Value, &s2)
		log.Info(msg.Topic, " partition:", msg.Partition, "offset:", msg.Offset, " value:", s2)

		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {

	var sync sync.WaitGroup
	sync.Add(2)
	go newConsumer("clientId1")
	go newConsumer("clientId2")
	sync.Wait()
}

func createTlsConfiguration() (t *tls.Config) {
	if *certFile != "" && *keyFile != "" && *caFile != "" {
		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
		if err != nil {
			log.Fatal(err)
		}

		caCert, err := ioutil.ReadFile(*caFile)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		t = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: *verifySsl,
		}
	}
	// will be nil by default if nothing is provided
	return t
}
