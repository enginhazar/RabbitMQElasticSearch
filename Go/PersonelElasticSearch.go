package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

type Personel struct {
	Sicilno    int32
	TcKimlikNo int64
	Ad         string
	Soyad      string
	Adres      string
	DogumTarih Tarih
}

type Tarih struct {
	time.Time
}

var elastic *elasticsearch.Client

func main() {
	client, err := newElasticClient()

	elastic = client

	if err != nil {
		fmt.Println("error connection")
		log.Fatalln("connection error")
		log.Fatalln(err)
		return
	}
	createIndex("personel")

	rabbitMqConsume()

}

func createIndex(index string) error {
	var indexAlias = index + "_alias"

	res, _err := elastic.Indices.Exists([]string{index})

	if _err != nil {
		log.Fatalln(_err)
		return _err
	}

	if res.StatusCode == 200 {
		return nil
	}

	if res.StatusCode == 400 {
		return fmt.Errorf("error Indices : %s", res.String())
	}

	res, _err = elastic.Indices.Create(index)
	if _err != nil {
		return fmt.Errorf("error create : %s", res.String())
	}

	if res.IsError() {
		return fmt.Errorf("İndex Create error %s", res.String())
	}

	res, _err = elastic.Indices.PutAlias([]string{index}, indexAlias)

	if _err != nil {
		log.Fatalln(_err)
		return _err
	}
	if res.IsError() {
		return fmt.Errorf("İndex Alias error %s", res.String())
	}
	fmt.Println("Index oluşturuldu")
	return nil
}

func rabbitMqConsume() {
	conn, err := amqp.Dial("amqp://admin:123456@localhost/")

	//amqp://guest:guest@localhost:5672/
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer ch.Close()

	// Kuyruk tanımlandı
	_, err = ch.QueueDeclare("Personel", false, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
	msgs, err := ch.Consume("Personel", "", true, false, false, false, nil)

	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)
	go func() {
		layout := `"2006-01-02T15:04:05"`
		for d := range msgs {
			var goPersonel Personel
			jsonveri := []byte(d.Body)
			err := json.Unmarshal(jsonveri, &goPersonel)
			if err != nil {
				log.Fatalln(err)
				log.Fatalln("Fatal Error")

			}
			//fmt.Println(gosicil.sicilno)
			fmt.Printf("Sicilno      : %d\n", goPersonel.Sicilno)
			fmt.Printf("TcKimlikNo   : %+v\n", goPersonel.TcKimlikNo)
			fmt.Printf("Adı          : %s\n", goPersonel.Ad)
			fmt.Printf("Soyadi       : %s\n", goPersonel.Soyad)
			fmt.Printf("Adres        : %s\n", goPersonel.Adres)
			fmt.Printf("Dogum Tarihi : %s\n", goPersonel.DogumTarih.Format(layout))

			//	fmt.Println("---------------------\n")

			addPersonelElastic(goPersonel)

		}
	}()

	log.Printf("Kuyruk Dinleniyor....")
	<-forever

}

// UnmarshalJSON / Unmarshall (decode) işlemi sırasında Tarih alanı formatı C# 'tan yazılan formata uygun şekilde parse edilmesi
// / için yazıldı
func (t *Tarih) UnmarshalJSON(b []byte) (err error) {
	layout := `"2006-01-02T15:04:05"`

	date, err := time.Parse(layout, string(b))
	if err != nil {
		log.Println(err)
		return err
	}
	t.Time = date
	return
}

func addPersonelElastic(personel Personel) {
	data, _err := json.Marshal(personel)
	var index = "personel"
	var indexAlias = index + "_alias"

	if _err != nil {
		log.Fatalln(_err)
		log.Fatalln("Fatal Error")
	}

	//res, err := elastic.Index(
	//	indexAlias,                      // Dokümanın ekleneceği indeks
	//	strings.NewReader(string(data)), // Doküman verisi
	//	elastic.Index.WithContext(context.Background()),
	//)

	req := esapi.IndexRequest{Index: indexAlias, DocumentID: string(personel.TcKimlikNo), Body: strings.NewReader(string(data))}
	res, err := req.Do(context.Background(), elastic.Transport)

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Fatal Error", err)
	}

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Fatal Error", err)
	}
	println(res.Body)
	res.Body.Close()

	if res.StatusCode == 409 {
		fmt.Println("status code 409")
		log.Fatalln("Fatal Error", err)

	}
	fmt.Printf("ElasticSearch Append %d", personel.TcKimlikNo)

}

func newElasticClient() (*elasticsearch.Client, error) {
	fmt.Printf("new client")
	cfg := elasticsearch.Config{

		Addresses: []string{"http://localhost:9200/"},
		Username:  "elastic",
		Password:  "123456",
		Logger:    &elastictransport.ColorLogger{Output: os.Stdout},
		//Logger:    &estransport.ColorLogger{os.Stdout, true, true},
	}
	fmt.Println("config")
	eclient, _err := elasticsearch.NewClient(cfg)

	if _err != nil {
		log.Fatalln("err")
		return nil, _err
	}

	return eclient, nil
}
