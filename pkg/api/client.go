package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MimiValsi/spacetraders/internal/database"
	"github.com/MimiValsi/spacetraders/pkg/model"
)

type Client struct {
	BaseURI      *url.URL
	Header       *http.Header
	HttpClient   *http.Client
	accountToken string

	DB *database.Queries
}

func NewClient(ctx context.Context, accountToken string, db *database.Queries) (*Client, error) {
	uri, err := url.Parse("https://api.spacetraders.io/v2/")
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURI: uri,
		Header: &http.Header{
			"Content-Type": {"application/json"},
		},
		accountToken: accountToken,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
		DB: db,
	}, nil
}

func (c *Client) GetAccount() error {
	url := url.URL{Path: "my/account"}
	uri := c.BaseURI.ResolveReference(&url)
	ctx := context.Background()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return err
	}

	agentToken, err := c.DB.GetAgentToken(ctx)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+agentToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	account := model.AccountData{}
	if err = decoder.Decode(&account); err != nil {
		return err
	}

	return c.DB.RegisterAccount(ctx, database.RegisterAccountParams{
		ID:        account.Data.Account.ID,
		Email:     account.Data.Account.Email,
		CreatedAt: account.Data.Account.CreatedAt,
	})
}

func (c *Client) Register(symbol, faction string) error {
	url := url.URL{Path: "register"}
	uri := c.BaseURI.ResolveReference(&url)

	data := strings.NewReader(fmt.Sprintf("{\"symbol\": \"%s\", \"faction\": \"%s\"}", symbol, faction))

	req, err := http.NewRequest("POST", uri.String(), data)
	if err != nil {
		return err
	}

	req.Header = *c.Header
	req.Header.Add("Authorization", "Bearer "+c.accountToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	//defer resp.Body.Close()

	r, m, nok := CheckIsError(resp.Body)
	if nok {
		fmt.Printf("Message: %s\n", m)
		return err
	}

	decoder := json.NewDecoder(r)
	agent := model.AgentRegister{}
	if err = decoder.Decode(&agent); err != nil {
		return err
	}

	return c.sendToDB(&agent)
}

func (c *Client) sendToDB(agent *model.AgentRegister) error {
	ctx := context.Background()

	agentID, err := c.DB.RegisterAgent(ctx, database.RegisterAgentParams{
		AccountID:       agent.Data.Agent.AccountID,
		Token:           agent.Data.Token,
		Symbol:          agent.Data.Agent.Symbol,
		Headquarters:    agent.Data.Agent.Headquarters,
		Credits:         agent.Data.Agent.Credits,
		StartingFaction: agent.Data.Agent.StartingFaction,
		ShipCount:       agent.Data.Agent.ShipCount,
	})
	if err != nil {
		return err
	}

	factionID, err := c.DB.RegisterFaction(ctx, database.RegisterFactionParams{
		Symbol:       agent.Data.Faction.Symbol,
		Name:         agent.Data.Faction.Name,
		Description:  agent.Data.Faction.Description,
		Headquarters: agent.Data.Faction.Headquarters,
		IsRecruiting: agent.Data.Faction.IsRecruiting,
		AgentID:      agentID,
	})
	if err != nil {
		return err
	}

	for i := range agent.Data.Faction.Traits {
		err = c.DB.RegisterTraits(ctx, database.RegisterTraitsParams{
			Symbol:      agent.Data.Faction.Traits[i].Symbol,
			Name:        agent.Data.Faction.Traits[i].Name,
			Description: agent.Data.Faction.Traits[i].Description,
			FactionID:   factionID,
		})
		if err != nil {
			return err
		}
	}

	contractID, err := c.DB.RegisterContract(ctx, database.RegisterContractParams{
		ID:               agent.Data.Contract.ID,
		FactionSymbol:    agent.Data.Contract.FactionSymbol,
		Type:             agent.Data.Contract.Type,
		Accepted:         agent.Data.Contract.Accepted,
		Fulfilled:        agent.Data.Contract.Fulfilled,
		DeadlineToAccept: agent.Data.Contract.DeadlineToAccept,
		AgentID:          agentID,
	})
	if err != nil {
		return err
	}

	termsID, err := c.DB.RegisterTerms(ctx, database.RegisterTermsParams{
		Deadline:   agent.Data.Contract.Terms.Deadline,
		ContractID: contractID,
	})
	if err != nil {
		return err
	}

	err = c.DB.RegisterPayment(ctx, database.RegisterPaymentParams{
		OnAccepted:  agent.Data.Contract.Terms.Payment.OnAccepted,
		OnFulfilled: agent.Data.Contract.Terms.Payment.OnFulfilled,
		TermID:      termsID,
	})
	if err != nil {
		return err
	}

	for i := range agent.Data.Contract.Terms.Deliver {
		err = c.DB.RegisterDeliver(ctx, database.RegisterDeliverParams{
			TradeSymbol:       agent.Data.Contract.Terms.Deliver[i].TradeSymbol,
			DestinationSymbol: agent.Data.Contract.Terms.Deliver[i].DestinationSymbol,
			UnitsRequired:     agent.Data.Contract.Terms.Deliver[i].UnitsRequired,
			UnitsFulfilled:    agent.Data.Contract.Terms.Deliver[i].UnitsFulfilled,
			TermID:            termsID,
		})
		if err != nil {
			return err
		}
	}

	for i := range agent.Data.Ships {
		shipID, err := c.DB.RegisterShip(ctx, database.RegisterShipParams{
			Symbol:  agent.Data.Ships[i].Symbol,
			AgentID: agentID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterRegistration(ctx, database.RegisterRegistrationParams{
			Name:          agent.Data.Ships[i].Registration.Name,
			FactionSymbol: agent.Data.Ships[i].Registration.FactionSymbol,
			Role:          agent.Data.Ships[i].Registration.Role,
			ShipID:        shipID,
		})
		if err != nil {
			return err
		}

		navID, err := c.DB.RegisterNav(ctx, database.RegisterNavParams{
			SystemSymbol:   agent.Data.Ships[i].Nav.SystemSymbol,
			WaypointSymbol: agent.Data.Ships[i].Nav.WaypointSymbol,
			Status:         agent.Data.Ships[i].Nav.Status,
			FlightMode:     agent.Data.Ships[i].Nav.FlightMode,
			ShipID:         shipID,
		})
		if err != nil {
			return err
		}

		routeID, err := c.DB.RegisterRoute(ctx, database.RegisterRouteParams{
			DepartureTime: agent.Data.Ships[i].Nav.Route.DepartureTime,
			Arrival:       agent.Data.Ships[i].Nav.Route.Arrival,
			NavID:         navID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterDestination(ctx, database.RegisterDestinationParams{
			Symbol:       agent.Data.Ships[i].Nav.Route.Destination.Symbol,
			Type:         agent.Data.Ships[i].Nav.Route.Destination.Type,
			SystemSymbol: agent.Data.Ships[i].Nav.Route.Destination.SystemSymbol,
			X:            agent.Data.Ships[i].Nav.Route.Destination.X,
			Y:            agent.Data.Ships[i].Nav.Route.Destination.Y,
			RouteID:      routeID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterOrigin(ctx, database.RegisterOriginParams{
			Symbol:       agent.Data.Ships[i].Nav.Route.Origin.Symbol,
			Type:         agent.Data.Ships[i].Nav.Route.Origin.Type,
			SystemSymbol: agent.Data.Ships[i].Nav.Route.Origin.SystemSymbol,
			X:            agent.Data.Ships[i].Nav.Route.Origin.X,
			Y:            agent.Data.Ships[i].Nav.Route.Origin.Y,
			RouteID:      routeID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterCrew(ctx, database.RegisterCrewParams{
			Current:  agent.Data.Ships[i].Crew.Current,
			Required: agent.Data.Ships[i].Crew.Required,
			Capacity: agent.Data.Ships[i].Crew.Capacity,
			Rotation: agent.Data.Ships[i].Crew.Rotation,
			Morale:   agent.Data.Ships[i].Crew.Morale,
			Wages:    agent.Data.Ships[i].Crew.Wages,
			ShipID:   shipID,
		})
		if err != nil {
			return err
		}

		frameReqID, err := c.DB.RegisterRequirements(ctx, database.RegisterRequirementsParams{
			Power: agent.Data.Ships[i].Frame.Requirements.Power,
			Crew:  agent.Data.Ships[i].Frame.Requirements.Crew,
			Slots: agent.Data.Ships[i].Frame.Requirements.Slots,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterFrame(ctx, database.RegisterFrameParams{
			Symbol:         agent.Data.Ships[i].Frame.Symbol,
			Name:           agent.Data.Ships[i].Frame.Name,
			Condition:      agent.Data.Ships[i].Frame.Condition,
			Integrity:      agent.Data.Ships[i].Frame.Integrity,
			Description:    agent.Data.Ships[i].Frame.Description,
			ModuleSlots:    agent.Data.Ships[i].Frame.ModuleSlots,
			MountingPoints: agent.Data.Ships[i].Frame.MountingPoints,
			FuelCapacity:   agent.Data.Ships[i].Frame.FuelCapacity,
			Quality:        agent.Data.Ships[i].Frame.Quality,
			RequirementID:  frameReqID,
			ShipID:         shipID,
		})
		if err != nil {
			return err
		}

		reactorReqID, err := c.DB.RegisterRequirements(ctx, database.RegisterRequirementsParams{
			Power: agent.Data.Ships[i].Reactor.Requirements.Power,
			Crew:  agent.Data.Ships[i].Reactor.Requirements.Crew,
			Slots: agent.Data.Ships[i].Reactor.Requirements.Slots,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterReactor(ctx, database.RegisterReactorParams{
			Symbol:        agent.Data.Ships[i].Reactor.Symbol,
			Name:          agent.Data.Ships[i].Reactor.Name,
			Condition:     agent.Data.Ships[i].Reactor.Condition,
			Integrity:     agent.Data.Ships[i].Reactor.Integrity,
			Description:   agent.Data.Ships[i].Reactor.Description,
			PowerOutput:   agent.Data.Ships[i].Reactor.PowerOutput,
			Quality:       agent.Data.Ships[i].Reactor.Quality,
			RequirementID: reactorReqID,
			ShipID:        shipID,
		})
		if err != nil {
			return err
		}

		engineReqID, err := c.DB.RegisterRequirements(ctx, database.RegisterRequirementsParams{
			Power: agent.Data.Ships[i].Engine.Requirements.Power,
			Crew:  agent.Data.Ships[i].Engine.Requirements.Crew,
			Slots: agent.Data.Ships[i].Engine.Requirements.Slots,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterEngine(ctx, database.RegisterEngineParams{
			Symbol:        agent.Data.Ships[i].Engine.Symbol,
			Name:          agent.Data.Ships[i].Engine.Name,
			Condition:     agent.Data.Ships[i].Engine.Condition,
			Integrity:     agent.Data.Ships[i].Engine.Integrity,
			Description:   agent.Data.Ships[i].Engine.Description,
			Speed:         agent.Data.Ships[i].Engine.Speed,
			Quality:       agent.Data.Ships[i].Engine.Quality,
			RequirementID: engineReqID,
			ShipID:        shipID,
		})
		if err != nil {
			return err
		}

		for j := range agent.Data.Ships[i].Modules {
			moduleReqID, err := c.DB.RegisterRequirements(ctx, database.RegisterRequirementsParams{
				Power: agent.Data.Ships[i].Modules[j].Requirements.Power,
				Crew:  agent.Data.Ships[i].Modules[j].Requirements.Crew,
				Slots: agent.Data.Ships[i].Modules[j].Requirements.Slots,
			})
			if err != nil {
				return err
			}

			err = c.DB.RegisterModule(ctx, database.RegisterModuleParams{
				Symbol:        agent.Data.Ships[i].Modules[j].Symbol,
				Name:          agent.Data.Ships[i].Modules[j].Name,
				Description:   agent.Data.Ships[i].Modules[j].Description,
				Capacity:      agent.Data.Ships[i].Modules[j].Capacity,
				Range:         agent.Data.Ships[i].Modules[j].Range,
				RequirementID: moduleReqID,
				ShipID:        shipID,
			})
			if err != nil {
				return err
			}

		}

		for j := range agent.Data.Ships[i].Mounts {
			mountReqID, err := c.DB.RegisterRequirements(ctx, database.RegisterRequirementsParams{
				Power: agent.Data.Ships[i].Mounts[j].Requirements.Power,
				Crew:  agent.Data.Ships[i].Mounts[j].Requirements.Crew,
				Slots: agent.Data.Ships[i].Mounts[j].Requirements.Slots,
			})
			if err != nil {
				return err
			}

			err = c.DB.RegisterMount(ctx, database.RegisterMountParams{
				Symbol:        agent.Data.Ships[i].Mounts[j].Symbol,
				Name:          agent.Data.Ships[i].Mounts[j].Name,
				Description:   agent.Data.Ships[i].Mounts[j].Description,
				Strength:      agent.Data.Ships[i].Mounts[j].Strength,
				Deposits:      agent.Data.Ships[i].Mounts[j].Deposits,
				RequirementID: mountReqID,
				ShipID:        shipID,
			})
			if err != nil {
				return err
			}

		}

		cargoID, err := c.DB.RegisterCargo(ctx, database.RegisterCargoParams{
			Capacity: agent.Data.Ships[i].Cargo.Capacity,
			Units:    agent.Data.Ships[i].Cargo.Units,
			ShipID:   shipID,
		})
		if err != nil {
			return err
		}

		for j := range agent.Data.Ships[i].Cargo.Inventory {
			err = c.DB.RegisterInventories(ctx, database.RegisterInventoriesParams{
				Symbol:  agent.Data.Ships[i].Cargo.Inventory[j].Symbol,
				Name:    agent.Data.Ships[i].Cargo.Inventory[j].Name,
				Units:   agent.Data.Ships[i].Cargo.Inventory[j].Units,
				CargoID: cargoID,
			})
			if err != nil {
				return err
			}
		}

		fuelID, err := c.DB.RegisterFuels(ctx, database.RegisterFuelsParams{
			Current:  agent.Data.Ships[i].Fuel.Current,
			Capacity: agent.Data.Ships[i].Fuel.Capacity,
			ShipID:   shipID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterConsumed(ctx, database.RegisterConsumedParams{
			Amount:   agent.Data.Ships[i].Fuel.Consumed.Amount,
			Timestmp: agent.Data.Ships[i].Fuel.Consumed.Timestamp,
			FuelID:   fuelID,
		})
		if err != nil {
			return err
		}

		err = c.DB.RegisterCooldown(ctx, database.RegisterCooldownParams{
			ShipSymbol:       agent.Data.Ships[i].Cooldown.ShipSymbol,
			TotalSeconds:     agent.Data.Ships[i].Cooldown.TotalSeconds,
			RemainingSeconds: agent.Data.Ships[i].Cooldown.RemainingSeconds,
			Expiration:       agent.Data.Ships[i].Cooldown.Expiration,
			ShipID:           shipID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
