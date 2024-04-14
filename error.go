package s10s

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	api "github.com/ruudiRatlos/s10s/api"
)

type APIError int

func (e APIError) Error() string {
	return fmt.Sprintf("API Error: %d", e)
}

func newAPIError(code int, msg string) error {
	return fmt.Errorf("%s: %w", msg, APIError(code))
}

func (e APIError) Is(target error) bool {
	var ae APIError
	if errors.As(target, &ae) {
		return int(e) == int(ae)
	}
	return false
}

type HTTPStatusError int

func (e HTTPStatusError) Error() string {
	return fmt.Sprintf("API Error: %d", e)
}

func newHTTPStatusError(code int, msg string) error {
	return fmt.Errorf("%s: %w", msg, HTTPStatusError(code))
}

func (e HTTPStatusError) Is(target error) bool {
	var he HTTPStatusError
	if errors.As(target, &he) {
		return int(e) == int(he)
	}
	return false
}

func enhanceErr(err error, r *http.Response) error {
	if r == nil || err == nil {
		return err
	}
	defer func() { _ = r.Body.Close() }()

	type JSONError struct {
		Error struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
			Data    struct {
				ShipSymbol        string        `json:"shipSymbol"`
				DestinationSymbol string        `json:"destinationSymbol"`
				Cooldown          *api.Cooldown `json:"cooldown"`
			} `json:"data"`
		} `json:"error"`
	}
	if r.StatusCode > 299 {
		err = errors.Join(err, newHTTPStatusError(r.StatusCode, r.Status))
	}

	dec := json.NewDecoder(r.Body)
	var jsonMsg JSONError
	jErr := dec.Decode(&jsonMsg)
	if jErr != nil {
		return errors.Join(jErr, err)
	}

	return errors.Join(newAPIError(jsonMsg.Error.Code, jsonMsg.Error.Message), err)
}

const ErrHTTPStatus429 HTTPStatusError = 429
const ErrHTTPStatus409 HTTPStatusError = 409

const ErrCooldownConflict APIError = 4000
const ErrWaypointNoAccess APIError = 4001
const ErrTokenEmpty APIError = 4100
const ErrTokenMissingSubject APIError = 4101
const ErrTokenInvalidSubject APIError = 4102
const ErrMissingTokenRequest APIError = 4103
const ErrInvalidTokenRequest APIError = 4104
const ErrInvalidTokenSubject APIError = 4105
const ErrAccountNotExists APIError = 4106
const ErrAgentNotExists APIError = 4107
const ErrAccountHasNoAgent APIError = 4108
const ErrRegisterAgentExists APIError = 4109
const ErrRegisterAgentSymbolReserved APIError = 4110
const ErrRegisterAgentConflictSymbol APIError = 4111
const ErrNavigateInTransit APIError = 4200
const ErrNavigateInvalidDestination APIError = 4201
const ErrNavigateOutsideSystem APIError = 4202
const ErrNavigateInsufficientFuel APIError = 4203
const ErrNavigateSameDestination APIError = 4204
const ErrShipExtractInvalidWaypoint APIError = 4205
const ErrShipExtractPermission APIError = 4206
const ErrShipJumpNoSystem APIError = 4207
const ErrShipJumpSameSystem APIError = 4208
const ErrShipJumpMissingModule APIError = 4210
const ErrShipJumpNoValidWaypoint APIError = 4211
const ErrShipJumpMissingAntimatter APIError = 4212
const ErrShipInTransit APIError = 4214
const ErrShipMissingSensorArrays APIError = 4215
const ErrPurchaseShipCredits APIError = 4216
const ErrShipCargoExceedsLimit APIError = 4217
const ErrShipCargoMissing APIError = 4218
const ErrShipCargoUnitCount APIError = 4219
const ErrShipSurveyVerification APIError = 4220
const ErrShipSurveyExpiration APIError = 4221
const ErrShipSurveyWaypointType APIError = 4222
const ErrShipSurveyOrbit APIError = 4223
const ErrShipSurveyExhausted APIError = 4224
const ErrShipRefuelDocked APIError = 4225
const ErrShipRefuelInvalidWaypoint APIError = 4226
const ErrShipMissingMounts APIError = 4227
const ErrShipCargoFull APIError = 4228
const ErrShipJumpFromGateToGate APIError = 4229
const ErrWaypointCharted APIError = 4230
const ErrShipTransferShipNotFound = 4231
const ErrShipTransferAgentConflict = 4232
const ErrShipTransferSameShipConflict = 4233
const ErrShipTransferLocationConflict = 4234
const ErrWarpInsideSystem APIError = 4235
const ErrShipNotInOrbit APIError = 4236
const ErrShipInvalidRefineryGood APIError = 4237
const ErrShipInvalidRefineryType APIError = 4238
const ErrShipMissingRefinery APIError = 4239
const ErrShipMissingSurveyor APIError = 4240
const ErrShipMissingWarpDrive APIError = 4241
const ErrShipMissingMineralProcessor APIError = 4242
const ErrShipMissingMiningLasers APIError = 4243
const ErrShipNotDocked APIError = 4244
const ErrPurchaseShipNotPresent APIError = 4245
const ErrShipMountNoShipyard APIError = 4246
const ErrShipMissingMount APIError = 4247
const ErrShipMountInsufficientCredits APIError = 4248
const ErrShipMissingPower APIError = 4249
const ErrShipMissingSlots APIError = 4250
const ErrShipMissingCrew APIError = 4252
const ErrShipExtractDestabilized APIError = 4253
const ErrShipJumpInvalidOrigin APIError = 4254
const ErrShipJumpInvalidWaypoint APIError = 4255
const ErrShipJumpOriginUnderConstruction APIError = 4256
const ErrShipMissingGasProcessor APIError = 4257
const ErrShipMissingGasSiphons APIError = 4258
const ErrShipSiphonInvalidWaypoint APIError = 4259
const ErrShipSiphonPermission APIError = 4260
const ErrWaypointNoYield APIError = 4261
const ErrShipJumpDestinationUnderConstruction APIError = 4262
const ErrAcceptContractNotAuthorized APIError = 4500
const ErrAcceptContractConflict APIError = 4501
const ErrFulfillContractDelivery APIError = 4502
const ErrContractDeadline APIError = 4503
const ErrContractFulfilled APIError = 4504
const ErrContractNotAccepted APIError = 4505
const ErrContractNotAuthorized APIError = 4506
const ErrShipDeliverTerms APIError = 4508
const ErrShipDeliverFulfilled APIError = 4509
const ErrShipDeliverInvalidLocation APIError = 4510
const ErrExistingContract APIError = 4511
const ErrMarketTradeInsufficientCredits APIError = 4600
const ErrMarketTradeNoPurchase APIError = 4601
const ErrMarketTradeNotSold APIError = 4602
const ErrMarketNotFound APIError = 4603
const ErrMarketTradeUnitLimit APIError = 4604
const ErrWaypointNoFaction APIError = 4700
const ErrConstructionMaterialNotRequired = 4800
const ErrConstructionMaterialFulfilled = 4801
const ErrShipConstructionInvalidLocation APIError = 4802
