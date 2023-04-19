package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestBaseSubscription_GetAccountAddress(t *testing.T) {
	type fields struct {
		AccountAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty account address",
			fields{
				AccountAddress: "",
			},
			nil,
		},
		{
			"20 bytes account address",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BaseSubscription{
				AccountAddress: tt.fields.AccountAddress,
			}
			if got := s.GetAccountAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSubscription_Validate(t *testing.T) {
	type fields struct {
		ID             uint64
		AccountAddress string
		Status         hubtypes.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero id",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"positive id",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"empty account address",
			fields{
				ID:             1000,
				AccountAddress: "",
			},
			true,
		},
		{
			"invalid account address",
			fields{
				ID:             1000,
				AccountAddress: "invalid",
			},
			true,
		},
		{
			"invalid prefix account address",
			fields{
				ID:             1000,
				AccountAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes account address",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgslawd5s",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"20 bytes account address",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"30 bytes account address",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"unspecified status",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusInactivePending,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"inactive status",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusInactive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"empty status at",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Time{},
			},
			true,
		},
		{
			"positive status at",
			fields{
				ID:             1000,
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BaseSubscription{
				ID:             tt.fields.ID,
				AccountAddress: tt.fields.AccountAddress,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNodeSubscription_GetNodeAddress(t *testing.T) {
	type fields struct {
		NodeAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty node address",
			fields{
				NodeAddress: "",
			},
			nil,
		},
		{
			"20 bytes node address",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodeSubscription{
				NodeAddress: tt.fields.NodeAddress,
			}
			if got := s.GetNodeAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNodeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSubscription_Type(t *testing.T) {
	v := &NodeSubscription{}
	require.Equal(t, TypeNode, v.Type())
}

func TestNodeSubscription_Validate(t *testing.T) {
	type fields struct {
		NodeAddress string
		Hours       int64
		Price       sdk.Coin
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty node address",
			fields{
				NodeAddress: "",
			},
			true,
		},
		{
			"invalid node address",
			fields{
				NodeAddress: "invalid",
			},
			true,
		},
		{
			"invalid prefix node address",
			fields{
				NodeAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes node address",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgse4wwrm",
				Hours:       1000,
			},
			false,
		},
		{
			"20 bytes node address",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
			},
			false,
		},
		{
			"30 bytes node address",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				Hours:       1000,
			},
			false,
		},
		{
			"negative hours",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       -1000,
			},
			true,
		},
		{
			"zero hours",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       0,
			},
			true,
		},
		{
			"positive hours",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
			},
			false,
		},
		{
			"empty price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{},
			},
			false,
		},
		{
			"empty denom price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: ""},
			},
			false,
		},
		{
			"invalid denom price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: "d"},
			},
			true,
		},
		{
			"nil amount price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: "one"},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
				Price:       sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodeSubscription{
				BaseSubscription: BaseSubscription{
					ID:             1000,
					AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
					Status:         hubtypes.StatusActive,
					StatusAt:       time.Now(),
				},
				NodeAddress: tt.fields.NodeAddress,
				Hours:       tt.fields.Hours,
				Price:       tt.fields.Price,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlanSubscription_Type(t *testing.T) {
	v := &PlanSubscription{}
	require.Equal(t, TypePlan, v.Type())
}

func TestPlanSubscription_Validate(t *testing.T) {
	type fields struct {
		PlanID uint64
		Denom  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero plan id",
			fields{
				PlanID: 0,
			},
			true,
		},
		{
			"positive plan id",
			fields{
				PlanID: 1000,
				Denom:  "one",
			},
			false,
		},
		{
			"empty denom",
			fields{
				PlanID: 1000,
				Denom:  "",
			},
			false,
		},
		{
			"invalid denom",
			fields{
				PlanID: 1000,
				Denom:  "d",
			},
			true,
		},
		{
			"one denom",
			fields{
				PlanID: 1000,
				Denom:  "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PlanSubscription{
				BaseSubscription: BaseSubscription{
					ID:             1000,
					AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
					Status:         hubtypes.StatusActive,
					StatusAt:       time.Now(),
				},
				PlanID: tt.fields.PlanID,
				Denom:  tt.fields.Denom,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
