package sparql

import (
	"bytes"
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Database struct {
	conf config.Config
}

func New(conf config.Config) (db *Database, err error) {
	return &Database{conf: conf}, nil
}

func (*Database) Disconnect() {
	panic("implement me")
}

func (*Database) GetDevice(ctx context.Context, id string) (device model.Device, exists bool, err error) {
	panic("implement me")
}

func (*Database) SetDevice(ctx context.Context, device model.Device) error {
	panic("implement me")
}

func (*Database) RemoveDevice(ctx context.Context, id string) error {
	panic("implement me")
}

func (*Database) GetDeviceType(ctx context.Context, id string) (deviceType model.DeviceType, exists bool, err error) {
	panic("implement me")
}

func (*Database) SetDeviceType(ctx context.Context, deviceType model.DeviceType) error {
	panic("implement me")
}

func (*Database) RemoveDeviceType(ctx context.Context, id string) error {
	panic("implement me")
}

func (this *Database) InsertData(triples string) (success bool, err error) {
	requestBody := []byte(triples)
	resp, err := http.Post(this.conf.RyaUrl+"/web.rya/loadrdf?format=N-Triples", "text/plain", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("ERROR: ", err)
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return true, nil
	} else {
		return false, errors.New("ERROR: Statuscode " + string(resp.StatusCode))
	}
}

func (*Database) ReadData() (body []byte, err error) {
	conf, err := config.Load("../config.json")
	if err != nil {
		log.Println("ERROR: unable to load to config", err)
		return nil, err
	}
	query := url.QueryEscape("construct { ?s ?p ?o.} where { ?s ?p ?o. }")
	resp, err := http.Get(conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}
	return body, nil
}

func (this *Database) GetConstructWithoutProperties(s string, p string, o string) (rdfxml string, err error) {
	query := this.getConstructQueryWithoutProperties(s, p, o)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetConstructWithProperties(subject string) (rdfxml string, err error) {
	query := this.getSubjectWithAllPropertiesQuery(subject)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func getTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
