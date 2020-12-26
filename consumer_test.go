package cls

import (
	"reflect"
	"testing"
)

func TestClSCleint_CreateConsumerGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		group   *ConsumerGroup
		topicID string
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
			if err := cls.CreateConsumerGroup(tt.args.group, tt.args.topicID); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.CreateConsumerGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_GetConsumerCursor(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID     string
		partitionID int
		from        string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCursor ConsumerCursor
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
			gotCursor, err := cls.GetConsumerCursor(tt.args.topicID, tt.args.partitionID, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetConsumerCursor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCursor, tt.wantCursor) {
				t.Errorf("ClSCleint.GetConsumerCursor() = %v, want %v", gotCursor, tt.wantCursor)
			}
		})
	}
}

func TestClSCleint_GetConsumerGroupCursors(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID       string
		partitionID   int
		consumerGroup string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCursors ConsumerGroupCursors
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
			gotCursors, err := cls.GetConsumerGroupCursors(tt.args.topicID, tt.args.partitionID, tt.args.consumerGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetConsumerGroupCursors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCursors, tt.wantCursors) {
				t.Errorf("ClSCleint.GetConsumerGroupCursors() = %v, want %v", gotCursors, tt.wantCursors)
			}
		})
	}
}

func TestClSCleint_GetConsumerData(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID     string
		partitionID int
		cursor      string
		count       int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList LogGroupList
		wantErr  bool
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
			gotList, err := cls.GetConsumerData(tt.args.topicID, tt.args.partitionID, tt.args.cursor, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetConsumerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("ClSCleint.GetConsumerData() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func TestClSCleint_GetConsumerHeartbeat(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		heartbeat *ConsumerHeartbeat
		topicID   string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantPartitionIDList []string
		wantErr             bool
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
			gotPartitionIDList, err := cls.GetConsumerHeartbeat(tt.args.heartbeat, tt.args.topicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetConsumerHeartbeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPartitionIDList, tt.wantPartitionIDList) {
				t.Errorf("ClSCleint.GetConsumerHeartbeat() = %v, want %v", gotPartitionIDList, tt.wantPartitionIDList)
			}
		})
	}
}

func TestClSCleint_GetConsumerGroups(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantGroups ConsumerGroups
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
			gotGroups, err := cls.GetConsumerGroups(tt.args.topicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.GetConsumerGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("ClSCleint.GetConsumerGroups() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}

func TestClSCleint_UpdateConsumerGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		group         *ConsumerGroup
		topicID       string
		consumerGroup string
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
			if err := cls.UpdateConsumerGroup(tt.args.group, tt.args.topicID, tt.args.consumerGroup); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateConsumerGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_UpdateConsumerGroupCursor(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		cursor        *ConsumerGroupCursor
		topicID       string
		consumerGroup string
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
			if err := cls.UpdateConsumerGroupCursor(tt.args.cursor, tt.args.topicID, tt.args.consumerGroup); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.UpdateConsumerGroupCursor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClSCleint_DeleteConsumerGroup(t *testing.T) {
	type fields struct {
		SecretId  string
		SecretKey string
		Host      string
	}
	type args struct {
		topicID       string
		consumerGroup string
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
			if err := cls.DeleteConsumerGroup(tt.args.topicID, tt.args.consumerGroup); (err != nil) != tt.wantErr {
				t.Errorf("ClSCleint.DeleteConsumerGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
