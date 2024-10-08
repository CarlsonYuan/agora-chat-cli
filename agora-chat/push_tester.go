/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"time"

	"github.com/ycj3/agora-chat-cli/http"
)

type PushManager struct {
	client *client
}

type fcmErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Details []struct {
		ErrorCode string `json:"errorCode"`
	}
}
type fcmResponse struct {
	Name string `json:"name,omitempty"`
}

type FcmData struct {
	fcmResponse
	FcmError *fcmErrorResponse `json:"error,omitempty"`
}

type PushNotification struct {
	Payload    string    `json:"payload"`
	Topic      string    `json:"topic"`
	Expiration time.Time `json:"expiration"`
	Priority   string    `json:"priority"`
	Token      string    `json:"token"`
}

type apnsResponse struct {
	TokenInvalidationTimestamp interface{}      `json:"tokenInvalidationTimestamp"`
	ApnsUniqueId               string           `json:"apnsUniqueId"`
	Accepted                   bool             `json:"accepted"`
	ApnsId                     string           `json:"apnsId"`
	RejectionReason            interface{}      `json:"rejectionReason"`
	PushNotification           PushNotification `json:"pushNotification"`
	StatusCode                 int              `json:"statusCode"`
}

type PushResultData struct {
	FcmData
	apnsResponse
}

type PushResult struct {
	PushStatus string         `json:"pushStatus"`
	Data       PushResultData `json:"data,omitempty"` // contains the response from the provider you are useing (e.g. FCM or APNs)
	Desc       string         `json:"desc,omitempty"`
	StatusCode int            `json:"statusCode,omitempty"`
}

type PushResponseResult struct {
	Response
	Data []PushResult `json:"data"`
}

type pushStrategy int

const (
	PushPrividerFirstThenAgoraChat pushStrategy = iota // 0
	OnlyAgoraChat                                      // 1
	OnlyPushPrivider                                   // 2 (Default)
	AgoraFirstThenPushPrivider                         // 3
	OnlyOnlineAgoraChat                                // 4
)

type PushMessage struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	SubTitle string `json:"sub_title"`
}

func (pm *PushManager) SyncPush(userID string, strategy pushStrategy, msg PushMessage) (PushResponseResult, error) {

	request := pm.syncPushRequest(userID, strategy, msg)
	res, err := pm.client.pushClient.Send(request)
	if err != nil {
		return PushResponseResult{}, fmt.Errorf("request failed: %w", err)
	}
	return res.Data, nil
}

func (pm *PushManager) syncPushRequest(userID string, strategy pushStrategy, msg PushMessage) http.Request {

	return http.Request{
		URL:            pm.syncPushURL(userID),
		Method:         http.MethodPOST,
		ResponseFormat: http.ResponseFormatJSON,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + pm.client.appToken,
		},
		Payload: &http.JSONPayload{
			Content: map[string]interface{}{
				"strategy":    strategy,
				"pushMessage": msg,
			},
		},
	}
}

func (pm *PushManager) syncPushURL(userID string) string {
	baseURL := pm.client.appConfig.BaseURL
	return fmt.Sprintf(baseURL+"/push/sync/"+"%s", userID)
}
