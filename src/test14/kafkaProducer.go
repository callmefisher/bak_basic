//MIT License
//
//Copyright (c) 2018 XiaYanji
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/Shopify/sarama"
	"github.com/qiniu/log.v1"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"util"
)

var (
	addr      = flag.String("addr", ":8080", "The address to bind to")
	brokers   = flag.String("brokers", "127.0.0.1:9092,127.0.0.1:9093", "The Kafka brokers to connect to, as a comma separated list")
	verbose   = flag.Bool("verbose", false, "Turn on Sarama logging")
	certFile  = flag.String("certificate", "", "The optional certificate file for client authentication")
	keyFile   = flag.String("key", "", "The optional key file for client authentication")
	caFile    = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
	verifySsl = flag.Bool("verify", false, "Optional verify ssl certificates chain")
)

type Server struct {
	SyncProducer sarama.SyncProducer
	AsyProducer  sarama.AsyncProducer
}

var server *Server

func init() {
	log.Info("init")

	flag.Parse()

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	server = &Server{
		SyncProducer: newSyncProducer(brokerList),
		AsyProducer:  newAsyProducer(brokerList),
	}

}

func (s *Server) ProduceSynMessage(value string) {
	partition, offset, err := s.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "topic1",
		Value: sarama.StringEncoder(value),
	})

	if err != nil {
		log.Errorf("Failed to  topic1 sync:, %s", err)
	} else {
		// The tuple (topic, partition, offset) can be used as a unique identifier
		// for a message in a Kafka cluster.
		log.Infof("topic1 /%d/%d   sync", partition, offset)
	}
}

func (s *Server) ProduceAsyMessage(value string) {

	started := time.Now()
	entry := &util.AccessLogEntry{
		Method:       value,
		Host:         "host",
		Path:         "/a/b",
		IP:           "10.0.26.3",
		ResponseTime: float64(started.Unix()),
	}

	// We will use the client's IP address as key. This will cause
	// all the access log entries of the same IP address to end up
	// on the same partition.

	s.AsyProducer.Input() <- &sarama.ProducerMessage{
		Topic: "topic2",
		Key:   sarama.StringEncoder(value),
		Value: entry,
	}
	log.Info(" topic2 async")
}

func (s *Server) Close() error {
	if err := s.SyncProducer.Close(); err != nil {
		log.Println("Failed to shut down syn producer cleanly", err)
	}

	if err := s.AsyProducer.Close(); err != nil {
		log.Println("Failed to shut down access log producer cleanly", err)
	}
	log.Info("producer close")
	return nil
}

func main() {

	defer func() {
		if err := server.Close(); err != nil {
			log.Println("Failed to close server", err)
		}
	}()

	if server == nil {
		log.Fatal("producer init fail")
	}

	var synCount sync.WaitGroup
	synCount.Add(2)
	go func() {

		defer func() {
			synCount.Done()
		}()
		for i := 0; i < 10000; i++ {
			value := " topic1 value" + strconv.Itoa(i)
			server.ProduceSynMessage(value)
		}

	}()

	go func() {

		defer func() {
			synCount.Done()
		}()

		for i := 0; i < 10000; i++ {
			value := " topic2 value" + strconv.Itoa(i)
			server.ProduceAsyMessage(value)
		}

	}()

	synCount.Wait()

	log.Info("hello world!")
}

func newSyncProducer(brokerList []string) sarama.SyncProducer {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama sync producer:", err)
	}

	return producer
}

func newAsyProducer(brokerList []string) sarama.AsyncProducer {

	config := sarama.NewConfig()
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	config.Producer.Flush.Messages = 100
	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama asy producer:", err)
	}
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
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
