// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"fmt"
)

func (r *Report) expectClanId() error {
	if r.ClanId = r.accept(lexClanId); r.ClanId == "" {
		return fmt.Errorf("expected clan id: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCloseParen() error {
	if !r.acceptDelimiter(')') {
		return fmt.Errorf("expected close paren: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectComma() error {
	if !r.acceptDelimiter(',') {
		return fmt.Errorf("expected comma: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCost() error {
	if r.Accounting.Cost = r.accept(lexCost); r.Accounting.Cost == "" {
		return fmt.Errorf("expected cost: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCredit() error {
	if r.Accounting.Credit = r.accept(lexCredit); r.Accounting.Credit == "" {
		return fmt.Errorf("expected credit: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCurrentHex() error {
	var hex string
	if hex = r.accept(lexCurrentHex); hex == "" {
		return fmt.Errorf("expected hex: got %q", r.nextToken())
	}
	if r.tribe == nil {
		r.tribe = &Tribe{Id: r.ClanId}
		r.Tribes = append(r.Tribes, r.tribe)
	}
	r.tribe.Hex = hex
	return nil
}

func (r *Report) expectCurrentMonth() error {
	if err := r.expectOpenParen(); err != nil {
		return err
	} else if r.Current.Month = r.accept(lexMonth); r.Current.Month == "" {
		return fmt.Errorf("expected month: got %q", r.nextToken())
	} else if err = r.expectCloseParen(); err != nil {
		return err
	}
	return nil
}

func (r *Report) expectCurrentSeason() error {
	if r.Current.Season = r.accept(lexSeason); r.Current.Season == "" {
		return fmt.Errorf("expected season: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCurrentTurn() error {
	r.Current = &Turn{}
	if r.Current.Turn = r.accept(lexCurrentTurn); r.Current.Turn == "" {
		return fmt.Errorf("expected turn: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectCurrentWeather() error {
	if r.Current.Weather = r.accept(lexWeather); r.Current.Weather == "" {
		return fmt.Errorf("expected weather: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectDesiredGood() error {
	bullet := fmt.Sprintf("(%d)", len(r.DesiredCommodities)+1)
	if !r.acceptLiteral(bullet) {
		return fmt.Errorf("expected %q: got %q", bullet, r.nextToken())
	}
	if goods := r.acceptGoods(); goods == "" {
		return fmt.Errorf("expected goods: got %q", r.nextToken())
	} else {
		r.DesiredCommodities = append(r.DesiredCommodities, goods)
	}
	return nil
}

func (r *Report) expectDesiredCommodities() error {
	if !r.acceptLiteral("Desired Commodities: ") {
		return fmt.Errorf("expected desired commodities: got %q", r.nextToken())
	}
	if r.acceptLiteral("No commodities allocated") {
		r.DesiredCommodities = nil
		return nil
	}
	if err := r.expectDesiredGood(); err != nil {
		return err
	}
	for r.acceptDelimiter(',') {
		if err := r.expectDesiredGood(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Report) expectGoodsTribe() error {
	r.tribe.GoodsTribe = ""
	if !r.acceptLiteral("Goods Tribe: ") {
		return fmt.Errorf("expected goods tribe: got %q", r.nextToken())
	}
	if r.acceptLiteral("No GT") {
		r.tribe.GoodsTribe = ""
	} else {
		return fmt.Errorf("!implemented")
	}
	return nil
}

func (r *Report) expectNextMonth() error {
	if err := r.expectOpenParen(); err != nil {
		return err
	} else if r.Next.Month = r.accept(lexMonth); r.Next.Month == "" {
		return fmt.Errorf("expected month: got %q", r.nextToken())
	} else if err = r.expectCloseParen(); err != nil {
		return err
	}
	return nil
}

func (r *Report) expectNextTurn() error {
	r.Next = &Turn{}
	if r.Next.Turn = r.accept(lexNextTurn); r.Next.Turn == "" {
		return fmt.Errorf("expected turn: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectOpenParen() error {
	if !r.acceptDelimiter('(') {
		return fmt.Errorf("expected open paren: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectPreviousHex() error {
	var hex string
	if hex = r.accept(lexPreviousHex); hex == "" {
		return fmt.Errorf("expected hex: got %q", r.nextToken())
	}
	r.tribe.StartingHex = hex
	return nil
}

func (r *Report) expectReceived() error {
	r.Accounting = &Accounting{}
	if r.Accounting.Received = r.accept(lexReceived); r.Accounting.Received == "" {
		return fmt.Errorf("expected received: got %q", r.nextToken())
	}
	return nil
}

func (r *Report) expectTurnDate() error {
	if r.Next.Due = r.accept(lexTurnDate); r.Next.Due == "" {
		return fmt.Errorf("expected date: got %q", r.nextToken())
	}
	return nil
}
