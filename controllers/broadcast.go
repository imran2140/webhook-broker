package controllers

import (
	"errors"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/imyousuf/webhook-broker/dispatcher"
	"github.com/imyousuf/webhook-broker/storage"
	"github.com/imyousuf/webhook-broker/storage/data"
	"github.com/julienschmidt/httprouter"
)

const (
	broadcastPath             = channelPath + "/broadcast"
	headerPriority            = "X-Broker-Message-Priority"
	headerChannelToken        = "X-Broker-Channel-Token"
	headerProducerToken       = "X-Broker-Producer-Token"
	headerProducerID          = "X-Broker-Producer-ID"
	defaultMessageContentType = "application/octet-stream"
)

var (
	errChannelTokenNotMatching  = errors.New("channel token does not match")
	errProducerTokenNotMatching = errors.New("producer token does not match")
	errProducerDoesNotExist     = errors.New("producer could not be found")
	errBodyCouldNotBeRead       = errors.New("body could not be read")
)

// BroadcastController receives new Message to broadcasted to a valid channel
type BroadcastController struct {
	MessageRepository  storage.MessageRepository
	ChannelRepository  storage.ChannelRepository
	ProducerRepository storage.ProducerRepository
	Dispatcher         dispatcher.MessageDispatcher
}

// NewBroadcastController creates a new instance of the controller responsible for broadcasting a message
func NewBroadcastController(channelRepo storage.ChannelRepository, msgRepo storage.MessageRepository, producerRepo storage.ProducerRepository, dispatcher dispatcher.MessageDispatcher) *BroadcastController {
	return &BroadcastController{ChannelRepository: channelRepo, MessageRepository: msgRepo, ProducerRepository: producerRepo, Dispatcher: dispatcher}
}

// Post Receives message to be broadcasted to a channel
func (broadcastController *BroadcastController) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	valid := true
	var producer *data.Producer
	channelID := params.ByName(channelIDPathParamKey)
	channel, err := broadcastController.ChannelRepository.Get(channelID)
	if err != nil {
		writeNotFound(w)
		valid = false
	} else if channel.Token != r.Header.Get(headerChannelToken) {
		writeStatus(w, http.StatusForbidden, errChannelTokenNotMatching)
		valid = false
	} else if producer, err = broadcastController.ProducerRepository.Get(r.Header.Get(headerProducerID)); err != nil {
		writeStatus(w, http.StatusUnauthorized, errProducerDoesNotExist)
		valid = false
	} else if producer.Token != r.Header.Get(headerProducerToken) {
		writeStatus(w, http.StatusForbidden, errProducerTokenNotMatching)
		valid = false
	}
	if !valid {
		return
	}
	contentType := r.Header.Get(headerContentType)
	if len(contentType) < 1 {
		contentType = defaultMessageContentType
	}
	priority, err := strconv.Atoi(r.Header.Get(headerPriority))
	if err != nil {
		priority = 0
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading body", err)
		writeErr(w, errBodyCouldNotBeRead)
		return
	}
	message, _ := data.NewMessage(channel, producer, string(body), contentType)
	message.Priority = uint(math.Abs(float64(priority)))
	if err = broadcastController.MessageRepository.Create(message); err == nil {
		go broadcastController.Dispatcher.Dispatch(message)
		writeStatus(w, http.StatusAccepted, nil)
	} else {
		writeErr(w, err)
	}

}

// GetPath returns the endpoint's path
func (broadcastController *BroadcastController) GetPath() string {
	return broadcastPath
}

// FormatAsRelativeLink Format as relative URL of this resource based on the params
func (broadcastController *BroadcastController) FormatAsRelativeLink(params ...httprouter.Param) string {
	return formatURL(params, broadcastPath, channelIDPathParamKey)
}