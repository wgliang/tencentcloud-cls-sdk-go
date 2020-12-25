package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_CreateLogSet(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logSetName string
		period     int
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
			got, err := cls.CreateLogSet(tt.args.logSetName, tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.CreateLogSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.CreateLogSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_GetLogSet(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logSetID string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantLogSet LogSet
		wantErr    bool
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
			gotLogSet, err := cls.GetLogSet(tt.args.logSetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogSet, tt.wantLogSet) {
				t.Errorf("ClSCleint.GetLogSet() = %v, want %v", gotLogSet, tt.wantLogSet)
			}
		})
	}
}

func TestClSCleint_GetLogSets(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	tests := []struct {
		name        string
		fields      fields
		wantLogSets LogSets
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
			gotLogSets, err := cls.GetLogSets()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogSets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogSets, tt.wantLogSets) {
				t.Errorf("ClSCleint.GetLogSets() = %v, want %v", gotLogSets, tt.wantLogSets)
			}
		})
	}
}

func TestClSCleint_UpdateLogSet(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logSetID   string
		logSetName string
		period     int
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
			if err := cls.UpdateLogSet(tt.args.logSetID, tt.args.logSetName, tt.args.period); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateLogSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_DeleteLogSet(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logSetID string
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
			if err := cls.DeleteLogSet(tt.args.logSetID); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DeleteLogSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
