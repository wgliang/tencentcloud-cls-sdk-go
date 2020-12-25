package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_CreateLogTopic(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopic LogTopic
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
			got, err := cls.CreateLogTopic(tt.args.logTopic)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.CreateLogTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClSCleint.CreateLogTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClSCleint_GetLogTopic(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantLogTopic LogTopic
		wantErr      bool
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
			gotLogTopic, err := cls.GetLogTopic(tt.args.logTopicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogTopic, tt.wantLogTopic) {
				t.Errorf("ClSCleint.GetLogTopic() = %v, want %v", gotLogTopic, tt.wantLogTopic)
			}
		})
	}
}

func TestClSCleint_GetLogTopics(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logsetId string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantLogTopics LogTopics
		wantErr       bool
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
			gotLogTopics, err := cls.GetLogTopics(tt.args.logsetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetLogTopics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogTopics, tt.wantLogTopics) {
				t.Errorf("ClSCleint.GetLogTopics() = %v, want %v", gotLogTopics, tt.wantLogTopics)
			}
		})
	}
}

func TestClSCleint_UpdateLogTopic(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopic LogTopic
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
			if err := cls.UpdateLogTopic(tt.args.logTopic); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateLogTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_DeleteLogTopic(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		logTopicID string
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
			if err := cls.DeleteLogTopic(tt.args.logTopicID); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DeleteLogTopic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
