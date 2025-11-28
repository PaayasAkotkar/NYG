package server

import (
	nygpostprotoc "app/nygprotoc"
	"context"
	"log"
)

type Services struct {
	nygpostprotoc.UnimplementedNYGPostProfileServer
}

func (*Services) RecieveNYGProfile(_ context.Context, npp *nygpostprotoc.NYGProfileParam) (*nygpostprotoc.TypeNYGProfile, error) {
	log.Println("in receive profile")
	token := GetProfile(npp.Id)
	return &nygpostprotoc.TypeNYGProfile{Id: token.Name, Name: token.Nickname, Nickname: token.Nickname}, nil
}

func (*Services) RecieveNYGDeck(_ context.Context, npp *nygpostprotoc.NYGDeckParam) (*nygpostprotoc.TypeNYGDeck, error) {
	log.Println("in receive deck")
	token := GetDeck(npp.Id)

	return &nygpostprotoc.TypeNYGDeck{Id: token.ID, First: token.Frist, Second: token.Second, IsDefault: token.IsDefault}, nil
}

func (*Services) RecieveNYGCredits(_ context.Context, npp *nygpostprotoc.NYGPlayerCreditsParam) (*nygpostprotoc.TypeNYGCredit, error) {
	log.Println("in receive credits")
	token := GetCredits(npp.Id)
	return &nygpostprotoc.TypeNYGCredit{Id: token.ID, Coins: token.Coin, BoardTheme: token.BoardTheme, GuessTheme: token.GuessTheme, Spurs: token.Spur, ClashProfile: &token.ClashProfile, OnevoneProfile: &token.OneVOneProfile, TwovtwoProfile: &token.TwoVTwoProfile, ImgURL: token.ImageURL, Freeze: token.Freeze, Nexus: token.Nexus}, nil
}
