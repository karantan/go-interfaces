package consumer

import "go-interfaces/good"

func SendWave(ds good.DiscordSource, channelID string) error {
	if err := good.SendMsg(ds, channelID, "sup :wave:"); err != nil {
		return err
	}
	if err := good.SendMsg(ds, channelID, "I'm a friendly bot"); err != nil {
		return err
	}
	return nil
}
