package Beacon

import (
	"github.com/tbalthazar/onesignal-go"
)

func sendNotification(message string) error {
	//Takes a string and sends it as a notification to all users of the app

	client := onesignal.NewClient(nil)
	client.AppKey = "YjYwNTU3OGItYTZjYi00YjYzLWEzOGMtMGNjYmMxYmE3NjVm"
	client.UserKey = "NGRiYTFlZTItMDg0MS00ZTJlLTgwNmQtOTU1NDk3Zjc4ZjE2"
	appID := "d3d1b2a0-f47a-4136-8838-7ffa7d9a6a30"

	listOpt := &onesignal.PlayerListOptions{
		AppID:  appID,
		Limit:  25,
		Offset: 0,
	}

	players,_ , err := client.Players.List(listOpt)
	if err != nil{
		return err
	}

	for _, p := range players.Players{
		notificationReq := &onesignal.NotificationRequest{
			AppID:            appID,
			Contents:         map[string]string{"en": message},
			IsIOS:            false,
			IncludePlayerIDs: []string{p.ID},
		}
		_, _, err := client.Notifications.Create(notificationReq)
		if err != nil{
			return err
		}
	}
	return nil
}