package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_CreateShipper(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipper *Shipper
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			got, err := cls.CreateShipper(tt.args.shipper)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.CreateShipper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.CreateShipper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_GetShipper(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipperID string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantShipper Shipper
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotShipper, err := cls.GetShipper(tt.args.shipperID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetShipper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShipper, tt.wantShipper) {
				t.Errorf("ClSCleint.GetShipper() = %v, want %v", gotShipper, tt.wantShipper)
			}
		})
	}
}

func TestClSCleint_GetTopicShipperList(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantShipperList ShipperList
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotShipperList, err := cls.GetTopicShipperList(tt.args.topicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetTopicShipperList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShipperList, tt.wantShipperList) {
				t.Errorf("ClSCleint.GetTopicShipperList() = %v, want %v", gotShipperList, tt.wantShipperList)
			}
		})
	}
}

func TestClSCleint_GetShipperList(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipperID string
		startTime string
		endTime   string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTasks Tasks
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			gotTasks, err := cls.GetShipperList(tt.args.shipperID, tt.args.startTime, tt.args.endTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetShipperList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("ClSCleint.GetShipperList() = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func TestClSCleint_UpdateShipper(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipper *Shipper
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			got, err := cls.UpdateShipper(tt.args.shipper)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateShipper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.UpdateShipper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_UpdateTask(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipper *Shipper
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			if err := cls.UpdateTask(tt.args.shipper); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_DeleteShipper(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		shipperID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := &ClSCleint{
				SecretId:  tt.fields.SecretId,
				SecretKey: tt.fields.SecretKey,
				Host:      tt.fields.Host,
			}
			if err := cls.DeleteShipper(tt.args.shipperID); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DeleteShipper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
